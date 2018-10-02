package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"gitlab.com/topivienonen/tableman/internal/mysql"
)

func main() {
	// Define command line flags
	if len(os.Args) <= 1 {
		fmt.Println("This program expects first argument to be a path to a directory containing the table definition files")
		os.Exit(1)
	}
	fmt.Println("Tableman is starting...")
	tableDefinitionDirectory := os.Args[1]

	dirStats, err := os.Stat(tableDefinitionDirectory)

	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	if !dirStats.IsDir() {
		fmt.Println("First argument is not a valid directory")
		os.Exit(3)
	}
	files, err := ioutil.ReadDir(tableDefinitionDirectory)
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}
	for i := range files {
		if files[i].IsDir() {
			continue
		}
		file, err := os.Open(
			fmt.Sprintf("%s%s", tableDefinitionDirectory, files[i].Name()))
		if err != nil {
			fmt.Println(err)
			continue
		}
		content, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println(err)
			continue
		}
		table := mysql.MySqlTable{}
		json.Unmarshal(content, &table)
		first_column := table.Columns[0]

		fmt.Println(table, first_column)
	}
}
