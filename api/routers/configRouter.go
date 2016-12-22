package routers

import (
	"github.com/projectjane/jane/api/common"
	"github.com/projectjane/jane/api/controllers"

	"github.com/gorilla/mux"
)

func setConfigRoutes(router *mux.Router) *mux.Router {
	configRouter := mux.NewRouter().StrictSlash(false)

	configRouter.Handle("/api/v1/config", common.MustAuth(controllers.GetConfig)).Methods("GET")
	configRouter.Handle("/api/v1/config", common.MustAuth(controllers.UpdateConfig)).Methods("PUT")
	configRouter.Handle("/api/v1/config", common.MustAuth(controllers.CreateConfig)).Methods("POST")

	router.PathPrefix("/api/v1/config").Handler(configRouter)

	return router
}
