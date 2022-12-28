package datamysql

import (
	"awesomeProject4/model"
	"awesomeProject4/privat_info"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

func Conect() {
	db, err := sqlx.Open("mysql", privat_info.DataBaseKey)
	if err != nil {
		panic(err)
	}
	Db = db
	//defer Db.Close() //Закрытие БД
	fmt.Println("DataBase_is_WORK")
}

func ExtractData(db *sqlx.DB, login string, password string) model.Is_login { //получение пользователя из бд
	var u model.User_DB
	res, err := db.Query("SELECT `first_name`,`last_name`,`id`, `is_admin` FROM `employees` WHERE `login`= ? AND `password` = ?;", login, password)
	if err != nil {
		panic(err)
	}
	for res.Next() {
		err = res.Scan(&u.FName, &u.LName, &u.Id, &u.Admin)
		if err != nil {
			panic(err)
		}
		//fmt.Println(fmt.Sprintf("in database have %s , %s ", u.FName, u.LName))
	}
	result := model.Is_login{
		Login: false,
		Admin: false,
	}
	fmt.Println(u.Id)
	if u.LName != "" && u.FName != "" {
		//privat_info.Login = true
		result.Login = true
		result.Id = u.Id
		if u.Admin == 1 {
			result.Admin = true
		}
	} else {
		//privat_info.Login = false
	}

	return result
	//fmt.Println(u.LName, " ", u.FName) //пример как вырывать параметры из запроса

}

func AddData(db *sqlx.DB) {
	result, err := db.Exec("insert into employees (id, first_name,last_name,post) values (?,?, ?, ?)", "8", "fuck", "fuck", "fuck")
	if err != nil {
		panic(err)
	}
	fmt.Println(result.LastInsertId()) // id добавленного объекта
	fmt.Println(result.RowsAffected()) // количество затронутых строк
}

func DelData(db *sqlx.DB) {
	result, err := db.Exec("delete from employees where id = ?", "8")
	if err != nil {
		panic(err)
	}
	fmt.Println(result.LastInsertId()) // id последнего удаленого объекта
	fmt.Println(result.RowsAffected()) // количество затронутых строк

}
func ExtractData_Projects() []model.Projects {

	var data model.Projects
	var data_mas []model.Projects
	res, err := Db.Query("SELECT `name`,`id`, `description` FROM `projects` ;")
	if err != nil {
		panic(err)
	}
	for res.Next() {
		err = res.Scan(&data.Title, &data.Id, &data.Text)
		if err != nil {
			panic(err)
		}
		//fmt.Println(fmt.Sprintf("in database have %s , %d, %s ", data.Title, data.Id, data.Text))
		data_mas = append(data_mas, data)
	}
	fmt.Println(data_mas)
	fmt.Println(data.Title, " ", data.Id, " ", data.Text) //пример как вырывать параметры из запроса
	return data_mas

}

func ExtractDataProject_info(id string) []model.Task { //получение пользователя из бд
	var u model.Task
	var u_mas []model.Task
	res, err := Db.Query("SELECT tasks.id, tasks.name, tasks.start_date, tasks.end_date, tasks.description, projects.colour, projects.text_colour FROM tasks INNER JOIN projects ON tasks.project_id = projects.id WHERE tasks.project_id = ? ;", id)
	if err != nil {
		panic(err)
	}
	for res.Next() {
		err = res.Scan(&u.Id, &u.Title, &u.Start, &u.End, &u.Text, &u.BackgroundColor, &u.TextColor)
		if err != nil {
			panic(err)
		}
		//fmt.Println(fmt.Sprintf("in database have %s , %s ", u.FName, u.LName))
		u_mas = append(u_mas, u)
	}
	fmt.Println(u_mas)
	fmt.Println(u.Id, " ", u.Title, " ", u.Start, " ", u.End, " ", u.BackgroundColor, " ", u.TextColor) //пример как вырывать параметры из запроса
	return u_mas

}
func ExtractDataProject(id string) model.Task { //получение пользователя из бд
	var u model.Task

	res, err := Db.Query("SELECT  id ,name, start_date, end_date,description, colour, text_colour FROM projects WHERE id = ? ;", id)
	if err != nil {
		panic(err)
	}
	for res.Next() {
		err = res.Scan(&u.Id, &u.Title, &u.Start, &u.End, &u.Text, &u.BackgroundColor, &u.TextColor)
		if err != nil {
			panic(err)
		}
		//fmt.Println(fmt.Sprintf("in database have %s , %s ", u.FName, u.LName))
	}
	fmt.Println(u)
	fmt.Println(u.Id, " ", u.Title, " ", u.Start, " ", u.End, " ", u.BackgroundColor, " ", u.TextColor) //пример как вырывать параметры из запроса
	return u

}
func ExtractDataTask(id string) model.Task { //получение пользователя из бд
	var u model.Task

	res, err := Db.Query("SELECT  id ,name, start_date, end_date,description FROM tasks WHERE id = ? ;", id)
	if err != nil {
		panic(err)
	}
	for res.Next() {
		err = res.Scan(&u.Id, &u.Title, &u.Start, &u.End, &u.Text)
		if err != nil {
			panic(err)
		}
		//баг с Text надо проверять значение столбца
		//fmt.Println(fmt.Sprintf("in database have %s , %s ", u.FName, u.LName))
	}
	fmt.Println(u)
	fmt.Println(u.Id, " ", u.Title, " ", u.Start, " ", u.End, " ") //пример как вырывать параметры из запроса
	return u

}
func ExtractDataUsers() []model.User_DB { //получение пользователя из бд
	var u model.User_DB
	var users []model.User_DB
	res, err := Db.Query("SELECT `id`, `first_name`,`last_name` FROM `employees`;")
	if err != nil {
		panic(err)
	}
	for res.Next() {
		err = res.Scan(&u.Id, &u.FName, &u.LName)
		if err != nil {
			panic(err)
		}
		//fmt.Println(fmt.Sprintf("in database have %s , %s ", u.FName, u.LName))
		users = append(users, u)
	}
	fmt.Println(users)
	fmt.Println(u.LName, " ", u.FName) //пример как вырывать параметры из запроса
	return users
}
func ExtractDataUsers_Task(id string) []model.User_DB { //получение пользователя из бд
	var u model.User_DB
	var users []model.User_DB
	res, err := Db.Query("SELECT `first_name`,`last_name`, `id` FROM `employees` where id in (select emp_id from task_for_emp where task_id = ?);", id)
	if err != nil {
		panic(err)
	}
	for res.Next() {
		err = res.Scan(&u.FName, &u.LName, &u.Id)
		if err != nil {
			panic(err)
		}
		//fmt.Println(fmt.Sprintf("in database have %s , %s ", u.FName, u.LName))
		users = append(users, u)
	}
	fmt.Println(users)
	fmt.Println(u.LName, " ", u.FName) //пример как вырывать параметры из запроса
	return users
}
func ExtractDataProject_info_ALL() []model.Task { //получение пользователя из бд
	var u model.Task
	var u_mas []model.Task
	res, err := Db.Query("SELECT tasks.id ,tasks.name,tasks.start_date,tasks.end_date, projects.colour, projects.text_colour, projects.id FROM tasks  INNER JOIN projects ON tasks.project_id = projects.id ;")
	if err != nil {
		panic(err)
	}
	for res.Next() {
		err = res.Scan(&u.Id, &u.Title, &u.Start, &u.End, &u.BackgroundColor, &u.TextColor, &u.Project_id)
		if err != nil {
			panic(err)
		}
		//fmt.Println(fmt.Sprintf("in database have %s , %s ", u.FName, u.LName))
		u_mas = append(u_mas, u)
	}
	fmt.Println(u_mas)
	fmt.Println(u.Id, " ", u.Title, " ", u.Start, " ", u.End, " ", u.BackgroundColor) //пример как вырывать параметры из запроса
	return u_mas

}
func ExtractDataProject_and_Task(id string) []model.Task { //получение пользователя из бд
	var u model.Task
	var u_mas []model.Task
	fmt.Println(id)
	res, err := Db.Query("SELECT tasks.id ,tasks.name,tasks.start_date,tasks.end_date, projects.colour,projects.name,projects.id  FROM tasks  INNER JOIN projects ON tasks.project_id = projects.id WHERE tasks.id in (select task_id from task_for_emp where emp_id =?) ;", id)
	if err != nil {
		panic(err)
	}
	for res.Next() {
		err = res.Scan(&u.Id, &u.Title, &u.Start, &u.End, &u.BackgroundColor, &u.Project_name, &u.Project_id)
		if err != nil {
			panic(err)
		}
		//fmt.Println(fmt.Sprintf("in database have %s , %s ", u.FName, u.LName))
		u_mas = append(u_mas, u)
	}
	fmt.Println(u_mas)
	fmt.Println(u.Id, " ", u.Title, " ", u.Start, " ", u.End, " ", u.BackgroundColor) //пример как вырывать параметры из запроса
	return u_mas

}
func AddPerson(login string, password string, first_name string, last_name string) {
	result, err := Db.Exec("insert into employees (login, password, first_name,last_name, is_admin ) values (?, ?, ?,?, 0)", login, password, first_name, last_name)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.LastInsertId()) // id добавленного объекта
	fmt.Println(result.RowsAffected()) // количество затронутых строк
}

func AddProject(name string, description string, colour string, text_colour string) {
	result, err := Db.Exec("insert into projects (name, description, colour , start_date, end_date, text_colour) values (?, ?, ?, sysdate(), sysdate(), ?)", name, description, colour, text_colour)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.LastInsertId()) // id добавленного объекта
	fmt.Println(result.RowsAffected()) // количество затронутых строк
}
func AddTask(name string, start_date string, end_date string, description string, id_project string, person []string) {
	result, err := Db.Exec("insert into tasks (name, start_date, end_date, description, project_id) values (?, STR_TO_DATE(?, '%Y-%m-%d'), STR_TO_DATE(?, '%Y-%m-%d'), ?, ?)", name, start_date, end_date, description, id_project)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.LastInsertId()) // id добавленного объекта
	fmt.Println(result.RowsAffected()) // количество затронутых строк

	var id_task string

	res, err := Db.Query("SELECT MAX(`id`) FROM tasks ;")
	if err != nil {
		panic(err)
	}
	for res.Next() {
		err = res.Scan(&id_task)
		if err != nil {
			panic(err)
		}
	}
	fmt.Println(id_task)
	for i := 0; i < len(person); i++ {
		result1, err1 := Db.Exec("insert into task_for_emp (task_id, emp_id) values (?,?)", id_task, person[i])
		if err1 != nil {
			panic(err)
		}
		fmt.Println(result1.LastInsertId()) // id добавленного объекта
		fmt.Println(result1.RowsAffected()) // количество затронутых строк
	}

}
func UpdateProject(id string, name string, description string, colour string, text_colour string) {
	result, err := Db.Exec("update  projects set name = ?, description = ?, colour = ?, text_colour = ? WHERE id = ?", name, description, colour, text_colour, id)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.LastInsertId()) // id добавленного объекта
	fmt.Println(result.RowsAffected()) // количество затронутых строк
}
func UpdateTask(id string, name string, start_date string, end_date string, description string, person []string) {
	result, err := Db.Exec("update  tasks set name=?, start_date=STR_TO_DATE(?, '%Y-%m-%d'), end_date=STR_TO_DATE(?, '%Y-%m-%d'), description =? where id = ?", name, start_date, end_date, description, id)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.LastInsertId()) // id добавленного объекта
	fmt.Println(result.RowsAffected()) // количество затронутых строк

	id_task := id

	res1, err1 := Db.Query("DELETE FROM task_for_emp WHERE task_id =? ;", id)
	if err1 != nil {
		panic(err)
	}
	for res1.Next() {
		err1 = res1.Scan(&id_task)
		if err1 != nil {
			panic(err)
		}
	}

	fmt.Println(id_task)
	for i := 0; i < len(person); i++ {
		result2, err2 := Db.Exec("insert into task_for_emp (task_id, emp_id) values (?,?)", id_task, person[i])
		if err2 != nil {
			panic(err)
		}
		fmt.Println(result2.LastInsertId()) // id добавленного объекта
		fmt.Println(result2.RowsAffected()) // количество затронутых строк
	}

}
