package main

import (
	// "awesomeProject4/datamysql"
	// "awesomeProject4/model"

	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	//подключение к БД
	// db, err := sql.Open("mysql", "")
	// if err != nil {
	// 	panic(err)
	// }

	// defer db.Close()
	// fmt.Println("DataBase_is_WORK")
	// //запросы к бд
	// datamysql.ExtractData(db)
	// datamysql.ExtractData(db)
	// datamysql.ExtractData(db)

	router := gin.Default()

	router.LoadHTMLGlob("html/*.html")   //шаблоны
	router.Static("/static", "./static") //css, js ...

	router.GET("/", projects_page) //роуты надо прописать в html пути со "static/"
	router.GET("/login", login_page)
	router.GET("/calendar", calendar)
	router.GET("/calendar/:id", project_calendar)
	router.GET("/project/:id", project_info)
	router.GET("/project/:id/task/:task_id", task_info)
	router.GET("/edit/:id/", edit_info)
	router.GET("/edit/", create_project)
	router.GET("/project/:id/task/", create_task)
	router.GET("/new_person", create_login)

	// api
	router.GET("/api/login", registration)
	router.GET("/api/tasks", get_all_calendar)
	router.GET("/api/tasks/:id", get_tasks)
	router.POST("/api/project/:id/task/:task_id", update_task)
	router.POST("/api/project/", add_project)
	router.POST("/api/project/:id", update_project)
	router.POST("/api/new_login", new_login)
	router.Run(":3001")

}

func login_page(c *gin.Context) {
	c.HTML(200, "login.html", nil)
}

type Projects struct {
	Title string

	Text string

	Id int
}

func projects_page(c *gin.Context) {
	var login = true
	if login {
		// Get data from DB
		projectsData := []Projects{
			{Title: "Hello world", Text: "1231", Id: 1},
			{Title: "Some Body", Text: "0933222222", Id: 2},
			{Title: "HAHAHAHAHA", Text: "HAHAHAHAHAH", Id: 3},
			{Title: "HAHAHAHAHA", Text: "HAHAHAHAHAH", Id: 4},
			{Title: "HAHAHAHAHA", Text: "HAHAHAHAHAH", Id: 4},
		}
		c.HTML(200, "projects.html", gin.H{
			"projects": projectsData,
			"admin": true,
		})
	} else {
		// redirect to /login
	}
}

func project_calendar(c *gin.Context) {
	var login = true
	if login {
		// Get data from DB

		id := c.Param("id")
		fmt.Println(id)
		c.HTML(200, "calendar.html", gin.H{
			"Id": id,
		})
	} else {
		// redirect to /login
	}
}

func registration(c *gin.Context) {
	// var user *model.User
	// decode := json.NewDecoder(c.Request.Body).Decode(&user)

	// fmt.Println(json.NewDecoder(c.Request.Body)) //провера состояния в постмане

	// if decode != nil {
	c.JSON(http.StatusOK, gin.H{
		"login": "true",
	})

	// c.HTML(200, "projects.html", nil)
	return
	//  }
	// return
}

type Task struct {
	Url             string `json:"url"`
	Title           string `json:"title"`
	Start           string `json:"start"`
	End             string `json:"end"`
	Text            string
	BackgroundColor string `json:"backgroundColor"`
}

func get_tasks(c *gin.Context) {
	tasks := []Task{
		Task{
			Url:             "../project/1/task/1",
			Title:           "event1",
			Start:           "2022-12-01",
			End:             "9999-12-17",
			BackgroundColor: "#F00",
		},
		Task{
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

func calendar(c *gin.Context) {
	var login = true
	if login {
		// Get data from DB
		c.HTML(200, "all_calendar.html", nil)
	} else {
		// redirect to /login
	}
}

func get_all_calendar(c *gin.Context) {
	tasks := []Task{
		Task{
			Url:             "../project/1/task/1",
			Title:           "event1",
			Start:           "2022-12-01",
			End:             "9999-12-17",
			BackgroundColor: "#F00",
		},
		Task{
			Url:             "../project/2/task/1",
			Title:           "event2",
			Start:           "2022-12-15",
			End:             "2022-12-17",
			BackgroundColor: "#0F0",
		},
	}
	c.JSON(http.StatusOK, tasks)
}

func project_info(c *gin.Context) {
	tasks := []Task{
		Task{
			Url:             "../project/1/task/1",
			Title:           "event1",
			Start:           "2022-12-01",
			End:             "9999-12-17",
			Text:            "Hello world",
			BackgroundColor: "#F00",
		},
		Task{
			Url:             "../project/1/task/2",
			Title:           "event2",
			Start:           "2022-12-15",
			End:             "2022-12-17",
			Text:            "Haahahaahahahahahahaha",
			BackgroundColor: "#F00",
		},
	}

	project_title := "Hello world"
	project_start := "22-02-2022"
	project_end := "22-02-2022"
	project_desc := "Lorem Ipsum hahahahahahahahhaa"
	c.HTML(200, "project_info.html", gin.H{
		"tasks":         tasks,
		"project_title": project_title,
		"project_start": project_start,
		"project_end":   project_end,
		"project_desc":  project_desc,
		"id":            1,
	})

}

type Person struct {
	Name string
}

func task_info(c *gin.Context) {
	persons := []Person{
		Person{
			Name: "Синицын Даниил",
		},
		Person{
			Name: "Иванов Иван Иванович",
		},
	}
	all_persons := []Person{
		Person{
			Name: "Синицын Даниил",
		},
		Person{
			Name: "Иванов Иван Иванович",
		},
		Person{
			Name: "HAHAHAHAHAHAHHAHA",
		},
	}
	c.HTML(200, "task_info.html", gin.H{
		"title":       "event1",
		"startDate":   "22-02-2022",
		"startEnd":    "22-12-2022",
		"text":        "AHAHHAHAHAHAHAHHAHAHAHAHAHAH",
		"person":      persons,
		"all_persons": all_persons,
		"id":          1,
	})
}

func update_task(c *gin.Context) {
	// update data base AHAHAHHAHAHAA
}

func edit_info(c *gin.Context) {
	project_title := "Hello world"
	project_desc := "Lorem Ipsum hahahahahahahahhaa"
	c.HTML(200, "edit_info.html", gin.H{
		"title": project_title,
		"text":  project_desc,
		"id":    1,
		"color": "#FFA000",
	})
}

func update_project(c *gin.Context) {
	// update data base AHAHAHHAHAHAA
}

func create_project(c *gin.Context) {
	c.HTML(200, "create_project.html", nil)
}

func add_project(c *gin.Context) {
	// update data base AHAHAHHAHAHAA
}

func create_task(c *gin.Context) {
	all_persons := []Person{
		Person{
			Name: "Синицын Даниил",
		},
		Person{
			Name: "Иванов Иван Иванович",
		},
		Person{
			Name: "HAHAHAHAHAHAHHAHA",
		},
	}
	c.HTML(200, "create_task.html", gin.H{
		"all_persons": all_persons,
		"id":          1,
	})
}

func create_login(c *gin.Context) {
	c.HTML(200, "create_login.html", gin.H{
		"admin": true,
	})
}


func new_login(c *gin.Context) {
	// Сохраняем данные. 
}
