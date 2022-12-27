package databasegorm

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

const (

)

var dbase *gorm.DB

type User struct {
	Id      int    `db:"id"`
	FName   string `db:"first_name"`
	LName   string `db:"last_name"`
	Post    int    `db:"post"`
	Proj_id int    `db:"project_id"`
}

func Init() *gorm.DB {
	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	//conect database
	db, err := gorm.Open(postgres.Open(psqlconn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(User{})
	fmt.Println("Connected!")
	return db
}
func GetDB() *gorm.DB {
	if dbase == nil {
		dbase = Init()
		var sleep = time.Duration(1)
		for dbase == nil {
			sleep = sleep * 2
			fmt.Printf("database not data sleep =%d", sleep)
			time.Sleep(sleep * time.Second)
			dbase = Init()
		}
	}
	return dbase
}
