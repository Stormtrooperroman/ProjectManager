package pages

import (
	"awesomeProject4/datamysql"
	"awesomeProject4/model"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
		c.Redirect(301, "/login")
	}
}

func Project_info(c *gin.Context) {
	_, err := c.Cookie("user")
	if err == nil {
		id := c.Param("id")
		tasks := datamysql.ExtractDataProject_info(id) //добавил для отрисовки
		project_info := datamysql.ExtractDataProject(id)
		for i := 0; i < len(tasks); i++ {
			tasks[i].Url = "../project/" + string(id) + "/task/" + string(tasks[i].Id)
		}

		returningResult := gin.H{
			"tasks":         tasks,
			"project_title": project_info.Title,
			"project_start": project_info.Start,
			"project_end":   project_info.End,
			"project_desc":  project_info.Text,
			"id":            id,
		}

		admin, _ := c.Cookie("admin")

		if admin == "true" {
			returningResult["admin"] = true
		}

		c.HTML(200, "project_info.html", returningResult)
	} else {
		c.Redirect(301, "/login")
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
		c.Redirect(301, "/login")
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
		c.Redirect(301, "/login")
	}
}

func Login_page(c *gin.Context) {
	c.HTML(200, "login.html", nil)
}

func Registration(c *gin.Context) {
	var user *model.Registr
	decode := json.NewDecoder(c.Request.Body).Decode(&user)
	is_logining := datamysql.ExtractData(datamysql.Db, string(user.Login), string(user.Pass))
	user_id := strconv.Itoa(is_logining.Id)

	if is_logining.Login == true { //отправка подтверждения логина
		c.SetCookie("user", user_id, 3600, "/", "localhost", false, false)
		returningResult := gin.H{
			"login": "true",
		}
		if is_logining.Admin {
			returningResult["admin"] = true
			c.SetCookie("admin", string("true"), 3600, "/", "localhost", false, false)
		}
		c.JSON(http.StatusOK, returningResult)
	}
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
		taskId := c.Param("task_id")
		all_persons := datamysql.ExtractDataUsers()
		task := datamysql.ExtractDataTask(taskId)
		persons := datamysql.ExtractDataUsers_Task(taskId)

		returningResult := gin.H{
			"title":       task.Title,
			"startDate":   task.Start,
			"endDate":     task.End,
			"text":        task.Text,
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
		c.Redirect(301, "/login")
	}
}

func Edit_info(c *gin.Context) {
	_, err := c.Cookie("user")
	if err == nil {
		id := c.Param("id")
		admin, _ := c.Cookie("admin")
		project_info := datamysql.ExtractDataProject(id)
		returningResult := gin.H{
			"title":      project_info.Title,
			"text":       project_info.Text,
			"id":         id,
			"color":      project_info.BackgroundColor,
			"text_color": project_info.TextColor,
		}
		if admin == "true" {
			returningResult["admin"] = true
		}
		c.HTML(200, "edit_info.html", returningResult)
	} else {
		c.Redirect(301, "/login")
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
		c.Redirect(301, "/login")
	}
}
func Create_task(c *gin.Context) {
	_, err := c.Cookie("user")
	if err == nil {
		all_persons := datamysql.ExtractDataUsers()

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
		c.Redirect(301, "/login")
	}
}
func Get_all_calendar(c *gin.Context) {
	tasks := datamysql.ExtractDataProject_info_ALL()
	for i := 0; i < len(tasks); i++ {
		tasks[i].Url = "../project/" + string(tasks[i].Project_id) + "/task/" + string(tasks[i].Id)
	}

	c.JSON(http.StatusOK, tasks)
}

func Get_tasks(c *gin.Context) {
	id := c.Param("id")
	tasks := datamysql.ExtractDataProject_info(id)
	for i := 0; i < len(tasks); i++ {
		tasks[i].Url = "../project/" + string(id) + "/task/" + string(tasks[i].Id)
	}

	c.JSON(http.StatusOK, tasks)
	return

}
func Person_tasks(c *gin.Context) {
	user, err := c.Cookie("user")
	if err == nil {
		tasks := datamysql.ExtractDataProject_and_Task(user)
		for i := 0; i < len(tasks); i++ {
			tasks[i].Url = "../project/" + string(tasks[i].Project_id) + "/task/" + string(tasks[i].Id)
		}
		returningResult := gin.H{"tasks": tasks}
		admin, _ := c.Cookie("admin")
		if admin == "true" {
			returningResult["admin"] = true
		}
		c.HTML(200, "all_tasks.html", returningResult)
	} else {
		c.Redirect(301, "/login")
	}
}

func NewPerson(c *gin.Context) {
	_, err := c.Cookie("user")
	if err == nil {
		returningResult := gin.H{}
		admin, _ := c.Cookie("admin")
		if admin == "true" {
			returningResult["admin"] = true
			c.HTML(200, "create_login.html", returningResult)
		} else {
			c.HTML(200, "projects.html", nil)
		}

	} else {
		c.Redirect(301, "/login")
	}
}
func CreateUser(c *gin.Context) {
	var user *model.CreateUser
	decode := json.NewDecoder(c.Request.Body).Decode(&user)
	fmt.Println(user.FName, " ", user.LName, " ", user.Login, " ", user.Password)
	if decode != nil {
		c.JSON(http.StatusOK, gin.H{
			"response": decode.Error(),
		})
	}
	datamysql.AddPerson(user.Login, user.Password, user.FName, user.LName)
}
func CreateProject(c *gin.Context) {
	var project *model.CreateProject
	decode := json.NewDecoder(c.Request.Body).Decode(&project)
	fmt.Println(project.Name, " ", project.Description, " ", project.Colour)
	if decode != nil {
		c.JSON(http.StatusOK, gin.H{
			"response": decode.Error(),
		})
	}
	datamysql.AddProject(project.Name, project.Description, project.Colour, project.TextColor)
}

func CreateTask(c *gin.Context) {
	var task *model.Task
	id := c.Param("id")
	decode := json.NewDecoder(c.Request.Body).Decode(&task)
	//fmt.Println(project.Name, " ", project.Description, " ", project.Colour)
	if decode != nil {
		c.JSON(http.StatusOK, gin.H{
			"response": decode.Error(),
		})
	}
	datamysql.AddTask(task.Title, task.Start, task.End, task.Text, id, task.Person_Mas)
}

func Update_project(c *gin.Context) {
	// update data base AHAHAHHAHAHAA
	id := c.Param("id")
	var project *model.CreateProject
	decode := json.NewDecoder(c.Request.Body).Decode(&project)
	fmt.Println(project.Name, " ", project.Description, " ", project.Colour)
	if decode != nil {
		c.JSON(http.StatusOK, gin.H{
			"response": decode.Error(),
		})
	}
	datamysql.UpdateProject(id, project.Name, project.Description, project.Colour, project.TextColor)
}

func Update_task(c *gin.Context) {
	var task *model.Task
	id := c.Param("task_id")
	decode := json.NewDecoder(c.Request.Body).Decode(&task)
	//fmt.Println(project.Name, " ", project.Description, " ", project.Colour)
	if decode != nil {
		c.JSON(http.StatusOK, gin.H{
			"response": decode.Error(),
		})
	}
	datamysql.UpdateTask(id, task.Title, task.Start, task.End, task.Text, task.Person_Mas)
}
