# Tableman

## Table definition format

```json
{
    "name": "mytable",
    "columns": [
        {
            "column_name": "mycolumn",
            "type": "string, tinyint, smallint, mediumint, integer, biginteger, decimal, double, enum, year, boolean, text, timestamp, date",
            "traits": [
                "autoincrement",
                "unsigned",
                "unique"
            ],
            "modifiers": {
                "charset": "utf8",
                "collation": "utf8_unicode_ci",
                "comment": "Comment text",
                "nullable": true,
                "indexed_as": "index_name"
            }
        }
    ],
    "constraints": [
        {
            "column_name": "mycolumn",
            "foreign_table": "table_name",
            "foreign_column": "column_name",
            "on_delete": "cascade etc..",
            "on_update": "cascade etc.."
        }
    ],
    "modifiers": {
        "primary_key": "column_name"
    }
}
```

