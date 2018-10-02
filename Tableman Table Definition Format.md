# Tableman

## Table definition format

```json
{
    "table_name": {
        "column_name": {
            "type": "string, tinyint, smallint, mediumint, integer, biginteger, decimal, double, enum, year, boolean, text, timestamp, date, uuid",
            "traits": [
                "autoincrement",
                "first",
                "unsigned",
                "usecurrent"
            ],
            "modifiers": {
                "after": "column_name",
                "charset": "utf8",
                "collation": "utf8_unicode_ci",
                "comment": "Comment text",
                "nullable": true,
                "stored_as": "MYSQL Expression",
                "virtual_as": "MYSQL Expression",
                "indexed_as": "index_name"
            }
        },
        "constraints": {
            "column_name": {
                "table": "table_name",
                "column": "column_name",
                "on_delete": "cascade etc..",
                "on_update": "cascade etc.."
            }
        }
    }
}
```

