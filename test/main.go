package main

import "github.com/gin-gonic/gin"

func main() {
	e := gin.Default()
	e.GET("/hello/:id", func(c *gin.Context) {
		c.Writer.WriteString("/hello/:id")
	})
	e.GET("/hello/:id/names", func(c *gin.Context) {
		c.Writer.WriteString("/hello/:id/names")
	})
	e.GET("/hello/names", func(c *gin.Context) {
		c.Writer.WriteString("/hello/names")
	})
	e.Run(":10086")
}
