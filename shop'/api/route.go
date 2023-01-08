package api

import "github.com/gin-gonic/gin"

func User() {
	m := gin.Default()
	x := m.Group("/user")
	{
		x.GET("/login", login)
		x.POST("/signup", signup)
	}
	m.Run()
}
