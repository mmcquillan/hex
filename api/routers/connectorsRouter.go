package routers

import (
	"github.com/projectjane/jane/api/common"
	"github.com/projectjane/jane/api/controllers"

	"github.com/gorilla/mux"
)

func setConnectorsRoutes(router *mux.Router) *mux.Router {
	connectorsRouter := mux.NewRouter().StrictSlash(false)

	connectorsRouter.Handle("/api/v1/connectors/{id}", common.MustAuth(controllers.GetConnectorByID)).Methods("GET")
	connectorsRouter.Handle("/api/v1/connectors/{id}", common.MustAuth(controllers.UpdateConnectorByID)).Methods("PUT")
	connectorsRouter.Handle("/api/v1/connectors", common.MustAuth(controllers.GetConnectors)).Methods("GET")
	connectorsRouter.Handle("/api/v1/connectors", common.MustAuth(controllers.CreateConnector)).Methods("POST")

	router.PathPrefix("/api/v1/connectors").Handler(connectorsRouter)

	return router
}
