package main

import (
	static "github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
)

func main() {
	router := gin.Default()
	mel := melody.New()

	router.Use(static.Serve("/", static.LocalFile("./public", true)))

	router.GET("/ws", func(c *gin.Context) {
		mel.HandleRequest(c.Writer, c.Request)
	})

	mel.HandleMessage(func(s *melody.Session, msg []byte) {
		mel.Broadcast(msg)
	})

	router.Run(":5000")
}