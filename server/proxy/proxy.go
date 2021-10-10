package proxy

import (
	"fmt"
	"log"
	"net/http"
	"proxy/proxy.go/utils"

	"proxy/proxy.go/repository"
)

type HttpProxy struct {
	Repo repository.RequestsRepo
}

func (p *HttpProxy) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Println("handler")
	if req.Method == http.MethodConnect {
		serveHTTPs(w, req)
		return
	}

	request, err := utils.ConvertHttpToRequest(req)
	if err != nil {
		log.Fatal("convert request error:", err)
	}
	err = p.Repo.SaveRequest(request)
	if err != nil {
		log.Fatal("save http request error:", err)
	}
	serveHTTP(w, req)
}

func serveHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("serve http")
	client := http.Client{}
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}
	r.Header.Del("Proxy-Connection")
	r.RequestURI = ""

	utils.DoRequest(w, r, client)
}

func serveHTTPs(w http.ResponseWriter, r *http.Request) {

}
