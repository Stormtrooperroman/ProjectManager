package model

type Registr struct {
	Id    int
	Login string
	Pass  string
}
type User_DB struct {
	Id    int    `json:"id"`
	FName string `json:"f_name"`
	LName string `json:"l_name"`
	Admin int    `json:"admin"`
}

type Projects struct {
	Title string

	Text string

	Id int
}

type Task struct {
	Id              string `json:"id"`
	Url             string `json:"url"`
	Title           string `json:"title"`
	Start           string `json:"start"`
	End             string `json:"end"`
	Text            string
	BackgroundColor string `json:"backgroundColor"`
	Project_name    string
	Project_id      string `json:"project_Id"`
}

type Person struct {
	Name string
}
type Is_login struct {
	Login bool
	Admin bool
	Id    int
}
