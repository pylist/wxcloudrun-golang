package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"wxcloudrun-golang/db"
	"wxcloudrun-golang/service"
)

func main() {
	if err := db.Init(); err != nil {
		panic(fmt.Sprintf("mysql init failed with %+v", err))
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	api := r.Group("api")
	api.GET("/user/list", service.UserList)
	api.GET("/user/info", service.UserInfo)
	api.POST("/message/template/send", service.SenTemplateMessage)
	log.Fatal(r.Run(":80")) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
