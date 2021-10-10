package main

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"proxy/proxy.go/server/proxy"
	"proxy/proxy.go/server/requestsApi"

	_ "github.com/lib/pq"

	"proxy/proxy.go/repository"
)

func initDB() *sql.DB {
	// подключение postgres
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", "127.0.0.1", "5432",
		"labzunova", "1111", "postgres")
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(3)

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func main() {
	db := initDB()
	repo := repository.NewRequestsRepo(db)

	// proxy starting
	proxyHandler := &proxy.HttpProxy{Repo: repo}
	proxy := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: proxyHandler,
	}
	log.Println("proxy server is starting, :8080")
	go func() {
		err := proxy.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()

	// web-api starting
	apiHandler := &requestsApi.WebApi{Repo: repo}

	mux := mux.NewRouter()
	mux.HandleFunc("/requests", apiHandler.HandleReturnAllRequests)
	mux.HandleFunc("/request/{id}", apiHandler.HandleOneRequest)
	mux.HandleFunc("/repeat/{id}", apiHandler.HandleRepeatRequest)
	//mux.HandleFunc("/scan/{id}", apiHandler.HandleScanRequest)

	log.Println("web api is starting, :8000")
	log.Fatal(http.ListenAndServe(":8000", mux))
}
