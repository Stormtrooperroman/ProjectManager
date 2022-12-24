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
	if privat_info.Admin == true {
		// Get data from DB
		projectsData := []model.Projects{
			{Title: "Hello world", Text: "1231", Id: 1},
			{Title: "Some Body", Text: "0933222222", Id: 2},
			{Title: "HAHAHAHAHA", Text: "HAHAHAHAHAH", Id: 3},
			{Title: "HAHAHAHAHA", Text: "HAHAHAHAHAH", Id: 4},
			{Title: "HAHAHAHAHA", Text: "HAHAHAHAHAH", Id: 4},
		}
		c.HTML(200, "projects.html", gin.H{
			"projects": projectsData,
		})
	} else {
		c.HTML(200, "login.html", nil)
	}
}

func Project_info(c *gin.Context) {
	if privat_info.Admin == true {
		tasks := []model.Task{
			model.Task{
				Url:             "../project/1/task/1",
				Title:           "event1",
				Start:           "2022-12-01",
				End:             "9999-12-17",
				Text:            "Hello world",
				BackgroundColor: "#F00",
			},
			model.Task{
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
	} else {
		c.HTML(200, "login.html", nil)
	}
}

func All_calendar(c *gin.Context) {
	if privat_info.Admin == true {
		// Get data from DB
		c.HTML(200, "all_calendar.html", nil)
	} else {
		c.HTML(200, "login.html", nil)
	}

}
func Project_calendar(c *gin.Context) {
	if privat_info.Admin == true {
		// Get data from DB

		id := c.Param("id")
		fmt.Println(id)
		c.HTML(200, "calendar.html", gin.H{
			"Id": id,
		})
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
func Task_info(c *gin.Context) {
	persons := []model.Person{
		model.Person{
			Name: "Синицын Даниил",
		},
		model.Person{
			Name: "Иванов Иван Иванович",
		},
	}
	all_persons := []model.Person{
		model.Person{
			Name: "Синицын Даниил",
		},
		model.Person{
			Name: "Иванов Иван Иванович",
		},
		model.Person{
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

func Edit_info(c *gin.Context) {
	project_title := "Hello world"
	project_desc := "Lorem Ipsum hahahahahahahahhaa"
	c.HTML(200, "edit_info.html", gin.H{
		"title": project_title,
		"text":  project_desc,
		"id":    1,
		"color": "#FFA000",
	})
}
func Create_project(c *gin.Context) {
	c.HTML(200, "create_project.html", nil)
}
func Create_task(c *gin.Context) {
	all_persons := []model.Person{
		model.Person{
			Name: "Синицын Даниил",
		},
		model.Person{
			Name: "Иванов Иван Иванович",
		},
		model.Person{
			Name: "HAHAHAHAHAHAHHAHA",
		},
	}
	c.HTML(200, "create_task.html", gin.H{
		"all_persons": all_persons,
		"id":          1,
	})
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
func Update_task(c *gin.Context) {
	// update data base AHAHAHHAHAHAA
}

func Add_project(c *gin.Context) {
	// update data base AHAHAHHAHAHAA
}
func Update_project(c *gin.Context) {
	// update data base AHAHAHHAHAHAA
}

////////////////////////////////////модели от Ромы///////////////////////////////////////////////////////////////////////
