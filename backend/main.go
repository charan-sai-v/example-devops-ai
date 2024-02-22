package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// to allow CORS
	r.Use(cors.Default())
	// routes
	r.GET("/", HelloWorld)
	// run server
	r.Run(":8000")
}

func HelloWorld(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello World!")
}
