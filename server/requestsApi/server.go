package requestsApi

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"proxy/proxy.go/repository"
	"proxy/proxy.go/utils"
	"strconv"
)

type WebApi struct {
	Repo repository.RequestsRepo
}

func (p *WebApi) HandleReturnAllRequests(w http.ResponseWriter, req *http.Request) {
	requests, err := p.Repo.LoadAllRequests()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := json.Marshal(requests)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(result)
	if err != nil {
		log.Fatalln("write() error", err)
		return
	}
}

func (p *WebApi) HandleOneRequest(w http.ResponseWriter, req *http.Request) {
	idString, ok := mux.Vars(req)["id"]
	if !ok {
		http.Error(w, "can't get ID", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "can't convert ID into int", http.StatusBadRequest)
		return
	}

	request, err := p.Repo.LoadOneRequest(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := json.Marshal(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(result)
	if err != nil {
		log.Fatalln("write() error", err)
		return
	}
}

func (p *WebApi) HandleRepeatRequest(w http.ResponseWriter, req *http.Request) {
	idString, ok := mux.Vars(req)["id"]
	if !ok {
		http.Error(w, "can't get ID", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "can't convert ID into int", http.StatusBadRequest)
		return
	}

	requestParams, err := p.Repo.LoadOneRequest(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	request, err := utils.ConvertRequestToHttp(requestParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = utils.RepeatRequest(w, &request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (p *WebApi) HandleScanRequest(w http.ResponseWriter, req *http.Request) {
	idString, ok := mux.Vars(req)["id"]
	if !ok {
		http.Error(w, "can't get ID", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "can't convert ID into int", http.StatusBadRequest)
		return
	}

	requestParams, err := p.Repo.LoadOneRequest(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.CheckXSS(w, requestParams)
}
