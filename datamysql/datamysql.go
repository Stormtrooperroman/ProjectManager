package datamysql

import (
	"awesomeProject4/model"
	"awesomeProject4/privat_info"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

func Conect() {
	db, err := sqlx.Open("mysql", privat_info.DataBaseKey)
	if err != nil {
		panic(err)
	}
	Db = db
	//defer Db.Close() //Закрытие БД
	fmt.Println("DataBase_is_WORK")
}

func ExtractData(db *sqlx.DB, login string, password string) { //получение пользователя из бд
	var u model.User_DB
	res, err := db.Query("SELECT `first_name`,`last_name`,`id` FROM `employees` WHERE `login`= ? AND `password` = ?;", login, password)
	if err != nil {
		panic(err)
	}
	for res.Next() {
		err = res.Scan(&u.FName, &u.LName, &u.Id)
		if err != nil {
			panic(err)
		}
		//fmt.Println(fmt.Sprintf("in database have %s , %s ", u.FName, u.LName))
	}
	if u.LName != "" && u.FName != "" {
		privat_info.Admin = true
	} else {
		privat_info.Admin = false
	}
	//fmt.Println(u.LName, " ", u.FName) //пример как вырывать параметры из запроса

}

func AddData(db *sqlx.DB) {
	result, err := db.Exec("insert into employees (id, first_name,last_name,post) values (?,?, ?, ?)", "8", "fuck", "fuck", "fuck")
	if err != nil {
		panic(err)
	}
	fmt.Println(result.LastInsertId()) // id добавленного объекта
	fmt.Println(result.RowsAffected()) // количество затронутых строк
}

func DelData(db *sqlx.DB) {
	result, err := db.Exec("delete from employees where id = ?", "8")
	if err != nil {
		panic(err)
	}
	fmt.Println(result.LastInsertId()) // id последнего удаленого объекта
	fmt.Println(result.RowsAffected()) // количество затронутых строк

}
