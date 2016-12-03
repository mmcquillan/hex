package routers

import (
	"github.com/projectjane/jane/api/common"
	"github.com/projectjane/jane/api/controllers"

	"github.com/gorilla/mux"
)

func setRoutesRoutes(router *mux.Router) *mux.Router {
	routesRouter := mux.NewRouter().StrictSlash(false)

	routesRouter.Handle("/api/v1/routes/{id}", common.MustAuth(controllers.GetRouteByID)).Methods("GET")
	routesRouter.Handle("/api/v1/routes/{id}", common.MustAuth(controllers.UpdateRouteByID)).Methods("PUT")
	routesRouter.Handle("/api/v1/routes", common.MustAuth(controllers.GetRoutes)).Methods("GET")
	routesRouter.Handle("/api/v1/routes", common.MustAuth(controllers.CreateRoute)).Methods("POST")

	router.PathPrefix("/api/v1/routes").Handler(routesRouter)

	return router
}
