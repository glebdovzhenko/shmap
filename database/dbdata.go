package shdb
import (
	"errors"
)


type DBData []DBTableData

func (db DBData) Name(tb_id int) (string, error) {
	if tb_id >= len(db) {
		return "", errors.New("tb_id > len(tables)")
	}

	return db[tb_id].Name, nil
}

func (db DBData) ColumnsLen(tb_id int) (int, error) {
	if tb_id >= len(db) {
		return 0, errors.New("tb_id > len(tables)")
	}

	return len(*db[tb_id].ColumnNames), nil
}

func (db DBData) ColumnName(tb_id int, col_id int) (string, error) {
	if tb_id >= len(db) {
		return "", errors.New("tb_id > len(tables)")
	}
    cl, _ := db.ColumnsLen(tb_id)
	if col_id >= cl {
		return "", errors.New("col_id > len(ColumnNames)")
	}

	return (*db[tb_id].ColumnNames)[col_id], nil
}

func (db DBData) RowsLen(tb_id int) (int, error) {
	if tb_id >= len(db) {
		return 0, errors.New("tb_id > len(tables)")
	}

	return len(*db[tb_id].Rows), nil
}

func (db DBData) RowPtr(tb_id int, row_id int) (*[]string, error) {
	if tb_id >= len(db) {
		return nil, errors.New("tb_id > len(tables)")
	}
    cl, _ := db.RowsLen(tb_id)
	if row_id >= cl {
		return nil, errors.New("col_id > len(ColumnNames)")
	}

	return &(*db[tb_id].Rows)[row_id], nil
}

func GetDBData() (*DBData)  {
	setupDB()
    db := &DBData{}

	tb_names := getTablesNames()

	for _, tb_name := range *tb_names {
		cs, rs := getTable(tb_name)
		*db = append(*db, DBTableData{Name: tb_name, ColumnNames: cs, Rows: rs})
	}

    return db
}
