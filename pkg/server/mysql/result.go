// Go MySQL Driver - A MySQL-Driver for Go's database/sql package
//
// Copyright 2012 The Go-MySQL-Driver Authors. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.

package mysql

type MysqlResult struct {
	Status       StatusFlag
	AffectedRows uint64
	InsertId     uint64
}

func (res *MysqlResult) LastInsertId() (uint64, error) {
	return res.InsertId, nil
}

func (res *MysqlResult) RowsAffected() (uint64, error) {
	return res.AffectedRows, nil
}
