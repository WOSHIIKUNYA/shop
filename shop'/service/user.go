package service

import (
	"demo/dao"
	"demo/modle"
)

func Login(user modle.LoginUser) bool {
	a := dao.Login(user)
	return a
}
func Check(user modle.User) bool {
	a := dao.Check(user)
	return a
}
func SignUp(user modle.User) {
	dao.SignUp(user)
}
