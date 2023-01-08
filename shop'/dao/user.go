package dao

import (
	"demo/modle"
	"fmt"
)

func Login(m modle.LoginUser) bool {
	k, err := database.Query("select phone,password from user")
	if err != nil {
		fmt.Println(err)
	}
	var j []modle.LoginUser
	for k.Next() {
		var x modle.LoginUser
		err = k.Scan(&x.Phone, &x.Password)
		j = append(j, x)
	}
	for P := 0; P < len(j); P++ {
		if m.Phone == j[P].Phone && m.Password == j[P].Password {
			return true
		}
	}
	return false
}
func Check(user modle.User) bool {
	k, err := database.Query("select phone,name from user")
	if err != nil {
		fmt.Println(err)
	}
	type x struct {
		Name  string `json:"Name"`
		Phone string `json:"Phone"`
	}
	var j []x
	for k.Next() {
		var z x
		_ = k.Scan(&z.Phone, &z.Name)
		j = append(j, z)
	}
	for p := 0; p < len(j); p++ {
		if user.Name == j[p].Name || user.Phone == j[p].Phone {
			return false
		}
	}
	return true
}
func SignUp(user modle.User) {
	_, err := database.Exec("insert into user (name ,password,question,answer,phone)values (?,?,?,?,?)", user.Name, user.Password, user.Question, user.Answer, user.Phone)
	if err != nil {
		fmt.Println(err)
	}
}
