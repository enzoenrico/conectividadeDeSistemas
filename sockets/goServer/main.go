package main

import (
	//routing lib
	static "github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	//serving lib
	"gopkg.in/olahol/melody.v1"
)

func main() {
	router:= gin.Default()
	mel :=melody.New()

	router.Use(static.Serve("/", static.LocalFile("./public", true)))

	router.GET("/ws", func(c *gin.Context){
		mel.handleRequest(c.Writer, c.Request)
	})

	mel.handleRequest(func(s *melody.Session, msg []byte){
		mel.Broadcast(msg)
	})

	router.Run(":9999")
}
