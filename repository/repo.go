package repository

import (
	"database/sql"
	"errors"
)

type requestsRepo struct {
	DB *sql.DB
}

func NewRequestsRepo(db *sql.DB) RequestsRepo {
	return &requestsRepo{
		DB: db,
	}
}

func (r requestsRepo) SaveRequest(req Request) error {
	_, err := r.DB.Exec(`INSERT INTO requests(host, path, method, headers, body, params, cookies) VALUES($1,$2,$3,$4,$5,$6,$7)`,
		req.Host, req.Path, req.Method, req.Headers, req.Body, req.Params, req.Cookies)
	return err
}

func (r requestsRepo) LoadAllRequests() ([]Request, error){
	requests := make([]Request, 0, 0)
	rows, err := r.DB.Query("select id, host, path, method, headers, body, params, cookies from requests")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		request := Request{}
		err := rows.Scan(&request.Id, &request.Host, &request.Path, &request.Method, &request.Headers, &request.Body, &request.Params, &request.Cookies)
		if err != nil {
			return nil, err
		}
		requests = append(requests, request)
	}
	return requests, nil
}

func (r requestsRepo) LoadOneRequest(id int) (Request, error) {
	req := Request{}
	err := r.DB.QueryRow(`select host, path, method, headers, body, params, cookies from requests where id = $1`, id).
		Scan(&req.Host, &req.Path, &req.Method, &req.Headers, &req.Body, &req.Params, &req.Cookies)
	if err == sql.ErrNoRows {
		return Request{}, errors.New("no such requrst ")
	}
	if err != nil {
		return Request{}, err
	}
	return req, nil
}

func (r requestsRepo) RepeatRequest(id int) error {
	panic("implement me")
}
