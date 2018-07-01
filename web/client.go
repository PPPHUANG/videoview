package main

import (
	"net/http"
	"log"
	"io/ioutil"
	"encoding/json"
	"io"
	"bytes"
	"github.com/julienschmidt/httprouter"
	"net/url"
	"net/http/httputil"
)

var httpClient *http.Client

func init() {
	httpClient = &http.Client{}
}

func request(b *ApiBody, w http.ResponseWriter, r *http.Request) {
	var resp *http.Response
	var err error
	switch http.MethodPost {
	case http.MethodGet:
		req, _ := http.NewRequest("GET", b.Url, nil)
		req.Header = r.Header
		resp, err = httpClient.Do(req)
		if err != nil {
			log.Printf("request error :%v", err)
			return
		}
		normalResponse(w, resp)

	case http.MethodPost:
		req, _ := http.NewRequest("POST", b.Url, bytes.NewBuffer([]byte(b.ReqBody)))
		req.Header = r.Header
		resp, err = httpClient.Do(req)
		if err != nil {
			log.Printf("request error :%v", err)
			return
		}
		normalResponse(w, resp)

	case http.MethodDelete:
		req, _ := http.NewRequest("Delete", b.Url, nil)
		req.Header = r.Header
		resp, err = httpClient.Do(req)
		if err != nil {
			log.Printf("request error :%v", err)
			return
		}
		normalResponse(w, resp)
	default:
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "Bad api request")
		return
	}

}

func normalResponse(w http.ResponseWriter, r *http.Response) {
	res, err := ioutil.ReadAll(r.Body)
	if err != nil {
		re, _ := json.Marshal(ErrorInternalFaults)
		w.WriteHeader(500)
		io.WriteString(w, string(re))
		return
	}
	w.WriteHeader(r.StatusCode)
	io.WriteString(w, string(res))
}

func proxyHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	u, _ := url.Parse("http://127.0.0.1:9000")
	proxy := httputil.NewSingleHostReverseProxy(u)
	proxy.ServeHTTP(w, r)
}