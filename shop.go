package api

import (
	"demo/modle"
	"demo/service"
	"github.com/gin-gonic/gin"
)

func sell(m *gin.Context) {
	var good modle.Good
	m.ShouldBind(&good)
	service.Sell(good, good.Picture)
	m.JSON(200, gin.H{
		"result": "yes",
	})
}
func AddCart(m *gin.Context) {
	var x modle.AddCart
	m.ShouldBind(&x)
	f := service.AddCart(x)
	m.JSON(200, gin.H{
		"id":      f.Id,
		"name":    f.Name,
		"price":   f.Price,
		"picture": f.Picture,
		"phone":   f.Phone,
	})
}
func DeleteCart(m *gin.Context) {
	var Id modle.DeleteCart
	m.ShouldBind(&Id)
	service.DeleteCart(Id.Id)
	m.JSON(200, gin.H{
		"result": true,
	})
}
func ClearCart(m *gin.Context) {
	var x modle.ClearCart
	m.ShouldBind(&x)
	z := service.ClearCart(x)
	if z == false {
		m.JSON(200, gin.H{
			"result": false,
			"reason": "have not enough money",
		})
		return
	}
	m.JSON(200, gin.H{
		"result": true,
	})
}
