package main

import (
	"log"
	"net/http"
	"net/url"
	"strings"
)

func login(req CheckRequest, client http.Client) {
	loginData := url.Values{}
	loginData.Set("action", "login")
	loginData.Set("username", req.Source.Username)
	loginData.Set("password", req.Source.Password)

	u, _ := url.Parse("https://qualysapi.qg3.apps.qualys.com/api/2.0/fo/session/")
	loginRequest, _ := http.NewRequest("POST", u.String(), strings.NewReader(loginData.Encode()))
	loginRequest.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	loginRequest.Header.Set("X-Requested-With", "Curl Sample")
	_, err := client.Do(loginRequest)
	if err != nil {
		log.Fatalf("Exiting with error: %s", err)
	}
}

func logout(client http.Client) {
	logoutData := url.Values{}
	logoutData.Set("action", "logout")

	u, _ := url.Parse("https://qualysapi.qg3.apps.qualys.com/api/2.0/fo/session/")
	logoutRequest, _ := http.NewRequest("POST", u.String(), strings.NewReader(logoutData.Encode()))
	logoutRequest.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	logoutRequest.Header.Set("X-Requested-With", "Curl Sample")
	_, err := client.Do(logoutRequest)
	if err != nil {
		log.Fatalf("Exiting with error: %s", err)
	}
}
