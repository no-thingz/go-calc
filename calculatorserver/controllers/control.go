package controllers

import (
	"encoding/json"
	"io"
	"net/http"
)

func GetListenAddress() string {
	port := "3030"
	return ":" + port
}

func RegisterControllers() {
	cc := newCalcController()

	http.Handle("/", *cc)
}

func encodeResponseAsJSON(data interface{}, w io.Writer) {
	enc := json.NewEncoder(w)
	enc.Encode(data)
}
