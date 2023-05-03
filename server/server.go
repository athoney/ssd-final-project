package main

import (
	"html/template"
	"time"

	ratelimit "github.com/JGLTechnologies/gin-rate-limit"
	"github.com/gin-gonic/gin"
	"hibp.ssd.com/api"
)

func keyFunc(c *gin.Context) string {
	return c.ClientIP()
}

func noescape(str string) template.HTML {
	return template.HTML(str)
}

func main() {
	r := gin.Default()
	r.NoRoute(api.FOF)

	// This makes it so each ip can only make 5 requests per second
	store := ratelimit.InMemoryStore(&ratelimit.InMemoryOptions{
		Rate:  time.Minute,
		Limit: 2,
	})
	mw := ratelimit.RateLimiter(store, &ratelimit.Options{
		ErrorHandler: api.ErrorHandler,
		KeyFunc:      keyFunc,
	})

	//Static Routes
	r.GET("/", api.Home)
	r.GET("/about", api.About)
	r.GET("/login", api.Login)
	r.POST("/login", mw, api.CheckSecurity)
	r.GET("/post", api.Post)
	r.GET("/contact", api.Contact)

	//Load HTML Templates
	r.Static("/assets", "../website/assets")
	r.LoadHTMLGlob("../website/*.html")

	r.Run()
}
