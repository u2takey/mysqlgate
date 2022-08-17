// Go MySQL Driver - A MySQL-Driver for Go's database/sql package
//
// Copyright 2012 The Go-MySQL-Driver Authors. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.

package mysql

import (
	"bytes"
	"encoding/binary"
	"time"

	"github.com/u2takey/mysqlgate/pkg/sql"
)

// Packets documentation:
// http://dev.mysql.com/doc/internals/en/client-server-protocol.html

// Read packet to buffer 'data'
func (mc *MysqlConn) readPacket() ([]byte, error) {
	var prevData []byte
	for {
		// read packet header
		data, err := mc.buf.readNext(4)
		if err != nil {
			errLog.Print(err)
			return nil, ErrInvalidConn
		}

		// packet length [24 bit]
		pktLen := int(uint32(data[0]) | uint32(data[1])<<8 | uint32(data[2])<<16)

		// check packet sync [8 bit]
		if data[3] != mc.sequence {
			if data[3] > mc.sequence {
				return nil, ErrPktSyncMul
			}
			return nil, ErrPktSync
		}
		mc.sequence++

		// packets with length 0 terminate a previous packet which is a
		// multiple of (2^24)-1 bytes long
		if pktLen == 0 {
			// there was no previous packet
			if prevData == nil {
				errLog.Print(ErrMalformPkt)
				return nil, ErrInvalidConn
			}

			return prevData, nil
		}

		// read packet body [pktLen bytes]
		data, err = mc.buf.readNext(pktLen)
		if err != nil {
			errLog.Print(err)
			return nil, ErrInvalidConn
		}

		// return data if this was the last packet
		if pktLen < maxPacketSize {
			// zero allocations for non-split packets
			if prevData == nil {
				return data, nil
			}

			return append(prevData, data...), nil
		}
		prevData = append(prevData, data...)
	}
}

// Write packet buffer 'data'
func (mc *MysqlConn) writePacket(data []byte) error {
	pktLen := len(data) - 4

	if pktLen > mc.maxAllowedPacket {
		return ErrPktTooLarge
	}

	//// Perform a stale connection check. We only perform this check for
	//// the first query on a connection that has been checked out of the
	//// connection pool: a fresh connection from the pool is more likely
	//// to be stale, and it has not performed any previous writes that
	//// could cause data corruption, so it's safe to return ErrBadConn
	//// if the check fails.
	//if mc.reset {
	//	mc.reset = false
	//	conn := mc.netConn
	//	if mc.rawConn != nil {
	//		conn = mc.rawConn
	//	}
	//	var err error
	//	// If this connection has a ReadTimeout which we've been setting on
	//	// reads, reset it to its default value before we attempt a non-blocking
	//	// read, otherwise the scheduler will just time us out before we can read
	//	if mc.cfg.ReadTimeout != 0 {
	//		err = conn.SetReadDeadline(time.Time{})
	//	}
	//	if err == nil && mc.cfg.CheckConnLiveness {
	//		err = connCheck(conn)
	//	}
	//	if err != nil {
	//		errLog.Print("closing bad idle connection: ", err)
	//		mc.Close()
	//		return driver.ErrBadConn
	//	}
	//}

	for {

		var size int
		if pktLen >= maxPacketSize {
			data[0] = 0xff
			data[1] = 0xff
			data[2] = 0xff
			size = maxPacketSize
		} else {
			data[0] = byte(pktLen)
			data[1] = byte(pktLen >> 8)
			data[2] = byte(pktLen >> 16)
			size = pktLen
		}
		data[3] = mc.sequence

		// Write packet
		if mc.writeTimeout > 0 {
			if err := mc.netConn.SetWriteDeadline(time.Now().Add(mc.writeTimeout)); err != nil {
				return err
			}
		}
		// fmt.Println("write", data[:4+size])
		n, err := mc.netConn.Write(data[:4+size])
		if err == nil && n == 4+size {
			mc.sequence++
			if size != maxPacketSize {
				return nil
			}
			pktLen -= size
			data = data[size:]
			continue
		}

		// Handle error
		if err == nil { // n != len(data)
			mc.cleanup()
			errLog.Print(ErrMalformPkt)
		} else {
			if n == 0 && pktLen == len(data)-4 {
				// only for the first loop iteration when nothing was written yet
				return errBadConnNoWrite
			}
			mc.cleanup()
			errLog.Print(err)
		}
		return ErrInvalidConn
	}
}

func (mc *MysqlConn) writeOK(r *MysqlResult) error {
	if r == nil {
		r = &MysqlResult{Status: mc.status}
	}
	data := make([]byte, 4, 32)
	data = append(data, IOK)

	data = appendLengthEncodedInteger(data, r.AffectedRows)
	data = appendLengthEncodedInteger(data, r.InsertId)

	if mc.capability&ClientProtocol41 > 0 {
		data = append(data, byte(r.Status), byte(r.Status>>8))
		data = append(data, 0, 0)
	}

	return mc.writePacket(data)
}

func (mc *MysqlConn) writeError(e error) error {
	var m *MySqlError
	var ok bool
	if m, ok = e.(*MySqlError); !ok {
		m = NewCustomError(ErUnknownError, e.Error())
	}

	data := make([]byte, 4, 16+len(m.Message))
	data = append(data, IERR)
	data = append(data, byte(m.Code), byte(m.Code>>8))

	if mc.capability&ClientProtocol41 > 0 {
		data = append(data, '#')
		data = append(data, m.State...)
	}
	data = append(data, m.Message...)

	return mc.writePacket(data)
}

func (mc *MysqlConn) writeEOF(status uint16) error {
	data := make([]byte, 4, 9)
	data = append(data, IEOF)
	if mc.capability&ClientProtocol41 > 0 {
		data = append(data, 0, 0)
		data = append(data, byte(status), byte(status>>8))
	}
	return mc.writePacket(data)
}

func (mc *MysqlConn) writeResultSet(r *sql.ExtendedRows) error {
	data := make([]byte, 4, 512)
	columns, err := r.Columns()
	if err != nil {
		return err
	}
	columnTypes, err := r.ColumnTypes()
	if err != nil {
		return err
	}
	// number of columns
	data = appendLengthEncodedInteger(data, uint64(len(columns)))
	err = mc.writePacket(data)
	if err != nil {
		return err
	}
	// column definitions terminated by an eof packet
	for _, v := range columnTypes {
		data = data[:4]
		data = append(data, v.RawType...)
		err = mc.writePacket(data)
		if err != nil {
			return err
		}
	}
	err = mc.writeEOF(r.Status)
	if err != nil {
		return err
	}

	// rows
	for r.Next() {
		rowData := make([]interface{}, len(columnTypes))
		for i := range columnTypes {
			var b []byte
			rowData[i] = &b
		}
		err = r.Scan(rowData...)
		if err != nil {
			return err
		}
		data = data[0:4]
		for i := range columnTypes {
			data = appendLengthEncodedString(data, *rowData[i].(*[]byte))
		}
		err = mc.writePacket(data)
		if err != nil {
			return err
		}
	}

	err = mc.writeEOF(r.Status)
	return err
}

/******************************************************************************
*                           Initialization Process                            *
******************************************************************************/

// Handshake Initialization Packet
// http://dev.mysql.com/doc/internals/en/connection-phase-packets.html#packet-Protocol::Handshake
func (mc *MysqlConn) writeHandshakePacket() (err error) {
	byteWriter := bytes.NewBuffer(nil)
	byteWriter.Write([]byte{0, 0, 0, 0})
	// protocol_version
	byteWriter.WriteByte(minProtocolVersion)
	// server_version
	byteWriter.WriteString("5.6.51")
	byteWriter.WriteByte(0)
	// connection_id
	byteWriter.Write(uint32ToBytes(mc.connectionId))
	// auth_plugin_data_part_1
	byteWriter.Write(mc.cfg.Salt[0:8])
	// filler_1
	byteWriter.WriteByte(0)
	// capability_flag_1
	capabilityByte := uint32ToBytes(uint32(defaultCapability))
	byteWriter.Write(capabilityByte[:2])
	// character_set
	byteWriter.WriteByte(collations[defaultCollation])
	// status_flags
	byteWriter.Write(uint16ToBytes(uint16(mc.status)))
	// capability_flags_2
	byteWriter.Write(capabilityByte[2:])
	// filter [0x15], for wireshark dump, value is 0x15
	byteWriter.WriteByte(0x15)
	// reserved 10 [00]
	byteWriter.Write([]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
	// auth_plugin_data_part_2
	byteWriter.Write(mc.cfg.Salt[8:])
	// filter [00]
	if defaultCapability&ClientPluginAuth > 0 {
		byteWriter.WriteString("mysql_native_password")
	}
	byteWriter.WriteByte(0)

	return mc.writePacket(byteWriter.Bytes())
}

// Client Authentication Packet
// http://dev.mysql.com/doc/internals/en/connection-phase-packets.html#packet-Protocol::HandshakeResponse
func (mc *MysqlConn) readHandshakeResponse() error {
	data, err := mc.readPacket()
	if err != nil {
		return err
	}
	pos := 0
	// capability_flags
	mc.capability = ClientFlag(binary.LittleEndian.Uint32(data[:4]))
	pos += 4
	// skip max_packet_size
	pos += 4
	// skip character_set
	pos += 1
	// skip reserved 23[00]
	pos += 23
	// username
	username := string(data[pos : pos+bytes.IndexByte(data[pos:], 0)])
	pos += len(username) + 1
	// auth-response
	var authResponse []byte
	if mc.capability&ClientPluginAuthLenEncClientData > 0 {
		authLen, _, numLen := readLengthEncodedInteger(data[pos:])
		pos += numLen
		authResponse = data[pos : pos+int(authLen)]
		pos += len(authResponse)
	} else if mc.capability&ClientSecureConn > 0 {
		authLen := uint64(data[pos])
		pos += 1
		authResponse = data[pos+1 : pos+1+int(authLen)]
		pos += len(authResponse)
	} else {
		authResponse = data[pos : pos+bytes.IndexByte(data[pos:], 0)]
		pos += len(authResponse) + 1
	}

	// check user
	if mc.cfg.User != username {
		return NewFormattedError(ErAccessDeniedError, username, mc.netConn.RemoteAddr().String(), "Yes")
	}

	// check password
	passwordHash := ScramblePassword(mc.cfg.Salt[:], mc.cfg.Passwd)
	passwordHash256 := ScrambleSHA256Password(mc.cfg.Salt[:], mc.cfg.Passwd)
	if len(authResponse) == 20 {
		if !bytes.Equal(authResponse, passwordHash) {
			return NewFormattedError(ErAccessDeniedError, username, mc.netConn.RemoteAddr().String(), "Yes")
		}
	} else if len(authResponse) == 32 {
		if !bytes.Equal(authResponse, passwordHash256) {
			return NewFormattedError(ErAccessDeniedError, username, mc.netConn.RemoteAddr().String(), "Yes")
		}
	} else {
		return NewFormattedError(ErNotSupportedAuthMode, username, mc.netConn.RemoteAddr().String(), "Yes")
	}

	// database
	if mc.capability&ClientConnectWithDB > 0 {
		if len(data[pos:]) == 0 {
			return nil
		}
		mc.database = string(data[pos : pos+bytes.IndexByte(data[pos:], 0)])
		pos += len(mc.database) + 1
	}
	// ignore auth plugin name and connect attrs
	return nil
}
