package main

import (
	"awesomeProject4/datamysql"
	"awesomeProject4/pages"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var router *gin.Engine

func main() {
	//подключение к БД
	datamysql.Conect()
	datamysql.ExtractData(datamysql.Db)
	//datamysql.DelData(db)

	router := gin.Default()

	router.LoadHTMLGlob("html/*.html")   //шаблоны
	router.Static("/static", "./static") //css, js ...
	//роуты надо прописать в html пути со "static/"
	router.GET("/", pages.Project)                // работает всё ок
	router.GET("/login", pages.Login_page)        // работает всё ок
	router.POST("/api/login", pages.Registration) // работает всё ок
	router.GET("/calendar", pages.All_calendar)   // Рома где Блять calendar.css , а так всё ок
	router.GET("/project", pages.Project_info)    // Рома где Блять info.css
	router.Run(":3001")

}
