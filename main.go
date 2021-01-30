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

	// echo
	r.GET("/echo/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.String(200, id)
	})

	r.GET("/echoall/*id", func(c *gin.Context) {
		id := c.Param("id")
		c.String(200, id)
	})

	r.Run(":8080")
}