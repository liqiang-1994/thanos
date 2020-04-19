package V1

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/liqiang/thanos/model"
)

func HelloHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	user, _ := c.Get("id")
	c.JSON(200, gin.H{
		"userId":   claims["id"],
		"userName": user.(*model.User).UserName,
		"text":     "Hello WOlr",
	})
}
