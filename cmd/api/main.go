package main

import (
	"fmt"
	"net/http"

	"github.com/acertainpoggerman/goapi/internal/handlers"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.Println("Starting GO API service...")
	fmt.Println(`
	
 ______   _____
/\  ___\ /\  __ \
\ \ \__ \\ \ \_\ \
 \ \_____\\ \_____\
  \/_____/ \/_____/
	
	`)

	err := http.ListenAndServe("localhost:8080", r)
	if err != nil {
		log.Error(err)
	}
}
