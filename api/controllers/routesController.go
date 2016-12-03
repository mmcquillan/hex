package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/projectjane/jane/models"
)

// GetRoutes Returns all routes, HTTP GET - /api/v1/routes
// swagger:route GET /api/v1/routes routes getRoutes
//		 Returns all routes
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Schemes: http, https
//
//     Responses:
//       200:
//       401:
//			 500:
func GetRoutes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	aliases, err := data.GetAliases()
	if err != nil {
		sendError(w, err, http.StatusInternalServerError, "Failed to get aliases.")
		return
	}

	aliasesJSON, err := json.Marshal(aliases)
	if err != nil {
		sendError(w, err, http.StatusInternalServerError, "Failed to marshal object.")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(aliasesJSON)
}

// GetRouteByID Returns a route by id, HTTP GET - /api/v1/routes/{id}
// swagger:route GET /api/v1/routes/{routeId} routes getRouteByID
//		 Returns a connector by id
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Schemes: http, https
//
//     Responses:
//       200:
//       401:
//			 500:
func GetRouteByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id := vars["id"]

	alias, err := data.GetAliasByID(id)
	if err != nil {
		sendError(w, err, http.StatusInternalServerError, "Failed to get alias.")
		return
	}

	aliasJSON, err := json.Marshal(alias)
	if err != nil {
		sendError(w, err, http.StatusInternalServerError, "Failed to marshal object.")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(aliasJSON)
}

// CreateRoute Creates a command alias, HTTP GET - /api/v1/routes
// swagger:route POST /api/v1/routes routes createRoute
//		 Creates a command alias
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Schemes: http, https
//
//     Responses:
//       200:
//       401:
//			 500:
func CreateRoute(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var alias models.Alias
	err := json.NewDecoder(r.Body).Decode(&alias)
	if err != nil {
		sendError(w, err, http.StatusInternalServerError, "Invalid alias format.")
		return
	}

	alias, err := data.CreateAlias(alias)
	if err != nil {
		sendError(w, err, http.StatusInternalServerError, "Failed to create alias.")
		return
	}

	aliasJSON, err := json.Marshal(alias)
	if err != nil {
		sendError(w, err, http.StatusInternalServerError, "Failed to marshal object.")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(aliasJSON)
}

// UpdateRouteByID Updates an alias, HTTP GET - /api/v1/routes/{id}
// swagger:route PUT /api/v1/routes/{routeId} routes updateRouteByID
//		 Updates an alias
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Schemes: http, https
//
//     Responses:
//       200:
//       401:
//			 500:
func UpdateRouteByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var alias models.Alias
	err := json.NewDecoder(r.Body).Decode(&alias)
	if err != nil {
		sendError(w, err, http.StatusInternalServerError, "Invalid connector format.")
		return
	}

	alias, err = data.UpdateAliasByID(alias)
	if err != nil {
		sendError(w, err, http.StatusInternalServerError, "Failed to update alias.")
		return
	}

	aliasJSON, err := json.Marshal(alias)
	if err != nil {
		sendError(w, err, http.StatusInternalServerError, "Failed to marshal object.")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(aliasJSON)
}

//RouteID Struct for routeID parameter
//swagger:parameters getRouteByID updateRouteByID
type RouteID struct {

	// in: path
	RouteID string `json:"routeId"`
}
