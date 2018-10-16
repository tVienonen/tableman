package mysql

import (
	"fmt"

	"github.com/tVienonen/tableman/internal/mysql/builder"
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
		args := make([]string, 0)
		if v, ok := m.Modifiers["column_arguments"]; ok {
			switch v.(type) {
			case []string:
				args = append(args, v.([]string)...)
			case []interface{}:
				tempArgs := v.([]interface{})
				for _, arg := range tempArgs {
					args = append(args, arg.(string))
				}
			}
		}
		b = b.Timestamp(args...)
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
