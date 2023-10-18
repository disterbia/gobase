package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nats-io/nats.go"
)

func main() {

	// Connect to NATS server
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		println(err)
	}
	defer nc.Close()

	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		// Publish an event
		if err := nc.Publish("hello", []byte("Hello, World!")); err != nil {
			println(err)
		}

		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})
	println("dmonster test")
	r.Run() // listen and serve on 0.0.0.0:8080

}
