package hibp

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
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

func CheckEmail(email string) string {
	req, err := http.NewRequest(
		http.MethodGet,
		"https://haveibeenpwned.com/api/v3/breachedaccount/"+email+"?truncateResponse=false",
		nil,
	)
	if err != nil {
		log.Fatalf("error creating HTTP request: %v", err)
	}

	// req.Header.Add("Accept", "application/json")
	req.Header.Add("hibp-api-key", os.Getenv("KEY"))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("error sending HTTP request: %v", err)
	}

	fmt.Println(res)
	return "hi"
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

func CheckPassword(password string) string {
	pw := password[0:5]
	//fmt.Println(pw)
	resp, err := http.Get("https://api.pwnedpasswords.com/range/" + pw)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	sb := string(body)
	s := strings.Split(sb, "\n")

	for i := 1; i < len(s); i++ {
		fmt.Println(s[i])
		fmt.Println(password)
		if strings.Contains(strings.ToUpper(password), s[i][0:35]) {

			//fmt.Printf("Password Found %s times", s[i][37:])
			return s[i][36:]
		}
	}
	fmt.Println("Password safe")
	return "0"
}
