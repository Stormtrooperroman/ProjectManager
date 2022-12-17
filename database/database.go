package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"os"
)

var db *sqlx.DB

const (

)

func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ConectDataBase() {
	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	//conect database
	db, err := sqlx.Open("postgres", os.Getenv(psqlconn))
	Check(err)

	defer db.Close()
	fmt.Println("Connected!")
}

type User struct {
	Id      int    `db:"id"`
	FName   string `db:"first_name"`
	LName   string `db:"last_name"`
	Post    int    `db:"post"`
	Proj_id int    `db:"project_id"`
}

func QueryRowDemo() { //проблема тут!!!
	sqlStr := "select id, first_name, last_name, post, project_id from employees"
	var users []User
	err := db.Select(&users, sqlStr, 0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	fmt.Printf("users:%#v\n", users)
}
func QueryDemo() {
	sqlStr := "select id, first_name, last_name, post, project_id from employees where id=1 "
	var u User

	err := db.Get(&u, sqlStr, 1)
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return
	}
	fmt.Printf("id:%d fname:%s lname:%s\n", u.Id, u.FName, u.FName)
}
