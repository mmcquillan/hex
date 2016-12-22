package api

import (
	"log"
	"net/http"

	"github.com/projectjane/jane/api/routers"
	"github.com/projectjane/jane/data"

	"github.com/rs/cors"
)

// Start Starts the api
func Start() {
	router := routers.InitRoutes()

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "OPTIONS", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type", "Access-Control-Allow-Origin", "Authorization"},
	})

	handler := c.Handler(router)

	//Need to make addr configurable or a standard port (80/443). Need to figure out ssl
	server := &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
		data.Database.Close()
	}
}
