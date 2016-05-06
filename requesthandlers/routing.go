package handlers

import (
	"net/http"
	"log"
	"github.com/rs/cors"
	"strconv"
)


func StartServer(portNumber int) (err error) {

	mux := http.NewServeMux()

	// Root: Show status of service
	mux.HandleFunc("/", StatusHandler)

	// Convert CX stream into Cytoscape JS
	mux.HandleFunc("/cx2cyjs", Cx2CyjsHandler)

	// Utility to convert network stored in NDEx into Cytoscape.js
	mux.HandleFunc("/ndex2cyjs/", Ndex2CyjsHandler)

	handler := cors.Default().Handler(mux)

	log.Println("Serving cxtool API on port ", portNumber)

	portNumStr := strconv.Itoa(portNumber)

	err = http.ListenAndServe(":" + portNumStr, handler)

	if err != nil {
		log.Fatal("Could not start API server: ", err)
	}

	return nil
}
