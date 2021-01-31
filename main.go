package main

import "github.com/gin-gonic/gin"

type User struct {
	ID   uint64
	Name string
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"name":"gin-note",
		})
	})

	// user
	users := []User{{ID:1,Name:"tom"}, {ID:2,Name:"tim"}}
	r.GET("/users", func(c *gin.Context) {
		c.JSON(200, users)
	})

	// /echo/123
	r.GET("/echo/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.String(200, id)
	})

	// /echoall/12/12
	r.GET("/echoall/*id", func(c *gin.Context) {
		id := c.Param("id")
		c.String(200, id)
	})

	// /q?q=aaa
	r.GET("/q", func(c *gin.Context) {
		c.String(200, c.Query("q"))
	})
	r.Run(":8080")
}