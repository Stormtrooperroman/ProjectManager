package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"os"
)

var db *sqlx.DB

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

func QueryRowDemo() { //проблема тут!!!
	sqlStr := "select * from employees"
	var u string
	err := db.Get(&u, sqlStr, 1)
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return
	}
	fmt.Printf("result", u)
}
