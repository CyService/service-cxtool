package handlers

import (
	"net/http"
	"strings"
	"log"
)

const (
	NDEX_URL = "http://dev2.ndexbio.org/rest/network/"
)


func Ndex2CyjsHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		get(w, r)
	} else {
		http.Error(w, "Request method must be GET.", 405)
	}
}

func get(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/ndex2cyjs/")

	target := NDEX_URL + id + "/asCX"
	log.Println("TARGET: ", target)

	resp, err := http.Get(target)

	if err != nil {
		log.Println(err)
		return
	}

	defer resp.Body.Close()

	log.Println("OK!: ", target)
	log.Println(resp)

	cx2cyjs.Convert(resp.Body, w)
}

