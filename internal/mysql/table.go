package mysql

import (
	"fmt"
)

type MySqlTable struct {
	Name        string                 `json:"name"`
	Columns     []MySqlColumn          `json:"columns"`
	Constraints []MySqlConstraint      `json:"constraints"`
	Modifiers   map[string]interface{} `json:"modifiers"`
}
type MySqlConstraint struct {
	ColumnName    string `json:"column_name"`
	ForeignTable  string `json:"foreign_table"`
	ForeignColumn string `json:"foreign_column"`
	OnDelete      string `json:"on_delete"`
	OnUpdate      string `json:"on_update"`
}

func (t *MySqlTable) BuildTableDefinition() string {
	definition := fmt.Sprintf("CREATE TABLE %s ()", t.Name)
	pos := len(definition) - 1
	// Add column declarations
	for i, column := range t.Columns {
		col := column.BuildColumnDefinition()
		if i < len(t.Columns)-1 {
			col += ","
		}
		if i > 0 {
			col = " " + col
		}
		definition = definition[:pos] + col + definition[pos:]
		pos += len(col)
	}
	// Add modifiers
	if val, ok := t.Modifiers["primary_key"]; ok {
		modifier := fmt.Sprintf(", PRIMARY KEY (%s)", val.(string))
		definition = definition[:pos] + modifier + definition[pos:]
		pos += len(modifier)
	}
	// Add constraints
	for _, constraint := range t.Constraints {
		var (
			onDelete = ""
			onUpdate = ""
		)
		if constraint.OnDelete != "" {
			onDelete = fmt.Sprintf("ON DELETE %s", constraint.OnDelete)
		}
		if constraint.OnUpdate != "" {
			onUpdate = fmt.Sprintf("ON UPDATE %s", constraint.OnUpdate)
		}
		constraintDefinition := fmt.Sprintf(", FOREIGN KEY (%s) REFERENCES %s(%s) %s %s",
			constraint.ColumnName,
			constraint.ForeignTable,
			constraint.ForeignColumn,
			onDelete,
			onUpdate)
		definition = definition[:pos] + constraintDefinition + definition[pos:]
		pos += len(constraintDefinition)
	}
	// Add indices
	for _, column := range t.Columns {
		if val, ok := column.Modifiers["indexed_as"]; ok {
			index := fmt.Sprintf(", INDEX %s (%s)", val.(string), column.Name)
			definition = definition[:pos] + index + definition[pos:]
			pos += len(index)
		}
	}
	// End statement
	definition += ";"
	return definition
}
