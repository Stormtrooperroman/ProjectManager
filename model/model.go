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
