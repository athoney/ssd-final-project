package main

import (
	"time"

	ratelimit "github.com/JGLTechnologies/gin-rate-limit"
	"github.com/gin-gonic/gin"
	"hibp.ssd.com/api"
)

func keyFunc(c *gin.Context) string {
	return c.ClientIP()
}

func errorHandler(c *gin.Context, info ratelimit.Info) {
	c.String(429, "Too many requests. Try again in "+time.Until(info.ResetTime).String())
}

func main() {
	r := gin.Default()
	r.NoRoute(api.FOF)

	// This makes it so each ip can only make 5 requests per second
	store := ratelimit.InMemoryStore(&ratelimit.InMemoryOptions{
		Rate:  time.Minute,
		Limit: 5,
	})
	mw := ratelimit.RateLimiter(store, &ratelimit.Options{
		ErrorHandler: errorHandler,
		KeyFunc:      keyFunc,
	})

	//Static Routes
	r.GET("/", api.Home)
	r.GET("/about", api.About)
	r.GET("/login", api.Login)
	r.POST("/login", api.LoginUser)
	r.POST("/login/email", mw, api.LoginUser)
	r.GET("/post", api.Post)
	r.GET("/contact", api.Contact)

	//Load HTML Templates
	r.Static("/assets", "../website/assets")
	r.LoadHTMLGlob("../website/*.html")

	r.Run()
}
