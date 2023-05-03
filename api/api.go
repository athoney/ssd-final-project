package api

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"net/http"
	"time"

	ratelimit "github.com/JGLTechnologies/gin-rate-limit"
	"github.com/gin-gonic/gin"
	"hibp.ssd.com/hibp"
)

func Home(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"home",
		gin.H{},
	)
}

func Login(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"login",
		gin.H{
			"pass": "0",
		},
	)
}

func CheckSecurity(c *gin.Context) {
	// Get email and password from form
	email := c.Request.PostFormValue("email")
	password := c.Request.PostFormValue("password")

	// Hash password
	h := sha1.New()
	h.Write([]byte(password))
	sha1_hash := hex.EncodeToString(h.Sum(nil))

	//Print info
	fmt.Printf("username: %s\n", email)
	fmt.Printf("password: %s\n", sha1_hash)

	// Check Password
	passInfo := hibp.CheckPassword(sha1_hash)

	// Check email
	emailInfo := hibp.CheckEmail(email)

	page := "home"

	if passInfo != "0" || emailInfo != "" {
		page = "summary"
	}

	c.HTML(
		http.StatusOK,
		page,
		gin.H{
			"status":   "OK",
			"password": passInfo,
			"email":    "",
		},
	)
}

func ErrorHandler(c *gin.Context, info ratelimit.Info) {
	// c.String(429, "Too many requests. Try again in "+time.Until(info.ResetTime).String())
	c.HTML(
		http.StatusOK,
		"home",
		gin.H{
			"Time": time.Until(info.ResetTime).String(),
		},
	)
}

func About(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"about",
		gin.H{},
	)
}

func Post(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"post",
		gin.H{},
	)
}

func Contact(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"contact",
		gin.H{},
	)
}

func FOF(c *gin.Context) {
	c.HTML(
		http.StatusNotFound,
		"fof",
		gin.H{},
	)
}
