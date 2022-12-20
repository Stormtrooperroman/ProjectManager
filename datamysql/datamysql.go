package datamysql

import (
	"awesomeProject4/model"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func ExtractData(db *sql.DB) {

	var u model.User_DB
	res, err := db.Query("SELECT `first_name`,`last_name`,`post` FROM `employees` WHERE `first_name`= 'Daniil';")
	if err != nil {
		panic(err)
	}
	for res.Next() {
		err = res.Scan(&u.FName, &u.LName, &u.Post)
		if err != nil {
			panic(err)
		}
		fmt.Println(fmt.Sprintf("in database have %s , %s , %s", u.FName, u.LName, u.Post))
	}
	fmt.Println(u.LName, " ", u.FName) //пример как вырывать параметры из запроса

}

func AddData(db *sql.DB) {
	result, err := db.Exec("insert into employees (id, first_name,last_name,post) values (?,?, ?, ?)", "8", "fuck", "fuck", "fuck")
	if err != nil {
		panic(err)
	}
	fmt.Println(result.LastInsertId()) // id добавленного объекта
	fmt.Println(result.RowsAffected()) // количество затронутых строк
}

func DelData(db *sql.DB) {
	result, err := db.Exec("delete from employees where id = ?", "8")
	if err != nil {
		panic(err)
	}
	fmt.Println(result.LastInsertId()) // id последнего удаленого объекта
	fmt.Println(result.RowsAffected()) // количество затронутых строк

}
