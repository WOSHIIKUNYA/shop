package api

import (
	"demo/modle"
	"demo/service"
	"github.com/gin-gonic/gin"
)

func send(m *gin.Context) {
	var x modle.Comment
	m.ShouldBind(&x)
	service.SendComment(x)
	m.JSON(200, gin.H{
		"result": "ok",
	})
}
func look(m *gin.Context) {
	var x modle.GoodId
	m.ShouldBind(&x)
	a := service.LookComment(x)
	for b := 0; b < len(a); b++ {
		m.JSON(200, gin.H{
			"id":     a[b].Id,
			"writer": a[b].Name,
			"score":  a[b].Score,
			"body":   a[b].Body,
		})
	}
}
func Delete(m *gin.Context) {
	var x modle.Writer
	m.ShouldBind(&x)
	z := service.Delete(x)
	if z == false {
		m.JSON(200, gin.H{
			"name": false,
		})
		return
	}
	m.JSON(200, gin.H{
		"name ": true,
	})
}
