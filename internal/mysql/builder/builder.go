package builder

import (
	"fmt"
	s "strings"
)

var supportedTraits = map[string]string{
	"autoincrement": "auto_increment",
	"unique":        "unique",
}
var supportedUnsignedTypes = []string{
	"tinyint", "smallint", "mediumint", "int", "bigint", "decimal", "double", "timestamp"}

type MySqlColumnDefinitionBuilder struct {
	definition string
	Name       string
	Size       string
}

func New(name string, size interface{}) MySqlColumnDefinitionBuilder {
	_size := ""
	if size != nil {
		switch v := size.(type) {
		case int:
			_size = fmt.Sprintf("(%d)", v)
		case string:
			_size = fmt.Sprintf("(%s)", v)
		}
	}
	return MySqlColumnDefinitionBuilder{
		Size: _size,
		Name: name}
}

func (b MySqlColumnDefinitionBuilder) String() MySqlColumnDefinitionBuilder {
	b.padDefinition()
	size := b.Size
	if size == "" {
		size = "255"
	}
	b.definition = fmt.Sprintf("`%s` varchar(%s)%s", b.Name, size, b.definition)
	return b
}
func (b MySqlColumnDefinitionBuilder) TinyInt() MySqlColumnDefinitionBuilder {
	b.padDefinition()
	b.definition = fmt.Sprintf("`%s` tinyint%s%s", b.Name, b.Size, b.definition)
	return b
}
func (b MySqlColumnDefinitionBuilder) SmallInt() MySqlColumnDefinitionBuilder {
	b.padDefinition()
	b.definition = fmt.Sprintf("`%s` smallint%s%s", b.Name, b.Size, b.definition)
	return b
}
func (b MySqlColumnDefinitionBuilder) MediumInt() MySqlColumnDefinitionBuilder {
	b.padDefinition()
	b.definition = fmt.Sprintf("`%s` mediumint%s%s", b.Name, b.Size, b.definition)
	return b
}
func (b MySqlColumnDefinitionBuilder) Int() MySqlColumnDefinitionBuilder {
	b.padDefinition()
	b.definition = fmt.Sprintf("`%s` int%s%s", b.Name, b.Size, b.definition)
	return b
}
func (b MySqlColumnDefinitionBuilder) BigInt() MySqlColumnDefinitionBuilder {
	b.padDefinition()
	b.definition = fmt.Sprintf("`%s` bigint%s%s", b.Name, b.Size, b.definition)
	return b
}
func (b MySqlColumnDefinitionBuilder) Decimal() MySqlColumnDefinitionBuilder {
	b.padDefinition()
	b.definition = fmt.Sprintf("`%s` decimal%s%s", b.Name, b.Size, b.definition)
	return b
}
func (b MySqlColumnDefinitionBuilder) Double() MySqlColumnDefinitionBuilder {
	b.padDefinition()
	b.definition = fmt.Sprintf("`%s` double%s%s", b.Name, b.Size, b.definition)
	return b
}
func (b MySqlColumnDefinitionBuilder) Enum() MySqlColumnDefinitionBuilder {
	b.padDefinition()
	b.definition = fmt.Sprintf("`%s` enum%s%s", b.Name, b.Size, b.definition)
	return b
}
func (b MySqlColumnDefinitionBuilder) Year() MySqlColumnDefinitionBuilder {
	b.padDefinition()
	b.definition = fmt.Sprintf("`%s` year%s%s", b.Name, b.Size, b.definition)
	return b
}
func (b MySqlColumnDefinitionBuilder) Bool() MySqlColumnDefinitionBuilder {
	b.padDefinition()
	b.definition = fmt.Sprintf("`%s` tinyint(1)%s", b.Name, b.definition)
	return b
}
func (b MySqlColumnDefinitionBuilder) Text() MySqlColumnDefinitionBuilder {
	b.padDefinition()
	b.definition = fmt.Sprintf("`%s` text%s%s", b.Name, b.Size, b.definition)
	return b
}
func (b MySqlColumnDefinitionBuilder) Timestamp() MySqlColumnDefinitionBuilder {
	b.padDefinition()
	b.definition = fmt.Sprintf("`%s` timestamp%s", b.Name, b.definition)
	return b
}
func (b MySqlColumnDefinitionBuilder) Date() MySqlColumnDefinitionBuilder {
	b.padDefinition()
	b.definition = fmt.Sprintf("`%s` date%s", b.Name, b.definition)
	return b
}

func (b MySqlColumnDefinitionBuilder) AddTraits(traits []string) MySqlColumnDefinitionBuilder {
	// unsigned
	for _, trait := range traits {
		if s.ToLower(trait) == "unsigned" {
			// Find position of the column type in definition string
			for _, unsignedType := range supportedUnsignedTypes {
				i := s.Index(b.definition, unsignedType)
				if i != -1 {
					endOfWord := i + len(unsignedType)
					b.definition = b.definition[:endOfWord] + " unsigned" + b.definition[endOfWord:]
					break
				}
			}
		}
	}
	// Add traits
	for _, trait := range traits {
		if val, ok := supportedTraits[trait]; ok {
			b.definition += fmt.Sprintf(" %s", val)
		}
	}
	return b
}
func (b MySqlColumnDefinitionBuilder) AddModifiers(modifiers map[string]interface{}) MySqlColumnDefinitionBuilder {
	// Add comment
	if val, ok := modifiers["comment"]; ok {
		b.definition += fmt.Sprintf(" comment '%s'", val.(string))
	}
	// Add charset
	if val, ok := modifiers["charset"]; ok {
		b.definition += fmt.Sprintf(" character set %s", val.(string))
	}
	// Add collation
	if val, ok := modifiers["collation"]; ok {
		b.definition += fmt.Sprintf(" collate %s", val.(string))
	}
	// Add nullable
	if val, ok := modifiers["nullable"]; ok && val.(bool) {
		b.definition += " null"
	} else {
		b.definition += " not null"
	}
	return b
}
func (b MySqlColumnDefinitionBuilder) Value() string {
	return b.definition
}
func (b *MySqlColumnDefinitionBuilder) padDefinition() {
	if len(b.definition) > 0 {
		b.definition = " " + b.definition
	}
}
