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

// {
//   "source": {
//     "username": "...",
//     "password": "..."
//     "server": "..."
//   },
//   "version": { "timestamp": YYYY-MM-DD[THH:MM:SSZ] }
// }

type Source struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Server    string `json:"server"`
	KBVersion string `json:"kb_version"`
}

type CheckRequest struct {
	Source  Source    `json:"source"`
	Version KBVersion `json:"version"`
}

type KBVersion struct {
	KBVersion string `json:"kb_version"`
}

func main() {
	req := CheckRequest{}
	err := json.NewDecoder(os.Stdin).Decode(&req)
	if err != nil {
		log.Fatalf("Exiting with error: %s", err)
	}
	// Login to Qualys
	jar, _ := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	client := http.Client{Jar: jar}

	login(req, client)

	timeSince := "1900-01-01T00:00:00Z"
	if req.Version != "" {
		timeSince = req.Version
	}
	kbData := url.Values{}
	kbData.Set("action", "list")
	kbData.Set("last_modified_after", timeSince)

	u, _ := url.Parse("https://qualysapi.qg3.apps.qualys.com/api/2.0/fo/session/")
	kbRequest, _ := http.NewRequest("POST", u.String(), strings.NewReader(kbData.Encode()))
	kbRequest.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	kbRequest.Header.Set("X-Requested-With", "Curl Sample")
	_, err := client.Do(kbRequest)

	logout(client)
	// Logout from Qualys
}
