package routers

import (
  "github.com/gorilla/mux"
  "github.com/projectjane/jane/api/controllers"
)

func SetBuildsRoutes(router *mux.Router) *mux.Router {
  buildsRouter := mux.NewRouter().StrictSlash(false)
  buildsRouter.HandleFunc("/api/v1/builds/start", controllers.StartBuild).Methods("POST")

  router.PathPrefix("/api/v1/builds").Handler(buildsRouter)

  return router
}
