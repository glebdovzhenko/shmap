package shdb

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func DefaultDB() *sql.DB {
	db, err := sql.Open("sqlite3", ".appdata/default.db")
	if err != nil {
		panic(err)
	}

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

	return db
}
