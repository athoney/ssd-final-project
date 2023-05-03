package hibp

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type response struct {
	name         string   `json:"Name"`
	title        string   `json:"Title"`
	domain       string   `json:"Domain"`
	breachDate   string   `json:"BreachDate"`
	AddedDate    string   `json:"AddedDate"`
	ModifiedDate string   `json:"ModifiedDate"`
	PwnCount     int      `json:"PwnCount"`
	Description  string   `json:"Description"`
	DataClasses  []string `json:"DataClasses"`
	IsVerified   bool     `json:"IsVerified"`
	IsFabricated bool     `json:"IsFabricated"`
	IsSensitive  bool     `json:"IsSensitive"`
	IsRetired    bool     `json:"IsRetired"`
	IsSpamList   bool     `json:"IsSpamList"`
	LogoPath     string   `json:"LogoPath"`
}

func CheckEmail(email string) []string {
	req, err := http.NewRequest(
		http.MethodGet,
		"https://haveibeenpwned.com/api/v3/breachedaccount/"+email+"?truncateResponse=false",
		nil,
	)
	if err != nil {
		log.Fatalf("error creating HTTP request: %v", err)
	}

	// req.Header.Add("Accept", "application/json")
	fmt.Println("Key:" + os.Getenv("KEY"))
	req.Header.Add("hibp-api-key", "aca6d5f51e824114bdc8a79924d69f5d")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("error sending HTTP request: %v", err)
	}

	decoder := json.NewDecoder(resp.Body)
	var emailInfo []response
	err = decoder.Decode(&emailInfo)

	var emailResp []string
	for _, breach := range emailInfo {
		// fmt.Printf(breach.Description)
		emailResp = append(emailResp, breach.Description)
	}

	return emailResp
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
		// fmt.Println(s[i])
		// fmt.Println(password)
		if strings.Contains(strings.ToUpper(password), s[i][0:35]) {

			//fmt.Printf("Password Found %s times", s[i][37:])
			return s[i][36:]
		}
	}
	fmt.Println("Password safe")
	return "0"
}
