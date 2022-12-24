package model

var U = Registr{"sib_coder", "210104"}

type Registr struct {
	Login string
	Pass  string
}
type User_DB struct {
	FName string `json:"f_name"`
	LName string `json:"l_name"`
}

type Projects struct {
	Title string

	Text string

	Id int
}

type Task struct {
	Url             string `json:"url"`
	Title           string `json:"title"`
	Start           string `json:"start"`
	End             string `json:"end"`
	Text            string
	BackgroundColor string `json:"backgroundColor"`
}

type Person struct {
	Name string
}
