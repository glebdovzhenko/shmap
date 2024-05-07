package shdb

import (
	"database/sql"
	"fmt"

	shcfg "github.com/glebdovzhenko/shmap/config"
	_ "github.com/mattn/go-sqlite3"
)

// Is it obvious this is the first time I'm writing SQL?
func defaultPopulate() {
	populateStudents()
}


func populateStudents() {
	app_cfg := shcfg.GetConfig()

	db, err := sql.Open(app_cfg.DBType, app_cfg.DBPath)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	names01 := [][]string{
		{"Пригода", "Рубан", "GSLV123"},
		{"Истистав", "Огафонов", "GSLV123"},
		{"Светомир", "Вельяминов", "GSLV123"},
		{"Ярополк", "Рожков", "GSLV123"},
		{"Станислав", "Серебряков", "GSLV123"},
		{"Вечезар", "Соколов", "GSLV123"},
		{"Горимир", "Медведев", "GSLV123"},
		{"Властислав", "Мудрый", "GSLV123"},
		{"Милолика", "Лубяная", "GSLV123"},
		{"Негослав", "Славинский", "GSLV123"},
		{"Бронислав", "Неплюй", "GSLV123"},
		{"Ведагор", "Добрый", "GSLV123"},
		{"Милонежа", "Оболенская", "GSLV123"},
		{"Житко", "Мещанин", "GSLV123"},
		{"Годун", "Аладьин", "GSLV123"},
		{"Миромира", "Плюснина", "GSLV123"},
		{"Путята", "Медведь", "GSLV123"},
		{"Асмуд", "Проскур", "GSLV123"},
		{"Станимир", "Сорокожердьев", "GSLV123"},
		{"Добронега", "Котова", "GSLV123"},
		{"Василько", "Заколюкин", "GSLV123"},
	}

	names02 := [][]string{
		{"Фабия", "Кокрофт", "GMID321"},
		{"Калпурний", "Фред", "GMID321"},
		{"Филокл", "Кальвейра", "GMID321"},
		{"Полибий", "Аттвуд", "GMID321"},
		{"Леда", "Риман", "GMID321"},
		{"Эвмен", "Албахари", "GMID321"},
		{"Арета", "Иогансон", "GMID321"},
		{"Эсселта", "Эмден", "GMID321"},
		{"Федра", "Ланде", "GMID321"},
		{"Филокл", "Жевре", "GMID321"},
		{"Юлий", "Фиори", "GMID321"},
		{"Випсаний", "Милнор", "GMID321"},
		{"Аппулей", "Шаферан", "GMID321"},
		{"Гердоний", "Аньези", "GMID321"},
		{"Адмет", "Дюваль", "GMID321"},
		{"Акилий", "Кальвино", "GMID321"},
		{"Теттий", "Скан", "GMID321"},
		{"Эвдаф", "Перро", "GMID321"},
		{"Медея", "Мамфорд", "GMID321"},
		{"Этра", "Деберейнер", "GMID321"},
		{"Лепидий", "Пизье", "GMID321"},
	}

	names03 := [][]string{
		{"Михайлова", "Амелия Кирилловна", "GREG212"},
		{"Волошин", "Иван Олегович", "GREG212"},
		{"Винокурова", "Анастасия Артёмовна", "GREG212"},
		{"Кириллова", "Полина Алексеевна", "GREG212"},
		{"Прохоров", "Иван Никитич", "GREG212"},
		{"Богданов", "Олег Максимович", "GREG212"},
		{"Андреев", "Даниил Артёмович", "GREG212"},
		{"Антонов", "Максим Андреевич", "GREG212"},
		{"Кузьмина", "Дарья Кирилловна", "GREG212"},
		{"Лукин", "Богдан Кириллович", "GREG212"},
		{"Климов", "Дмитрий Данилович", "GREG212"},
		{"Николаев", "Ярослав Иванович", "GREG212"},
		{"Иванов", "Матвей Даниилович", "GREG212"},
		{"Кириллова", "Александра Эмировна", "GREG212"},
		{"Петрова", "Мария Мироновна", "GREG212"},
		{"Иванов", "Артём Ярославович", "GREG212"},
		{"Сергеев", "Константин Иванович", "GREG212"},
		{"Кузьмина", "Милана Платоновна", "GREG212"},
		{"Корчагин", "Александр Александрович", "GREG212"},
		{"Зыков", "Савелий Фёдорович", "GREG212"},
		{"Майорова", "Кира Артёмовна", "GREG212"},
	}

    _, err = db.Exec("CREATE TABLE usual (id integer, name varchar(32), surname varchar(32), sgroup varchar(32))")
    if err != nil {
        panic(err)
    }

	for ii, vv := range names03 {
		_, err = db.Exec(fmt.Sprintf("INSERT INTO usual VALUES (%d, \"%s\", \"%s\", \"%s\")", ii, vv[1], vv[0], vv[2]))
		if err != nil {
			panic(err)
		}
	}

    _, err = db.Exec("CREATE TABLE souffrant (id integer, name varchar(32), surname varchar(32), sgroup varchar(32))")
    if err != nil {
        panic(err)
    }

	for ii, vv := range names02 {
		_, err = db.Exec(fmt.Sprintf("INSERT INTO souffrant VALUES (%d, \"%s\", \"%s\", \"%s\")", ii, vv[0], vv[1], vv[2]))
		if err != nil {
			panic(err)
		}
	}
    _, err = db.Exec("CREATE TABLE svarog (id integer, name varchar(32), surname varchar(32), sgroup varchar(32))")
    if err != nil {
        panic(err)
    }

	for ii, vv := range names01 {
		_, err = db.Exec(fmt.Sprintf("INSERT INTO svarog VALUES (%d, \"%s\", \"%s\", \"%s\")", ii, vv[0], vv[1], vv[2]))
		if err != nil {
			panic(err)
		}
	}


}

