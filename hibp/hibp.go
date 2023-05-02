package hibp

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"golang.org/x/time/rate"
)

type APIConn struct {
	rateLimiter *rate.Limiter
}

func Open() *APIConn {
	return &APIConn{
		rateLimiter: rate.NewLimiter(rate.Every(time.Minute), 3),
	}
}

func CheckEmail(email string) {
	print("Hello world")
}

func (c *APIConn) read(ctx context.Context, email string) (string, error) {
	key := os.Getenv("KEY")
	if err := c.rateLimiter.Wait(ctx); err != nil {
		return "", err
	}

	resp, err := http.Get("https://haveibeenpwned.com/api/v3/breachedaccount/alicia.thoney@gmail.com?truncateResponse=false?hibp-api-key=" + key)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Print(resp)
	//DoWork
	return "Read", nil
}

func CheckPassword(password string) {
	fmt.Println("Hello word")
}
