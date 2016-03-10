package routers

import (
  "github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
  router := mux.NewRouter().StrictSlash(false)

  router = SetBuildsRoutes(router)

  return router
}
