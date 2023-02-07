package dao

import (
	"demo/modle"
	"fmt"
)

func Login(m modle.Login) bool {
	k, err := database.Query("select Password,phone from user")
	if err != nil {
		fmt.Println(err)
	}
	var j []modle.Login
	for k.Next() {
		var x modle.Login
		err = k.Scan(&x.Password, &x.Phone)
		if err != nil {
			fmt.Println(err)
		}
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
func Seek(b modle.Phone) modle.Seek {
	x, err := database.Query("select phone,question,answer from user")
	if err != nil {
		fmt.Println(err)
	}
	var z []modle.Seek
	for x.Next() {
		var c modle.Seek
		x.Scan(&c.Phone, &c.Question, &c.Answer)
		z = append(z, c)
	}
	for a := 0; a < len(z); a++ {
		if z[a].Phone == b.Phone {
			return z[a]
			fmt.Println("good")
		}
	}
	var v modle.Seek
	return v
}
func Reset(m modle.LoginUser) {
	_, err := database.Exec("update user set password=? where phone=?", m.Password, m.Phone)
	if err != nil {
		fmt.Println(err)
	}
}
