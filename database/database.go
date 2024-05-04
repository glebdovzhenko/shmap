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


func GetTable(t_name string) {
    app_cfg := shcfg.GetConfig()

    db, err := sql.Open(app_cfg.DBType, app_cfg.DBLoc)
    if err != nil {
        panic(err)
    }
    defer db.Close()
    
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
    var cns []string
    for _, ct := range  cts{
        cns = append(cns, ct.Name())
        fmt.Printf("%v\n", ct.ScanType())
    }
    fmt.Printf("%v\n", cns)

}
