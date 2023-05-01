package hibp

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func CheckEmail(email string) {
	fmt.Println("Hello word")
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
			return s[i][37:]
		}
	}
	fmt.Println("Password safe")
	return "0"
}
