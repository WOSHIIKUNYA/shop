package api

import (
	"demo/modle"
	"demo/service"
	"github.com/gin-gonic/gin"
)

func login(m *gin.Context) {
	var a modle.LoginUser
	m.ShouldBind(&a)
	x := service.Login(a)
	if x == false {
		m.String(200, "false")
		return
	}
	if x == true {
		m.String(200, "true")
		m.SetCookie("phone", a.Phone, 2000, "./", "localhost", false, true)
	}

}
func signup(m *gin.Context) {
	var a modle.User
	m.ShouldBind(&a)
	z := service.Check(a)
	if z == false {
		m.String(200, "false")
		return
	}
	service.SignUp(a)
	m.String(200, "true")
}
