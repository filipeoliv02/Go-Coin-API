package main

import (
	"fmt"
	"github.com/filipeoliv02/goapi/internal/handlers"
	"net/http"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)
	fmt.Println("Starting GO API service...")
	// insert cool ascii art here
	err := http.ListenAndServe("localhost:8080", r)

	if err != nil {
		log.Fatal(err)
	}

}
