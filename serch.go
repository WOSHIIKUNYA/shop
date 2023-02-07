package api

import (
	"demo/modle"
	"demo/service"
	"github.com/gin-gonic/gin"
)

func mule(m *gin.Context) {
	var name modle.SearchName
	m.ShouldBind(&name)
	Goods := service.SearchMule(name)
	if Goods.Name == "" {
		m.JSON(200, gin.H{
			"result": false,
		})
		return
	}
	m.JSON(200, gin.H{
		"id":      Goods.Id,
		"name":    Goods.Name,
		"kind":    Goods.Kind,
		"price":   Goods.Price,
		"picture": Goods.Picture,
	})
}
func tag(m *gin.Context) {
	var name modle.Tag
	m.ShouldBind(&name)
	x := service.SearchTage(name)
	for b := 0; b < len(x); b++ {
		m.JSON(200, gin.H{
			"id":      x[b].Id,
			"name":    x[b].Name,
			"kind":    x[b].Kind,
			"price":   x[b].Price,
			"picture": x[b].Picture,
		})
	}
}
func SearchPicture(m *gin.Context) {
	var id modle.Id
	m.ShouldBind(&id)
	if id.Id == 0 {
		m.JSON(300, gin.H{
			"result": "入参错误",
		})
		return
	}
	x := service.SearchPicture(id)
	m.JSON(200, gin.H{"picture": x})
}
func SearchPrice(m *gin.Context) {
	var kind modle.Tag
	m.ShouldBind(&kind)
	x := service.SearchPrice(kind)
	for b := 0; b < len(x); b++ {
		m.JSON(200, gin.H{
			"id":      x[b].Id,
			"name":    x[b].Name,
			"kind":    x[b].Kind,
			"price":   x[b].Price,
			"picture": x[b].Picture,
		})
	}
}
