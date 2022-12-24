package pages

import (
	"awesomeProject4/datamysql"
	"awesomeProject4/model"
	"awesomeProject4/privat_info"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Project(c *gin.Context) {
	if privat_info.Admin == true {
		c.HTML(200, "projects.html", nil)
	} else {
		c.HTML(200, "login.html", nil)
	}
}

func Project_info(c *gin.Context) {
	if privat_info.Admin == true {
		c.HTML(200, "project_info.html", nil)
	} else {
		c.HTML(200, "login.html", nil)
	}
}

func All_calendar(c *gin.Context) {
	if privat_info.Admin == true {
		c.HTML(200, "all_calendar.html", nil)
	} else {
		c.HTML(200, "login.html", nil)
	}

}

func Login_page(c *gin.Context) {
	c.HTML(200, "login.html", nil)
}

func Registration(c *gin.Context) {
	var user *model.Registr
	decode := json.NewDecoder(c.Request.Body).Decode(&user)
	fmt.Println(user)                       // структура с json внутри
	fmt.Println(user.Login, " ", user.Pass) //провера состояния в постмане
	datamysql.ExtractData(datamysql.Db, string(user.Login), string(user.Pass))
	fmt.Println(privat_info.Admin)
	if decode != nil {
		c.JSON(http.StatusOK, gin.H{
			"response": decode.Error(),
		})
		return
	}

}
