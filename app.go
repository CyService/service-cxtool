package main

import (
	"log"
	"os"
	"strconv"
	"flag"
	elsa "github.com/cyService/elsa-client-go/reg"
	handlers "github.com/cyService/service-cxtool/requesthandlers"
)

func main() {

	elsaLocation := flag.String("agent", "http://localhost:8080/registration", "Agent URL")
	flag.Parse()

	var servicePort int
	portFlag := flag.Lookup("port")

	if portFlag == nil {
		servicePort = elsa.DefPort
	} else {
		var portErr error
		servicePort, portErr = strconv.Atoi(portFlag.Value.String())
		if portErr != nil {
			log.Fatal("Could not start API server: ", portErr.Error())
			os.Exit(1)
		}
	}

	// Asynchronously register this service to Submit Agent
	go elsa.Register(*elsaLocation)

	// Start API server
	serverErr := handlers.StartServer(servicePort)

	if serverErr != nil {
		log.Fatal("Could not start API server: ", serverErr.Error())
		os.Exit(1)
	}
}
