package main

import (
	"log"
	"net/http"
	"onecause/apphandlers"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// We are adding a PathPrefix so that we can control the version of the apis easily
	apiv1 := r.PathPrefix("/apiv1").Subrouter()

	// [ROUTING]
	apiv1.HandleFunc("/", apphandlers.HealthHandler).Methods(http.MethodGet)
	apiv1.HandleFunc("/login", apphandlers.LoginHandler) //.Methods(http.MethodPost)

	// Logged Router helps us to get the logs of api request.
	// These logs should be stored in common place like AWS cloud watch and analysed
	// later to improve user experience
	loggedRouter := handlers.LoggingHandler(os.Stdout, apiv1)

	// We are useing CompressHandler to compress the payload using gzip for client
	// It helps in low latency since the data is compressed and there is less data to send
	compressedHandler := handlers.CompressHandler(loggedRouter)

	// We are implementing Recovery Handler in case of unseen panics in the code
	// we do not want our server to crash we want it to handle the panic and report
	// us the error. For it to report us the error we have to built our own handler
	// just for time sake i used default
	recoveryHandler := handlers.RecoveryHandler()(compressedHandler)

	// CORS for local development
	originsOk := handlers.AllowedOrigins([]string{"http://localhost:4200"})
	c := handlers.CORS(originsOk)(recoveryHandler)

	log.Fatal(http.ListenAndServe(":8080", c))
}
