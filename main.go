package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strings"

	"golang.org/x/net/publicsuffix"
)

type source struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type request struct {
	Source source `json:"source"`
}

func main() {
	req := request{}
	err := json.NewDecoder(os.Stdin).Decode(&req)
	if err != nil {
		log.Fatalf("Exiting with error: %s", err)
	}
	// Login to Qualys
	jar, _ := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	client := http.Client{Jar: jar}

	loginData := url.Values{}
	loginData.Set("action", "login")
	loginData.Set("username", req.Source.Username)
	loginData.Set("password", req.Source.Password)

	u, _ := url.Parse("https://qualysapi.qg3.apps.qualys.com/api/2.0/fo/session/")
	loginRequest, _ := http.NewRequest("POST", u.String(), strings.NewReader(loginData.Encode()))
	loginRequest.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	loginRequest.Header.Set("X-Requested-With", "Curl Sample")
	_, err = client.Do(loginRequest)
	if err != nil {
		log.Fatalf("Exiting with error: %s", err)
	}

	// Logout from Qualys
	logoutData := url.Values{}
	logoutData.Set("action", "logout")
	logoutRequest, _ := http.NewRequest("POST", u.String(), strings.NewReader(logoutData.Encode()))
	logoutRequest.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	logoutRequest.Header.Set("X-Requested-With", "Curl Sample")
	_, err = client.Do(logoutRequest)
	if err != nil {
		log.Fatalf("Exiting with error: %s", err)
	}
}
