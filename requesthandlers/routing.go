package handlers

import (
	"net/http"
	"log"
	"github.com/rs/cors"
	"strconv"
)


func StartServer(portNumber int) (err error) {

	mux := http.NewServeMux()
	mux.HandleFunc("/", StatusHandler)
	mux.HandleFunc("/cx2cyjs", Cx2CyjsHandler)
	mux.HandleFunc("/ndex2cyjs/", Ndex2CyjsHandler)

	handler := cors.Default().Handler(mux)

	log.Println("Serving API on port ", portNumber)

	portNumStr := strconv.Itoa(portNumber)

	err = http.ListenAndServe(":" + portNumStr, handler)

	if err != nil {
		log.Fatal("Could not start API server: ", err)
	}

	return nil
}
