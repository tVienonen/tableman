package mysql

import (
	"fmt"
	s "strings"
)

type MySqlConstraint struct {
	ColumnName    string `json:"column_name"`
	ForeignTable  string `json:"foreign_table"`
	ForeignColumn string `json:"foreign_column"`
	OnDelete      string `json:"on_delete"`
	OnUpdate      string `json:"on_update"`
}

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
	var definition string
	size := ""
	if val, ok := m.Modifiers["size"]; ok {
		switch v := val.(type) {
		case int:
			size = fmt.Sprintf("(%d)", v)
		case string:
			size = fmt.Sprintf("(%s)", v)
		}
	}
	switch m.Type {
	case "string":
		definition = fmt.Sprintf("%s varchar(255)", m.Name)
		break
	case "tinyint":
		definition = fmt.Sprintf("%s tinyint%s", m.Name, size)
		break
	case "smallint":
		definition = fmt.Sprintf("%s smallint%s", m.Name, size)
		break
	case "mediumint":
		definition = fmt.Sprintf("%s mediumint%s", m.Name, size)
		break
	case "int":
		definition = fmt.Sprintf("%s int%s", m.Name, size)
		break
	case "integer":
		definition = fmt.Sprintf("%s int%s", m.Name, size)
		break
	case "bigint":
		definition = fmt.Sprintf("%s bigint%s", m.Name, size)
		break
	case "decimal":
		definition = fmt.Sprintf("%s decimal%s", m.Name, size)
		break
	case "double":
		definition = fmt.Sprintf("%s double%s", m.Name, size)
		break
	case "enum":
		definition = fmt.Sprintf("%s enum%s", m.Name, size)
		break
	case "year":
		definition = fmt.Sprintf("%s year%s", m.Name, size)
		break
	case "boolean":
		definition = fmt.Sprintf("%s tinyint(1)", m.Name)
		break
	case "text":
		definition = fmt.Sprintf("%s text%s", m.Name, size)
		break
	case "timestamp":
		definition = fmt.Sprintf("%s timestamp", m.Name)
		break
	case "date":
		definition = fmt.Sprintf("%s date", m.Name)
		break
	default:
		panic(fmt.Sprintf("Unsupported type %s", m.Type))
	}
	// unsigned
	for _, trait := range m.Traits {
		if s.ToLower(trait) == "unsigned" {
			definition += " unsigned"
		}
	}
	// Add traits
	for _, trait := range m.Traits {
		if val, ok := supportedTraits[trait]; ok {
			definition += fmt.Sprintf(" %s", val)
		}
	}
	// Add comment
	if val, ok := m.Modifiers["comment"]; ok {
		definition += fmt.Sprintf(" comment '%s'", val.(string))
	}
	// Add charset
	if val, ok := m.Modifiers["charset"]; ok {
		definition += fmt.Sprintf(" character set %s", val.(string))
	}
	// Add collation
	if val, ok := m.Modifiers["collation"]; ok {
		definition += fmt.Sprintf(" collate %s", val.(string))
	}
	// Add nullable
	if val, ok := m.Modifiers["nullable"]; ok && val.(bool) {
		definition += " null"
	} else {
		definition += " not null"
	}
	return definition
}
