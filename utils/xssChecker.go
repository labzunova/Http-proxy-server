package utils

import (
	"io/ioutil"
	"log"
	"net/http"
	"proxy/proxy.go/repository"
	"strings"
)

const XSS = "'\"><img src onerror=alert()>"

func CheckXSS(w http.ResponseWriter, request repository.Request) {
	var isVulnerable bool
	requestQuery, err := ConvertRequestToHttp(request)
	if err != nil {
		return
	}
	initialQuery := requestQuery.URL.RawQuery

	for key, value := range requestQuery.URL.Query() {
		requestQuery.URL.RawQuery = strings.ReplaceAll(requestQuery.URL.RawQuery, value[0], XSS)
		resp, err := http.DefaultTransport.RoundTrip(&requestQuery)
		if err != nil {
			http.Error(w, "roundtrip() error", http.StatusInternalServerError)
			return
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, "readall() error", http.StatusInternalServerError)
			return
		}
		if strings.Contains(string(body), XSS) {
			_, err = w.Write([]byte(key + " is vulnerable\n"))
			if err != nil {
				log.Fatalln("write() error")
			}
			isVulnerable = true
		} else {
			_, err = w.Write([]byte(key + " is vulnerable\n"))
			if err != nil {
				log.Fatalln("write() error")
			}
			isVulnerable = true
		}
		requestQuery.URL.RawQuery = initialQuery
	}

	if !isVulnerable {
		_, err = w.Write([]byte("Vulnerable aren't find"))
		if err != nil {
			log.Fatalln("write() error")
		}
	}
}
