package controllers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"regexp"
)

type calculateNumber struct {
	NumA float64 `json:"a"`
	NumB float64 `json:"b"`
}

var (
	calculateNumbers []*calculateNumber
)

func GetListenAddress() string {
	port := "3000"
	return ":" + port
}

func getProxyURL() string {
	return "http://localhost:3030"
}

func RegisterControllers() {

	http.HandleFunc("/", handleRequestAndRedirect)

}

func handleRequestAndRedirect(w http.ResponseWriter, r *http.Request) {
	var validPath = regexp.MustCompile(`^/calculator\.(sum|sub|mul|div)$`)
	m := validPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Invalid Endpoint"))
		return
	}

	url := getProxyURL()
	err := validateRequestBody(r)
	if err != nil {
		log.Printf("Error parse request body: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid JSON body request"))
		return
	}
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Allow POST Method only"))
		return
	}
	serveReverseProxy(url, w, r)
}

func requestBodyDecoder(r *http.Request) *json.Decoder {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
	}
	r.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	return json.NewDecoder(ioutil.NopCloser(bytes.NewBuffer(body)))
}

func validateRequestBody(r *http.Request) error {
	dec := requestBodyDecoder(r)
	var cn calculateNumber
	err := dec.Decode(&cn)
	return err
}

func serveReverseProxy(target string, w http.ResponseWriter, r *http.Request) {
	url, _ := url.Parse(target)

	proxy := httputil.NewSingleHostReverseProxy(url)

	r.URL.Host = url.Host
	r.URL.Scheme = url.Scheme
	r.Header.Set("X-Forwarded-Host", r.Header.Get("Host"))
	r.Host = url.Host

	proxy.ServeHTTP(w, r)
}
