package mysql

import "testing"

var cases = []struct {
	Expected    string
	TestColumn  *MySqlColumn
	ShouldPanic bool
}{
	{
		Expected: "`testColumn` varchar(255) not null",
		TestColumn: &MySqlColumn{
			Name:      "testColumn",
			Traits:    make([]string, 0),
			Type:      "string",
			Modifiers: make(map[string]interface{}, 0)},
	},
	{
		Expected: "`testColumn` varchar(800) not null",
		TestColumn: &MySqlColumn{
			Name:      "testColumn",
			Traits:    []string{},
			Type:      "string",
			Modifiers: map[string]interface{}{"size": float64(800)}},
	},
	{
		Expected: "`testColumn2` tinyint(3) not null",
		TestColumn: &MySqlColumn{
			Name:   "testColumn2",
			Traits: make([]string, 0),
			Type:   "tinyint",
			Modifiers: map[string]interface{}{
				"size": 3}},
	},
	{
		Expected: "`testColumn3` int not null",
		TestColumn: &MySqlColumn{
			Name:      "testColumn3",
			Traits:    make([]string, 0),
			Type:      "int",
			Modifiers: make(map[string]interface{}, 0)},
	},
	{
		Expected: "`testColumn3` int not null",
		TestColumn: &MySqlColumn{
			Name:      "testColumn3",
			Traits:    make([]string, 0),
			Type:      "integer",
			Modifiers: make(map[string]interface{}, 0)},
	},
	{
		Expected: "`testColumn3` smallint not null",
		TestColumn: &MySqlColumn{
			Name:      "testColumn3",
			Traits:    make([]string, 0),
			Type:      "smallint",
			Modifiers: make(map[string]interface{}, 0)},
	},
	{
		Expected: "`testColumn3` mediumint not null",
		TestColumn: &MySqlColumn{
			Name:      "testColumn3",
			Traits:    make([]string, 0),
			Type:      "mediumint",
			Modifiers: make(map[string]interface{}, 0)},
	},
	{
		Expected: "Unsupported type unsupported_type",
		TestColumn: &MySqlColumn{
			Name:      "testColumn3",
			Traits:    make([]string, 0),
			Type:      "unsupported_type",
			Modifiers: make(map[string]interface{}, 0)},
		ShouldPanic: true,
	},
	{
		Expected: "`testColumn3` int unsigned auto_increment not null",
		TestColumn: &MySqlColumn{
			Name:      "testColumn3",
			Traits:    []string{"autoincrement", "unsigned"},
			Type:      "int",
			Modifiers: make(map[string]interface{}, 0)},
	},
	{
		Expected: "`testColumn3` bigint unsigned auto_increment not null",
		TestColumn: &MySqlColumn{
			Name:      "testColumn3",
			Traits:    []string{"autoincrement", "unsigned"},
			Type:      "bigint",
			Modifiers: make(map[string]interface{}, 0)},
	},
	{
		Expected: "`testColumn3` decimal not null",
		TestColumn: &MySqlColumn{
			Name:      "testColumn3",
			Traits:    []string{},
			Type:      "decimal",
			Modifiers: make(map[string]interface{}, 0)},
	},
	{
		Expected: "`testColumn3` double not null",
		TestColumn: &MySqlColumn{
			Name:      "testColumn3",
			Traits:    []string{},
			Type:      "double",
			Modifiers: make(map[string]interface{}, 0)},
	},
	{
		Expected: "`testColumn3` enum('A', 'B') not null",
		TestColumn: &MySqlColumn{
			Name:   "testColumn3",
			Traits: []string{},
			Type:   "enum",
			Modifiers: map[string]interface{}{
				"size": "'A', 'B'"}},
	},
	{
		Expected: "`testColumn3` year not null",
		TestColumn: &MySqlColumn{
			Name:      "testColumn3",
			Traits:    []string{},
			Type:      "year",
			Modifiers: map[string]interface{}{}},
	},
	{
		Expected: "`testColumn3` tinyint(1) not null",
		TestColumn: &MySqlColumn{
			Name:      "testColumn3",
			Traits:    []string{},
			Type:      "boolean",
			Modifiers: map[string]interface{}{}},
	},
	{
		Expected: "`testColumn3` text not null",
		TestColumn: &MySqlColumn{
			Name:      "testColumn3",
			Traits:    []string{},
			Type:      "text",
			Modifiers: map[string]interface{}{}},
	},
	{
		Expected: "`testColumn3` timestamp not null",
		TestColumn: &MySqlColumn{
			Name:      "testColumn3",
			Traits:    []string{},
			Type:      "timestamp",
			Modifiers: map[string]interface{}{}},
	},
	{
		Expected: "`testColumn3` date not null",
		TestColumn: &MySqlColumn{
			Name:      "testColumn3",
			Traits:    []string{},
			Type:      "date",
			Modifiers: map[string]interface{}{}},
	},
}

func TestBuildDefinition(t *testing.T) {
	for _, testCase := range cases {
		func() {
			if testCase.ShouldPanic {
				defer func() {
					if r := recover(); r != nil {
						if r.(string) != testCase.Expected {
							t.Errorf("Expected %s to be equal to %s", r.(string), testCase.Expected)
						}
					}
				}()
			}
			result := testCase.TestColumn.BuildColumnDefinition()
			if testCase.Expected != result {
				t.Errorf("Expected %s to be equal to %s", result, testCase.Expected)
			}
		}()
	}
}
