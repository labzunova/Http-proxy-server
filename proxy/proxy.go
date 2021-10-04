package proxy

import (
	"io"
	"net/http"
)

type HttpProxy struct {}

func (p *HttpProxy) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	req.Header.Del("Proxy-Connection")
	req.RequestURI = ""

	client := http.Client{}
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}

	res, err := client.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
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
}