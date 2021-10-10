package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"proxy/proxy.go/repository"
)

func ConvertHttpToRequest(r *http.Request) (repository.Request, error) {
	result := repository.Request{
		Method: r.Method,
		Host: r.Host,
		Path: r.URL.Path,
	}

	headers, err := json.Marshal(r.Header)
	if err != nil {
		return result, err
	}
	result.Headers = string(headers)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return result, err
	}
	result.Body = string(body)

	params, err := json.Marshal(r.URL.Query())
	if err != nil {
		return result, err
	}
	result.Params = string(params)

	for _, cookie := range r.Cookies() {
		result.Cookies += cookie.Name + ": " + cookie.Value + "\n"
	}

	//r.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	//
	//if r.URL.Scheme == "" {
	//	r.URL.Scheme = "https"
	//}

	return result, nil
}

func ConvertRequestToHttp(request repository.Request) (http.Request, error) {
	httpRequest := http.Request{
		URL: &url.URL{
			//Scheme: request.Scheme,
			Host:   request.Host,
			Path:   request.Path,
		},
		Method: request.Method,
		Host:   request.Host,
		Body: ioutil.NopCloser(strings.NewReader(request.Body)),
	}

	var headers http.Header
	err := json.Unmarshal([]byte(request.Headers), &headers)
	if err != nil {
		return http.Request{}, err
	}

	var params url.Values
	err = json.Unmarshal([]byte(request.Params), &params)
	if err != nil {
		return http.Request{}, err
	}

	query := httpRequest.URL.Query()
	for key, values := range params {
		for _, value := range values {
			query.Add(key, value)
		}
	}
	httpRequest.URL.RawQuery = query.Encode()

	httpRequest.Header = http.Header{}
	for key, values := range headers {
		for _, value := range values {
			httpRequest.Header.Add(key, value)
		}
	}

	return httpRequest, nil
}
