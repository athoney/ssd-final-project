package api

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"net/http"

	"hibp.ssd.com/hibp"

	"github.com/gin-gonic/gin"
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
		gin.H{},
	)
}

func LoginUser(c *gin.Context) {
	username := c.Request.PostFormValue("email")
	password := c.Request.PostFormValue("password")
	checkEmail := c.Request.PostFormValue("checkEmail")
	checkPassword := c.Request.PostFormValue("checkPassword")
	h := sha1.New()
	h.Write([]byte(password))
	sha1_hash := hex.EncodeToString(h.Sum(nil))
	fmt.Printf("username: %s\n", username)
	fmt.Printf("password: %s\n", sha1_hash)
	fmt.Printf("check email: %s\n", checkEmail)
	fmt.Printf("check password: %s\n", checkPassword)
	if checkEmail == "on" {
		hibp.CheckEmail(username)
	}

	if checkPassword == "on" {
		hibp.CheckEmail(username)
	}

	// val := query.NewUser(conn, username, string(hashedPassword))
	c.HTML(
		http.StatusOK,
		"login",
		gin.H{
			"status": "OK",
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

func Admin(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"admin",
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
