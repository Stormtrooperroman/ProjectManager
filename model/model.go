package model

type User struct {
	Id    uint
	Login string
	Pass  string
}
type User_DB struct {
	FName string `json:"f_name"`
	LName string `json:"l_name"`
	Post  string `json:"post"`
}
