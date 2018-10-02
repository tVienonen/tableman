package internal

// ColumnBuilder builds a column definition
type ColumnBuilder interface {
	// Expects Column structure to create a valid column definition string
	BuildColumnDefinition() string
}
type TableBuilder interface {
	BuildTableDefinition() string
}
