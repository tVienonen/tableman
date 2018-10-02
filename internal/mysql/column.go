package mysql

import (
	"fmt"

	"gitlab.com/topivienonen/tableman/internal/mysql/builder"
)

type MySqlColumn struct {
	Name      string                 `json:"column_name"`
	Type      string                 `json:"type"`
	Traits    []string               `json:"traits"`
	Modifiers map[string]interface{} `json:"modifiers"`
}

var supportedTraits = map[string]string{
	"autoincrement": "auto_increment",
	"unique":        "unique",
}

func (m *MySqlColumn) BuildColumnDefinition() string {
	b := builder.New(m.Name, m.Modifiers["size"])
	switch m.Type {
	case "string":
		b = b.String()
		break
	case "tinyint":
		b = b.TinyInt()
		break
	case "smallint":
		b = b.SmallInt()
		break
	case "mediumint":
		b = b.MediumInt()
		break
	case "int":
		b = b.Int()
		break
	case "integer":
		b = b.Int()
		break
	case "bigint":
		b = b.BigInt()
		break
	case "decimal":
		b = b.Decimal()
		break
	case "double":
		b = b.Double()
		break
	case "enum":
		b = b.Enum()
		break
	case "year":
		b = b.Year()
		break
	case "boolean":
		b = b.Bool()
		break
	case "text":
		b = b.Text()
		break
	case "timestamp":
		b = b.Timestamp()
		break
	case "date":
		b = b.Date()
		break
	default:
		panic(fmt.Sprintf("Unsupported type %s", m.Type))
	}
	return b.
		AddTraits(m.Traits).
		AddModifiers(m.Modifiers).
		Value()
}
