package handlers

import (
	"net/http"
	"encoding/json"
)


type Status struct {
	Name        string `json:"name"`
	Version     string `json:"version"`
	Build       string `json:"build"`
	Description string `json:"description"`
	Documents   string `json:"documents"`
}


func StatusHandler(w http.ResponseWriter, r *http.Request) {

	serviceStatus := Status{
		Name:"Cxtool service",
		Version:"v1",
		Build: "05-06-2016",
		Description:"Converts CX format into Cytoscape.js compatible JSON.",
		Documents: "https://github.com/cyService/service-cxtool",
	}

	if r.Method == GET {
		json.NewEncoder(w).Encode(serviceStatus)
	} else {
		http.Error(w, "Request method must be GET.", 405)
	}
}
