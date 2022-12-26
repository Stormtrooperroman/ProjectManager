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

func Project_page(c *gin.Context) {
	_, err := c.Cookie("user")

	if err == nil {
		// Get data from DB
		projectsData := datamysql.ExtractData_Projects()
		returningResult := gin.H{
			"projects": projectsData,
		}
		admin, _ := c.Cookie("admin")

		if admin == "true" {
			returningResult["admin"] = true
		}
		c.HTML(200, "projects.html", returningResult)
	} else {
		c.HTML(200, "login.html", nil)
	}
}

func Project_info(c *gin.Context) {
	_, err := c.Cookie("user")
	if err == nil {
		tasks := datamysql.ExtractDataProject_info() //добавил для отрисовки
		//tasks := []model.Task{
		//	model.Task{
		//		Url:             "../project/1/task/1",
		//		Title:           "event1",
		//		Start:           "2022-12-01",
		//		End:             "9999-12-17",
		//		Text:            "Hello world",
		//		BackgroundColor: "#F00",
		//	},
		//	model.Task{
		//		Url:             "../project/1/task/2",
		//		Title:           "event2",
		//		Start:           "2022-12-15",
		//		End:             "2022-12-17",
		//		Text:            "Haahahaahahahahahahaha",
		//		BackgroundColor: "#F00",
		//	},
		//}
		//
		//project_title := "Hello world"
		//project_start := "22-02-2022"
		//project_end := "22-02-2022"
		//project_desc := "Lorem Ipsum hahahahahahahahhaa"

		//projectsData := datamysql.ExtractData_Projects()
		//returningResult := gin.H{
		//	"tasks":         tasks,
		//	"project_title": project_title,
		//	"project_start": project_start,
		//	"project_end":   project_end,
		//	"project_desc":  project_desc,
		//	"id":            id,
		//}
		id := c.Param("id")
		admin, _ := c.Cookie("admin")

		if admin == "true" {
			returningResult["admin"] = true
		}

		c.HTML(200, "project_info.html", returningResult)
	} else {
		c.HTML(200, "login.html", nil)
	}
}

func All_calendar(c *gin.Context) {
	_, err := c.Cookie("user")
	if err == nil {
		// Get data from DB
		returningResult := gin.H{}
		admin, _ := c.Cookie("admin")

		if admin == "true" {
			returningResult["admin"] = true
		}
		c.HTML(200, "all_calendar.html", returningResult)
	} else {
		c.HTML(200, "login.html", nil)
	}

}
func Project_calendar(c *gin.Context) {
	_, err := c.Cookie("user")
	if err == nil {
		// Get data from DB

		id := c.Param("id")
		fmt.Println(id)

		returningResult := gin.H{
			"Id": id,
		}
		admin, _ := c.Cookie("admin")

		if admin == "true" {
			returningResult["admin"] = true
		}

		c.HTML(200, "calendar.html", returningResult)
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
	//fmt.Println(user)                       // структура с json внутри
	fmt.Println(user.Login, " ", user.Pass, " ", user.Id) //провера состояния в постмане
	is_logining := datamysql.ExtractData(datamysql.Db, string(user.Login), string(user.Pass))

	if is_logining.Login == true { //отправка подтверждения логина
		c.SetCookie("user", string(is_logining.Id), 3600, "/", "localhost", false, false)
		returningResult := gin.H{
			"login": "true",
		}
		if is_logining.Admin {
			returningResult["admin"] = true
			c.SetCookie("admin", string("true"), 3600, "/", "localhost", false, false)
		}
		c.JSON(http.StatusOK, returningResult)
	}
	fmt.Println(privat_info.Login)
	if decode != nil {
		c.JSON(http.StatusOK, gin.H{
			"response": decode.Error(),
		})
		return
	}

}
func Task_info(c *gin.Context) {
	_, err := c.Cookie("user")
	if err == nil {

		//persons := []model.Person{
		//	model.Person{
		//		Name: "Синицын Даниил",
		//	},
		//	model.Person{
		//		Name: "Иванов Иван Иванович",
		//	},
		//}
		//all_persons := []model.Person{
		//	model.Person{
		//		Name: "Синицын Даниил",
		//	},
		//	model.Person{
		//		Name: "Иванов Иван Иванович",
		//	},
		//	model.Person{
		//		Name: "HAHAHAHAHAHAHHAHA",
		//	},
		//}

		taskId := c.Param("task_id")

		returningResult := gin.H{
			"title":       "event1",
			"startDate":   "22-02-2022",
			"startEnd":    "22-12-2022",
			"text":        "AHAHHAHAHAHAHAHHAHAHAHAHAHAH",
			"person":      persons,
			"all_persons": all_persons,
			"id":          taskId,
		}
		admin, _ := c.Cookie("admin")

		if admin == "true" {
			returningResult["admin"] = true
		}

		c.HTML(200, "task_info.html", returningResult)
	} else {
		c.HTML(200, "login.html", nil)
	}
}

func Edit_info(c *gin.Context) {
	_, err := c.Cookie("user")
	if err == nil {
		id := c.Param("id")
		//project_title := "Hello world"
		//project_desc := "Lorem Ipsum hahahahahahahahhaa"
		returningResult := gin.H{
			"title": project_title,
			"text":  project_desc,
			"id":    id,
			"color": "#FFA000",
		}
		if admin == "true" {
			returningResult["admin"] = true
		}
		c.HTML(200, "edit_info.html", returningResult)
	} else {
		c.HTML(200, "login.html", nil)
	}
}
func Create_project(c *gin.Context) {
	_, err := c.Cookie("user")
	if err == nil {
		returningResult := gin.H{}
		admin, _ := c.Cookie("admin")

		if admin == "true" {
			returningResult["admin"] = true
		}
		c.HTML(200, "create_project.html", returningResult)
	} else {
		c.HTML(200, "login.html", nil)
	}
}
func Create_task(c *gin.Context) {
	_, err := c.Cookie("user")
	if err == nil {

		//all_persons := []model.Person{
		//	model.Person{
		//		Name: "Синицын Даниил",
		//	},
		//	model.Person{
		//		Name: "Иванов Иван Иванович",
		//	},
		//	model.Person{
		//		Name: "HAHAHAHAHAHAHHAHA",
		//	},
		//}

		id := c.Param("id")
		returningResult := gin.H{
			"all_persons": all_persons,
			"id":          id,
		}
		admin, _ := c.Cookie("admin")
		if admin == "true" {
			returningResult["admin"] = true
		}
		c.HTML(200, "create_task.html", returningResult)
	} else {
		c.HTML(200, "login.html", nil)
	}
}
func Get_all_calendar(c *gin.Context) {

	tasks := []model.Task{
		model.Task{
			Url:             "../project/1/task/1",
			Title:           "event1",
			Start:           "2022-12-01",
			End:             "9999-12-17",
			BackgroundColor: "#F00",
		},
		model.Task{
			Url:             "../project/2/task/1",
			Title:           "event2",
			Start:           "2022-12-15",
			End:             "2022-12-17",
			BackgroundColor: "#0F0",
		},
	}
	c.JSON(http.StatusOK, tasks)
}
func Get_tasks(c *gin.Context) {
	tasks := []model.Task{
		model.Task{
			Url:             "../project/1/task/1",
			Title:           "event1",
			Start:           "2022-12-01",
			End:             "9999-12-17",
			BackgroundColor: "#F00",
		},
		model.Task{
			Url:             "../project/1/task/2",
			Title:           "event2",
			Start:           "2022-12-15",
			End:             "2022-12-17",
			BackgroundColor: "#F00",
		},
	}
	c.JSON(http.StatusOK, tasks)
	return

}
func Person_tasks(c *gin.Context) {
	_, err := c.Cookie("user")
	if err == nil {
		returningResult := gin.H{}
		admin, _ := c.Cookie("admin")
		if admin == "true" {
			returningResult["admin"] = true
		}
		c.HTML(200, "all_tasks.html", returningResult)
	} else {
		c.HTML(200, "login.html", nil)
	}
}

func Update_task(c *gin.Context) {
	// update data base AHAHAHHAHAHAA
}

func Add_project(c *gin.Context) {
	// update data base AHAHAHHAHAHAA
}
func Update_project(c *gin.Context) {
	// update data base AHAHAHHAHAHAA
}
