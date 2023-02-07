package api

import (
	"demo/modle"
	"demo/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func login(m *gin.Context) {
	var a modle.Login
	m.ShouldBind(&a)
	x := service.Login(a)
	if x == false {
		m.JSON(200, gin.H{
			"result": false,
		})
		return
	} else {
		m.SetCookie("Phone", a.Phone, 36000, "/", "43.139.37.159", false, true)
		m.JSON(200, gin.H{
			"result": true,
			"Phone":  a.Phone,
		})
	}
}
func signup(m *gin.Context) {
	var a modle.User
	m.ShouldBind(&a)
	z := service.Check(a)
	if a.Name == "" {
		m.String(200, "密码没有")
		return
	}
	if a.Phone == "" {
		m.String(200, "电话没有")
		return
	}
	if a.Password == "" {
		m.String(200, "密码没有")
		return
	}
	if a.Answer == "" {
		m.String(200, "密码问题没有")
		return
	}
	if a.Question == "" {
		fmt.Println(200, "密保问题没有")
		return
	}
	if z == false {
		m.JSON(200, gin.H{"result": false})
		return
	}
	service.SignUp(a)
	m.JSON(200, gin.H{"result": true})
}
func CheckLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("Phone")
		if err != nil {
			fmt.Println(err)
		}
		if cookie != "" {
			c.Next()
			return
		}
		c.JSON(http.StatusUnauthorized, gin.H{"error": "err"})
		return
	}
}
func SeekLogin(m *gin.Context) {
	var a modle.Phone
	m.ShouldBind(&a)
	if a.Phone == "" {
		m.JSON(300, gin.H{
			"result": "入参错误",
		})
		return
	}
	fmt.Println(a.Phone)
	x := service.Seek(a)
	if x.Phone == "" {
		m.JSON(200, gin.H{
			"phone": false,
		})
		return
	}
	m.SetCookie("Answer", x.Answer, 3600, "/", "43.139.37.159", false, true)
	m.SetCookie("Question", x.Question, 3600, "/", "43.139.37.159", false, true)
	m.JSON(200, gin.H{
		"phone": true,
	})
}
func answer(m *gin.Context) {
	r, err := m.Cookie("Answer")
	if err != nil {
		fmt.Println(err)
	}
	x, err1 := m.Cookie("Question")
	if err1 != nil {
		fmt.Println(err)
	}
	var Answer modle.Answer
	m.ShouldBind(&Answer)
	if Answer.Answer == "" {
		m.JSON(200, gin.H{
			"Question": x,
		})
		return
	} else {
		if Answer.Answer == r {
			m.JSON(200, gin.H{
				"result": true,
			})
			return
		} else {
			m.JSON(200, gin.H{
				"result": false,
			})
		}
	}
}
func Reset(m *gin.Context) {
	var a modle.LoginUser
	m.ShouldBind(&a)
	service.Reset(a)
	m.JSON(200, gin.H{
		"result": true,
	})
}
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method               //请求方法
		origin := c.Request.Header.Get("Origin") //请求头部
		var headerKeys []string                  // 声明请求头keys
		for k, _ := range c.Request.Header {
			headerKeys = append(headerKeys, k)
		}
		headerStr := strings.Join(headerKeys, ", ")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}
		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Origin", "*")                                       // 这是允许访问所有域
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE") //服务器支持的所有跨域请求的方法,为了避免浏览次请求的多次'预检'请求
			//  header的类型
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			//				允许跨域设置																										可以返回其他子段
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar") // 跨域关键设置 让浏览器可以解析
			c.Header("Access-Control-Max-Age", "172800")                                                                                                                                                           // 缓存请求信息 单位为秒
			c.Header("Access-Control-Allow-Credentials", "true")                                                                                                                                                   //	跨域请求是否需要带cookie信息 默认设置为true
			c.Set("content-type", "application/json")                                                                                                                                                              // 设置返回格式是json
		}

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}
		// 处理请求
		c.Next()
	}
}
