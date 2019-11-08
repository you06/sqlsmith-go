package util


import (
	"fmt"
	"time"
	"github.com/pingcap/parser/mysql"
	tidbTypes "github.com/pingcap/tidb/types"
)

func GenerateDateItemString(columnType string) string {
	d := GenerateDataItem(columnType)
	switch c := d.(type) {
	case string:
		return c
	case int:
		return fmt.Sprintf("\"%d\"", c)
	case time.Time:
		return c.Format("2006-01-02 15:04:05")
	case tidbTypes.Time:
		return c.String()
	case float64:
		return fmt.Sprintf("%f", c)
	}
	return "not implement data transfer"
}

func GenerateDataItem(columnType string) interface{} {
	var res interface{}
	switch columnType {
	case "varchar":
		res = GenerateStringItem()
	case "text":
		res = GenerateStringItem()
	case "int":
		res = GenerateIntItem()
	case "timestamp":
		res = GenerateTiDBDateItem()
	case "float":
		res = GenerateFloatItem()
	}
	return res
}

func GenerateStringItem() string {
	return RdString(Rd(100))
}

func GenerateIntItem() int {
	return Rd(2147483647)
}

func GenerateFloatItem() float64 {
	return float64(Rd(100000)) * RdFloat64()
}

func GenerateDateItem() time.Time {
	return RdDate()
}

func GenerateTiDBDateItem() tidbTypes.Time {
	return tidbTypes.Time{
		Time: tidbTypes.FromGoTime(GenerateDateItem()),
		Type: mysql.TypeDatetime,
	}
}
