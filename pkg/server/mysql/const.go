// Go MySQL Driver - A MySQL-Driver for Go's database/sql package
//
// Copyright 2012 The Go-MySQL-Driver Authors. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.

package mysql

const (
	defaultAuthPlugin       = "mysql_native_password"
	defaultMaxAllowedPacket = 4 << 20 // 4 MiB
	minProtocolVersion      = 10
	maxPacketSize           = 1<<24 - 1
	timeFormat              = "2006-01-02 15:04:05.999999"
	defaultCapability       = ClientLongPassword | ClientLongFlag | ClientConnectWithDB |
		ClientProtocol41 | ClientTransactions | ClientSecureConn
)

// MySQL constants documentation:
// http://dev.mysql.com/doc/internals/en/client-server-protocol.html

const (
	IOK           byte = 0x00
	IAuthMoreData byte = 0x01
	ILocalInFile  byte = 0xfb
	IEOF          byte = 0xfe
	IERR          byte = 0xff
)

// https://dev.mysql.com/doc/internals/en/capability-flags.html#packet-Protocol::CapabilityFlags
type ClientFlag uint32

const (
	ClientLongPassword ClientFlag = 1 << iota
	ClientFoundRows
	ClientLongFlag
	ClientConnectWithDB
	ClientNoSchema
	ClientCompress
	ClientODBC
	ClientLocalFiles
	ClientIgnoreSpace
	ClientProtocol41
	ClientInteractive
	ClientSSL
	ClientIgnoreSIGPIPE
	ClientTransactions
	ClientReserved
	ClientSecureConn
	ClientMultiStatements
	ClientMultiResults
	ClientPSMultiResults
	ClientPluginAuth
	ClientConnectAttrs
	ClientPluginAuthLenEncClientData
	ClientCanHandleExpiredPasswords
	ClientSessionTrack
	ClientDeprecateEOF
)

const (
	ComQuit byte = iota + 1
	ComInitDB
	ComQuery
	ComFieldList
	ComCreateDB
	ComDropDB
	ComRefresh
	ComShutdown
	ComStatistics
	ComProcessInfo
	ComConnect
	ComProcessKill
	ComDebug
	ComPing
	ComTime
	ComDelayedInsert
	ComChangeUser
	ComBinlogDump
	ComTableDump
	ComConnectOut
	ComRegisterSlave
	ComStmtPrepare
	ComStmtExecute
	ComStmtSendLongData
	ComStmtClose
	ComStmtReset
	ComSetOption
	ComStmtFetch
)

// https://dev.mysql.com/doc/internals/en/com-query-response.html#packet-Protocol::ColumnType
type fieldType byte

const (
	fieldTypeDecimal fieldType = iota
	fieldTypeTiny
	fieldTypeShort
	fieldTypeLong
	fieldTypeFloat
	fieldTypeDouble
	fieldTypeNULL
	fieldTypeTimestamp
	fieldTypeLongLong
	fieldTypeInt24
	fieldTypeDate
	fieldTypeTime
	fieldTypeDateTime
	fieldTypeYear
	fieldTypeNewDate
	fieldTypeVarChar
	fieldTypeBit
)
const (
	fieldTypeJSON fieldType = iota + 0xf5
	fieldTypeNewDecimal
	fieldTypeEnum
	fieldTypeSet
	fieldTypeTinyBLOB
	fieldTypeMediumBLOB
	fieldTypeLongBLOB
	fieldTypeBLOB
	fieldTypeVarString
	fieldTypeString
	fieldTypeGeometry
)

type fieldFlag uint16

const (
	flagNotNULL fieldFlag = 1 << iota
	flagPriKey
	flagUniqueKey
	flagMultipleKey
	flagBLOB
	flagUnsigned
	flagZeroFill
	flagBinary
	flagEnum
	flagAutoIncrement
	flagTimestamp
	flagSet
	flagUnknown1
	flagUnknown2
	flagUnknown3
	flagUnknown4
)

// http://dev.mysql.com/doc/internals/en/status-flags.html
type StatusFlag uint16

const (
	StatusInTrans StatusFlag = 1 << iota
	StatusInAutocommit
	StatusReserved // Not in documentation
	StatusMoreResultsExists
	StatusNoGoodIndexUsed
	StatusNoIndexUsed
	StatusCursorExists
	StatusLastRowSent
	StatusDbDropped
	StatusNoBackslashEscapes
	StatusMetadataChanged
	StatusQueryWasSlow
	StatusPsOutParams
	StatusInTransReadonly
	StatusSessionStateChanged
)

const (
	cachingSha2PasswordRequestPublicKey          = 2
	cachingSha2PasswordFastAuthSuccess           = 3
	cachingSha2PasswordPerformFullAuthentication = 4
)

const (
	DefaultMysqlState = "HY000"
)

var MySQLState = map[uint16]string{
	ErDupKey:                              "23000",
	ErOutofmemory:                         "hy001",
	ErOutOfSortmemory:                     "hy001",
	ErConCountError:                       "08004",
	ErBadHostError:                        "08s01",
	ErHandshakeError:                      "08s01",
	ErDbaccessDeniedError:                 "42000",
	ErAccessDeniedError:                   "28000",
	ErNoDbError:                           "3d000",
	ErUnknownComError:                     "08s01",
	ErBadNullError:                        "23000",
	ErBadDbError:                          "42000",
	ErTableExistsError:                    "42s01",
	ErBadTableError:                       "42s02",
	ErNonUniqError:                        "23000",
	ErServerShutdown:                      "08s01",
	ErBadFieldError:                       "42s22",
	ErWrongFieldWithGroup:                 "42000",
	ErWrongSumSelect:                      "42000",
	ErWrongGroupField:                     "42000",
	ErWrongValueCount:                     "21s01",
	ErTooLongIdent:                        "42000",
	ErDupFieldname:                        "42s21",
	ErDupKeyname:                          "42000",
	ErDupEntry:                            "23000",
	ErWrongFieldSpec:                      "42000",
	ErParseError:                          "42000",
	ErEmptyQuery:                          "42000",
	ErNonuniqTable:                        "42000",
	ErInvalidDefault:                      "42000",
	ErMultiplePriKey:                      "42000",
	ErTooManyKeys:                         "42000",
	ErTooManyKeyParts:                     "42000",
	ErTooLongKey:                          "42000",
	ErKeyColumnDoesNotExits:               "42000",
	ErBlobUsedAsKey:                       "42000",
	ErTooBigFieldlength:                   "42000",
	ErWrongAutoKey:                        "42000",
	ErForcingClose:                        "08s01",
	ErIpsockError:                         "08s01",
	ErNoSuchIndex:                         "42s12",
	ErWrongFieldTerminators:               "42000",
	ErBlobsAndNoTerminated:                "42000",
	ErCantRemoveAllFields:                 "42000",
	ErCantDropFieldOrKey:                  "42000",
	ErBlobCantHaveDefault:                 "42000",
	ErWrongDbName:                         "42000",
	ErWrongTableName:                      "42000",
	ErTooBigSelect:                        "42000",
	ErUnknownProcedure:                    "42000",
	ErWrongParamcountToProcedure:          "42000",
	ErUnknownTable:                        "42s02",
	ErFieldSpecifiedTwice:                 "42000",
	ErUnsupportedExtension:                "42000",
	ErTableMustHaveColumns:                "42000",
	ErUnknownCharacterSet:                 "42000",
	ErTooBigRowsize:                       "42000",
	ErWrongOuterJoin:                      "42000",
	ErNullColumnInIndex:                   "42000",
	ErPasswordAnonymousUser:               "42000",
	ErPasswordNotAllowed:                  "42000",
	ErPasswordNoMatch:                     "42000",
	ErWrongValueCountOnRow:                "21s01",
	ErInvalidUseOfNull:                    "22004",
	ErRegexpError:                         "42000",
	ErMixOfGroupFuncAndFields:             "42000",
	ErNonexistingGrant:                    "42000",
	ErTableaccessDeniedError:              "42000",
	ErColumnaccessDeniedError:             "42000",
	ErIllegalGrantForTable:                "42000",
	ErGrantWrongHostOrUser:                "42000",
	ErNoSuchTable:                         "42s02",
	ErNonexistingTableGrant:               "42000",
	ErNotAllowedCommand:                   "42000",
	ErSyntaxError:                         "42000",
	ErAbortingConnection:                  "08s01",
	ErNetPacketTooLarge:                   "08s01",
	ErNetReadErrorFromPipe:                "08s01",
	ErNetFcntlError:                       "08s01",
	ErNetPacketsOutOfOrder:                "08s01",
	ErNetUncompressError:                  "08s01",
	ErNetReadError:                        "08s01",
	ErNetReadInterrupted:                  "08s01",
	ErNetErrorOnWrite:                     "08s01",
	ErNetWriteInterrupted:                 "08s01",
	ErTooLongString:                       "42000",
	ErTableCantHandleBlob:                 "42000",
	ErTableCantHandleAutoIncrement:        "42000",
	ErWrongColumnName:                     "42000",
	ErWrongKeyColumn:                      "42000",
	ErDupUnique:                           "23000",
	ErBlobKeyWithoutLength:                "42000",
	ErPrimaryCantHaveNull:                 "42000",
	ErTooManyRows:                         "42000",
	ErRequiresPrimaryKey:                  "42000",
	ErKeyDoesNotExits:                     "42000",
	ErCheckNoSuchTable:                    "42000",
	ErCheckNotImplemented:                 "42000",
	ErCantDoThisDuringAnTransaction:       "25000",
	ErNewAbortingConnection:               "08s01",
	ErMasterNetRead:                       "08s01",
	ErMasterNetWrite:                      "08s01",
	ErTooManyUserConnections:              "42000",
	ErReadOnlyTransaction:                 "25000",
	ErNoPermissionToCreateUser:            "42000",
	ErLockDeadlock:                        "40001",
	ErNoReferencedRow:                     "23000",
	ErRowIsReferenced:                     "23000",
	ErConnectToMaster:                     "08s01",
	ErWrongNumberOfColumnsInSelect:        "21000",
	ErUserLimitReached:                    "42000",
	ErSpecificAccessDeniedError:           "42000",
	ErNoDefault:                           "42000",
	ErWrongValueForVar:                    "42000",
	ErWrongTypeForVar:                     "42000",
	ErCantUseOptionHere:                   "42000",
	ErNotSupportedYet:                     "42000",
	ErWrongFkDef:                          "42000",
	ErOperandColumns:                      "21000",
	ErSubqueryNo1Row:                      "21000",
	ErIllegalReference:                    "42s22",
	ErDerivedMustHaveAlias:                "42000",
	ErSelectReduced:                       "01000",
	ErTablenameNotAllowedHere:             "42000",
	ErNotSupportedAuthMode:                "08004",
	ErSpatialCantHaveNull:                 "42000",
	ErCollationCharsetMismatch:            "42000",
	ErWarnTooFewRecords:                   "01000",
	ErWarnTooManyRecords:                  "01000",
	ErWarnNullToNotnull:                   "22004",
	ErWarnDataOutOfRange:                  "22003",
	WarnDataTruncated:                     "01000",
	ErWrongNameForIndex:                   "42000",
	ErWrongNameForCatalog:                 "42000",
	ErUnknownStorageEngine:                "42000",
	ErTruncatedWrongValue:                 "22007",
	ErSpNoRecursiveCreate:                 "2f003",
	ErSpAlreadyExists:                     "42000",
	ErSpDoesNotExist:                      "42000",
	ErSpLilabelMismatch:                   "42000",
	ErSpLabelRedefine:                     "42000",
	ErSpLabelMismatch:                     "42000",
	ErSpUninitVar:                         "01000",
	ErSpBadselect:                         "0a000",
	ErSpBadreturn:                         "42000",
	ErSpBadstatement:                      "0a000",
	ErUpdateLogDeprecatedIgnored:          "42000",
	ErUpdateLogDeprecatedTranslated:       "42000",
	ErQueryInterrupted:                    "70100",
	ErSpWrongNoOfArgs:                     "42000",
	ErSpCondMismatch:                      "42000",
	ErSpNoreturn:                          "42000",
	ErSpNoreturnend:                       "2f005",
	ErSpBadCursorQuery:                    "42000",
	ErSpBadCursorSelect:                   "42000",
	ErSpCursorMismatch:                    "42000",
	ErSpCursorAlreadyOpen:                 "24000",
	ErSpCursorNotOpen:                     "24000",
	ErSpUndeclaredVar:                     "42000",
	ErSpFetchNoData:                       "02000",
	ErSpDupParam:                          "42000",
	ErSpDupVar:                            "42000",
	ErSpDupCond:                           "42000",
	ErSpDupCurs:                           "42000",
	ErSpSubselectNyi:                      "0a000",
	ErStmtNotAllowedInSfOrTrg:             "0a000",
	ErSpVarcondAfterCurshndlr:             "42000",
	ErSpCursorAfterHandler:                "42000",
	ErSpCaseNotFound:                      "20000",
	ErDivisionByZero:                      "22012",
	ErIllegalValueForType:                 "22007",
	ErProcaccessDeniedError:               "42000",
	ErXaerNota:                            "xae04",
	ErXaerInval:                           "xae05",
	ErXaerRmfail:                          "xae07",
	ErXaerOutside:                         "xae09",
	ErXaerRmerr:                           "xae03",
	ErXaRbrollback:                        "xa100",
	ErNonexistingProcGrant:                "42000",
	ErDataTooLong:                         "22001",
	ErSpBadSqlstate:                       "42000",
	ErCantCreateUserWithGrant:             "42000",
	ErSpDupHandler:                        "42000",
	ErSpNotVarArg:                         "42000",
	ErSpNoRetset:                          "0a000",
	ErCantCreateGeometryObject:            "22003",
	ErTooBigScale:                         "42000",
	ErTooBigPrecision:                     "42000",
	ErMBiggerThanD:                        "42000",
	ErTooLongBody:                         "42000",
	ErTooBigDisplaywidth:                  "42000",
	ErXaerDupid:                           "xae08",
	ErDatetimeFunctionOverflow:            "22008",
	ErRowIsReferenced2:                    "23000",
	ErNoReferencedRow2:                    "23000",
	ErSpBadVarShadow:                      "42000",
	ErSpWrongName:                         "42000",
	ErSpNoAggregate:                       "42000",
	ErMaxPreparedStmtCountReached:         "42000",
	ErNonGroupingFieldUsed:                "42000",
	ErForeignDuplicateKeyOldUnused:        "23000",
	ErCantChangeTxCharacteristics:         "25001",
	ErWrongParamcountToNativeFct:          "42000",
	ErWrongParametersToNativeFct:          "42000",
	ErWrongParametersToStoredFct:          "42000",
	ErDupEntryWithKeyName:                 "23000",
	ErXaRbtimeout:                         "xa106",
	ErXaRbdeadlock:                        "xa102",
	ErFuncInexistentNameCollision:         "42000",
	ErDupSignalSet:                        "42000",
	ErSignalWarn:                          "01000",
	ErSignalNotFound:                      "02000",
	ErSignalException:                     "hy000",
	ErResignalWithoutActiveHandler:        "0k000",
	ErSpatialMustHaveGeomCol:              "42000",
	ErDataOutOfRange:                      "22003",
	ErAccessDeniedNoPasswordError:         "28000",
	ErTruncateIllegalFk:                   "42000",
	ErDaInvalidConditionNumber:            "35000",
	ErForeignDuplicateKeyWithChildInfo:    "23000",
	ErForeignDuplicateKeyWithoutChildInfo: "23000",
	ErCantExecuteInReadOnlyTransaction:    "25006",
	ErAlterOperationNotSupported:          "0a000",
	ErAlterOperationNotSupportedReason:    "0a000",
	ErDupUnknownInIndex:                   "23000",
}

const (
	ErErrorFirst                                            uint16 = 1000
	ErHashchk                                                      = 1000
	ErNisamchk                                                     = 1001
	ErNo                                                           = 1002
	ErYes                                                          = 1003
	ErCantCreateFile                                               = 1004
	ErCantCreateTable                                              = 1005
	ErCantCreateDb                                                 = 1006
	ErDbCreateExists                                               = 1007
	ErDbDropExists                                                 = 1008
	ErDbDropDelete                                                 = 1009
	ErDbDropRmdir                                                  = 1010
	ErCantDeleteFile                                               = 1011
	ErCantFindSystemRec                                            = 1012
	ErCantGetStat                                                  = 1013
	ErCantGetWd                                                    = 1014
	ErCantLock                                                     = 1015
	ErCantOpenFile                                                 = 1016
	ErFileNotFound                                                 = 1017
	ErCantReadDir                                                  = 1018
	ErCantSetWd                                                    = 1019
	ErCheckread                                                    = 1020
	ErDiskFull                                                     = 1021
	ErDupKey                                                       = 1022
	ErErrorOnClose                                                 = 1023
	ErErrorOnRead                                                  = 1024
	ErErrorOnRename                                                = 1025
	ErErrorOnWrite                                                 = 1026
	ErFileUsed                                                     = 1027
	ErFilsortAbort                                                 = 1028
	ErFormNotFound                                                 = 1029
	ErGetErrno                                                     = 1030
	ErIllegalHa                                                    = 1031
	ErKeyNotFound                                                  = 1032
	ErNotFormFile                                                  = 1033
	ErNotKeyfile                                                   = 1034
	ErOldKeyfile                                                   = 1035
	ErOpenAsReadonly                                               = 1036
	ErOutofmemory                                                  = 1037
	ErOutOfSortmemory                                              = 1038
	ErUnexpectedEof                                                = 1039
	ErConCountError                                                = 1040
	ErOutOfResources                                               = 1041
	ErBadHostError                                                 = 1042
	ErHandshakeError                                               = 1043
	ErDbaccessDeniedError                                          = 1044
	ErAccessDeniedError                                            = 1045
	ErNoDbError                                                    = 1046
	ErUnknownComError                                              = 1047
	ErBadNullError                                                 = 1048
	ErBadDbError                                                   = 1049
	ErTableExistsError                                             = 1050
	ErBadTableError                                                = 1051
	ErNonUniqError                                                 = 1052
	ErServerShutdown                                               = 1053
	ErBadFieldError                                                = 1054
	ErWrongFieldWithGroup                                          = 1055
	ErWrongGroupField                                              = 1056
	ErWrongSumSelect                                               = 1057
	ErWrongValueCount                                              = 1058
	ErTooLongIdent                                                 = 1059
	ErDupFieldname                                                 = 1060
	ErDupKeyname                                                   = 1061
	ErDupEntry                                                     = 1062
	ErWrongFieldSpec                                               = 1063
	ErParseError                                                   = 1064
	ErEmptyQuery                                                   = 1065
	ErNonuniqTable                                                 = 1066
	ErInvalidDefault                                               = 1067
	ErMultiplePriKey                                               = 1068
	ErTooManyKeys                                                  = 1069
	ErTooManyKeyParts                                              = 1070
	ErTooLongKey                                                   = 1071
	ErKeyColumnDoesNotExits                                        = 1072
	ErBlobUsedAsKey                                                = 1073
	ErTooBigFieldlength                                            = 1074
	ErWrongAutoKey                                                 = 1075
	ErReady                                                        = 1076
	ErNormalShutdown                                               = 1077
	ErGotSignal                                                    = 1078
	ErShutdownComplete                                             = 1079
	ErForcingClose                                                 = 1080
	ErIpsockError                                                  = 1081
	ErNoSuchIndex                                                  = 1082
	ErWrongFieldTerminators                                        = 1083
	ErBlobsAndNoTerminated                                         = 1084
	ErTextfileNotReadable                                          = 1085
	ErFileExistsError                                              = 1086
	ErLoadInfo                                                     = 1087
	ErAlterInfo                                                    = 1088
	ErWrongSubKey                                                  = 1089
	ErCantRemoveAllFields                                          = 1090
	ErCantDropFieldOrKey                                           = 1091
	ErInsertInfo                                                   = 1092
	ErUpdateTableUsed                                              = 1093
	ErNoSuchThread                                                 = 1094
	ErKillDeniedError                                              = 1095
	ErNoTablesUsed                                                 = 1096
	ErTooBigSet                                                    = 1097
	ErNoUniqueLogfile                                              = 1098
	ErTableNotLockedForWrite                                       = 1099
	ErTableNotLocked                                               = 1100
	ErBlobCantHaveDefault                                          = 1101
	ErWrongDbName                                                  = 1102
	ErWrongTableName                                               = 1103
	ErTooBigSelect                                                 = 1104
	ErUnknownError                                                 = 1105
	ErUnknownProcedure                                             = 1106
	ErWrongParamcountToProcedure                                   = 1107
	ErWrongParametersToProcedure                                   = 1108
	ErUnknownTable                                                 = 1109
	ErFieldSpecifiedTwice                                          = 1110
	ErInvalidGroupFuncUse                                          = 1111
	ErUnsupportedExtension                                         = 1112
	ErTableMustHaveColumns                                         = 1113
	ErRecordFileFull                                               = 1114
	ErUnknownCharacterSet                                          = 1115
	ErTooManyTables                                                = 1116
	ErTooManyFields                                                = 1117
	ErTooBigRowsize                                                = 1118
	ErStackOverrun                                                 = 1119
	ErWrongOuterJoin                                               = 1120
	ErNullColumnInIndex                                            = 1121
	ErCantFindUdf                                                  = 1122
	ErCantInitializeUdf                                            = 1123
	ErUdfNoPaths                                                   = 1124
	ErUdfExists                                                    = 1125
	ErCantOpenLibrary                                              = 1126
	ErCantFindDlEntry                                              = 1127
	ErFunctionNotDefined                                           = 1128
	ErHostIsBlocked                                                = 1129
	ErHostNotPrivileged                                            = 1130
	ErPasswordAnonymousUser                                        = 1131
	ErPasswordNotAllowed                                           = 1132
	ErPasswordNoMatch                                              = 1133
	ErUpdateInfo                                                   = 1134
	ErCantCreateThread                                             = 1135
	ErWrongValueCountOnRow                                         = 1136
	ErCantReopenTable                                              = 1137
	ErInvalidUseOfNull                                             = 1138
	ErRegexpError                                                  = 1139
	ErMixOfGroupFuncAndFields                                      = 1140
	ErNonexistingGrant                                             = 1141
	ErTableaccessDeniedError                                       = 1142
	ErColumnaccessDeniedError                                      = 1143
	ErIllegalGrantForTable                                         = 1144
	ErGrantWrongHostOrUser                                         = 1145
	ErNoSuchTable                                                  = 1146
	ErNonexistingTableGrant                                        = 1147
	ErNotAllowedCommand                                            = 1148
	ErSyntaxError                                                  = 1149
	ErDelayedCantChangeLock                                        = 1150
	ErTooManyDelayedThreads                                        = 1151
	ErAbortingConnection                                           = 1152
	ErNetPacketTooLarge                                            = 1153
	ErNetReadErrorFromPipe                                         = 1154
	ErNetFcntlError                                                = 1155
	ErNetPacketsOutOfOrder                                         = 1156
	ErNetUncompressError                                           = 1157
	ErNetReadError                                                 = 1158
	ErNetReadInterrupted                                           = 1159
	ErNetErrorOnWrite                                              = 1160
	ErNetWriteInterrupted                                          = 1161
	ErTooLongString                                                = 1162
	ErTableCantHandleBlob                                          = 1163
	ErTableCantHandleAutoIncrement                                 = 1164
	ErDelayedInsertTableLocked                                     = 1165
	ErWrongColumnName                                              = 1166
	ErWrongKeyColumn                                               = 1167
	ErWrongMrgTable                                                = 1168
	ErDupUnique                                                    = 1169
	ErBlobKeyWithoutLength                                         = 1170
	ErPrimaryCantHaveNull                                          = 1171
	ErTooManyRows                                                  = 1172
	ErRequiresPrimaryKey                                           = 1173
	ErNoRaidCompiled                                               = 1174
	ErUpdateWithoutKeyInSafeMode                                   = 1175
	ErKeyDoesNotExits                                              = 1176
	ErCheckNoSuchTable                                             = 1177
	ErCheckNotImplemented                                          = 1178
	ErCantDoThisDuringAnTransaction                                = 1179
	ErErrorDuringCommit                                            = 1180
	ErErrorDuringRollback                                          = 1181
	ErErrorDuringFlushLogs                                         = 1182
	ErErrorDuringCheckpoint                                        = 1183
	ErNewAbortingConnection                                        = 1184
	ErDumpNotImplemented                                           = 1185
	ErFlushMasterBinlogClosed                                      = 1186
	ErIndexRebuild                                                 = 1187
	ErMaster                                                       = 1188
	ErMasterNetRead                                                = 1189
	ErMasterNetWrite                                               = 1190
	ErFtMatchingKeyNotFound                                        = 1191
	ErLockOrActiveTransaction                                      = 1192
	ErUnknownSystemVariable                                        = 1193
	ErCrashedOnUsage                                               = 1194
	ErCrashedOnRepair                                              = 1195
	ErWarningNotCompleteRollback                                   = 1196
	ErTransCacheFull                                               = 1197
	ErSlaveMustStop                                                = 1198
	ErSlaveNotRunning                                              = 1199
	ErBadSlave                                                     = 1200
	ErMasterInfo                                                   = 1201
	ErSlaveThread                                                  = 1202
	ErTooManyUserConnections                                       = 1203
	ErSetConstantsOnly                                             = 1204
	ErLockWaitTimeout                                              = 1205
	ErLockTableFull                                                = 1206
	ErReadOnlyTransaction                                          = 1207
	ErDropDbWithReadLock                                           = 1208
	ErCreateDbWithReadLock                                         = 1209
	ErWrongArguments                                               = 1210
	ErNoPermissionToCreateUser                                     = 1211
	ErUnionTablesInDifferentDir                                    = 1212
	ErLockDeadlock                                                 = 1213
	ErTableCantHandleFt                                            = 1214
	ErCannotAddForeign                                             = 1215
	ErNoReferencedRow                                              = 1216
	ErRowIsReferenced                                              = 1217
	ErConnectToMaster                                              = 1218
	ErQueryOnMaster                                                = 1219
	ErErrorWhenExecutingCommand                                    = 1220
	ErWrongUsage                                                   = 1221
	ErWrongNumberOfColumnsInSelect                                 = 1222
	ErCantUpdateWithReadlock                                       = 1223
	ErMixingNotAllowed                                             = 1224
	ErDupArgument                                                  = 1225
	ErUserLimitReached                                             = 1226
	ErSpecificAccessDeniedError                                    = 1227
	ErLocalVariable                                                = 1228
	ErGlobalVariable                                               = 1229
	ErNoDefault                                                    = 1230
	ErWrongValueForVar                                             = 1231
	ErWrongTypeForVar                                              = 1232
	ErVarCantBeRead                                                = 1233
	ErCantUseOptionHere                                            = 1234
	ErNotSupportedYet                                              = 1235
	ErMasterFatalErrorReadingBinlog                                = 1236
	ErSlaveIgnoredTable                                            = 1237
	ErIncorrectGlobalLocalVar                                      = 1238
	ErWrongFkDef                                                   = 1239
	ErKeyRefDoNotMatchTableRef                                     = 1240
	ErOperandColumns                                               = 1241
	ErSubqueryNo1Row                                               = 1242
	ErUnknownStmtHandler                                           = 1243
	ErCorruptHelpDb                                                = 1244
	ErCyclicReference                                              = 1245
	ErAutoConvert                                                  = 1246
	ErIllegalReference                                             = 1247
	ErDerivedMustHaveAlias                                         = 1248
	ErSelectReduced                                                = 1249
	ErTablenameNotAllowedHere                                      = 1250
	ErNotSupportedAuthMode                                         = 1251
	ErSpatialCantHaveNull                                          = 1252
	ErCollationCharsetMismatch                                     = 1253
	ErSlaveWasRunning                                              = 1254
	ErSlaveWasNotRunning                                           = 1255
	ErTooBigForUncompress                                          = 1256
	ErZlibZMemError                                                = 1257
	ErZlibZBufError                                                = 1258
	ErZlibZDataError                                               = 1259
	ErCutValueGroupConcat                                          = 1260
	ErWarnTooFewRecords                                            = 1261
	ErWarnTooManyRecords                                           = 1262
	ErWarnNullToNotnull                                            = 1263
	ErWarnDataOutOfRange                                           = 1264
	WarnDataTruncated                                              = 1265
	ErWarnUsingOtherHandler                                        = 1266
	ErCantAggregate2collations                                     = 1267
	ErDropUser                                                     = 1268
	ErRevokeGrants                                                 = 1269
	ErCantAggregate3collations                                     = 1270
	ErCantAggregateNcollations                                     = 1271
	ErVariableIsNotStruct                                          = 1272
	ErUnknownCollation                                             = 1273
	ErSlaveIgnoredSslParams                                        = 1274
	ErServerIsInSecureAuthMode                                     = 1275
	ErWarnFieldResolved                                            = 1276
	ErBadSlaveUntilCond                                            = 1277
	ErMissingSkipSlave                                             = 1278
	ErUntilCondIgnored                                             = 1279
	ErWrongNameForIndex                                            = 1280
	ErWrongNameForCatalog                                          = 1281
	ErWarnQcResize                                                 = 1282
	ErBadFtColumn                                                  = 1283
	ErUnknownKeyCache                                              = 1284
	ErWarnHostnameWontWork                                         = 1285
	ErUnknownStorageEngine                                         = 1286
	ErWarnDeprecatedSyntax                                         = 1287
	ErNonUpdatableTable                                            = 1288
	ErFeatureDisabled                                              = 1289
	ErOptionPreventsStatement                                      = 1290
	ErDuplicatedValueInType                                        = 1291
	ErTruncatedWrongValue                                          = 1292
	ErTooMuchAutoTimestampCols                                     = 1293
	ErInvalidOnUpdate                                              = 1294
	ErUnsupportedPs                                                = 1295
	ErGetErrmsg                                                    = 1296
	ErGetTemporaryErrmsg                                           = 1297
	ErUnknownTimeZone                                              = 1298
	ErWarnInvalidTimestamp                                         = 1299
	ErInvalidCharacterString                                       = 1300
	ErWarnAllowedPacketOverflowed                                  = 1301
	ErConflictingDeclarations                                      = 1302
	ErSpNoRecursiveCreate                                          = 1303
	ErSpAlreadyExists                                              = 1304
	ErSpDoesNotExist                                               = 1305
	ErSpDropFailed                                                 = 1306
	ErSpStoreFailed                                                = 1307
	ErSpLilabelMismatch                                            = 1308
	ErSpLabelRedefine                                              = 1309
	ErSpLabelMismatch                                              = 1310
	ErSpUninitVar                                                  = 1311
	ErSpBadselect                                                  = 1312
	ErSpBadreturn                                                  = 1313
	ErSpBadstatement                                               = 1314
	ErUpdateLogDeprecatedIgnored                                   = 1315
	ErUpdateLogDeprecatedTranslated                                = 1316
	ErQueryInterrupted                                             = 1317
	ErSpWrongNoOfArgs                                              = 1318
	ErSpCondMismatch                                               = 1319
	ErSpNoreturn                                                   = 1320
	ErSpNoreturnend                                                = 1321
	ErSpBadCursorQuery                                             = 1322
	ErSpBadCursorSelect                                            = 1323
	ErSpCursorMismatch                                             = 1324
	ErSpCursorAlreadyOpen                                          = 1325
	ErSpCursorNotOpen                                              = 1326
	ErSpUndeclaredVar                                              = 1327
	ErSpWrongNoOfFetchArgs                                         = 1328
	ErSpFetchNoData                                                = 1329
	ErSpDupParam                                                   = 1330
	ErSpDupVar                                                     = 1331
	ErSpDupCond                                                    = 1332
	ErSpDupCurs                                                    = 1333
	ErSpCantAlter                                                  = 1334
	ErSpSubselectNyi                                               = 1335
	ErStmtNotAllowedInSfOrTrg                                      = 1336
	ErSpVarcondAfterCurshndlr                                      = 1337
	ErSpCursorAfterHandler                                         = 1338
	ErSpCaseNotFound                                               = 1339
	ErFparserTooBigFile                                            = 1340
	ErFparserBadHeader                                             = 1341
	ErFparserEofInComment                                          = 1342
	ErFparserErrorInParameter                                      = 1343
	ErFparserEofInUnknownParameter                                 = 1344
	ErViewNoExplain                                                = 1345
	ErFrmUnknownType                                               = 1346
	ErWrongObject                                                  = 1347
	ErNonupdateableColumn                                          = 1348
	ErViewSelectDerived                                            = 1349
	ErViewSelectClause                                             = 1350
	ErViewSelectVariable                                           = 1351
	ErViewSelectTmptable                                           = 1352
	ErViewWrongList                                                = 1353
	ErWarnViewMerge                                                = 1354
	ErWarnViewWithoutKey                                           = 1355
	ErViewInvalid                                                  = 1356
	ErSpNoDropSp                                                   = 1357
	ErSpGotoInHndlr                                                = 1358
	ErTrgAlreadyExists                                             = 1359
	ErTrgDoesNotExist                                              = 1360
	ErTrgOnViewOrTempTable                                         = 1361
	ErTrgCantChangeRow                                             = 1362
	ErTrgNoSuchRowInTrg                                            = 1363
	ErNoDefaultForField                                            = 1364
	ErDivisionByZero                                               = 1365
	ErTruncatedWrongValueForField                                  = 1366
	ErIllegalValueForType                                          = 1367
	ErViewNonupdCheck                                              = 1368
	ErViewCheckFailed                                              = 1369
	ErProcaccessDeniedError                                        = 1370
	ErRelayLogFail                                                 = 1371
	ErPasswdLength                                                 = 1372
	ErUnknownTargetBinlog                                          = 1373
	ErIoErrLogIndexRead                                            = 1374
	ErBinlogPurgeProhibited                                        = 1375
	ErFseekFail                                                    = 1376
	ErBinlogPurgeFatalErr                                          = 1377
	ErLogInUse                                                     = 1378
	ErLogPurgeUnknownErr                                           = 1379
	ErRelayLogInit                                                 = 1380
	ErNoBinaryLogging                                              = 1381
	ErReservedSyntax                                               = 1382
	ErWsasFailed                                                   = 1383
	ErDiffGroupsProc                                               = 1384
	ErNoGroupForProc                                               = 1385
	ErOrderWithProc                                                = 1386
	ErLoggingProhibitChangingOf                                    = 1387
	ErNoFileMapping                                                = 1388
	ErWrongMagic                                                   = 1389
	ErPsManyParam                                                  = 1390
	ErKeyPart0                                                     = 1391
	ErViewChecksum                                                 = 1392
	ErViewMultiupdate                                              = 1393
	ErViewNoInsertFieldList                                        = 1394
	ErViewDeleteMergeView                                          = 1395
	ErCannotUser                                                   = 1396
	ErXaerNota                                                     = 1397
	ErXaerInval                                                    = 1398
	ErXaerRmfail                                                   = 1399
	ErXaerOutside                                                  = 1400
	ErXaerRmerr                                                    = 1401
	ErXaRbrollback                                                 = 1402
	ErNonexistingProcGrant                                         = 1403
	ErProcAutoGrantFail                                            = 1404
	ErProcAutoRevokeFail                                           = 1405
	ErDataTooLong                                                  = 1406
	ErSpBadSqlstate                                                = 1407
	ErStartup                                                      = 1408
	ErLoadFromFixedSizeRowsToVar                                   = 1409
	ErCantCreateUserWithGrant                                      = 1410
	ErWrongValueForType                                            = 1411
	ErTableDefChanged                                              = 1412
	ErSpDupHandler                                                 = 1413
	ErSpNotVarArg                                                  = 1414
	ErSpNoRetset                                                   = 1415
	ErCantCreateGeometryObject                                     = 1416
	ErFailedRoutineBreakBinlog                                     = 1417
	ErBinlogUnsafeRoutine                                          = 1418
	ErBinlogCreateRoutineNeedSuper                                 = 1419
	ErExecStmtWithOpenCursor                                       = 1420
	ErStmtHasNoOpenCursor                                          = 1421
	ErCommitNotAllowedInSfOrTrg                                    = 1422
	ErNoDefaultForViewField                                        = 1423
	ErSpNoRecursion                                                = 1424
	ErTooBigScale                                                  = 1425
	ErTooBigPrecision                                              = 1426
	ErMBiggerThanD                                                 = 1427
	ErWrongLockOfSystemTable                                       = 1428
	ErConnectToForeignDataSource                                   = 1429
	ErQueryOnForeignDataSource                                     = 1430
	ErForeignDataSourceDoesntExist                                 = 1431
	ErForeignDataStringInvalidCantCreate                           = 1432
	ErForeignDataStringInvalid                                     = 1433
	ErCantCreateFederatedTable                                     = 1434
	ErTrgInWrongSchema                                             = 1435
	ErStackOverrunNeedMore                                         = 1436
	ErTooLongBody                                                  = 1437
	ErWarnCantDropDefaultKeycache                                  = 1438
	ErTooBigDisplaywidth                                           = 1439
	ErXaerDupid                                                    = 1440
	ErDatetimeFunctionOverflow                                     = 1441
	ErCantUpdateUsedTableInSfOrTrg                                 = 1442
	ErViewPreventUpdate                                            = 1443
	ErPsNoRecursion                                                = 1444
	ErSpCantSetAutocommit                                          = 1445
	ErMalformedDefiner                                             = 1446
	ErViewFrmNoUser                                                = 1447
	ErViewOtherUser                                                = 1448
	ErNoSuchUser                                                   = 1449
	ErForbidSchemaChange                                           = 1450
	ErRowIsReferenced2                                             = 1451
	ErNoReferencedRow2                                             = 1452
	ErSpBadVarShadow                                               = 1453
	ErTrgNoDefiner                                                 = 1454
	ErOldFileFormat                                                = 1455
	ErSpRecursionLimit                                             = 1456
	ErSpProcTableCorrupt                                           = 1457
	ErSpWrongName                                                  = 1458
	ErTableNeedsUpgrade                                            = 1459
	ErSpNoAggregate                                                = 1460
	ErMaxPreparedStmtCountReached                                  = 1461
	ErViewRecursive                                                = 1462
	ErNonGroupingFieldUsed                                         = 1463
	ErTableCantHandleSpkeys                                        = 1464
	ErNoTriggersOnSystemSchema                                     = 1465
	ErRemovedSpaces                                                = 1466
	ErAutoincReadFailed                                            = 1467
	ErUsername                                                     = 1468
	ErHostname                                                     = 1469
	ErWrongStringLength                                            = 1470
	ErNonInsertableTable                                           = 1471
	ErAdminWrongMrgTable                                           = 1472
	ErTooHighLevelOfNestingForSelect                               = 1473
	ErNameBecomesEmpty                                             = 1474
	ErAmbiguousFieldTerm                                           = 1475
	ErForeignServerExists                                          = 1476
	ErForeignServerDoesntExist                                     = 1477
	ErIllegalHaCreateOption                                        = 1478
	ErPartitionRequiresValuesError                                 = 1479
	ErPartitionWrongValuesError                                    = 1480
	ErPartitionMaxvalueError                                       = 1481
	ErPartitionSubpartitionError                                   = 1482
	ErPartitionSubpartMixError                                     = 1483
	ErPartitionWrongNoPartError                                    = 1484
	ErPartitionWrongNoSubpartError                                 = 1485
	ErWrongExprInPartitionFuncError                                = 1486
	ErNoConstExprInRangeOrListError                                = 1487
	ErFieldNotFoundPartError                                       = 1488
	ErListOfFieldsOnlyInHashError                                  = 1489
	ErInconsistentPartitionInfoError                               = 1490
	ErPartitionFuncNotAllowedError                                 = 1491
	ErPartitionsMustBeDefinedError                                 = 1492
	ErRangeNotIncreasingError                                      = 1493
	ErInconsistentTypeOfFunctionsError                             = 1494
	ErMultipleDefConstInListPartError                              = 1495
	ErPartitionEntryError                                          = 1496
	ErMixHandlerError                                              = 1497
	ErPartitionNotDefinedError                                     = 1498
	ErTooManyPartitionsError                                       = 1499
	ErSubpartitionError                                            = 1500
	ErCantCreateHandlerFile                                        = 1501
	ErBlobFieldInPartFuncError                                     = 1502
	ErUniqueKeyNeedAllFieldsInPf                                   = 1503
	ErNoPartsError                                                 = 1504
	ErPartitionMgmtOnNonpartitioned                                = 1505
	ErForeignKeyOnPartitioned                                      = 1506
	ErDropPartitionNonExistent                                     = 1507
	ErDropLastPartition                                            = 1508
	ErCoalesceOnlyOnHashPartition                                  = 1509
	ErReorgHashOnlyOnSameNo                                        = 1510
	ErReorgNoParamError                                            = 1511
	ErOnlyOnRangeListPartition                                     = 1512
	ErAddPartitionSubpartError                                     = 1513
	ErAddPartitionNoNewPartition                                   = 1514
	ErCoalescePartitionNoPartition                                 = 1515
	ErReorgPartitionNotExist                                       = 1516
	ErSameNamePartition                                            = 1517
	ErNoBinlogError                                                = 1518
	ErConsecutiveReorgPartitions                                   = 1519
	ErReorgOutsideRange                                            = 1520
	ErPartitionFunctionFailure                                     = 1521
	ErPartStateError                                               = 1522
	ErLimitedPartRange                                             = 1523
	ErPluginIsNotLoaded                                            = 1524
	ErWrongValue                                                   = 1525
	ErNoPartitionForGivenValue                                     = 1526
	ErFilegroupOptionOnlyOnce                                      = 1527
	ErCreateFilegroupFailed                                        = 1528
	ErDropFilegroupFailed                                          = 1529
	ErTablespaceAutoExtendError                                    = 1530
	ErWrongSizeNumber                                              = 1531
	ErSizeOverflowError                                            = 1532
	ErAlterFilegroupFailed                                         = 1533
	ErBinlogRowLoggingFailed                                       = 1534
	ErBinlogRowWrongTableDef                                       = 1535
	ErBinlogRowRbrToSbr                                            = 1536
	ErEventAlreadyExists                                           = 1537
	ErEventStoreFailed                                             = 1538
	ErEventDoesNotExist                                            = 1539
	ErEventCantAlter                                               = 1540
	ErEventDropFailed                                              = 1541
	ErEventIntervalNotPositiveOrTooBig                             = 1542
	ErEventEndsBeforeStarts                                        = 1543
	ErEventExecTimeInThePast                                       = 1544
	ErEventOpenTableFailed                                         = 1545
	ErEventNeitherMExprNorMAt                                      = 1546
	ErObsoleteColCountDoesntMatchCorrupted                         = 1547
	ErObsoleteCannotLoadFromTable                                  = 1548
	ErEventCannotDelete                                            = 1549
	ErEventCompileError                                            = 1550
	ErEventSameName                                                = 1551
	ErEventDataTooLong                                             = 1552
	ErDropIndexFk                                                  = 1553
	ErWarnDeprecatedSyntaxWithVer                                  = 1554
	ErCantWriteLockLogTable                                        = 1555
	ErCantLockLogTable                                             = 1556
	ErForeignDuplicateKeyOldUnused                                 = 1557
	ErColCountDoesntMatchPleaseUpdate                              = 1558
	ErTempTablePreventsSwitchOutOfRbr                              = 1559
	ErStoredFunctionPreventsSwitchBinlogFormat                     = 1560
	ErNdbCantSwitchBinlogFormat                                    = 1561
	ErPartitionNoTemporary                                         = 1562
	ErPartitionConstDomainError                                    = 1563
	ErPartitionFunctionIsNotAllowed                                = 1564
	ErDdlLogError                                                  = 1565
	ErNullInValuesLessThan                                         = 1566
	ErWrongPartitionName                                           = 1567
	ErCantChangeTxCharacteristics                                  = 1568
	ErDupEntryAutoincrementCase                                    = 1569
	ErEventModifyQueueError                                        = 1570
	ErEventSetVarError                                             = 1571
	ErPartitionMergeError                                          = 1572
	ErCantActivateLog                                              = 1573
	ErRbrNotAvailable                                              = 1574
	ErBase64DecodeError                                            = 1575
	ErEventRecursionForbidden                                      = 1576
	ErEventsDbError                                                = 1577
	ErOnlyIntegersAllowed                                          = 1578
	ErUnsuportedLogEngine                                          = 1579
	ErBadLogStatement                                              = 1580
	ErCantRenameLogTable                                           = 1581
	ErWrongParamcountToNativeFct                                   = 1582
	ErWrongParametersToNativeFct                                   = 1583
	ErWrongParametersToStoredFct                                   = 1584
	ErNativeFctNameCollision                                       = 1585
	ErDupEntryWithKeyName                                          = 1586
	ErBinlogPurgeEmfile                                            = 1587
	ErEventCannotCreateInThePast                                   = 1588
	ErEventCannotAlterInThePast                                    = 1589
	ErSlaveIncident                                                = 1590
	ErNoPartitionForGivenValueSilent                               = 1591
	ErBinlogUnsafeStatement                                        = 1592
	ErSlaveFatalError                                              = 1593
	ErSlaveRelayLogReadFailure                                     = 1594
	ErSlaveRelayLogWriteFailure                                    = 1595
	ErSlaveCreateEventFailure                                      = 1596
	ErSlaveMasterComFailure                                        = 1597
	ErBinlogLoggingImpossible                                      = 1598
	ErViewNoCreationCtx                                            = 1599
	ErViewInvalidCreationCtx                                       = 1600
	ErSrInvalidCreationCtx                                         = 1601
	ErTrgCorruptedFile                                             = 1602
	ErTrgNoCreationCtx                                             = 1603
	ErTrgInvalidCreationCtx                                        = 1604
	ErEventInvalidCreationCtx                                      = 1605
	ErTrgCantOpenTable                                             = 1606
	ErCantCreateSroutine                                           = 1607
	ErNeverUsed                                                    = 1608
	ErNoFormatDescriptionEventBeforeBinlogStatement                = 1609
	ErSlaveCorruptEvent                                            = 1610
	ErLoadDataInvalidColumn                                        = 1611
	ErLogPurgeNoFile                                               = 1612
	ErXaRbtimeout                                                  = 1613
	ErXaRbdeadlock                                                 = 1614
	ErNeedReprepare                                                = 1615
	ErDelayedNotSupported                                          = 1616
	WarnNoMasterInfo                                               = 1617
	WarnOptionIgnored                                              = 1618
	WarnPluginDeleteBuiltin                                        = 1619
	WarnPluginBusy                                                 = 1620
	ErVariableIsReadonly                                           = 1621
	ErWarnEngineTransactionRollback                                = 1622
	ErSlaveHeartbeatFailure                                        = 1623
	ErSlaveHeartbeatValueOutOfRange                                = 1624
	ErNdbReplicationSchemaError                                    = 1625
	ErConflictFnParseError                                         = 1626
	ErExceptionsWriteError                                         = 1627
	ErTooLongTableComment                                          = 1628
	ErTooLongFieldComment                                          = 1629
	ErFuncInexistentNameCollision                                  = 1630
	ErDatabaseName                                                 = 1631
	ErTableName                                                    = 1632
	ErPartitionName                                                = 1633
	ErSubpartitionName                                             = 1634
	ErTemporaryName                                                = 1635
	ErRenamedName                                                  = 1636
	ErTooManyConcurrentTrxs                                        = 1637
	WarnNonAsciiSeparatorNotImplemented                            = 1638
	ErDebugSyncTimeout                                             = 1639
	ErDebugSyncHitLimit                                            = 1640
	ErDupSignalSet                                                 = 1641
	ErSignalWarn                                                   = 1642
	ErSignalNotFound                                               = 1643
	ErSignalException                                              = 1644
	ErResignalWithoutActiveHandler                                 = 1645
	ErSignalBadConditionType                                       = 1646
	WarnCondItemTruncated                                          = 1647
	ErCondItemTooLong                                              = 1648
	ErUnknownLocale                                                = 1649
	ErSlaveIgnoreServerIds                                         = 1650
	ErQueryCacheDisabled                                           = 1651
	ErSameNamePartitionField                                       = 1652
	ErPartitionColumnListError                                     = 1653
	ErWrongTypeColumnValueError                                    = 1654
	ErTooManyPartitionFuncFieldsError                              = 1655
	ErMaxvalueInValuesIn                                           = 1656
	ErTooManyValuesError                                           = 1657
	ErRowSinglePartitionFieldError                                 = 1658
	ErFieldTypeNotAllowedAsPartitionField                          = 1659
	ErPartitionFieldsTooLong                                       = 1660
	ErBinlogRowEngineAndStmtEngine                                 = 1661
	ErBinlogRowModeAndStmtEngine                                   = 1662
	ErBinlogUnsafeAndStmtEngine                                    = 1663
	ErBinlogRowInjectionAndStmtEngine                              = 1664
	ErBinlogStmtModeAndRowEngine                                   = 1665
	ErBinlogRowInjectionAndStmtMode                                = 1666
	ErBinlogMultipleEnginesAndSelfLoggingEngine                    = 1667
	ErBinlogUnsafeLimit                                            = 1668
	ErBinlogUnsafeInsertDelayed                                    = 1669
	ErBinlogUnsafeSystemTable                                      = 1670
	ErBinlogUnsafeAutoincColumns                                   = 1671
	ErBinlogUnsafeUdf                                              = 1672
	ErBinlogUnsafeSystemVariable                                   = 1673
	ErBinlogUnsafeSystemFunction                                   = 1674
	ErBinlogUnsafeNontransAfterTrans                               = 1675
	ErMessageAndStatement                                          = 1676
	ErSlaveConversionFailed                                        = 1677
	ErSlaveCantCreateConversion                                    = 1678
	ErInsideTransactionPreventsSwitchBinlogFormat                  = 1679
	ErPathLength                                                   = 1680
	ErWarnDeprecatedSyntaxNoReplacement                            = 1681
	ErWrongNativeTableStructure                                    = 1682
	ErWrongPerfschemaUsage                                         = 1683
	ErWarnISSkippedTable                                           = 1684
	ErInsideTransactionPreventsSwitchBinlogDirect                  = 1685
	ErStoredFunctionPreventsSwitchBinlogDirect                     = 1686
	ErSpatialMustHaveGeomCol                                       = 1687
	ErTooLongIndexComment                                          = 1688
	ErLockAborted                                                  = 1689
	ErDataOutOfRange                                               = 1690
	ErWrongSpvarTypeInLimit                                        = 1691
	ErBinlogUnsafeMultipleEnginesAndSelfLoggingEngine              = 1692
	ErBinlogUnsafeMixedStatement                                   = 1693
	ErInsideTransactionPreventsSwitchSqlLogBin                     = 1694
	ErStoredFunctionPreventsSwitchSqlLogBin                        = 1695
	ErFailedReadFromParFile                                        = 1696
	ErValuesIsNotIntTypeError                                      = 1697
	ErAccessDeniedNoPasswordError                                  = 1698
	ErSetPasswordAuthPlugin                                        = 1699
	ErGrantPluginUserExists                                        = 1700
	ErTruncateIllegalFk                                            = 1701
	ErPluginIsPermanent                                            = 1702
	ErSlaveHeartbeatValueOutOfRangeMin                             = 1703
	ErSlaveHeartbeatValueOutOfRangeMax                             = 1704
	ErStmtCacheFull                                                = 1705
	ErMultiUpdateKeyConflict                                       = 1706
	ErTableNeedsRebuild                                            = 1707
	WarnOptionBelowLimit                                           = 1708
	ErIndexColumnTooLong                                           = 1709
	ErErrorInTriggerBody                                           = 1710
	ErErrorInUnknownTriggerBody                                    = 1711
	ErIndexCorrupt                                                 = 1712
	ErUndoRecordTooBig                                             = 1713
	ErBinlogUnsafeInsertIgnoreSelect                               = 1714
	ErBinlogUnsafeInsertSelectUpdate                               = 1715
	ErBinlogUnsafeReplaceSelect                                    = 1716
	ErBinlogUnsafeCreateIgnoreSelect                               = 1717
	ErBinlogUnsafeCreateReplaceSelect                              = 1718
	ErBinlogUnsafeUpdateIgnore                                     = 1719
	ErPluginNoUninstall                                            = 1720
	ErPluginNoInstall                                              = 1721
	ErBinlogUnsafeWriteAutoincSelect                               = 1722
	ErBinlogUnsafeCreateSelectAutoinc                              = 1723
	ErBinlogUnsafeInsertTwoKeys                                    = 1724
	ErTableInFkCheck                                               = 1725
	ErUnsupportedEngine                                            = 1726
	ErBinlogUnsafeAutoincNotFirst                                  = 1727
	ErCannotLoadFromTableV2                                        = 1728
	ErMasterDelayValueOutOfRange                                   = 1729
	ErOnlyFdAndRbrEventsAllowedInBinlogStatement                   = 1730
	ErPartitionExchangeDifferentOption                             = 1731
	ErPartitionExchangePartTable                                   = 1732
	ErPartitionExchangeTempTable                                   = 1733
	ErPartitionInsteadOfSubpartition                               = 1734
	ErUnknownPartition                                             = 1735
	ErTablesDifferentMetadata                                      = 1736
	ErRowDoesNotMatchPartition                                     = 1737
	ErBinlogCacheSizeGreaterThanMax                                = 1738
	ErWarnIndexNotApplicable                                       = 1739
	ErPartitionExchangeForeignKey                                  = 1740
	ErNoSuchKeyValue                                               = 1741
	ErRplInfoDataTooLong                                           = 1742
	ErNetworkReadEventChecksumFailure                              = 1743
	ErBinlogReadEventChecksumFailure                               = 1744
	ErBinlogStmtCacheSizeGreaterThanMax                            = 1745
	ErCantUpdateTableInCreateTableSelect                           = 1746
	ErPartitionClauseOnNonpartitioned                              = 1747
	ErRowDoesNotMatchGivenPartitionSet                             = 1748
	ErNoSuchPartition_Unused                                       = 1749
	ErChangeRplInfoRepositoryFailure                               = 1750
	ErWarningNotCompleteRollbackWithCreatedTempTable               = 1751
	ErWarningNotCompleteRollbackWithDroppedTempTable               = 1752
	ErMtsFeatureIsNotSupported                                     = 1753
	ErMtsUpdatedDbsGreaterMax                                      = 1754
	ErMtsCantParallel                                              = 1755
	ErMtsInconsistentData                                          = 1756
	ErFulltextNotSupportedWithPartitioning                         = 1757
	ErDaInvalidConditionNumber                                     = 1758
	ErInsecurePlainText                                            = 1759
	ErInsecureChangeMaster                                         = 1760
	ErForeignDuplicateKeyWithChildInfo                             = 1761
	ErForeignDuplicateKeyWithoutChildInfo                          = 1762
	ErSqlthreadWithSecureSlave                                     = 1763
	ErTableHasNoFt                                                 = 1764
	ErVariableNotSettableInSfOrTrigger                             = 1765
	ErVariableNotSettableInTransaction                             = 1766
	ErGtidNextIsNotInGtidNextList                                  = 1767
	ErCantChangeGtidNextInTransactionWhenGtidNextListIsNull        = 1768
	ErSetStatementCannotInvokeFunction                             = 1769
	ErGtidNextCantBeAutomaticIfGtidNextListIsNonNull               = 1770
	ErSkippingLoggedTransaction                                    = 1771
	ErMalformedGtidSetSpecification                                = 1772
	ErMalformedGtidSetEncoding                                     = 1773
	ErMalformedGtidSpecification                                   = 1774
	ErGnoExhausted                                                 = 1775
	ErBadSlaveAutoPosition                                         = 1776
	ErAutoPositionRequiresGtidModeOn                               = 1777
	ErCantDoImplicitCommitInTrxWhenGtidNextIsSet                   = 1778
	ErGtidMode2Or3RequiresEnforceGtidConsistencyOn                 = 1779
	ErGtidModeRequiresBinlog                                       = 1780
	ErCantSetGtidNextToGtidWhenGtidModeIsOff                       = 1781
	ErCantSetGtidNextToAnonymousWhenGtidModeIsOn                   = 1782
	ErCantSetGtidNextListToNonNullWhenGtidModeIsOff                = 1783
	ErFoundGtidEventWhenGtidModeIsOff                              = 1784
	ErGtidUnsafeNonTransactionalTable                              = 1785
	ErGtidUnsafeCreateSelect                                       = 1786
	ErGtidUnsafeCreateDropTemporaryTableInTransaction              = 1787
	ErGtidModeCanOnlyChangeOneStepAtATime                          = 1788
	ErMasterHasPurgedRequiredGtids                                 = 1789
	ErCantSetGtidNextWhenOwningGtid                                = 1790
	ErUnknownExplainFormat                                         = 1791
	ErCantExecuteInReadOnlyTransaction                             = 1792
	ErTooLongTablePartitionComment                                 = 1793
	ErSlaveConfiguration                                           = 1794
	ErInnodbFtLimit                                                = 1795
	ErInnodbNoFtTempTable                                          = 1796
	ErInnodbFtWrongDocidColumn                                     = 1797
	ErInnodbFtWrongDocidIndex                                      = 1798
	ErInnodbOnlineLogTooBig                                        = 1799
	ErUnknownAlterAlgorithm                                        = 1800
	ErUnknownAlterLock                                             = 1801
	ErMtsChangeMasterCantRunWithGaps                               = 1802
	ErMtsRecoveryFailure                                           = 1803
	ErMtsResetWorkers                                              = 1804
	ErColCountDoesntMatchCorruptedV2                               = 1805
	ErSlaveSilentRetryTransaction                                  = 1806
	ErDiscardFkChecksRunning                                       = 1807
	ErTableSchemaMismatch                                          = 1808
	ErTableInSystemTablespace                                      = 1809
	ErIoReadError                                                  = 1810
	ErIoWriteError                                                 = 1811
	ErTablespaceMissing                                            = 1812
	ErTablespaceExists                                             = 1813
	ErTablespaceDiscarded                                          = 1814
	ErInternalError                                                = 1815
	ErInnodbImportError                                            = 1816
	ErInnodbIndexCorrupt                                           = 1817
	ErInvalidYearColumnLength                                      = 1818
	ErNotValidPassword                                             = 1819
	ErMustChangePassword                                           = 1820
	ErFkNoIndexChild                                               = 1821
	ErFkNoIndexParent                                              = 1822
	ErFkFailAddSystem                                              = 1823
	ErFkCannotOpenParent                                           = 1824
	ErFkIncorrectOption                                            = 1825
	ErFkDupName                                                    = 1826
	ErPasswordFormat                                               = 1827
	ErFkColumnCannotDrop                                           = 1828
	ErFkColumnCannotDropChild                                      = 1829
	ErFkColumnNotNull                                              = 1830
	ErDupIndex                                                     = 1831
	ErFkColumnCannotChange                                         = 1832
	ErFkColumnCannotChangeChild                                    = 1833
	ErFkCannotDeleteParent                                         = 1834
	ErMalformedPacket                                              = 1835
	ErReadOnlyMode                                                 = 1836
	ErGtidNextTypeUndefinedGroup                                   = 1837
	ErVariableNotSettableInSp                                      = 1838
	ErCantSetGtidPurgedWhenGtidModeIsOff                           = 1839
	ErCantSetGtidPurgedWhenGtidExecutedIsNotEmpty                  = 1840
	ErCantSetGtidPurgedWhenOwnedGtidsIsNotEmpty                    = 1841
	ErGtidPurgedWasChanged                                         = 1842
	ErGtidExecutedWasChanged                                       = 1843
	ErBinlogStmtModeAndNoReplTables                                = 1844
	ErAlterOperationNotSupported                                   = 1845
	ErAlterOperationNotSupportedReason                             = 1846
	ErAlterOperationNotSupportedReasonCopy                         = 1847
	ErAlterOperationNotSupportedReasonPartition                    = 1848
	ErAlterOperationNotSupportedReasonFkRename                     = 1849
	ErAlterOperationNotSupportedReasonColumnType                   = 1850
	ErAlterOperationNotSupportedReasonFkCheck                      = 1851
	ErAlterOperationNotSupportedReasonIgnore                       = 1852
	ErAlterOperationNotSupportedReasonNopk                         = 1853
	ErAlterOperationNotSupportedReasonAutoinc                      = 1854
	ErAlterOperationNotSupportedReasonHiddenFts                    = 1855
	ErAlterOperationNotSupportedReasonChangeFts                    = 1856
	ErAlterOperationNotSupportedReasonFts                          = 1857
	ErSqlSlaveSkipCounterNotSettableInGtidMode                     = 1858
	ErDupUnknownInIndex                                            = 1859
	ErIdentCausesTooLongPath                                       = 1860
	ErAlterOperationNotSupportedReasonNotNull                      = 1861
	ErMustChangePasswordLogin                                      = 1862
	ErRowInWrongPartition                                          = 1863
	ErErrorLast                                                    = 1863
)

var MySQLErrName = map[uint16]string{
	ErHashchk:                                  "hashchk",
	ErNisamchk:                                 "isamchk",
	ErNo:                                       "NO",
	ErYes:                                      "YES",
	ErCantCreateFile:                           "Can't create file '%-.200s' (errno: %d - %s)",
	ErCantCreateTable:                          "Can't create table '%-.200s' (errno: %d)",
	ErCantCreateDb:                             "Can't create database '%-.192s' (errno: %d)",
	ErDbCreateExists:                           "Can't create database '%-.192s'; database exists",
	ErDbDropExists:                             "Can't drop database '%-.192s'; database doesn't exist",
	ErDbDropDelete:                             "Error dropping database (can't delete '%-.192s', errno: %d)",
	ErDbDropRmdir:                              "Error dropping database (can't rmdir '%-.192s', errno: %d)",
	ErCantDeleteFile:                           "Error on delete of '%-.192s' (errno: %d - %s)",
	ErCantFindSystemRec:                        "Can't read record in system table",
	ErCantGetStat:                              "Can't get status of '%-.200s' (errno: %d - %s)",
	ErCantGetWd:                                "Can't get working directory (errno: %d - %s)",
	ErCantLock:                                 "Can't lock file (errno: %d - %s)",
	ErCantOpenFile:                             "Can't open file: '%-.200s' (errno: %d - %s)",
	ErFileNotFound:                             "Can't find file: '%-.200s' (errno: %d - %s)",
	ErCantReadDir:                              "Can't read dir of '%-.192s' (errno: %d - %s)",
	ErCantSetWd:                                "Can't change dir to '%-.192s' (errno: %d - %s)",
	ErCheckread:                                "Record has changed since last read in table '%-.192s'",
	ErDiskFull:                                 "Disk full (%s); waiting for someone to free some space... (errno: %d - %s)",
	ErDupKey:                                   "Can't write; duplicate key in table '%-.192s'",
	ErErrorOnClose:                             "Error on close of '%-.192s' (errno: %d - %s)",
	ErErrorOnRead:                              "Error reading file '%-.200s' (errno: %d - %s)",
	ErErrorOnRename:                            "Error on rename of '%-.210s' to '%-.210s' (errno: %d - %s)",
	ErErrorOnWrite:                             "Error writing file '%-.200s' (errno: %d - %s)",
	ErFileUsed:                                 "'%-.192s' is locked against change",
	ErFilsortAbort:                             "Sort aborted",
	ErFormNotFound:                             "View '%-.192s' doesn't exist for '%-.192s'",
	ErGetErrno:                                 "Got error %d from storage engine",
	ErIllegalHa:                                "Table storage engine for '%-.192s' doesn't have this option",
	ErKeyNotFound:                              "Can't find record in '%-.192s'",
	ErNotFormFile:                              "Incorrect information in file: '%-.200s'",
	ErNotKeyfile:                               "Incorrect key file for table '%-.200s'; try to repair it",
	ErOldKeyfile:                               "Old key file for table '%-.192s'; repair it!",
	ErOpenAsReadonly:                           "Table '%-.192s' is read only",
	ErOutofmemory:                              "Out of memory; restart server and try again (needed %d bytes)",
	ErOutOfSortmemory:                          "Out of sort memory, consider increasing server sort buffer size",
	ErUnexpectedEof:                            "Unexpected EOF found when reading file '%-.192s' (errno: %d - %s)",
	ErConCountError:                            "Too many connections",
	ErOutOfResources:                           "Out of memory; check if mysqld or some other process uses all available memory; if not, you may have to use 'ulimit' to allow mysqld to use more memory or you can add more swap space",
	ErBadHostError:                             "Can't get hostname for your address",
	ErHandshakeError:                           "Bad handshake",
	ErDbaccessDeniedError:                      "Access denied for user '%-.48s'@'%-.64s' to database '%-.192s'",
	ErAccessDeniedError:                        "Access denied for user '%-.48s'@'%-.64s' (using password: %s)",
	ErNoDbError:                                "No database selected",
	ErUnknownComError:                          "Unknown command",
	ErBadNullError:                             "Column '%-.192s' cannot be null",
	ErBadDbError:                               "Unknown database '%-.192s'",
	ErTableExistsError:                         "Table '%-.192s' already exists",
	ErBadTableError:                            "Unknown table '%-.100s'",
	ErNonUniqError:                             "Column '%-.192s' in %-.192s is ambiguous",
	ErServerShutdown:                           "Server shutdown in progress",
	ErBadFieldError:                            "Unknown column '%-.192s' in '%-.192s'",
	ErWrongFieldWithGroup:                      "'%-.192s' isn't in GROUP BY",
	ErWrongGroupField:                          "Can't group on '%-.192s'",
	ErWrongSumSelect:                           "Statement has sum functions and columns in same statement",
	ErWrongValueCount:                          "Column count doesn't match value count",
	ErTooLongIdent:                             "Identifier name '%-.100s' is too long",
	ErDupFieldname:                             "Duplicate column name '%-.192s'",
	ErDupKeyname:                               "Duplicate key name '%-.192s'",
	ErDupEntry:                                 "Duplicate entry '%-.192s' for key %d",
	ErWrongFieldSpec:                           "Incorrect column specifier for column '%-.192s'",
	ErParseError:                               "%s near '%-.80s' at line %d",
	ErEmptyQuery:                               "Query was empty",
	ErNonuniqTable:                             "Not unique table/alias: '%-.192s'",
	ErInvalidDefault:                           "Invalid default value for '%-.192s'",
	ErMultiplePriKey:                           "Multiple primary key defined",
	ErTooManyKeys:                              "Too many keys specified; max %d keys allowed",
	ErTooManyKeyParts:                          "Too many key parts specified; max %d parts allowed",
	ErTooLongKey:                               "Specified key was too long; max key length is %d bytes",
	ErKeyColumnDoesNotExits:                    "Key column '%-.192s' doesn't exist in table",
	ErBlobUsedAsKey:                            "BLOB column '%-.192s' can't be used in key specification with the used table type",
	ErTooBigFieldlength:                        "Column length too big for column '%-.192s' (max = %lu); use BLOB or TEXT instead",
	ErWrongAutoKey:                             "Incorrect table definition; there can be only one auto column and it must be defined as a key",
	ErReady:                                    "%s: ready for connections.\nVersion: '%s'  socket: '%s'  port: %d",
	ErNormalShutdown:                           "%s: Normal shutdown\n",
	ErGotSignal:                                "%s: Got signal %d. Aborting!\n",
	ErShutdownComplete:                         "%s: Shutdown complete\n",
	ErForcingClose:                             "%s: Forcing close of thread %ld  user: '%-.48s'\n",
	ErIpsockError:                              "Can't create IP socket",
	ErNoSuchIndex:                              "Table '%-.192s' has no index like the one used in CREATE INDEX; recreate the table",
	ErWrongFieldTerminators:                    "Field separator argument is not what is expected; check the manual",
	ErBlobsAndNoTerminated:                     "You can't use fixed rowlength with BLOBs; please use 'fields terminated by'",
	ErTextfileNotReadable:                      "The file '%-.128s' must be in the database directory or be readable by all",
	ErFileExistsError:                          "File '%-.200s' already exists",
	ErLoadInfo:                                 "Records: %ld  Deleted: %ld  Skipped: %ld  Warnings: %ld",
	ErAlterInfo:                                "Records: %ld  Duplicates: %ld",
	ErWrongSubKey:                              "Incorrect prefix key; the used key part isn't a string, the used length is longer than the key part, or the storage engine doesn't support unique prefix keys",
	ErCantRemoveAllFields:                      "You can't delete all columns with ALTER TABLE; use DROP TABLE instead",
	ErCantDropFieldOrKey:                       "Can't DROP '%-.192s'; check that column/key exists",
	ErInsertInfo:                               "Records: %ld  Duplicates: %ld  Warnings: %ld",
	ErUpdateTableUsed:                          "You can't specify target table '%-.192s' for update in FROM clause",
	ErNoSuchThread:                             "Unknown thread id: %lu",
	ErKillDeniedError:                          "You are not owner of thread %lu",
	ErNoTablesUsed:                             "No tables used",
	ErTooBigSet:                                "Too many strings for column %-.192s and SET",
	ErNoUniqueLogfile:                          "Can't generate a unique log-filename %-.200s.(1-999)\n",
	ErTableNotLockedForWrite:                   "Table '%-.192s' was locked with a READ lock and can't be updated",
	ErTableNotLocked:                           "Table '%-.192s' was not locked with LOCK TABLES",
	ErBlobCantHaveDefault:                      "BLOB/TEXT column '%-.192s' can't have a default value",
	ErWrongDbName:                              "Incorrect database name '%-.100s'",
	ErWrongTableName:                           "Incorrect table name '%-.100s'",
	ErTooBigSelect:                             "The SELECT would examine more than MAX_JOIN_SIZE rows; check your WHERE and use SET SQL_BIG_SELECTS=1 or SET MAX_JOIN_SIZE=# if the SELECT is okay",
	ErUnknownError:                             "Unknown error",
	ErUnknownProcedure:                         "Unknown procedure '%-.192s'",
	ErWrongParamcountToProcedure:               "Incorrect parameter count to procedure '%-.192s'",
	ErWrongParametersToProcedure:               "Incorrect parameters to procedure '%-.192s'",
	ErUnknownTable:                             "Unknown table '%-.192s' in %-.32s",
	ErFieldSpecifiedTwice:                      "Column '%-.192s' specified twice",
	ErInvalidGroupFuncUse:                      "Invalid use of group function",
	ErUnsupportedExtension:                     "Table '%-.192s' uses an extension that doesn't exist in this MySQL version",
	ErTableMustHaveColumns:                     "A table must have at least 1 column",
	ErRecordFileFull:                           "The table '%-.192s' is full",
	ErUnknownCharacterSet:                      "Unknown character set: '%-.64s'",
	ErTooManyTables:                            "Too many tables; MySQL can only use %d tables in a join",
	ErTooManyFields:                            "Too many columns",
	ErTooBigRowsize:                            "Row size too large. The maximum row size for the used table type, not counting BLOBs, is %ld. This includes storage overhead, check the manual. You have to change some columns to TEXT or BLOBs",
	ErStackOverrun:                             "Thread stack overrun:  Used: %ld of a %ld stack.  Use 'mysqld --thread_stack=#' to specify a bigger stack if needed",
	ErWrongOuterJoin:                           "Cross dependency found in OUTER JOIN; examine your ON conditions",
	ErNullColumnInIndex:                        "Table handler doesn't support NULL in given index. Please change column '%-.192s' to be NOT NULL or use another handler",
	ErCantFindUdf:                              "Can't load function '%-.192s'",
	ErCantInitializeUdf:                        "Can't initialize function '%-.192s'; %-.80s",
	ErUdfNoPaths:                               "No paths allowed for shared library",
	ErUdfExists:                                "Function '%-.192s' already exists",
	ErCantOpenLibrary:                          "Can't open shared library '%-.192s' (errno: %d %-.128s)",
	ErCantFindDlEntry:                          "Can't find symbol '%-.128s' in library",
	ErFunctionNotDefined:                       "Function '%-.192s' is not defined",
	ErHostIsBlocked:                            "Host '%-.64s' is blocked because of many connection errors; unblock with 'mysqladmin flush-hosts'",
	ErHostNotPrivileged:                        "Host '%-.64s' is not allowed to connect to this MySQL server",
	ErPasswordAnonymousUser:                    "You are using MySQL as an anonymous user and anonymous users are not allowed to change passwords",
	ErPasswordNotAllowed:                       "You must have privileges to update tables in the mysql database to be able to change passwords for others",
	ErPasswordNoMatch:                          "Can't find any matching row in the user table",
	ErUpdateInfo:                               "Rows matched: %ld  Changed: %ld  Warnings: %ld",
	ErCantCreateThread:                         "Can't create a new thread (errno %d); if you are not out of available memory, you can consult the manual for a possible OS-dependent bug",
	ErWrongValueCountOnRow:                     "Column count doesn't match value count at row %ld",
	ErCantReopenTable:                          "Can't reopen table: '%-.192s'",
	ErInvalidUseOfNull:                         "Invalid use of NULL value",
	ErRegexpError:                              "Got error '%-.64s' from regexp",
	ErMixOfGroupFuncAndFields:                  "Mixing of GROUP columns (MIN(),MAX(),COUNT(),...) with no GROUP columns is illegal if there is no GROUP BY clause",
	ErNonexistingGrant:                         "There is no such grant defined for user '%-.48s' on host '%-.64s'",
	ErTableaccessDeniedError:                   "%-.128s command denied to user '%-.48s'@'%-.64s' for table '%-.64s'",
	ErColumnaccessDeniedError:                  "%-.16s command denied to user '%-.48s'@'%-.64s' for column '%-.192s' in table '%-.192s'",
	ErIllegalGrantForTable:                     "Illegal GRANT/REVOKE command; please consult the manual to see which privileges can be used",
	ErGrantWrongHostOrUser:                     "The host or user argument to GRANT is too long",
	ErNoSuchTable:                              "Table '%-.192s.%-.192s' doesn't exist",
	ErNonexistingTableGrant:                    "There is no such grant defined for user '%-.48s' on host '%-.64s' on table '%-.192s'",
	ErNotAllowedCommand:                        "The used command is not allowed with this MySQL version",
	ErSyntaxError:                              "You have an error in your SQL syntax; check the manual that corresponds to your MySQL server version for the right syntax to use",
	ErDelayedCantChangeLock:                    "Delayed insert thread couldn't get requested lock for table %-.192s",
	ErTooManyDelayedThreads:                    "Too many delayed threads in use",
	ErAbortingConnection:                       "Aborted connection %ld to db: '%-.192s' user: '%-.48s' (%-.64s)",
	ErNetPacketTooLarge:                        "Got a packet bigger than 'max_allowed_packet' bytes",
	ErNetReadErrorFromPipe:                     "Got a read error from the connection pipe",
	ErNetFcntlError:                            "Got an error from fcntl()",
	ErNetPacketsOutOfOrder:                     "Got packets out of order",
	ErNetUncompressError:                       "Couldn't uncompress communication packet",
	ErNetReadError:                             "Got an error reading communication packets",
	ErNetReadInterrupted:                       "Got timeout reading communication packets",
	ErNetErrorOnWrite:                          "Got an error writing communication packets",
	ErNetWriteInterrupted:                      "Got timeout writing communication packets",
	ErTooLongString:                            "Result string is longer than 'max_allowed_packet' bytes",
	ErTableCantHandleBlob:                      "The used table type doesn't support BLOB/TEXT columns",
	ErTableCantHandleAutoIncrement:             "The used table type doesn't support AUTO_INCREMENT columns",
	ErDelayedInsertTableLocked:                 "INSERT DELAYED can't be used with table '%-.192s' because it is locked with LOCK TABLES",
	ErWrongColumnName:                          "Incorrect column name '%-.100s'",
	ErWrongKeyColumn:                           "The used storage engine can't index column '%-.192s'",
	ErWrongMrgTable:                            "Unable to open underlying table which is differently defined or of non-MyISAM type or doesn't exist",
	ErDupUnique:                                "Can't write, because of unique constraint, to table '%-.192s'",
	ErBlobKeyWithoutLength:                     "BLOB/TEXT column '%-.192s' used in key specification without a key length",
	ErPrimaryCantHaveNull:                      "All parts of a PRIMARY KEY must be NOT NULL; if you need NULL in a key, use UNIQUE instead",
	ErTooManyRows:                              "Result consisted of more than one row",
	ErRequiresPrimaryKey:                       "This table type requires a primary key",
	ErNoRaidCompiled:                           "This version of MySQL is not compiled with RAID support",
	ErUpdateWithoutKeyInSafeMode:               "You are using safe update mode and you tried to update a table without a WHERE that uses a KEY column",
	ErKeyDoesNotExits:                          "Key '%-.192s' doesn't exist in table '%-.192s'",
	ErCheckNoSuchTable:                         "Can't open table",
	ErCheckNotImplemented:                      "The storage engine for the table doesn't support %s",
	ErCantDoThisDuringAnTransaction:            "You are not allowed to execute this command in a transaction",
	ErErrorDuringCommit:                        "Got error %d during COMMIT",
	ErErrorDuringRollback:                      "Got error %d during ROLLBACK",
	ErErrorDuringFlushLogs:                     "Got error %d during FLUSH_LOGS",
	ErErrorDuringCheckpoint:                    "Got error %d during CHECKPOINT",
	ErNewAbortingConnection:                    "Aborted connection %ld to db: '%-.192s' user: '%-.48s' host: '%-.64s' (%-.64s)",
	ErDumpNotImplemented:                       "The storage engine for the table does not support binary table dump",
	ErFlushMasterBinlogClosed:                  "Binlog closed, cannot RESET MASTER",
	ErIndexRebuild:                             "Failed rebuilding the index of  dumped table '%-.192s'",
	ErMaster:                                   "Error from master: '%-.64s'",
	ErMasterNetRead:                            "Net error reading from master",
	ErMasterNetWrite:                           "Net error writing to master",
	ErFtMatchingKeyNotFound:                    "Can't find FULLTEXT index matching the column list",
	ErLockOrActiveTransaction:                  "Can't execute the given command because you have active locked tables or an active transaction",
	ErUnknownSystemVariable:                    "Unknown system variable '%-.64s'",
	ErCrashedOnUsage:                           "Table '%-.192s' is marked as crashed and should be repaired",
	ErCrashedOnRepair:                          "Table '%-.192s' is marked as crashed and last (automatic?) repair failed",
	ErWarningNotCompleteRollback:               "Some non-transactional changed tables couldn't be rolled back",
	ErTransCacheFull:                           "Multi-statement transaction required more than 'max_binlog_cache_size' bytes of storage; increase this mysqld variable and try again",
	ErSlaveMustStop:                            "This operation cannot be performed with a running slave; run STOP SLAVE first",
	ErSlaveNotRunning:                          "This operation requires a running slave; configure slave and do START SLAVE",
	ErBadSlave:                                 "The server is not configured as slave; fix in config file or with CHANGE MASTER TO",
	ErMasterInfo:                               "Could not initialize master info structure; more error messages can be found in the MySQL error log",
	ErSlaveThread:                              "Could not create slave thread; check system resources",
	ErTooManyUserConnections:                   "User %-.64s already has more than 'max_user_connections' active connections",
	ErSetConstantsOnly:                         "You may only use constant expressions with SET",
	ErLockWaitTimeout:                          "Lock wait timeout exceeded; try restarting transaction",
	ErLockTableFull:                            "The total number of locks exceeds the lock table size",
	ErReadOnlyTransaction:                      "Update locks cannot be acquired during a READ UNCOMMITTED transaction",
	ErDropDbWithReadLock:                       "DROP DATABASE not allowed while thread is holding global read lock",
	ErCreateDbWithReadLock:                     "CREATE DATABASE not allowed while thread is holding global read lock",
	ErWrongArguments:                           "Incorrect arguments to %s",
	ErNoPermissionToCreateUser:                 "'%-.48s'@'%-.64s' is not allowed to create new users",
	ErUnionTablesInDifferentDir:                "Incorrect table definition; all MERGE tables must be in the same database",
	ErLockDeadlock:                             "Deadlock found when trying to get lock; try restarting transaction",
	ErTableCantHandleFt:                        "The used table type doesn't support FULLTEXT indexes",
	ErCannotAddForeign:                         "Cannot add foreign key constraint",
	ErNoReferencedRow:                          "Cannot add or update a child row: a foreign key constraint fails",
	ErRowIsReferenced:                          "Cannot delete or update a parent row: a foreign key constraint fails",
	ErConnectToMaster:                          "Error connecting to master: %-.128s",
	ErQueryOnMaster:                            "Error running query on master: %-.128s",
	ErErrorWhenExecutingCommand:                "Error when executing command %s: %-.128s",
	ErWrongUsage:                               "Incorrect usage of %s and %s",
	ErWrongNumberOfColumnsInSelect:             "The used SELECT statements have a different number of columns",
	ErCantUpdateWithReadlock:                   "Can't execute the query because you have a conflicting read lock",
	ErMixingNotAllowed:                         "Mixing of transactional and non-transactional tables is disabled",
	ErDupArgument:                              "Option '%s' used twice in statement",
	ErUserLimitReached:                         "User '%-.64s' has exceeded the '%s' resource (current value: %ld)",
	ErSpecificAccessDeniedError:                "Access denied; you need (at least one of) the %-.128s privilege(s) for this operation",
	ErLocalVariable:                            "Variable '%-.64s' is a SESSION variable and can't be used with SET GLOBAL",
	ErGlobalVariable:                           "Variable '%-.64s' is a GLOBAL variable and should be set with SET GLOBAL",
	ErNoDefault:                                "Variable '%-.64s' doesn't have a default value",
	ErWrongValueForVar:                         "Variable '%-.64s' can't be set to the value of '%-.200s'",
	ErWrongTypeForVar:                          "Incorrect argument type to variable '%-.64s'",
	ErVarCantBeRead:                            "Variable '%-.64s' can only be set, not read",
	ErCantUseOptionHere:                        "Incorrect usage/placement of '%s'",
	ErNotSupportedYet:                          "This version of MySQL doesn't yet support '%s'",
	ErMasterFatalErrorReadingBinlog:            "Got fatal error %d from master when reading data from binary log: '%-.320s'",
	ErSlaveIgnoredTable:                        "Slave SQL thread ignored the query because of replicate-*-table rules",
	ErIncorrectGlobalLocalVar:                  "Variable '%-.192s' is a %s variable",
	ErWrongFkDef:                               "Incorrect foreign key definition for '%-.192s': %s",
	ErKeyRefDoNotMatchTableRef:                 "Key reference and table reference don't match",
	ErOperandColumns:                           "Operand should contain %d column(s)",
	ErSubqueryNo1Row:                           "Subquery returns more than 1 row",
	ErUnknownStmtHandler:                       "Unknown prepared statement handler (%.*s) given to %s",
	ErCorruptHelpDb:                            "Help database is corrupt or does not exist",
	ErCyclicReference:                          "Cyclic reference on subqueries",
	ErAutoConvert:                              "Converting column '%s' from %s to %s",
	ErIllegalReference:                         "Reference '%-.64s' not supported (%s)",
	ErDerivedMustHaveAlias:                     "Every derived table must have its own alias",
	ErSelectReduced:                            "Select %u was reduced during optimization",
	ErTablenameNotAllowedHere:                  "Table '%-.192s' from one of the SELECTs cannot be used in %-.32s",
	ErNotSupportedAuthMode:                     "Client does not support authentication protocol requested by server; consider upgrading MySQL client",
	ErSpatialCantHaveNull:                      "All parts of a SPATIAL index must be NOT NULL",
	ErCollationCharsetMismatch:                 "COLLATION '%s' is not valid for CHARACTER SET '%s'",
	ErSlaveWasRunning:                          "Slave is already running",
	ErSlaveWasNotRunning:                       "Slave already has been stopped",
	ErTooBigForUncompress:                      "Uncompressed data size too large; the maximum size is %d (probably, length of uncompressed data was corrupted)",
	ErZlibZMemError:                            "ZLIB: Not enough memory",
	ErZlibZBufError:                            "ZLIB: Not enough room in the output buffer (probably, length of uncompressed data was corrupted)",
	ErZlibZDataError:                           "ZLIB: Input data corrupted",
	ErCutValueGroupConcat:                      "Row %u was cut by GROUP_CONCAT()",
	ErWarnTooFewRecords:                        "Row %ld doesn't contain data for all columns",
	ErWarnTooManyRecords:                       "Row %ld was truncated; it contained more data than there were input columns",
	ErWarnNullToNotnull:                        "Column set to default value; NULL supplied to NOT NULL column '%s' at row %ld",
	ErWarnDataOutOfRange:                       "Out of range value for column '%s' at row %ld",
	WarnDataTruncated:                          "Data truncated for column '%s' at row %ld",
	ErWarnUsingOtherHandler:                    "Using storage engine %s for table '%s'",
	ErCantAggregate2collations:                 "Illegal mix of collations (%s,%s) and (%s,%s) for operation '%s'",
	ErDropUser:                                 "Cannot drop one or more of the requested users",
	ErRevokeGrants:                             "Can't revoke all privileges for one or more of the requested users",
	ErCantAggregate3collations:                 "Illegal mix of collations (%s,%s), (%s,%s), (%s,%s) for operation '%s'",
	ErCantAggregateNcollations:                 "Illegal mix of collations for operation '%s'",
	ErVariableIsNotStruct:                      "Variable '%-.64s' is not a variable component (can't be used as XXXX.variable_name)",
	ErUnknownCollation:                         "Unknown collation: '%-.64s'",
	ErSlaveIgnoredSslParams:                    "SSL parameters in CHANGE MASTER are ignored because this MySQL slave was compiled without SSL support; they can be used later if MySQL slave with SSL is started",
	ErServerIsInSecureAuthMode:                 "Server is running in --secure-auth mode, but '%s'@'%s' has a password in the old format; please change the password to the new format",
	ErWarnFieldResolved:                        "Field or reference '%-.192s%s%-.192s%s%-.192s' of SELECT #%d was resolved in SELECT #%d",
	ErBadSlaveUntilCond:                        "Incorrect parameter or combination of parameters for START SLAVE UNTIL",
	ErMissingSkipSlave:                         "It is recommended to use --skip-slave-start when doing step-by-step replication with START SLAVE UNTIL; otherwise, you will get problems if you get an unexpected slave's mysqld restart",
	ErUntilCondIgnored:                         "SQL thread is not to be started so UNTIL options are ignored",
	ErWrongNameForIndex:                        "Incorrect index name '%-.100s'",
	ErWrongNameForCatalog:                      "Incorrect catalog name '%-.100s'",
	ErWarnQcResize:                             "Query cache failed to set size %lu; new query cache size is %lu",
	ErBadFtColumn:                              "Column '%-.192s' cannot be part of FULLTEXT index",
	ErUnknownKeyCache:                          "Unknown key cache '%-.100s'",
	ErWarnHostnameWontWork:                     "MySQL is started in --skip-name-resolve mode; you must restart it without this switch for this grant to work",
	ErUnknownStorageEngine:                     "Unknown storage engine '%s'",
	ErWarnDeprecatedSyntax:                     "'%s' is deprecated and will be removed in a future release. Please use %s instead",
	ErNonUpdatableTable:                        "The target table %-.100s of the %s is not updatable",
	ErFeatureDisabled:                          "The '%s' feature is disabled; you need MySQL built with '%s' to have it working",
	ErOptionPreventsStatement:                  "The MySQL server is running with the %s option so it cannot execute this statement",
	ErDuplicatedValueInType:                    "Column '%-.100s' has duplicated value '%-.64s' in %s",
	ErTruncatedWrongValue:                      "Truncated incorrect %-.32s value: '%-.128s'",
	ErTooMuchAutoTimestampCols:                 "Incorrect table definition; there can be only one TIMESTAMP column with CURRENT_TIMESTAMP in DEFAULT or ON UPDATE clause",
	ErInvalidOnUpdate:                          "Invalid ON UPDATE clause for '%-.192s' column",
	ErUnsupportedPs:                            "This command is not supported in the prepared statement protocol yet",
	ErGetErrmsg:                                "Got error %d '%-.100s' from %s",
	ErGetTemporaryErrmsg:                       "Got temporary error %d '%-.100s' from %s",
	ErUnknownTimeZone:                          "Unknown or incorrect time zone: '%-.64s'",
	ErWarnInvalidTimestamp:                     "Invalid TIMESTAMP value in column '%s' at row %ld",
	ErInvalidCharacterString:                   "Invalid %s character string: '%.64s'",
	ErWarnAllowedPacketOverflowed:              "Result of %s() was larger than max_allowed_packet (%ld) - truncated",
	ErConflictingDeclarations:                  "Conflicting declarations: '%s%s' and '%s%s'",
	ErSpNoRecursiveCreate:                      "Can't create a %s from within another stored routine",
	ErSpAlreadyExists:                          "%s %s already exists",
	ErSpDoesNotExist:                           "%s %s does not exist",
	ErSpDropFailed:                             "Failed to DROP %s %s",
	ErSpStoreFailed:                            "Failed to CREATE %s %s",
	ErSpLilabelMismatch:                        "%s with no matching label: %s",
	ErSpLabelRedefine:                          "Redefining label %s",
	ErSpLabelMismatch:                          "End-label %s without match",
	ErSpUninitVar:                              "Referring to uninitialized variable %s",
	ErSpBadselect:                              "PROCEDURE %s can't return a result set in the given context",
	ErSpBadreturn:                              "RETURN is only allowed in a FUNCTION",
	ErSpBadstatement:                           "%s is not allowed in stored procedures",
	ErUpdateLogDeprecatedIgnored:               "The update log is deprecated and replaced by the binary log; SET SQL_LOG_UPDATE has been ignored.",
	ErUpdateLogDeprecatedTranslated:            "The update log is deprecated and replaced by the binary log; SET SQL_LOG_UPDATE has been translated to SET SQL_LOG_BIN.",
	ErQueryInterrupted:                         "Query execution was interrupted",
	ErSpWrongNoOfArgs:                          "Incorrect number of arguments for %s %s; expected %u, got %u",
	ErSpCondMismatch:                           "Undefined CONDITION: %s",
	ErSpNoreturn:                               "No RETURN found in FUNCTION %s",
	ErSpNoreturnend:                            "FUNCTION %s ended without RETURN",
	ErSpBadCursorQuery:                         "Cursor statement must be a SELECT",
	ErSpBadCursorSelect:                        "Cursor SELECT must not have INTO",
	ErSpCursorMismatch:                         "Undefined CURSOR: %s",
	ErSpCursorAlreadyOpen:                      "Cursor is already open",
	ErSpCursorNotOpen:                          "Cursor is not open",
	ErSpUndeclaredVar:                          "Undeclared variable: %s",
	ErSpWrongNoOfFetchArgs:                     "Incorrect number of FETCH variables",
	ErSpFetchNoData:                            "No data - zero rows fetched, selected, or processed",
	ErSpDupParam:                               "Duplicate parameter: %s",
	ErSpDupVar:                                 "Duplicate variable: %s",
	ErSpDupCond:                                "Duplicate condition: %s",
	ErSpDupCurs:                                "Duplicate cursor: %s",
	ErSpCantAlter:                              "Failed to ALTER %s %s",
	ErSpSubselectNyi:                           "Subquery value not supported",
	ErStmtNotAllowedInSfOrTrg:                  "%s is not allowed in stored function or trigger",
	ErSpVarcondAfterCurshndlr:                  "Variable or condition declaration after cursor or handler declaration",
	ErSpCursorAfterHandler:                     "Cursor declaration after handler declaration",
	ErSpCaseNotFound:                           "Case not found for CASE statement",
	ErFparserTooBigFile:                        "Configuration file '%-.192s' is too big",
	ErFparserBadHeader:                         "Malformed file type header in file '%-.192s'",
	ErFparserEofInComment:                      "Unexpected end of file while parsing comment '%-.200s'",
	ErFparserErrorInParameter:                  "Error while parsing parameter '%-.192s' (line: '%-.192s')",
	ErFparserEofInUnknownParameter:             "Unexpected end of file while skipping unknown parameter '%-.192s'",
	ErViewNoExplain:                            "EXPLAIN/SHOW can not be issued; lacking privileges for underlying table",
	ErFrmUnknownType:                           "File '%-.192s' has unknown type '%-.64s' in its header",
	ErWrongObject:                              "'%-.192s.%-.192s' is not %s",
	ErNonupdateableColumn:                      "Column '%-.192s' is not updatable",
	ErViewSelectDerived:                        "View's SELECT contains a subquery in the FROM clause",
	ErViewSelectClause:                         "View's SELECT contains a '%s' clause",
	ErViewSelectVariable:                       "View's SELECT contains a variable or parameter",
	ErViewSelectTmptable:                       "View's SELECT refers to a temporary table '%-.192s'",
	ErViewWrongList:                            "View's SELECT and view's field list have different column counts",
	ErWarnViewMerge:                            "View merge algorithm can't be used here for now (assumed undefined algorithm)",
	ErWarnViewWithoutKey:                       "View being updated does not have complete key of underlying table in it",
	ErViewInvalid:                              "View '%-.192s.%-.192s' references invalid table(s) or column(s) or function(s) or definer/invoker of view lack rights to use them",
	ErSpNoDropSp:                               "Can't drop or alter a %s from within another stored routine",
	ErSpGotoInHndlr:                            "GOTO is not allowed in a stored procedure handler",
	ErTrgAlreadyExists:                         "Trigger already exists",
	ErTrgDoesNotExist:                          "Trigger does not exist",
	ErTrgOnViewOrTempTable:                     "Trigger's '%-.192s' is view or temporary table",
	ErTrgCantChangeRow:                         "Updating of %s row is not allowed in %strigger",
	ErTrgNoSuchRowInTrg:                        "There is no %s row in %s trigger",
	ErNoDefaultForField:                        "Field '%-.192s' doesn't have a default value",
	ErDivisionByZero:                           "Division by 0",
	ErTruncatedWrongValueForField:              "Incorrect %-.32s value: '%-.128s' for column '%.192s' at row %ld",
	ErIllegalValueForType:                      "Illegal %s '%-.192s' value found during parsing",
	ErViewNonupdCheck:                          "CHECK OPTION on non-updatable view '%-.192s.%-.192s'",
	ErViewCheckFailed:                          "CHECK OPTION failed '%-.192s.%-.192s'",
	ErProcaccessDeniedError:                    "%-.16s command denied to user '%-.48s'@'%-.64s' for routine '%-.192s'",
	ErRelayLogFail:                             "Failed purging old relay logs: %s",
	ErPasswdLength:                             "Password hash should be a %d-digit hexadecimal number",
	ErUnknownTargetBinlog:                      "Target log not found in binlog index",
	ErIoErrLogIndexRead:                        "I/O error reading log index file",
	ErBinlogPurgeProhibited:                    "Server configuration does not permit binlog purge",
	ErFseekFail:                                "Failed on fseek()",
	ErBinlogPurgeFatalErr:                      "Fatal error during log purge",
	ErLogInUse:                                 "A purgeable log is in use, will not purge",
	ErLogPurgeUnknownErr:                       "Unknown error during log purge",
	ErRelayLogInit:                             "Failed initializing relay log position: %s",
	ErNoBinaryLogging:                          "You are not using binary logging",
	ErReservedSyntax:                           "The '%-.64s' syntax is reserved for purposes internal to the MySQL server",
	ErWsasFailed:                               "WSAStartup Failed",
	ErDiffGroupsProc:                           "Can't handle procedures with different groups yet",
	ErNoGroupForProc:                           "Select must have a group with this procedure",
	ErOrderWithProc:                            "Can't use ORDER clause with this procedure",
	ErLoggingProhibitChangingOf:                "Binary logging and replication forbid changing the global server %s",
	ErNoFileMapping:                            "Can't map file: %-.200s, errno: %d",
	ErWrongMagic:                               "Wrong magic in %-.64s",
	ErPsManyParam:                              "Prepared statement contains too many placeholders",
	ErKeyPart0:                                 "Key part '%-.192s' length cannot be 0",
	ErViewChecksum:                             "View text checksum failed",
	ErViewMultiupdate:                          "Can not modify more than one base table through a join view '%-.192s.%-.192s'",
	ErViewNoInsertFieldList:                    "Can not insert into join view '%-.192s.%-.192s' without fields list",
	ErViewDeleteMergeView:                      "Can not delete from join view '%-.192s.%-.192s'",
	ErCannotUser:                               "Operation %s failed for %.256s",
	ErXaerNota:                                 "XAER_NOTA: Unknown XID",
	ErXaerInval:                                "XAER_INVAL: Invalid arguments (or unsupported command)",
	ErXaerRmfail:                               "XAER_RMFAIL: The command cannot be executed when global transaction is in the  %.64s state",
	ErXaerOutside:                              "XAER_OUTSIDE: Some work is done outside global transaction",
	ErXaerRmerr:                                "XAER_RMERR: Fatal error occurred in the transaction branch - check your data for consistency",
	ErXaRbrollback:                             "XA_RBROLLBACK: Transaction branch was rolled back",
	ErNonexistingProcGrant:                     "There is no such grant defined for user '%-.48s' on host '%-.64s' on routine '%-.192s'",
	ErProcAutoGrantFail:                        "Failed to grant EXECUTE and ALTER ROUTINE privileges",
	ErProcAutoRevokeFail:                       "Failed to revoke all privileges to dropped routine",
	ErDataTooLong:                              "Data too long for column '%s' at row %ld",
	ErSpBadSqlstate:                            "Bad SQLSTATE: '%s'",
	ErStartup:                                  "%s: ready for connections.\nVersion: '%s'  socket: '%s'  port: %d  %s",
	ErLoadFromFixedSizeRowsToVar:               "Can't load value from file with fixed size rows to variable",
	ErCantCreateUserWithGrant:                  "You are not allowed to create a user with GRANT",
	ErWrongValueForType:                        "Incorrect %-.32s value: '%-.128s' for function %-.32s",
	ErTableDefChanged:                          "Table definition has changed, please retry transaction",
	ErSpDupHandler:                             "Duplicate handler declared in the same block",
	ErSpNotVarArg:                              "OUT or INOUT argument %d for routine %s is not a variable or NEW pseudo-variable in BEFORE trigger",
	ErSpNoRetset:                               "Not allowed to return a result set from a %s",
	ErCantCreateGeometryObject:                 "Cannot get geometry object from data you send to the GEOMETRY field",
	ErFailedRoutineBreakBinlog:                 "A routine failed and has neither NO SQL nor READS SQL DATA in its declaration and binary logging is enabled; if non-transactional tables were updated, the binary log will miss their changes",
	ErBinlogUnsafeRoutine:                      "This function has none of DETERMINISTIC, NO SQL, or READS SQL DATA in its declaration and binary logging is enabled (you *might* want to use the less safe log_bin_trust_function_creators variable)",
	ErBinlogCreateRoutineNeedSuper:             "You do not have the SUPER privilege and binary logging is enabled (you *might* want to use the less safe log_bin_trust_function_creators variable)",
	ErExecStmtWithOpenCursor:                   "You can't execute a prepared statement which has an open cursor associated with it. Reset the statement to re-execute it.",
	ErStmtHasNoOpenCursor:                      "The statement (%lu) has no open cursor.",
	ErCommitNotAllowedInSfOrTrg:                "Explicit or implicit commit is not allowed in stored function or trigger.",
	ErNoDefaultForViewField:                    "Field of view '%-.192s.%-.192s' underlying table doesn't have a default value",
	ErSpNoRecursion:                            "Recursive stored functions and triggers are not allowed.",
	ErTooBigScale:                              "Too big scale %d specified for column '%-.192s'. Maximum is %lu.",
	ErTooBigPrecision:                          "Too big precision %d specified for column '%-.192s'. Maximum is %lu.",
	ErMBiggerThanD:                             "For float(M,D), double(M,D) or decimal(M,D), M must be >= D (column '%-.192s').",
	ErWrongLockOfSystemTable:                   "You can't combine write-locking of system tables with other tables or lock types",
	ErConnectToForeignDataSource:               "Unable to connect to foreign data source: %.64s",
	ErQueryOnForeignDataSource:                 "There was a problem processing the query on the foreign data source. Data source error: %-.64s",
	ErForeignDataSourceDoesntExist:             "The foreign data source you are trying to reference does not exist. Data source error:  %-.64s",
	ErForeignDataStringInvalidCantCreate:       "Can't create federated table. The data source connection string '%-.64s' is not in the correct format",
	ErForeignDataStringInvalid:                 "The data source connection string '%-.64s' is not in the correct format",
	ErCantCreateFederatedTable:                 "Can't create federated table. Foreign data src error:  %-.64s",
	ErTrgInWrongSchema:                         "Trigger in wrong schema",
	ErStackOverrunNeedMore:                     "Thread stack overrun:  %ld bytes used of a %ld byte stack, and %ld bytes needed.  Use 'mysqld --thread_stack=#' to specify a bigger stack.",
	ErTooLongBody:                              "Routine body for '%-.100s' is too long",
	ErWarnCantDropDefaultKeycache:              "Cannot drop default keycache",
	ErTooBigDisplaywidth:                       "Display width out of range for column '%-.192s' (max = %lu)",
	ErXaerDupid:                                "XAER_DUPID: The XID already exists",
	ErDatetimeFunctionOverflow:                 "Datetime function: %-.32s field overflow",
	ErCantUpdateUsedTableInSfOrTrg:             "Can't update table '%-.192s' in stored function/trigger because it is already used by statement which invoked this stored function/trigger.",
	ErViewPreventUpdate:                        "The definition of table '%-.192s' prevents operation %.192s on table '%-.192s'.",
	ErPsNoRecursion:                            "The prepared statement contains a stored routine call that refers to that same statement. It's not allowed to execute a prepared statement in such a recursive manner",
	ErSpCantSetAutocommit:                      "Not allowed to set autocommit from a stored function or trigger",
	ErMalformedDefiner:                         "Definer is not fully qualified",
	ErViewFrmNoUser:                            "View '%-.192s'.'%-.192s' has no definer information (old table format). Current user is used as definer. Please recreate the view!",
	ErViewOtherUser:                            "You need the SUPER privilege for creation view with '%-.192s'@'%-.192s' definer",
	ErNoSuchUser:                               "The user specified as a definer ('%-.64s'@'%-.64s') does not exist",
	ErForbidSchemaChange:                       "Changing schema from '%-.192s' to '%-.192s' is not allowed.",
	ErRowIsReferenced2:                         "Cannot delete or update a parent row: a foreign key constraint fails (%.192s)",
	ErNoReferencedRow2:                         "Cannot add or update a child row: a foreign key constraint fails (%.192s)",
	ErSpBadVarShadow:                           "Variable '%-.64s' must be quoted with `...`, or renamed",
	ErTrgNoDefiner:                             "No definer attribute for trigger '%-.192s'.'%-.192s'. The trigger will be activated under the authorization of the invoker, which may have insufficient privileges. Please recreate the trigger.",
	ErOldFileFormat:                            "'%-.192s' has an old format, you should re-create the '%s' object(s)",
	ErSpRecursionLimit:                         "Recursive limit %d (as set by the max_sp_recursion_depth variable) was exceeded for routine %.192s",
	ErSpProcTableCorrupt:                       "Failed to load routine %-.192s. The table mysql.proc is missing, corrupt, or contains bad data (internal code %d)",
	ErSpWrongName:                              "Incorrect routine name '%-.192s'",
	ErTableNeedsUpgrade:                        "Table upgrade required. Please do \"REPAIR TABLE `%-.32s`\" or dump/reload to fix it!",
	ErSpNoAggregate:                            "AGGREGATE is not supported for stored functions",
	ErMaxPreparedStmtCountReached:              "Can't create more than max_prepared_stmt_count statements (current value: %lu)",
	ErViewRecursive:                            "`%-.192s`.`%-.192s` contains view recursion",
	ErNonGroupingFieldUsed:                     "Non-grouping field '%-.192s' is used in %-.64s clause",
	ErTableCantHandleSpkeys:                    "The used table type doesn't support SPATIAL indexes",
	ErNoTriggersOnSystemSchema:                 "Triggers can not be created on system tables",
	ErRemovedSpaces:                            "Leading spaces are removed from name '%s'",
	ErAutoincReadFailed:                        "Failed to read auto-increment value from storage engine",
	ErUsername:                                 "user name",
	ErHostname:                                 "host name",
	ErWrongStringLength:                        "String '%-.70s' is too long for %s (should be no longer than %d)",
	ErNonInsertableTable:                       "The target table %-.100s of the %s is not insertable-into",
	ErAdminWrongMrgTable:                       "Table '%-.64s' is differently defined or of non-MyISAM type or doesn't exist",
	ErTooHighLevelOfNestingForSelect:           "Too high level of nesting for select",
	ErNameBecomesEmpty:                         "Name '%-.64s' has become ''",
	ErAmbiguousFieldTerm:                       "First character of the FIELDS TERMINATED string is ambiguous; please use non-optional and non-empty FIELDS ENCLOSED BY",
	ErForeignServerExists:                      "The foreign server, %s, you are trying to create already exists.",
	ErForeignServerDoesntExist:                 "The foreign server name you are trying to reference does not exist. Data source error:  %-.64s",
	ErIllegalHaCreateOption:                    "Table storage engine '%-.64s' does not support the create option '%.64s'",
	ErPartitionRequiresValuesError:             "Syntax error: %-.64s PARTITIONING requires definition of VALUES %-.64s for each partition",
	ErPartitionWrongValuesError:                "Only %-.64s PARTITIONING can use VALUES %-.64s in partition definition",
	ErPartitionMaxvalueError:                   "MAXVALUE can only be used in last partition definition",
	ErPartitionSubpartitionError:               "Subpartitions can only be hash partitions and by key",
	ErPartitionSubpartMixError:                 "Must define subpartitions on all partitions if on one partition",
	ErPartitionWrongNoPartError:                "Wrong number of partitions defined, mismatch with previous setting",
	ErPartitionWrongNoSubpartError:             "Wrong number of subpartitions defined, mismatch with previous setting",
	ErWrongExprInPartitionFuncError:            "Constant, random or timezone-dependent expressions in (sub)partitioning function are not allowed",
	ErNoConstExprInRangeOrListError:            "Expression in RANGE/LIST VALUES must be constant",
	ErFieldNotFoundPartError:                   "Field in list of fields for partition function not found in table",
	ErListOfFieldsOnlyInHashError:              "List of fields is only allowed in KEY partitions",
	ErInconsistentPartitionInfoError:           "The partition info in the frm file is not consistent with what can be written into the frm file",
	ErPartitionFuncNotAllowedError:             "The %-.192s function returns the wrong type",
	ErPartitionsMustBeDefinedError:             "For %-.64s partitions each partition must be defined",
	ErRangeNotIncreasingError:                  "VALUES LESS THAN value must be strictly increasing for each partition",
	ErInconsistentTypeOfFunctionsError:         "VALUES value must be of same type as partition function",
	ErMultipleDefConstInListPartError:          "Multiple definition of same constant in list partitioning",
	ErPartitionEntryError:                      "Partitioning can not be used stand-alone in query",
	ErMixHandlerError:                          "The mix of handlers in the partitions is not allowed in this version of MySQL",
	ErPartitionNotDefinedError:                 "For the partitioned engine it is necessary to define all %-.64s",
	ErTooManyPartitionsError:                   "Too many partitions (including subpartitions) were defined",
	ErSubpartitionError:                        "It is only possible to mix RANGE/LIST partitioning with HASH/KEY partitioning for subpartitioning",
	ErCantCreateHandlerFile:                    "Failed to create specific handler file",
	ErBlobFieldInPartFuncError:                 "A BLOB field is not allowed in partition function",
	ErUniqueKeyNeedAllFieldsInPf:               "A %-.192s must include all columns in the table's partitioning function",
	ErNoPartsError:                             "Number of %-.64s = 0 is not an allowed value",
	ErPartitionMgmtOnNonpartitioned:            "Partition management on a not partitioned table is not possible",
	ErForeignKeyOnPartitioned:                  "Foreign key clause is not yet supported in conjunction with partitioning",
	ErDropPartitionNonExistent:                 "Error in list of partitions to %-.64s",
	ErDropLastPartition:                        "Cannot remove all partitions, use DROP TABLE instead",
	ErCoalesceOnlyOnHashPartition:              "COALESCE PARTITION can only be used on HASH/KEY partitions",
	ErReorgHashOnlyOnSameNo:                    "REORGANIZE PARTITION can only be used to reorganize partitions not to change their numbers",
	ErReorgNoParamError:                        "REORGANIZE PARTITION without parameters can only be used on auto-partitioned tables using HASH PARTITIONs",
	ErOnlyOnRangeListPartition:                 "%-.64s PARTITION can only be used on RANGE/LIST partitions",
	ErAddPartitionSubpartError:                 "Trying to Add partition(s) with wrong number of subpartitions",
	ErAddPartitionNoNewPartition:               "At least one partition must be added",
	ErCoalescePartitionNoPartition:             "At least one partition must be coalesced",
	ErReorgPartitionNotExist:                   "More partitions to reorganize than there are partitions",
	ErSameNamePartition:                        "Duplicate partition name %-.192s",
	ErNoBinlogError:                            "It is not allowed to shut off binlog on this command",
	ErConsecutiveReorgPartitions:               "When reorganizing a set of partitions they must be in consecutive order",
	ErReorgOutsideRange:                        "Reorganize of range partitions cannot change total ranges except for last partition where it can extend the range",
	ErPartitionFunctionFailure:                 "Partition function not supported in this version for this handler",
	ErPartStateError:                           "Partition state cannot be defined from CREATE/ALTER TABLE",
	ErLimitedPartRange:                         "The %-.64s handler only supports 32 bit integers in VALUES",
	ErPluginIsNotLoaded:                        "Plugin '%-.192s' is not loaded",
	ErWrongValue:                               "Incorrect %-.32s value: '%-.128s'",
	ErNoPartitionForGivenValue:                 "Table has no partition for value %-.64s",
	ErFilegroupOptionOnlyOnce:                  "It is not allowed to specify %s more than once",
	ErCreateFilegroupFailed:                    "Failed to create %s",
	ErDropFilegroupFailed:                      "Failed to drop %s",
	ErTablespaceAutoExtendError:                "The handler doesn't support autoextend of tablespaces",
	ErWrongSizeNumber:                          "A size parameter was incorrectly specified, either number or on the form 10M",
	ErSizeOverflowError:                        "The size number was correct but we don't allow the digit part to be more than 2 billion",
	ErAlterFilegroupFailed:                     "Failed to alter: %s",
	ErBinlogRowLoggingFailed:                   "Writing one row to the row-based binary log failed",
	ErBinlogRowWrongTableDef:                   "Table definition on master and slave does not match: %s",
	ErBinlogRowRbrToSbr:                        "Slave running with --log-slave-updates must use row-based binary logging to be able to replicate row-based binary log events",
	ErEventAlreadyExists:                       "Event '%-.192s' already exists",
	ErEventStoreFailed:                         "Failed to store event %s. Error code %d from storage engine.",
	ErEventDoesNotExist:                        "Unknown event '%-.192s'",
	ErEventCantAlter:                           "Failed to alter event '%-.192s'",
	ErEventDropFailed:                          "Failed to drop %s",
	ErEventIntervalNotPositiveOrTooBig:         "INTERVAL is either not positive or too big",
	ErEventEndsBeforeStarts:                    "ENDS is either invalid or before STARTS",
	ErEventExecTimeInThePast:                   "Event execution time is in the past. Event has been disabled",
	ErEventOpenTableFailed:                     "Failed to open mysql.event",
	ErEventNeitherMExprNorMAt:                  "No datetime expression provided",
	ErObsoleteColCountDoesntMatchCorrupted:     "Column count of mysql.%s is wrong. Expected %d, found %d. The table is probably corrupted",
	ErObsoleteCannotLoadFromTable:              "Cannot load from mysql.%s. The table is probably corrupted",
	ErEventCannotDelete:                        "Failed to delete the event from mysql.event",
	ErEventCompileError:                        "Error during compilation of event's body",
	ErEventSameName:                            "Same old and new event name",
	ErEventDataTooLong:                         "Data for column '%s' too long",
	ErDropIndexFk:                              "Cannot drop index '%-.192s': needed in a foreign key constraint",
	ErWarnDeprecatedSyntaxWithVer:              "The syntax '%s' is deprecated and will be removed in MySQL %s. Please use %s instead",
	ErCantWriteLockLogTable:                    "You can't write-lock a log table. Only read access is possible",
	ErCantLockLogTable:                         "You can't use locks with log tables.",
	ErForeignDuplicateKeyOldUnused:             "Upholding foreign key constraints for table '%.192s', entry '%-.192s', key %d would lead to a duplicate entry",
	ErColCountDoesntMatchPleaseUpdate:          "Column count of mysql.%s is wrong. Expected %d, found %d. Created with MySQL %d, now running %d. Please use mysql_upgrade to fix this error.",
	ErTempTablePreventsSwitchOutOfRbr:          "Cannot switch out of the row-based binary log format when the session has open temporary tables",
	ErStoredFunctionPreventsSwitchBinlogFormat: "Cannot change the binary logging format inside a stored function or trigger",
	ErNdbCantSwitchBinlogFormat:                "The NDB cluster engine does not support changing the binlog format on the fly yet",
	ErPartitionNoTemporary:                     "Cannot create temporary table with partitions",
	ErPartitionConstDomainError:                "Partition constant is out of partition function domain",
	ErPartitionFunctionIsNotAllowed:            "This partition function is not allowed",
	ErDdlLogError:                              "Error in DDL log",
	ErNullInValuesLessThan:                     "Not allowed to use NULL value in VALUES LESS THAN",
	ErWrongPartitionName:                       "Incorrect partition name",
	ErCantChangeTxCharacteristics:              "Transaction characteristics can't be changed while a transaction is in progress",
	ErDupEntryAutoincrementCase:                "ALTER TABLE causes auto_increment resequencing, resulting in duplicate entry '%-.192s' for key '%-.192s'",
	ErEventModifyQueueError:                    "Internal scheduler error %d",
	ErEventSetVarError:                         "Error during starting/stopping of the scheduler. Error code %u",
	ErPartitionMergeError:                      "Engine cannot be used in partitioned tables",
	ErCantActivateLog:                          "Cannot activate '%-.64s' log",
	ErRbrNotAvailable:                          "The server was not built with row-based replication",
	ErBase64DecodeError:                        "Decoding of base64 string failed",
	ErEventRecursionForbidden:                  "Recursion of EVENT DDL statements is forbidden when body is present",
	ErEventsDbError:                            "Cannot proceed because system tables used by Event Scheduler were found damaged at server start",
	ErOnlyIntegersAllowed:                      "Only integers allowed as number here",
	ErUnsuportedLogEngine:                      "This storage engine cannot be used for log tables\"",
	ErBadLogStatement:                          "You cannot '%s' a log table if logging is enabled",
	ErCantRenameLogTable:                       "Cannot rename '%s'. When logging enabled, rename to/from log table must rename two tables: the log table to an archive table and another table back to '%s'",
	ErWrongParamcountToNativeFct:               "Incorrect parameter count in the call to native function '%-.192s'",
	ErWrongParametersToNativeFct:               "Incorrect parameters in the call to native function '%-.192s'",
	ErWrongParametersToStoredFct:               "Incorrect parameters in the call to stored function '%-.192s'",
	ErNativeFctNameCollision:                   "This function '%-.192s' has the same name as a native function",
	ErDupEntryWithKeyName:                      "Duplicate entry '%-.64s' for key '%-.192s'",
	ErBinlogPurgeEmfile:                        "Too many files opened, please execute the command again",
	ErEventCannotCreateInThePast:               "Event execution time is in the past and ON COMPLETION NOT PRESERVE is set. The event was dropped immediately after creation.",
	ErEventCannotAlterInThePast:                "Event execution time is in the past and ON COMPLETION NOT PRESERVE is set. The event was not changed. Specify a time in the future.",
	ErSlaveIncident:                            "The incident %s occured on the master. Message: %-.64s",
	ErNoPartitionForGivenValueSilent:           "Table has no partition for some existing values",
	ErBinlogUnsafeStatement:                    "Unsafe statement written to the binary log using statement format since BINLOG_FORMAT = STATEMENT. %s",
	ErSlaveFatalError:                          "Fatal error: %s",
	ErSlaveRelayLogReadFailure:                 "Relay log read failure: %s",
	ErSlaveRelayLogWriteFailure:                "Relay log write failure: %s",
	ErSlaveCreateEventFailure:                  "Failed to create %s",
	ErSlaveMasterComFailure:                    "Master command %s failed: %s",
	ErBinlogLoggingImpossible:                  "Binary logging not possible. Message: %s",
	ErViewNoCreationCtx:                        "View `%-.64s`.`%-.64s` has no creation context",
	ErViewInvalidCreationCtx:                   "Creation context of view `%-.64s`.`%-.64s' is invalid",
	ErSrInvalidCreationCtx:                     "Creation context of stored routine `%-.64s`.`%-.64s` is invalid",
	ErTrgCorruptedFile:                         "Corrupted TRG file for table `%-.64s`.`%-.64s`",
	ErTrgNoCreationCtx:                         "Triggers for table `%-.64s`.`%-.64s` have no creation context",
	ErTrgInvalidCreationCtx:                    "Trigger creation context of table `%-.64s`.`%-.64s` is invalid",
	ErEventInvalidCreationCtx:                  "Creation context of event `%-.64s`.`%-.64s` is invalid",
	ErTrgCantOpenTable:                         "Cannot open table for trigger `%-.64s`.`%-.64s`",
	ErCantCreateSroutine:                       "Cannot create stored routine `%-.64s`. Check warnings",
	ErNeverUsed:                                "Ambiguous slave modes combination. %s",
	ErNoFormatDescriptionEventBeforeBinlogStatement:         "The BINLOG statement of type `%s` was not preceded by a format description BINLOG statement.",
	ErSlaveCorruptEvent:                                     "Corrupted replication event was detected",
	ErLoadDataInvalidColumn:                                 "Invalid column reference (%-.64s) in LOAD DATA",
	ErLogPurgeNoFile:                                        "Being purged log %s was not found",
	ErXaRbtimeout:                                           "XA_RBTIMEOUT: Transaction branch was rolled back: took too long",
	ErXaRbdeadlock:                                          "XA_RBDEADLOCK: Transaction branch was rolled back: deadlock was detected",
	ErNeedReprepare:                                         "Prepared statement needs to be re-prepared",
	ErDelayedNotSupported:                                   "DELAYED option not supported for table '%-.192s'",
	WarnNoMasterInfo:                                        "The master info structure does not exist",
	WarnOptionIgnored:                                       "<%-.64s> option ignored",
	WarnPluginDeleteBuiltin:                                 "Built-in plugins cannot be deleted",
	WarnPluginBusy:                                          "Plugin is busy and will be uninstalled on shutdown",
	ErVariableIsReadonly:                                    "%s variable '%s' is read-only. Use SET %s to assign the value",
	ErWarnEngineTransactionRollback:                         "Storage engine %s does not support rollback for this statement. Transaction rolled back and must be restarted",
	ErSlaveHeartbeatFailure:                                 "Unexpected master's heartbeat data: %s",
	ErSlaveHeartbeatValueOutOfRange:                         "The requested value for the heartbeat period is either negative or exceeds the maximum allowed (%s seconds).",
	ErNdbReplicationSchemaError:                             "Bad schema for mysql.ndb_replication table. Message: %-.64s",
	ErConflictFnParseError:                                  "Error in parsing conflict function. Message: %-.64s",
	ErExceptionsWriteError:                                  "Write to exceptions table failed. Message: %-.128s\"",
	ErTooLongTableComment:                                   "Comment for table '%-.64s' is too long (max = %lu)",
	ErTooLongFieldComment:                                   "Comment for field '%-.64s' is too long (max = %lu)",
	ErFuncInexistentNameCollision:                           "FUNCTION %s does not exist. Check the 'Function Name Parsing and Resolution' section in the Reference Manual",
	ErDatabaseName:                                          "Database",
	ErTableName:                                             "Table",
	ErPartitionName:                                         "Partition",
	ErSubpartitionName:                                      "Subpartition",
	ErTemporaryName:                                         "Temporary",
	ErRenamedName:                                           "Renamed",
	ErTooManyConcurrentTrxs:                                 "Too many active concurrent transactions",
	WarnNonAsciiSeparatorNotImplemented:                     "Non-ASCII separator arguments are not fully supported",
	ErDebugSyncTimeout:                                      "debug sync point wait timed out",
	ErDebugSyncHitLimit:                                     "debug sync point hit limit reached",
	ErDupSignalSet:                                          "Duplicate condition information item '%s'",
	ErSignalWarn:                                            "Unhandled user-defined warning condition",
	ErSignalNotFound:                                        "Unhandled user-defined not found condition",
	ErSignalException:                                       "Unhandled user-defined exception condition",
	ErResignalWithoutActiveHandler:                          "RESIGNAL when handler not active",
	ErSignalBadConditionType:                                "SIGNAL/RESIGNAL can only use a CONDITION defined with SQLSTATE",
	WarnCondItemTruncated:                                   "Data truncated for condition item '%s'",
	ErCondItemTooLong:                                       "Data too long for condition item '%s'",
	ErUnknownLocale:                                         "Unknown locale: '%-.64s'",
	ErSlaveIgnoreServerIds:                                  "The requested server id %d clashes with the slave startup option --replicate-same-server-id",
	ErQueryCacheDisabled:                                    "Query cache is disabled; restart the server with query_cache_type=1 to enable it",
	ErSameNamePartitionField:                                "Duplicate partition field name '%-.192s'",
	ErPartitionColumnListError:                              "Inconsistency in usage of column lists for partitioning",
	ErWrongTypeColumnValueError:                             "Partition column values of incorrect type",
	ErTooManyPartitionFuncFieldsError:                       "Too many fields in '%-.192s'",
	ErMaxvalueInValuesIn:                                    "Cannot use MAXVALUE as value in VALUES IN",
	ErTooManyValuesError:                                    "Cannot have more than one value for this type of %-.64s partitioning",
	ErRowSinglePartitionFieldError:                          "Row expressions in VALUES IN only allowed for multi-field column partitioning",
	ErFieldTypeNotAllowedAsPartitionField:                   "Field '%-.192s' is of a not allowed type for this type of partitioning",
	ErPartitionFieldsTooLong:                                "The total length of the partitioning fields is too large",
	ErBinlogRowEngineAndStmtEngine:                          "Cannot execute statement: impossible to write to binary log since both row-incapable engines and statement-incapable engines are involved.",
	ErBinlogRowModeAndStmtEngine:                            "Cannot execute statement: impossible to write to binary log since BINLOG_FORMAT = ROW and at least one table uses a storage engine limited to statement-based logging.",
	ErBinlogUnsafeAndStmtEngine:                             "Cannot execute statement: impossible to write to binary log since statement is unsafe, storage engine is limited to statement-based logging, and BINLOG_FORMAT = MIXED. %s",
	ErBinlogRowInjectionAndStmtEngine:                       "Cannot execute statement: impossible to write to binary log since statement is in row format and at least one table uses a storage engine limited to statement-based logging.",
	ErBinlogStmtModeAndRowEngine:                            "Cannot execute statement: impossible to write to binary log since BINLOG_FORMAT = STATEMENT and at least one table uses a storage engine limited to row-based logging.%s",
	ErBinlogRowInjectionAndStmtMode:                         "Cannot execute statement: impossible to write to binary log since statement is in row format and BINLOG_FORMAT = STATEMENT.",
	ErBinlogMultipleEnginesAndSelfLoggingEngine:             "Cannot execute statement: impossible to write to binary log since more than one engine is involved and at least one engine is self-logging.",
	ErBinlogUnsafeLimit:                                     "The statement is unsafe because it uses a LIMIT clause. This is unsafe because the set of rows included cannot be predicted.",
	ErBinlogUnsafeInsertDelayed:                             "The statement is unsafe because it uses INSERT DELAYED. This is unsafe because the times when rows are inserted cannot be predicted.",
	ErBinlogUnsafeSystemTable:                               "The statement is unsafe because it uses the general log, slow query log, or performance_schema table(s). This is unsafe because system tables may differ on slaves.",
	ErBinlogUnsafeAutoincColumns:                            "Statement is unsafe because it invokes a trigger or a stored function that inserts into an AUTO_INCREMENT column. Inserted values cannot be logged correctly.",
	ErBinlogUnsafeUdf:                                       "Statement is unsafe because it uses a UDF which may not return the same value on the slave.",
	ErBinlogUnsafeSystemVariable:                            "Statement is unsafe because it uses a system variable that may have a different value on the slave.",
	ErBinlogUnsafeSystemFunction:                            "Statement is unsafe because it uses a system function that may return a different value on the slave.",
	ErBinlogUnsafeNontransAfterTrans:                        "Statement is unsafe because it accesses a non-transactional table after accessing a transactional table within the same transaction.",
	ErMessageAndStatement:                                   "%s Statement: %s",
	ErSlaveConversionFailed:                                 "Column %d of table '%-.192s.%-.192s' cannot be converted from type '%-.32s' to type '%-.32s'",
	ErSlaveCantCreateConversion:                             "Can't create conversion table for table '%-.192s.%-.192s'",
	ErInsideTransactionPreventsSwitchBinlogFormat:           "Cannot modify @@session.binlog_format inside a transaction",
	ErPathLength:                                            "The path specified for %.64s is too long.",
	ErWarnDeprecatedSyntaxNoReplacement:                     "'%s' is deprecated and will be removed in a future release.",
	ErWrongNativeTableStructure:                             "Native table '%-.64s'.'%-.64s' has the wrong structure",
	ErWrongPerfschemaUsage:                                  "Invalid performance_schema usage.",
	ErWarnISSkippedTable:                                    "Table '%s'.'%s' was skipped since its definition is being modified by concurrent DDL statement",
	ErInsideTransactionPreventsSwitchBinlogDirect:           "Cannot modify @@session.binlog_direct_non_transactional_updates inside a transaction",
	ErStoredFunctionPreventsSwitchBinlogDirect:              "Cannot change the binlog direct flag inside a stored function or trigger",
	ErSpatialMustHaveGeomCol:                                "A SPATIAL index may only contain a geometrical type column",
	ErTooLongIndexComment:                                   "Comment for index '%-.64s' is too long (max = %lu)",
	ErLockAborted:                                           "Wait on a lock was aborted due to a pending exclusive lock",
	ErDataOutOfRange:                                        "%s value is out of range in '%s'",
	ErWrongSpvarTypeInLimit:                                 "A variable of a non-integer based type in LIMIT clause",
	ErBinlogUnsafeMultipleEnginesAndSelfLoggingEngine:       "Mixing self-logging and non-self-logging engines in a statement is unsafe.",
	ErBinlogUnsafeMixedStatement:                            "Statement accesses nontransactional table as well as transactional or temporary table, and writes to any of them.",
	ErInsideTransactionPreventsSwitchSqlLogBin:              "Cannot modify @@session.sql_log_bin inside a transaction",
	ErStoredFunctionPreventsSwitchSqlLogBin:                 "Cannot change the sql_log_bin inside a stored function or trigger",
	ErFailedReadFromParFile:                                 "Failed to read from the .par file",
	ErValuesIsNotIntTypeError:                               "VALUES value for partition '%-.64s' must have type INT",
	ErAccessDeniedNoPasswordError:                           "Access denied for user '%-.48s'@'%-.64s'",
	ErSetPasswordAuthPlugin:                                 "SET PASSWORD has no significance for users authenticating via plugins",
	ErGrantPluginUserExists:                                 "GRANT with IDENTIFIED WITH is illegal because the user %-.*s already exists",
	ErTruncateIllegalFk:                                     "Cannot truncate a table referenced in a foreign key constraint (%.192s)",
	ErPluginIsPermanent:                                     "Plugin '%s' is force_plus_permanent and can not be unloaded",
	ErSlaveHeartbeatValueOutOfRangeMin:                      "The requested value for the heartbeat period is less than 1 millisecond. The value is reset to 0, meaning that heartbeating will effectively be disabled.",
	ErSlaveHeartbeatValueOutOfRangeMax:                      "The requested value for the heartbeat period exceeds the value of `slave_net_timeout' seconds. A sensible value for the period should be less than the timeout.",
	ErStmtCacheFull:                                         "Multi-row statements required more than 'max_binlog_stmt_cache_size' bytes of storage; increase this mysqld variable and try again",
	ErMultiUpdateKeyConflict:                                "Primary key/partition key update is not allowed since the table is updated both as '%-.192s' and '%-.192s'.",
	ErTableNeedsRebuild:                                     "Table rebuild required. Please do \"ALTER TABLE `%-.32s` FORCE\" or dump/reload to fix it!",
	WarnOptionBelowLimit:                                    "The value of '%s' should be no less than the value of '%s'",
	ErIndexColumnTooLong:                                    "Index column size too large. The maximum column size is %lu bytes.",
	ErErrorInTriggerBody:                                    "Trigger '%-.64s' has an error in its body: '%-.256s'",
	ErErrorInUnknownTriggerBody:                             "Unknown trigger has an error in its body: '%-.256s'",
	ErIndexCorrupt:                                          "Index %s is corrupted",
	ErUndoRecordTooBig:                                      "Undo log record is too big.",
	ErBinlogUnsafeInsertIgnoreSelect:                        "INSERT IGNORE... SELECT is unsafe because the order in which rows are retrieved by the SELECT determines which (if any) rows are ignored. This order cannot be predicted and may differ on master and the slave.",
	ErBinlogUnsafeInsertSelectUpdate:                        "INSERT... SELECT... ON DUPLICATE KEY UPDATE is unsafe because the order in which rows are retrieved by the SELECT determines which (if any) rows are updated. This order cannot be predicted and may differ on master and the slave.",
	ErBinlogUnsafeReplaceSelect:                             "REPLACE... SELECT is unsafe because the order in which rows are retrieved by the SELECT determines which (if any) rows are replaced. This order cannot be predicted and may differ on master and the slave.",
	ErBinlogUnsafeCreateIgnoreSelect:                        "CREATE... IGNORE SELECT is unsafe because the order in which rows are retrieved by the SELECT determines which (if any) rows are ignored. This order cannot be predicted and may differ on master and the slave.",
	ErBinlogUnsafeCreateReplaceSelect:                       "CREATE... REPLACE SELECT is unsafe because the order in which rows are retrieved by the SELECT determines which (if any) rows are replaced. This order cannot be predicted and may differ on master and the slave.",
	ErBinlogUnsafeUpdateIgnore:                              "UPDATE IGNORE is unsafe because the order in which rows are updated determines which (if any) rows are ignored. This order cannot be predicted and may differ on master and the slave.",
	ErPluginNoUninstall:                                     "Plugin '%s' is marked as not dynamically uninstallable. You have to stop the server to uninstall it.",
	ErPluginNoInstall:                                       "Plugin '%s' is marked as not dynamically installable. You have to stop the server to install it.",
	ErBinlogUnsafeWriteAutoincSelect:                        "Statements writing to a table with an auto-increment column after selecting from another table are unsafe because the order in which rows are retrieved determines what (if any) rows will be written. This order cannot be predicted and may differ on master and the slave.",
	ErBinlogUnsafeCreateSelectAutoinc:                       "CREATE TABLE... SELECT...  on a table with an auto-increment column is unsafe because the order in which rows are retrieved by the SELECT determines which (if any) rows are inserted. This order cannot be predicted and may differ on master and the slave.",
	ErBinlogUnsafeInsertTwoKeys:                             "INSERT... ON DUPLICATE KEY UPDATE  on a table with more than one UNIQUE KEY is unsafe",
	ErTableInFkCheck:                                        "Table is being used in foreign key check.",
	ErUnsupportedEngine:                                     "Storage engine '%s' does not support system tables. [%s.%s]",
	ErBinlogUnsafeAutoincNotFirst:                           "INSERT into autoincrement field which is not the first part in the composed primary key is unsafe.",
	ErCannotLoadFromTableV2:                                 "Cannot load from %s.%s. The table is probably corrupted",
	ErMasterDelayValueOutOfRange:                            "The requested value %u for the master delay exceeds the maximum %u",
	ErOnlyFdAndRbrEventsAllowedInBinlogStatement:            "Only Format_description_log_event and row events are allowed in BINLOG statements (but %s was provided)",
	ErPartitionExchangeDifferentOption:                      "Non matching attribute '%-.64s' between partition and table",
	ErPartitionExchangePartTable:                            "Table to exchange with partition is partitioned: '%-.64s'",
	ErPartitionExchangeTempTable:                            "Table to exchange with partition is temporary: '%-.64s'",
	ErPartitionInsteadOfSubpartition:                        "Subpartitioned table, use subpartition instead of partition",
	ErUnknownPartition:                                      "Unknown partition '%-.64s' in table '%-.64s'",
	ErTablesDifferentMetadata:                               "Tables have different definitions",
	ErRowDoesNotMatchPartition:                              "Found a row that does not match the partition",
	ErBinlogCacheSizeGreaterThanMax:                         "Option binlog_cache_size (%lu) is greater than max_binlog_cache_size (%lu); setting binlog_cache_size equal to max_binlog_cache_size.",
	ErWarnIndexNotApplicable:                                "Cannot use %-.64s access on index '%-.64s' due to type or collation conversion on field '%-.64s'",
	ErPartitionExchangeForeignKey:                           "Table to exchange with partition has foreign key references: '%-.64s'",
	ErNoSuchKeyValue:                                        "Key value '%-.192s' was not found in table '%-.192s.%-.192s'",
	ErRplInfoDataTooLong:                                    "Data for column '%s' too long",
	ErNetworkReadEventChecksumFailure:                       "Replication event checksum verification failed while reading from network.",
	ErBinlogReadEventChecksumFailure:                        "Replication event checksum verification failed while reading from a log file.",
	ErBinlogStmtCacheSizeGreaterThanMax:                     "Option binlog_stmt_cache_size (%lu) is greater than max_binlog_stmt_cache_size (%lu); setting binlog_stmt_cache_size equal to max_binlog_stmt_cache_size.",
	ErCantUpdateTableInCreateTableSelect:                    "Can't update table '%-.192s' while '%-.192s' is being created.",
	ErPartitionClauseOnNonpartitioned:                       "PARTITION () clause on non partitioned table",
	ErRowDoesNotMatchGivenPartitionSet:                      "Found a row not matching the given partition set",
	ErNoSuchPartition_Unused:                                "partition '%-.64s' doesn't exist",
	ErChangeRplInfoRepositoryFailure:                        "Failure while changing the type of replication repository: %s.",
	ErWarningNotCompleteRollbackWithCreatedTempTable:        "The creation of some temporary tables could not be rolled back.",
	ErWarningNotCompleteRollbackWithDroppedTempTable:        "Some temporary tables were dropped, but these operations could not be rolled back.",
	ErMtsFeatureIsNotSupported:                              "%s is not supported in multi-threaded slave mode. %s",
	ErMtsUpdatedDbsGreaterMax:                               "The number of modified databases exceeds the maximum %d; the database names will not be included in the replication event metadata.",
	ErMtsCantParallel:                                       "Cannot execute the current event group in the parallel mode. Encountered event %s, relay-log name %s, position %s which prevents execution of this event group in parallel mode. Reason: %s.",
	ErMtsInconsistentData:                                   "%s",
	ErFulltextNotSupportedWithPartitioning:                  "FULLTEXT index is not supported for partitioned tables.",
	ErDaInvalidConditionNumber:                              "Invalid condition number",
	ErInsecurePlainText:                                     "Sending passwords in plain text without SSL/TLS is extremely insecure.",
	ErInsecureChangeMaster:                                  "Storing MySQL user name or password information in the master.info repository is not secure and is therefore not recommended. Please see the MySQL Manual for more about this issue and possible alternatives.",
	ErForeignDuplicateKeyWithChildInfo:                      "Foreign key constraint for table '%.192s', record '%-.192s' would lead to a duplicate entry in table '%.192s', key '%.192s'",
	ErForeignDuplicateKeyWithoutChildInfo:                   "Foreign key constraint for table '%.192s', record '%-.192s' would lead to a duplicate entry in a child table",
	ErSqlthreadWithSecureSlave:                              "Setting authentication options is not possible when only the Slave SQL Thread is being started.",
	ErTableHasNoFt:                                          "The table does not have FULLTEXT index to support this query",
	ErVariableNotSettableInSfOrTrigger:                      "The system variable %.200s cannot be set in stored functions or triggers.",
	ErVariableNotSettableInTransaction:                      "The system variable %.200s cannot be set when there is an ongoing transaction.",
	ErGtidNextIsNotInGtidNextList:                           "The system variable @@SESSION.GTID_NEXT has the value %.200s, which is not listed in @@SESSION.GTID_NEXT_LIST.",
	ErCantChangeGtidNextInTransactionWhenGtidNextListIsNull: "When @@SESSION.GTID_NEXT_LIST == NULL, the system variable @@SESSION.GTID_NEXT cannot change inside a transaction.",
	ErSetStatementCannotInvokeFunction:                      "The statement 'SET %.200s' cannot invoke a stored function.",
	ErGtidNextCantBeAutomaticIfGtidNextListIsNonNull:        "The system variable @@SESSION.GTID_NEXT cannot be 'AUTOMATIC' when @@SESSION.GTID_NEXT_LIST is non-NULL.",
	ErSkippingLoggedTransaction:                             "Skipping transaction %.200s because it has already been executed and logged.",
	ErMalformedGtidSetSpecification:                         "Malformed GTID set specification '%.200s'.",
	ErMalformedGtidSetEncoding:                              "Malformed GTID set encoding.",
	ErMalformedGtidSpecification:                            "Malformed GTID specification '%.200s'.",
	ErGnoExhausted:                                          "Impossible to generate Global Transaction Identifier: the integer component reached the maximal value. Restart the server with a new server_uuid.",
	ErBadSlaveAutoPosition:                                  "Parameters MASTER_LOG_FILE, MASTER_LOG_POS, RELAY_LOG_FILE and RELAY_LOG_POS cannot be set when MASTER_AUTO_POSITION is active.",
	ErAutoPositionRequiresGtidModeOn:                        "CHANGE MASTER TO MASTER_AUTO_POSITION = 1 can only be executed when @@GLOBAL.GTID_MODE = ON.",
	ErCantDoImplicitCommitInTrxWhenGtidNextIsSet:            "Cannot execute statements with implicit commit inside a transaction when @@SESSION.GTID_NEXT != AUTOMATIC or @@SESSION.GTID_NEXT_LIST != NULL.",
	ErGtidMode2Or3RequiresEnforceGtidConsistencyOn:          "@@GLOBAL.GTID_MODE = ON or UPGRADE_STEP_2 requires @@GLOBAL.ENFORCE_GTID_CONSISTENCY = 1.",
	ErGtidModeRequiresBinlog:                                "@@GLOBAL.GTID_MODE = ON or UPGRADE_STEP_1 or UPGRADE_STEP_2 requires --log-bin and --log-slave-updates.",
	ErCantSetGtidNextToGtidWhenGtidModeIsOff:                "@@SESSION.GTID_NEXT cannot be set to UUID:NUMBER when @@GLOBAL.GTID_MODE = OFF.",
	ErCantSetGtidNextToAnonymousWhenGtidModeIsOn:            "@@SESSION.GTID_NEXT cannot be set to ANONYMOUS when @@GLOBAL.GTID_MODE = ON.",
	ErCantSetGtidNextListToNonNullWhenGtidModeIsOff:         "@@SESSION.GTID_NEXT_LIST cannot be set to a non-NULL value when @@GLOBAL.GTID_MODE = OFF.",
	ErFoundGtidEventWhenGtidModeIsOff:                       "Found a Gtid_log_event or Previous_gtids_log_event when @@GLOBAL.GTID_MODE = OFF.",
	ErGtidUnsafeNonTransactionalTable:                       "When @@GLOBAL.ENFORCE_GTID_CONSISTENCY = 1, updates to non-transactional tables can only be done in either autocommitted statements or single-statement transactions, and never in the same statement as updates to transactional tables.",
	ErGtidUnsafeCreateSelect:                                "CREATE TABLE ... SELECT is forbidden when @@GLOBAL.ENFORCE_GTID_CONSISTENCY = 1.",
	ErGtidUnsafeCreateDropTemporaryTableInTransaction:       "When @@GLOBAL.ENFORCE_GTID_CONSISTENCY = 1, the statements CREATE TEMPORARY TABLE and DROP TEMPORARY TABLE can be executed in a non-transactional context only, and require that AUTOCOMMIT = 1.",
	ErGtidModeCanOnlyChangeOneStepAtATime:                   "The value of @@GLOBAL.GTID_MODE can only change one step at a time: OFF <-> UPGRADE_STEP_1 <-> UPGRADE_STEP_2 <-> ON. Also note that this value must be stepped up or down simultaneously on all servers; see the Manual for instructions.",
	ErMasterHasPurgedRequiredGtids:                          "The slave is connecting using CHANGE MASTER TO MASTER_AUTO_POSITION = 1, but the master has purged binary logs containing GTIDs that the slave requires.",
	ErCantSetGtidNextWhenOwningGtid:                         "@@SESSION.GTID_NEXT cannot be changed by a client that owns a GTID. The client owns %s. Ownership is released on COMMIT or ROLLBACK.",
	ErUnknownExplainFormat:                                  "Unknown EXPLAIN format name: '%s'",
	ErCantExecuteInReadOnlyTransaction:                      "Cannot execute statement in a READ ONLY transaction.",
	ErTooLongTablePartitionComment:                          "Comment for table partition '%-.64s' is too long (max = %lu)",
	ErSlaveConfiguration:                                    "Slave is not configured or failed to initialize properly. You must at least set --server-id to enable either a master or a slave. Additional error messages can be found in the MySQL error log.",
	ErInnodbFtLimit:                                         "InnoDB presently supports one FULLTEXT index creation at a time",
	ErInnodbNoFtTempTable:                                   "Cannot create FULLTEXT index on temporary InnoDB table",
	ErInnodbFtWrongDocidColumn:                              "Column '%-.192s' is of wrong type for an InnoDB FULLTEXT index",
	ErInnodbFtWrongDocidIndex:                               "Index '%-.192s' is of wrong type for an InnoDB FULLTEXT index",
	ErInnodbOnlineLogTooBig:                                 "Creating index '%-.192s' required more than 'innodb_online_alter_log_max_size' bytes of modification log. Please try again.",
	ErUnknownAlterAlgorithm:                                 "Unknown ALGORITHM '%s'",
	ErUnknownAlterLock:                                      "Unknown LOCK type '%s'",
	ErMtsChangeMasterCantRunWithGaps:                        "CHANGE MASTER cannot be executed when the slave was stopped with an error or killed in MTS mode. Consider using RESET SLAVE or START SLAVE UNTIL.",
	ErMtsRecoveryFailure:                                    "Cannot recover after SLAVE errored out in parallel execution mode. Additional error messages can be found in the MySQL error log.",
	ErMtsResetWorkers:                                       "Cannot clean up worker info tables. Additional error messages can be found in the MySQL error log.",
	ErColCountDoesntMatchCorruptedV2:                        "Column count of %s.%s is wrong. Expected %d, found %d. The table is probably corrupted",
	ErSlaveSilentRetryTransaction:                           "Slave must silently retry current transaction",
	ErDiscardFkChecksRunning:                                "There is a foreign key check running on table '%-.192s'. Cannot discard the table.",
	ErTableSchemaMismatch:                                   "Schema mismatch (%s)",
	ErTableInSystemTablespace:                               "Table '%-.192s' in system tablespace",
	ErIoReadError:                                           "IO Read error: (%lu, %s) %s",
	ErIoWriteError:                                          "IO Write error: (%lu, %s) %s",
	ErTablespaceMissing:                                     "Tablespace is missing for table '%-.192s'",
	ErTablespaceExists:                                      "Tablespace for table '%-.192s' exists. Please DISCARD the tablespace before IMPORT.",
	ErTablespaceDiscarded:                                   "Tablespace has been discarded for table '%-.192s'",
	ErInternalError:                                         "Internal error: %s",
	ErInnodbImportError:                                     "ALTER TABLE '%-.192s' IMPORT TABLESPACE failed with error %lu : '%s'",
	ErInnodbIndexCorrupt:                                    "Index corrupt: %s",
	ErInvalidYearColumnLength:                               "YEAR(%lu) column type is deprecated. Creating YEAR(4) column instead.",
	ErNotValidPassword:                                      "Your password does not satisfy the current policy requirements",
	ErMustChangePassword:                                    "You must SET PASSWORD before executing this statement",
	ErFkNoIndexChild:                                        "Failed to add the foreign key constaint. Missing index for constraint '%s' in the foreign table '%s'",
	ErFkNoIndexParent:                                       "Failed to add the foreign key constaint. Missing index for constraint '%s' in the referenced table '%s'",
	ErFkFailAddSystem:                                       "Failed to add the foreign key constraint '%s' to system tables",
	ErFkCannotOpenParent:                                    "Failed to open the referenced table '%s'",
	ErFkIncorrectOption:                                     "Failed to add the foreign key constraint on table '%s'. Incorrect options in FOREIGN KEY constraint '%s'",
	ErFkDupName:                                             "Duplicate foreign key constraint name '%s'",
	ErPasswordFormat:                                        "The password hash doesn't have the expected format. Check if the correct password algorithm is being used with the PASSWORD() function.",
	ErFkColumnCannotDrop:                                    "Cannot drop column '%-.192s': needed in a foreign key constraint '%-.192s'",
	ErFkColumnCannotDropChild:                               "Cannot drop column '%-.192s': needed in a foreign key constraint '%-.192s' of table '%-.192s'",
	ErFkColumnNotNull:                                       "Column '%-.192s' cannot be NOT NULL: needed in a foreign key constraint '%-.192s' SET NULL",
	ErDupIndex:                                              "Duplicate index '%-.64s' defined on the table '%-.64s.%-.64s'. This is deprecated and will be disallowed in a future release.",
	ErFkColumnCannotChange:                                  "Cannot change column '%-.192s': used in a foreign key constraint '%-.192s'",
	ErFkColumnCannotChangeChild:                             "Cannot change column '%-.192s': used in a foreign key constraint '%-.192s' of table '%-.192s'",
	ErFkCannotDeleteParent:                                  "Cannot delete rows from table which is parent in a foreign key constraint '%-.192s' of table '%-.192s'",
	ErMalformedPacket:                                       "Malformed communication packet.",
	ErReadOnlyMode:                                          "Running in read-only mode",
	ErGtidNextTypeUndefinedGroup:                            "When @@SESSION.GTID_NEXT is set to a GTID, you must explicitly set it again after a COMMIT or ROLLBACK. If you see this error message in the slave SQL thread, it means that a table in the current transaction is transactional on the master and non-transactional on the slave. In a client connection, it means that you executed SET @@SESSION.GTID_NEXT before a transaction and forgot to set @@SESSION.GTID_NEXT to a different identifier or to 'AUTOMATIC' after COMMIT or ROLLBACK. Current @@SESSION.GTID_NEXT is '%s'.",
	ErVariableNotSettableInSp:                               "The system variable %.200s cannot be set in stored procedures.",
	ErCantSetGtidPurgedWhenGtidModeIsOff:                    "@@GLOBAL.GTID_PURGED can only be set when @@GLOBAL.GTID_MODE = ON.",
	ErCantSetGtidPurgedWhenGtidExecutedIsNotEmpty:           "@@GLOBAL.GTID_PURGED can only be set when @@GLOBAL.GTID_EXECUTED is empty.",
	ErCantSetGtidPurgedWhenOwnedGtidsIsNotEmpty:             "@@GLOBAL.GTID_PURGED can only be set when there are no ongoing transactions (not even in other clients).",
	ErGtidPurgedWasChanged:                                  "@@GLOBAL.GTID_PURGED was changed from '%s' to '%s'.",
	ErGtidExecutedWasChanged:                                "@@GLOBAL.GTID_EXECUTED was changed from '%s' to '%s'.",
	ErBinlogStmtModeAndNoReplTables:                         "Cannot execute statement: impossible to write to binary log since BINLOG_FORMAT = STATEMENT, and both replicated and non replicated tables are written to.",
	ErAlterOperationNotSupported:                            "%s is not supported for this operation. Try %s.",
	ErAlterOperationNotSupportedReason:                      "%s is not supported. Reason: %s. Try %s.",
	ErAlterOperationNotSupportedReasonCopy:                  "COPY algorithm requires a lock",
	ErAlterOperationNotSupportedReasonPartition:             "Partition specific operations do not yet support LOCK/ALGORITHM",
	ErAlterOperationNotSupportedReasonFkRename:              "Columns participating in a foreign key are renamed",
	ErAlterOperationNotSupportedReasonColumnType:            "Cannot change column type INPLACE",
	ErAlterOperationNotSupportedReasonFkCheck:               "Adding foreign keys needs foreign_key_checks=OFF",
	ErAlterOperationNotSupportedReasonIgnore:                "Creating unique indexes with IGNORE requires COPY algorithm to remove duplicate rows",
	ErAlterOperationNotSupportedReasonNopk:                  "Dropping a primary key is not allowed without also adding a new primary key",
	ErAlterOperationNotSupportedReasonAutoinc:               "Adding an auto-increment column requires a lock",
	ErAlterOperationNotSupportedReasonHiddenFts:             "Cannot replace hidden FTS_DOC_ID with a user-visible one",
	ErAlterOperationNotSupportedReasonChangeFts:             "Cannot drop or rename FTS_DOC_ID",
	ErAlterOperationNotSupportedReasonFts:                   "Fulltext index creation requires a lock",
	ErSqlSlaveSkipCounterNotSettableInGtidMode:              "sql_slave_skip_counter can not be set when the server is running with @@GLOBAL.GTID_MODE = ON. Instead, for each transaction that you want to skip, generate an empty transaction with the same GTID as the transaction",
	ErDupUnknownInIndex:                                     "Duplicate entry for key '%-.192s'",
	ErIdentCausesTooLongPath:                                "Long database name and identifier for object resulted in path length exceeding %d characters. Path: '%s'.",
	ErAlterOperationNotSupportedReasonNotNull:               "cannot silently convert NULL values, as required in this SQL_MODE",
	ErMustChangePasswordLogin:                               "Your password has expired. To log in you must change it using a client that supports expired passwords.",
	ErRowInWrongPartition:                                   "Found a row in wrong partition %s",
}
