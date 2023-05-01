package main

import (
	"github.com/gin-gonic/gin"
	"hibp.ssd.com/api"
)

func main() {
	r := gin.Default()
	r.NoRoute(api.FOF)

	//Static Routes
	r.GET("/", api.Home)
	r.GET("/about", api.About)
	r.GET("/login", api.Login)
	r.POST("/login", api.LoginUser)
	r.GET("/post", api.Post)
	r.GET("/contact", api.Contact)

	//Load HTML Templates
	r.Static("/assets", "../website/assets")
	r.LoadHTMLGlob("../website/*.html")

	r.Run()
}
