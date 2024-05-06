// shdb package provides an inteface for 
// 1. Initializing a dummy database
// 2. Reading from the database.
package shdb

import (
	"database/sql"
	"fmt"

	shcfg "github.com/glebdovzhenko/shmap/config"
	_ "github.com/mattn/go-sqlite3"
)

type DBTableData struct {
	Name        string
	ColumnNames *[]string
	Rows        *[][]string
}

func getTable(t_name string) (*[]string, *[][]string) {
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

	return &column_names, &row_data
}

func getTablesNames() *[]string{
	app_cfg := shcfg.GetConfig()

	// opening DB
	db, err := sql.Open(app_cfg.DBType, app_cfg.DBLoc)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// querying all data from the table
	rows, err := db.Query("SELECT NAME FROM sqlite_master WHERE type='table'")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var tb_names []string
	for rows.Next() {
		pt := new(string)
		rows.Scan(pt)
		tb_names = append(tb_names, *pt)
	}

    return &tb_names
}

func GetDBData() *[]DBTableData {
    tb_names := getTablesNames()
    var result []DBTableData
    
    for _, tb_name := range *tb_names {
        cs, rs := getTable(tb_name)
        result = append(result, DBTableData{Name: tb_name, ColumnNames: cs, Rows: rs})
    }

    return &result
}
