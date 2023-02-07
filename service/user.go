package service

import (
	"demo/dao"
	"demo/modle"
)

func Login(x modle.Login) bool {
	a := dao.Login(x)
	return a
}
func Check(user modle.User) bool {
	a := dao.Check(user)
	return a
}
func SignUp(user modle.User) {
	dao.SignUp(user)
}
func Seek(a modle.Phone) modle.Seek {
	m := dao.Seek(a)
	return m
}
func Reset(a modle.LoginUser) {
	dao.Reset(a)
}
