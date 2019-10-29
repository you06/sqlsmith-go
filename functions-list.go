package sqlsmith


import "github.com/pingcap/parser/ast"

type functionClass struct {
	name string
	minArg int
	maxArg int
}

var funcs = map[string]*functionClass{
	// common functions
	ast.Coalesce: &functionClass{ast.Coalesce, 1, -1},
	ast.IsNull:   &functionClass{ast.IsNull, 1, 1},
	ast.Greatest: &functionClass{ast.Greatest, 2, -1},
	ast.Least:    &functionClass{ast.Least, 2, -1},
	ast.Interval: &functionClass{ast.Interval, 2, -1},

	// math functions
	ast.Abs:      &functionClass{ast.Abs, 1, 1},
	ast.Acos:     &functionClass{ast.Acos, 1, 1},
	ast.Asin:     &functionClass{ast.Asin, 1, 1},
	ast.Atan:     &functionClass{ast.Atan, 1, 2},
	ast.Atan2:    &functionClass{ast.Atan2, 2, 2},
	ast.Ceil:     &functionClass{ast.Ceil, 1, 1},
	ast.Ceiling:  &functionClass{ast.Ceiling, 1, 1},
	ast.Conv:     &functionClass{ast.Conv, 3, 3},
	ast.Cos:      &functionClass{ast.Cos, 1, 1},
	ast.Cot:      &functionClass{ast.Cot, 1, 1},
	ast.CRC32:    &functionClass{ast.CRC32, 1, 1},
	ast.Degrees:  &functionClass{ast.Degrees, 1, 1},
	ast.Exp:      &functionClass{ast.Exp, 1, 1},
	ast.Floor:    &functionClass{ast.Floor, 1, 1},
	ast.Ln:       &functionClass{ast.Ln, 1, 1},
	ast.Log:      &functionClass{ast.Log, 1, 2},
	ast.Log2:     &functionClass{ast.Log2, 1, 1},
	ast.Log10:    &functionClass{ast.Log10, 1, 1},
	ast.PI:       &functionClass{ast.PI, 0, 0},
	ast.Pow:      &functionClass{ast.Pow, 2, 2},
	ast.Power:    &functionClass{ast.Power, 2, 2},
	ast.Radians:  &functionClass{ast.Radians, 1, 1},
	ast.Rand:     &functionClass{ast.Rand, 0, 1},
	ast.Round:    &functionClass{ast.Round, 1, 2},
	ast.Sign:     &functionClass{ast.Sign, 1, 1},
	ast.Sin:      &functionClass{ast.Sin, 1, 1},
	ast.Sqrt:     &functionClass{ast.Sqrt, 1, 1},
	ast.Tan:      &functionClass{ast.Tan, 1, 1},
	ast.Truncate: &functionClass{ast.Truncate, 2, 2},

	// time functions
	ast.AddDate:          &functionClass{ast.AddDate, 3, 3},
	ast.DateAdd:          &functionClass{ast.DateAdd, 3, 3},
	ast.SubDate:          &functionClass{ast.SubDate, 3, 3},
	ast.DateSub:          &functionClass{ast.DateSub, 3, 3},
	ast.AddTime:          &functionClass{ast.AddTime, 2, 2},
	ast.ConvertTz:        &functionClass{ast.ConvertTz, 3, 3},
	ast.Curdate:          &functionClass{ast.Curdate, 0, 0},
	ast.CurrentDate:      &functionClass{ast.CurrentDate, 0, 0},
	ast.CurrentTime:      &functionClass{ast.CurrentTime, 0, 1},
	ast.CurrentTimestamp: &functionClass{ast.CurrentTimestamp, 0, 1},
	ast.Curtime:          &functionClass{ast.Curtime, 0, 1},
	ast.Date:             &functionClass{ast.Date, 1, 1},
	ast.DateLiteral:      &functionClass{ast.DateLiteral, 1, 1},
	ast.DateFormat:       &functionClass{ast.DateFormat, 2, 2},
	ast.DateDiff:         &functionClass{ast.DateDiff, 2, 2},
	ast.Day:              &functionClass{ast.Day, 1, 1},
	ast.DayName:          &functionClass{ast.DayName, 1, 1},
	ast.DayOfMonth:       &functionClass{ast.DayOfMonth, 1, 1},
	ast.DayOfWeek:        &functionClass{ast.DayOfWeek, 1, 1},
	ast.DayOfYear:        &functionClass{ast.DayOfYear, 1, 1},
	ast.Extract:          &functionClass{ast.Extract, 2, 2},
	ast.FromDays:         &functionClass{ast.FromDays, 1, 1},
	ast.FromUnixTime:     &functionClass{ast.FromUnixTime, 1, 2},
	ast.GetFormat:        &functionClass{ast.GetFormat, 2, 2},
	ast.Hour:             &functionClass{ast.Hour, 1, 1},
	ast.LocalTime:        &functionClass{ast.LocalTime, 0, 1},
	ast.LocalTimestamp:   &functionClass{ast.LocalTimestamp, 0, 1},
	ast.MakeDate:         &functionClass{ast.MakeDate, 2, 2},
	ast.MakeTime:         &functionClass{ast.MakeTime, 3, 3},
	ast.MicroSecond:      &functionClass{ast.MicroSecond, 1, 1},
	ast.Minute:           &functionClass{ast.Minute, 1, 1},
	ast.Month:            &functionClass{ast.Month, 1, 1},
	ast.MonthName:        &functionClass{ast.MonthName, 1, 1},
	ast.Now:              &functionClass{ast.Now, 0, 1},
	ast.PeriodAdd:        &functionClass{ast.PeriodAdd, 2, 2},
	ast.PeriodDiff:       &functionClass{ast.PeriodDiff, 2, 2},
	ast.Quarter:          &functionClass{ast.Quarter, 1, 1},
	ast.SecToTime:        &functionClass{ast.SecToTime, 1, 1},
	ast.Second:           &functionClass{ast.Second, 1, 1},
	ast.StrToDate:        &functionClass{ast.StrToDate, 2, 2},
	ast.SubTime:          &functionClass{ast.SubTime, 2, 2},
	ast.Sysdate:          &functionClass{ast.Sysdate, 0, 1},
	ast.Time:             &functionClass{ast.Time, 1, 1},
	ast.TimeLiteral:      &functionClass{ast.TimeLiteral, 1, 1},
	ast.TimeFormat:       &functionClass{ast.TimeFormat, 2, 2},
	ast.TimeToSec:        &functionClass{ast.TimeToSec, 1, 1},
	ast.TimeDiff:         &functionClass{ast.TimeDiff, 2, 2},
	ast.Timestamp:        &functionClass{ast.Timestamp, 1, 2},
	ast.TimestampLiteral: &functionClass{ast.TimestampLiteral, 1, 2},
	ast.TimestampAdd:     &functionClass{ast.TimestampAdd, 3, 3},
	ast.TimestampDiff:    &functionClass{ast.TimestampDiff, 3, 3},
	ast.ToDays:           &functionClass{ast.ToDays, 1, 1},
	ast.ToSeconds:        &functionClass{ast.ToSeconds, 1, 1},
	ast.UnixTimestamp:    &functionClass{ast.UnixTimestamp, 0, 1},
	ast.UTCDate:          &functionClass{ast.UTCDate, 0, 0},
	ast.UTCTime:          &functionClass{ast.UTCTime, 0, 1},
	ast.UTCTimestamp:     &functionClass{ast.UTCTimestamp, 0, 1},
	ast.Week:             &functionClass{ast.Week, 1, 2},
	ast.Weekday:          &functionClass{ast.Weekday, 1, 1},
	ast.WeekOfYear:       &functionClass{ast.WeekOfYear, 1, 1},
	ast.Year:             &functionClass{ast.Year, 1, 1},
	ast.YearWeek:         &functionClass{ast.YearWeek, 1, 2},
	ast.LastDay:          &functionClass{ast.LastDay, 1, 1},

	// string functions
	ast.ASCII:           &functionClass{ast.ASCII, 1, 1},
	ast.Bin:             &functionClass{ast.Bin, 1, 1},
	ast.Concat:          &functionClass{ast.Concat, 1, -1},
	ast.ConcatWS:        &functionClass{ast.ConcatWS, 2, -1},
	ast.Convert:         &functionClass{ast.Convert, 2, 2},
	ast.Elt:             &functionClass{ast.Elt, 2, -1},
	ast.ExportSet:       &functionClass{ast.ExportSet, 3, 5},
	ast.Field:           &functionClass{ast.Field, 2, -1},
	ast.Format:          &functionClass{ast.Format, 2, 3},
	ast.FromBase64:      &functionClass{ast.FromBase64, 1, 1},
	ast.InsertFunc:      &functionClass{ast.InsertFunc, 4, 4},
	ast.Instr:           &functionClass{ast.Instr, 2, 2},
	ast.Lcase:           &functionClass{ast.Lcase, 1, 1},
	ast.Left:            &functionClass{ast.Left, 2, 2},
	ast.Right:           &functionClass{ast.Right, 2, 2},
	ast.Length:          &functionClass{ast.Length, 1, 1},
	ast.LoadFile:        &functionClass{ast.LoadFile, 1, 1},
	ast.Locate:          &functionClass{ast.Locate, 2, 3},
	ast.Lower:           &functionClass{ast.Lower, 1, 1},
	ast.Lpad:            &functionClass{ast.Lpad, 3, 3},
	ast.LTrim:           &functionClass{ast.LTrim, 1, 1},
	ast.Mid:             &functionClass{ast.Mid, 3, 3},
	ast.MakeSet:         &functionClass{ast.MakeSet, 2, -1},
	ast.Oct:             &functionClass{ast.Oct, 1, 1},
	ast.OctetLength:     &functionClass{ast.OctetLength, 1, 1},
	ast.Ord:             &functionClass{ast.Ord, 1, 1},
	ast.Position:        &functionClass{ast.Position, 2, 2},
	ast.Quote:           &functionClass{ast.Quote, 1, 1},
	ast.Repeat:          &functionClass{ast.Repeat, 2, 2},
	ast.Replace:         &functionClass{ast.Replace, 3, 3},
	ast.Reverse:         &functionClass{ast.Reverse, 1, 1},
	ast.RTrim:           &functionClass{ast.RTrim, 1, 1},
	ast.Space:           &functionClass{ast.Space, 1, 1},
	ast.Strcmp:          &functionClass{ast.Strcmp, 2, 2},
	ast.Substring:       &functionClass{ast.Substring, 2, 3},
	ast.Substr:          &functionClass{ast.Substr, 2, 3},
	ast.SubstringIndex:  &functionClass{ast.SubstringIndex, 3, 3},
	ast.ToBase64:        &functionClass{ast.ToBase64, 1, 1},
	ast.Trim:            &functionClass{ast.Trim, 1, 3},
	ast.Upper:           &functionClass{ast.Upper, 1, 1},
	ast.Ucase:           &functionClass{ast.Ucase, 1, 1},
	ast.Hex:             &functionClass{ast.Hex, 1, 1},
	ast.Unhex:           &functionClass{ast.Unhex, 1, 1},
	ast.Rpad:            &functionClass{ast.Rpad, 3, 3},
	ast.BitLength:       &functionClass{ast.BitLength, 1, 1},
	ast.CharFunc:        &functionClass{ast.CharFunc, 2, -1},
	ast.CharLength:      &functionClass{ast.CharLength, 1, 1},
	ast.CharacterLength: &functionClass{ast.CharacterLength, 1, 1},
	ast.FindInSet:       &functionClass{ast.FindInSet, 2, 2},

	// information functions
	ast.ConnectionID: &functionClass{ast.ConnectionID, 0, 0},
	ast.CurrentUser:  &functionClass{ast.CurrentUser, 0, 0},
	ast.CurrentRole:  &functionClass{ast.CurrentRole, 0, 0},
	ast.Database:     &functionClass{ast.Database, 0, 0},
	// This function is a synonym for DATABASE().
	// See http://dev.mysql.com/doc/refman/5.7/en/information-functions.html#function_schema
	ast.Schema:       &functionClass{ast.Schema, 0, 0},
	ast.FoundRows:    &functionClass{ast.FoundRows, 0, 0},
	ast.LastInsertId: &functionClass{ast.LastInsertId, 0, 1},
	ast.User:         &functionClass{ast.User, 0, 0},
	ast.Version:      &functionClass{ast.Version, 0, 0},
	ast.Benchmark:    &functionClass{ast.Benchmark, 2, 2},
	ast.Charset:      &functionClass{ast.Charset, 1, 1},
	ast.Coercibility: &functionClass{ast.Coercibility, 1, 1},
	ast.Collation:    &functionClass{ast.Collation, 1, 1},
	ast.RowCount:     &functionClass{ast.RowCount, 0, 0},
	ast.SessionUser:  &functionClass{ast.SessionUser, 0, 0},
	ast.SystemUser:   &functionClass{ast.SystemUser, 0, 0},

	// control functions
	ast.If:     &functionClass{ast.If, 3, 3},
	ast.Ifnull: &functionClass{ast.Ifnull, 2, 2},

	// miscellaneous functions
	ast.AnyValue:        &functionClass{ast.AnyValue, 1, 1},
	ast.InetAton:        &functionClass{ast.InetAton, 1, 1},
	ast.InetNtoa:        &functionClass{ast.InetNtoa, 1, 1},
	ast.Inet6Aton:       &functionClass{ast.Inet6Aton, 1, 1},
	ast.Inet6Ntoa:       &functionClass{ast.Inet6Ntoa, 1, 1},
	ast.IsFreeLock:      &functionClass{ast.IsFreeLock, 1, 1},
	ast.IsIPv4:          &functionClass{ast.IsIPv4, 1, 1},
	ast.IsIPv4Compat:    &functionClass{ast.IsIPv4Compat, 1, 1},
	ast.IsIPv4Mapped:    &functionClass{ast.IsIPv4Mapped, 1, 1},
	ast.IsIPv6:          &functionClass{ast.IsIPv6, 1, 1},
	ast.IsUsedLock:      &functionClass{ast.IsUsedLock, 1, 1},
	ast.MasterPosWait:   &functionClass{ast.MasterPosWait, 2, 4},
	ast.NameConst:       &functionClass{ast.NameConst, 2, 2},
	// ast.ReleaseAllLocks: &functionClass{ast.ReleaseAllLocks, 0, 0},
	ast.UUID:            &functionClass{ast.UUID, 0, 0},
	// ast.UUIDShort:       &functionClass{ast.UUIDShort, 0, 0},

	// encryption and compression functions
	ast.AesDecrypt:               &functionClass{ast.AesDecrypt, 2, 3},
	ast.AesEncrypt:               &functionClass{ast.AesEncrypt, 2, 3},
	ast.Compress:                 &functionClass{ast.Compress, 1, 1},
	ast.Decode:                   &functionClass{ast.Decode, 2, 2},
	ast.DesDecrypt:               &functionClass{ast.DesDecrypt, 1, 2},
	ast.DesEncrypt:               &functionClass{ast.DesEncrypt, 1, 2},
	ast.Encode:                   &functionClass{ast.Encode, 2, 2},
	ast.Encrypt:                  &functionClass{ast.Encrypt, 1, 2},
	ast.MD5:                      &functionClass{ast.MD5, 1, 1},
	ast.OldPassword:              &functionClass{ast.OldPassword, 1, 1},
	ast.PasswordFunc:             &functionClass{ast.PasswordFunc, 1, 1},
	ast.RandomBytes:              &functionClass{ast.RandomBytes, 1, 1},
	ast.SHA1:                     &functionClass{ast.SHA1, 1, 1},
	ast.SHA:                      &functionClass{ast.SHA, 1, 1},
	ast.SHA2:                     &functionClass{ast.SHA2, 2, 2},
	ast.Uncompress:               &functionClass{ast.Uncompress, 1, 1},
	ast.UncompressedLength:       &functionClass{ast.UncompressedLength, 1, 1},
	ast.ValidatePasswordStrength: &functionClass{ast.ValidatePasswordStrength, 1, 1},

	// json functions
	ast.JSONType:          &functionClass{ast.JSONType, 1, 1},
	ast.JSONExtract:       &functionClass{ast.JSONExtract, 2, -1},
	ast.JSONUnquote:       &functionClass{ast.JSONUnquote, 1, 1},
	ast.JSONSet:           &functionClass{ast.JSONSet, 3, -1},
	ast.JSONInsert:        &functionClass{ast.JSONInsert, 3, -1},
	ast.JSONReplace:       &functionClass{ast.JSONReplace, 3, -1},
	ast.JSONRemove:        &functionClass{ast.JSONRemove, 2, -1},
	ast.JSONMerge:         &functionClass{ast.JSONMerge, 2, -1},
	ast.JSONObject:        &functionClass{ast.JSONObject, 0, -1},
	ast.JSONArray:         &functionClass{ast.JSONArray, 0, -1},
	ast.JSONContains:      &functionClass{ast.JSONContains, 2, 3},
	ast.JSONContainsPath:  &functionClass{ast.JSONContainsPath, 3, -1},
	ast.JSONValid:         &functionClass{ast.JSONValid, 1, 1},
	ast.JSONArrayAppend:   &functionClass{ast.JSONArrayAppend, 3, -1},
	ast.JSONArrayInsert:   &functionClass{ast.JSONArrayInsert, 3, -1},
	ast.JSONMergePatch:    &functionClass{ast.JSONMergePatch, 2, -1},
	ast.JSONMergePreserve: &functionClass{ast.JSONMergePreserve, 2, -1},
	ast.JSONPretty:        &functionClass{ast.JSONPretty, 1, 1},
	ast.JSONQuote:         &functionClass{ast.JSONQuote, 1, 1},
	ast.JSONSearch:        &functionClass{ast.JSONSearch, 3, -1},
	ast.JSONStorageSize:   &functionClass{ast.JSONStorageSize, 1, 1},
	ast.JSONDepth:         &functionClass{ast.JSONDepth, 1, 1},
	ast.JSONKeys:          &functionClass{ast.JSONKeys, 1, 2},
	ast.JSONLength:        &functionClass{ast.JSONLength, 1, 2},

	// TiDB internal function.
	ast.TiDBDecodeKey: &functionClass{ast.TiDBDecodeKey, 1, 1},
	// This function is used to show tidb-server version info.
	ast.TiDBVersion:    &functionClass{ast.TiDBVersion, 0, 0},
	ast.TiDBIsDDLOwner: &functionClass{ast.TiDBIsDDLOwner, 0, 0},
	ast.TiDBParseTso:   &functionClass{ast.TiDBParseTso, 1, 1},
	ast.TiDBDecodePlan: &functionClass{ast.TiDBDecodePlan, 1, 1},
}
