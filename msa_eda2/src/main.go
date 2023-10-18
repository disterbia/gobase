package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/nats-io/nats.go"
)

func main() {
	// Connect to NATS server
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		println(err)
		log.Fatal(err)
	}
	defer nc.Close()

	// Subscribe to "hello" subject
	_, err = nc.Subscribe("hello", func(m *nats.Msg) {
		println("-=-==")
		log.Printf("Received a message: %s\n", string(m.Data))
	})
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World 2!",
		})
	})
	r.Run(":8081") // listen and serve on 0.0.0.0:8081
}
