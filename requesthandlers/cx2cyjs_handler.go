package handlers

import (
	"net/http"
	"encoding/json"
	"github.com/cyService/cxtool/converter"
	"log"
	"errors"
)


const (
	GET = "GET"
	POST = "POST"
)

type Message struct {
	Code    int `json:"code"`
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}

var cx2cyjs converter.Converter


func init() {
	cx2cyjs = converter.Cx2Cyjs{}
	log.Println("Converter created.")
}


func Cx2CyjsHandler(w http.ResponseWriter, r *http.Request) {

	method := r.Method
	switch method {
	case POST:
		post(w, r)
	default:
		unsupported(w, r)
	}
}


func post(w http.ResponseWriter, r *http.Request) {
	cx2cyjs.Convert(r.Body, w)
}


func unsupported(w http.ResponseWriter, r *http.Request) {
	code := 405
	res := getErrorMsg(code, "You need to POST your data to use this service.",
		errors.New("Invalid HTTP method used: " + r.Method))
	log.Println("Unsupported method call from: ", r.RemoteAddr)
	http.Error(w, res, code)
}


func getErrorMsg(code int, msg string, err error) (jsonMsg string) {

	message := Message{
		Code: code,
		Message: msg,
	}

	if err != nil {
		message.Error = err.Error()
	}

	result, err := json.Marshal(message)

	if err != nil {
		// TODO: What should I return?
	}

	return string(result)
}
