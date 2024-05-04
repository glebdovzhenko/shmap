package shdb

import (
	"database/sql"
	"fmt"

	shcfg "github.com/glebdovzhenko/shmap/config"
	_ "github.com/mattn/go-sqlite3"
)

// Is it obvious this is the first time I'm writing SQL?
func DefaultPopulate() {
	populateFolders()
	populateStudents()
}

func populateFolders() {
	app_cfg := shcfg.GetConfig()

	db, err := sql.Open(app_cfg.DBType, app_cfg.DBLoc)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.Exec("DROP TABLE folders")
	_, err = db.Exec("CREATE TABLE folders (id integer, name varchar(32))")
	if err != nil {
		panic(err)
	}

	names := []string{"f1", "f2", "f3", "f4", "f5"}
	for ii, vv := range names {
		_, err = db.Exec(fmt.Sprintf("INSERT INTO folders VALUES (%d, \"%s\")", ii, vv))
		if err != nil {
			panic(err)
		}
	}
}

func populateStudents() {
	app_cfg := shcfg.GetConfig()

	db, err := sql.Open(app_cfg.DBType, app_cfg.DBLoc)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.Exec("DROP TABLE students")
	_, err = db.Exec("CREATE TABLE students (id integer, name varchar(32), surname varchar(32), sgroup varchar(32))")
	if err != nil {
		panic(err)
	}

	names := [][]string{
		{"Ivan", "Ivanov", "TST01"},
		{"Petr", "Petrov", "TST01"},
		{"Sidor", "Sidorov", "TST01"},
		{"Test", "Testovich", "TST01"},
		{"Lorem", "Ipsum", "TST01"},
		{"Dolor", "Sit", "TST01"},
		{"Amet", "Consectetur", "TST01"},
		{"Adipiscing", "Elit", "TST01"},
	}

	for ii, vv := range names {
		_, err = db.Exec(fmt.Sprintf("INSERT INTO students VALUES (%d, \"%s\", \"%s\", \"%s\")", ii, vv[0], vv[1], vv[2]))
		if err != nil {
			panic(err)
		}
	}
}

func GetTable(t_name string) ([]string, [][]string) {
	app_cfg := shcfg.GetConfig()

	// opening DB
	db, err := sql.Open(app_cfg.DBType, app_cfg.DBLoc)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// querying all data from the table
	rows, err := db.Query(fmt.Sprintf("SELECT * FROM %s", t_name))
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	// getting column types
	cts, err := rows.ColumnTypes()
	if err != nil {
		panic(err)
	}

	// from column types getting names & data types
	var column_names []string
	for _, ct := range cts {
		column_names = append(column_names, ct.Name())
		fmt.Printf("%v\n", ct.ScanType())
	}

	// initializing row_data
	var row_data [][]string

	// writing row_data from sql object
	// TODO: is there a better way?
	for rows.Next() {
		pointers := make([]interface{}, len(column_names))
		container := make([]string, len(column_names))
		for i, _ := range pointers {
			pointers[i] = &container[i]
		}

		rows.Scan(pointers...)
		row_data = append(row_data, container)
	}

	return column_names, row_data
}
