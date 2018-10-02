package mysql

import (
	"testing"
)

func TestBuildTableDefinition(t *testing.T) {
	expected := "CREATE TABLE testTable (id int unsigned auto_increment not null, testColumn2 varchar(255) not null, PRIMARY KEY (id));"
	testColumn := &MySqlColumn{
		Name:      "id",
		Traits:    []string{"autoincrement", "unsigned"},
		Type:      "int",
		Modifiers: map[string]interface{}{}}
	testColumn2 := &MySqlColumn{
		Name:      "testColumn2",
		Traits:    make([]string, 0),
		Type:      "string",
		Modifiers: make(map[string]interface{}, 0)}

	table := &MySqlTable{
		Columns:     []MySqlColumn{*testColumn, *testColumn2},
		Constraints: make([]MySqlConstraint, 0),
		Name:        "testTable",
		Modifiers: map[string]interface{}{
			"primary_key": "id"}}
	res := table.BuildTableDefinition()
	if expected != res {
		t.Errorf("Expected %s to be equal to %s", res, expected)
	}
}
func TestBuildTableDefinition2(t *testing.T) {
	expected := "CREATE TABLE testTable2 (id int unsigned auto_increment not null, testColumn2 varchar(255) not null, PRIMARY KEY (id), FOREIGN KEY (id) REFERENCES testTable(id) ON DELETE cascade ON UPDATE set null);"
	testColumn := &MySqlColumn{
		Name:      "id",
		Traits:    []string{"autoincrement", "unsigned"},
		Type:      "int",
		Modifiers: map[string]interface{}{}}
	testColumn2 := &MySqlColumn{
		Name:      "testColumn2",
		Traits:    make([]string, 0),
		Type:      "string",
		Modifiers: make(map[string]interface{}, 0)}
	table := &MySqlTable{
		Columns: []MySqlColumn{*testColumn, *testColumn2},
		Constraints: []MySqlConstraint{
			{
				ColumnName:    "id",
				ForeignColumn: "id",
				ForeignTable:  "testTable",
				OnDelete:      "cascade",
				OnUpdate:      "set null",
			},
		},
		Name: "testTable2",
		Modifiers: map[string]interface{}{
			"primary_key": "id"}}
	res := table.BuildTableDefinition()
	if expected != res {
		t.Errorf("Expected %s to be equal to %s", res, expected)
	}
}
func TestBuildTableDefinition3(t *testing.T) {
	expected := "CREATE TABLE testTable2 (id int unsigned auto_increment not null, testColumn2 varchar(255) not null, PRIMARY KEY (id), FOREIGN KEY (id) REFERENCES testTable(id) ON DELETE cascade ON UPDATE set null, INDEX testColumn2_ind (testColumn2));"
	testColumn := &MySqlColumn{
		Name:      "id",
		Traits:    []string{"autoincrement", "unsigned"},
		Type:      "int",
		Modifiers: map[string]interface{}{}}
	testColumn2 := &MySqlColumn{
		Name:   "testColumn2",
		Traits: make([]string, 0),
		Type:   "string",
		Modifiers: map[string]interface{}{
			"indexed_as": "testColumn2_ind",
		}}
	table := &MySqlTable{
		Columns: []MySqlColumn{*testColumn, *testColumn2},
		Constraints: []MySqlConstraint{
			{
				ColumnName:    "id",
				ForeignColumn: "id",
				ForeignTable:  "testTable",
				OnDelete:      "cascade",
				OnUpdate:      "set null",
			},
		},
		Name: "testTable2",
		Modifiers: map[string]interface{}{
			"primary_key": "id"}}
	res := table.BuildTableDefinition()
	if expected != res {
		t.Errorf("Expected %s to be equal to %s", res, expected)
	}
}
