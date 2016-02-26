package handlers

import (
	"net/http"
	"strings"
	"log"
	"io/ioutil"
//	"bytes"
	"encoding/json"
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
	log.Println("Calling NDEx API: ", target)

	resp, err := http.Get(target)

	if err != nil {
		msg := getErrorString(resp, "Failed to fetch data from NDEx API.")
		http.Error(w, msg, 500)
		return
	}

	if resp.StatusCode != 200 {
		msg := getErrorString(resp, "NDEx API returns abnormal response.  Check status of " + target)
		http.Error(w, msg, resp.StatusCode)
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		msg := getErrorString(resp, "NDEx response contains invalid data.  Check output of " + target)
		http.Error(w, msg, 500)
		return
	}

	bs := string(body[:])
	cyjsReader := strings.NewReader(bs)
	cx2cyjs.Convert(cyjsReader, w)
}

func getErrorString(resp *http.Response, m string) string {
	ndexErr := Error{
		Message: m,
		Status: resp.Status,
		StatusCode:resp.StatusCode}
	ers, _ := json.MarshalIndent(ndexErr,"", "    ")
	return string(ers)
}
