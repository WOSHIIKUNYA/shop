package api

import (
	"github.com/gin-gonic/gin"
)

func User() {
	m := gin.Default()
	m.Use(Cors())
	m.MaxMultipartMemory = 8 << 20
	r := m.Group("/search")
	{
		r.POST("/mule", mule)
		r.POST("/picture", SearchPicture)
		r.POST("price", SearchPrice)
		r.POST("/tag", tag)
	}
	x := m.Group("/user")
	{
		x.POST("/login", login)
		x.POST("/signup", signup)
		x.POST("/seek", SeekLogin)
		x.POST("answer", answer)
		x.PUT("/reset", Reset)
	}
	y := m.Group("/shop", CheckLogin())
	{
		y.POST("/sell", sell)
		y.POST("/add", AddCart)
		y.DELETE("/delete", DeleteCart)
		y.DELETE("/clear", ClearCart)
	}
	z := m.Group("/comment", CheckLogin())
	{
		z.POST("/send", send)
		z.POST("/look", look)
		z.DELETE("/delete", Delete)
	}
	p := m.Group("/person", CheckLogin())
	{
		p.PUT("/recharge", recharge)
		p.PUT("/changeaddress", ChangeAddress)
		p.PUT("/change", change)
	}
	m.Run()
}
