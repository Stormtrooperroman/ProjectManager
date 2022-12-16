package main

import (
	"awesomeProject4/database"
	"awesomeProject4/model"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var router *gin.Engine

func main() {
	database.ConectDataBase()
	router := gin.Default()

	router.LoadHTMLGlob("html/*.html")   //шаблоны
	router.Static("/static", "./static") //css, js ...

	router.GET("/", login_page) //роуты надо прописать в html пути со "static/"
	router.GET("/login", login_page)
	router.POST("/api/login", registration)
	router.Run(":3000")

}

func login_page(c *gin.Context) {
	c.HTML(200, "login.html", nil)
}

func registration(c *gin.Context) {
	var user *model.User
	decode := json.NewDecoder(c.Request.Body).Decode(&user)

	fmt.Println(user.Pass) //провера состояния в постмане

	if decode != nil {
		c.JSON(http.StatusOK, gin.H{
			"response": decode.Error(),
		})
		return
	}

}
