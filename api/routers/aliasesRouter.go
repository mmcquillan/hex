package routers

import (
	"github.com/projectjane/jane/api/common"
	"github.com/projectjane/jane/api/controllers"

	"github.com/gorilla/mux"
)

func setAliasesRoutes(router *mux.Router) *mux.Router {
	aliasesRouter := mux.NewRouter().StrictSlash(false)

	aliasesRouter.Handle("/api/v1/aliases/{id}", common.MustAuth(controllers.GetAliasByID)).Methods("GET")
	aliasesRouter.Handle("/api/v1/aliases/{id}", common.MustAuth(controllers.UpdateAliasByID)).Methods("PUT")
	aliasesRouter.Handle("/api/v1/aliases", common.MustAuth(controllers.GetAliases)).Methods("GET")
	aliasesRouter.Handle("/api/v1/aliases", common.MustAuth(controllers.CreateAlias)).Methods("POST")

	router.PathPrefix("/api/v1/aliases").Handler(aliasesRouter)

	return router
}
