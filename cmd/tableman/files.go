package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	p "path"
	"strings"

	"github.com/tVienonen/tableman/internal/mysql"
)

func GetTableDefinitionFiles(path string) []mysql.MySqlTable {
	if err := validatePath(path); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var (
		fileNames []string
		err       error
	)
	if fileNames, err = getFileNames(path); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	tables := make([]mysql.MySqlTable, 0)
	for _, filename := range fileNames {
		file, err := os.Open(p.Join(path, filename))
		if err != nil {
			fmt.Println(err)
			continue
		}
		defer file.Close()
		content, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println(err)
			continue
		}
		table := mysql.MySqlTable{}
		json.Unmarshal(content, &table)
		tables = append(tables, table)
	}
	return tables
}
func getFileNames(path string) ([]string, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return []string{}, err
	}
	res := func() []string {
		fileNames := make([]string, len(files))
		for i, file := range files {
			if !file.IsDir() && strings.HasSuffix(file.Name(), ".json") {
				fileNames[i] = file.Name()
			}
		}
		return fileNames
	}()
	return res, nil
}
func validatePath(path string) error {
	dirStats, err := os.Stat(path)

	if err != nil {
		return err
	}
	if !dirStats.IsDir() {
		return errors.New("First argument is not a valid directory")
	}
	return nil
}
