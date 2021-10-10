package utils

import (
	"io"
	"net/http"
	"time"
)

func DoRequest(w http.ResponseWriter, r *http.Request, client http.Client) error {
	res, err := client.Do(r)
	if err != nil { 
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}

	w.WriteHeader(res.StatusCode)
	_, err = io.Copy(w, res.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	for header, values := range res.Header {
		for _, value := range values {
			w.Header().Add(header, value)
		}
	}
	return nil
}

func RepeatRequest(w http.ResponseWriter, r *http.Request) error {
	client := http.Client{
		Timeout: time.Second * 10,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	return DoRequest(w, r, client)
}

