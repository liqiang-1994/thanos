package main

import (
	"github.com/gin-gonic/gin"
)



func main() {
	r := gin.Default()
	r.GET("/echo/:name", func(c *gin.Context) {
		age := c.Query("age")
		name :=c.Param("name")
		c.JSON(200, gin.H{
			"OK": "Hi" + name + age,
		})
	})
	r.Run()
}