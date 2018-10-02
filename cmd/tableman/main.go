package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Define command line flags
	if len(os.Args) <= 1 {
		fmt.Println("This program expects first argument to be a path to a directory containing the table definition files")
		os.Exit(1)
	}
	fmt.Println("Tableman is starting...")
	tableDefinitionDirectory := os.Args[1]
	tables := GetTableDefinitionFiles(tableDefinitionDirectory)

	user := os.Getenv("TABLEMAN_USER")
	pw := os.Getenv("TABLEMAN_PASS")
	db := os.Getenv("TABLEMAN_DB_NAME")
	host := os.Getenv("TABLEMAN_HOST")
	port := "3306"
	if val, ok := os.LookupEnv("TABLEMAN_PORT"); ok {
		port = val
	}

	dbClient, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		user,
		pw,
		host,
		port,
		db))
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	defer dbClient.Close()
	for _, table := range tables {
		definition := table.BuildTableDefinition()
		_, err := dbClient.Exec(definition)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("Created table %s successfully\n", table.Name)
		}
	}
}
