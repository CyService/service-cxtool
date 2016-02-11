package handlers

import (
	"net/http"
	"strings"
	"log"
	"io/ioutil"
//	"bytes"
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
		http.Error(w, "could not access NDEx", 500)
	}

	defer resp.Body.Close()

	// Read it
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Could not access NDEx service.", 500)
	}

	log.Println("GOT body: ", target)
	bs := string(body[:])
	cyjsReader := strings.NewReader(bs)

	cx2cyjs.Convert(cyjsReader, w)
}

