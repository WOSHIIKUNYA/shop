package api

import (
	"demo/modle"
	"demo/service"
	"fmt"
	"github.com/gin-gonic/gin"
)

func See(m *gin.Context) {

}
func recharge(m *gin.Context) {
	var money modle.Recharge
	m.ShouldBind(&money)
	x := service.Recharge(money)
	m.JSON(200, gin.H{
		"result": "ok",
		"money":  x,
	})
}
func ChangeAddress(m *gin.Context) {
	var x modle.ChangeAddress
	m.ShouldBind(&x)
	service.ChangeAddress(x)
	m.JSON(200, gin.H{
		"result":  "ok",
		"address": x.Address,
	})
}
func change(m *gin.Context) {
	var z modle.Change
	m.ShouldBind(&z)
	x, err := m.FormFile("Picture")
	if err != nil {
		fmt.Println(err)
	}
	err = m.SaveUploadedFile(x, "./"+x.Filename)
	if err != nil {
		fmt.Println(err)
	}
	service.Change(z, x.Filename)
	m.JSON(200, gin.H{
		"result": true,
	})
}
