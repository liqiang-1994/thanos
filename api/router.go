package api

import (
	"github.com/gin-gonic/gin"
	"github.com/liqiang/thanos/api/http/V1"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	var auth = V1.Auth()
	r.POST("/login", auth.LoginHandler)
	apiV1 := r.Group("/api")
	apiV1.GET("/refresh_token", auth.RefreshHandler)
	apiV1.Use(auth.MiddlewareFunc())
	{
		//add api
		apiV1.GET("hello", V1.HelloHandler)
	}

	return r
}
