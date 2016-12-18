package routers

import "github.com/gorilla/mux"

// InitRoutes Initializes the routes for the API
func InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)

	router := setConfigRoutes(router)

	return router
}
