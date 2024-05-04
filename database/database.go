package shdb

import (
	"database/sql"
	"fmt"

	shcfg "github.com/glebdovzhenko/shmap/config"
	_ "github.com/mattn/go-sqlite3"
)

// Is it obvious this is the first time I'm writing SQL?
func DefaultPopulate() {
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

	_, err = db.Exec("INSERT INTO folders VALUES (1, \"f1\")")
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("INSERT INTO folders VALUES (2, \"f2\")")
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("INSERT INTO folders VALUES (3, \"f3\")")
	if err != nil {
		panic(err)
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
