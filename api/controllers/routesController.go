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

	routes, err := data.GetRoutes()
	if err != nil {
		sendError(w, err, http.StatusInternalServerError, "Failed to get routes.")
		return
	}

	routesJSON, err := json.Marshal(routes)
	if err != nil {
		sendError(w, err, http.StatusInternalServerError, "Failed to marshal object.")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(routesJSON)
}

// GetRouteByID Returns a route by id, HTTP GET - /api/v1/routes/{id}
// swagger:route GET /api/v1/routes/{routeId} routes getRouteByID
//		 Returns a route by id
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
	routeID := vars["id"]

	route, err := data.GetRouteByID(routeID)
	if err != nil {
		sendError(w, err, http.StatusInternalServerError, "Failed to get route.")
		return
	}

	routeJSON, err := json.Marshal(route)
	if err != nil {
		sendError(w, err, http.StatusInternalServerError, "Failed to marshal object.")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(routeJSON)
}

// CreateRoute Creates a command alias, HTTP GET - /api/v1/routes
// swagger:route POST /api/v1/routes routes createRoute
//		 Creates a route
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

	var route models.Route
	err := json.NewDecoder(r.Body).Decode(&route)
	if err != nil {
		sendError(w, err, http.StatusInternalServerError, "Invalid route format.")
		return
	}

	route, err := data.CreateRoute(route)
	if err != nil {
		sendError(w, err, http.StatusInternalServerError, "Failed to create route.")
		return
	}

	routeJSON, err := json.Marshal(route)
	if err != nil {
		sendError(w, err, http.StatusInternalServerError, "Failed to marshal object.")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(routeJSON)
}

// UpdateRouteByID Updates a route, HTTP GET - /api/v1/routes/{id}
// swagger:route PUT /api/v1/routes/{routeId} routes updateRouteByID
//		 Updates a route
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

	var route models.Route
	err := json.NewDecoder(r.Body).Decode(&route)
	if err != nil {
		sendError(w, err, http.StatusInternalServerError, "Invalid route format.")
		return
	}

	route, err = data.UpdateAliasByID(route)
	if err != nil {
		sendError(w, err, http.StatusInternalServerError, "Failed to update route.")
		return
	}

	routeJSON, err := json.Marshal(route)
	if err != nil {
		sendError(w, err, http.StatusInternalServerError, "Failed to marshal object.")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(routeJSON)
}

// DeleteRouteByID Deletes a route, HTTP DELETE - /api/v1/routes/{id}
// swagger:route DELETE /api/v1/routes/{routeId} routes deleteRouteByID
//		 Deletes a route
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
func DeleteRouteByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	routeID := vars["id"]

	err := data.DeleteRouteByID(routeID)
	if err != nil {
		sendError(w, err, http.StatusInternalServerError, "Failed to delete route.")
		return
	}

	respJSON, err := json.Marshal("Successfully deleted route")
	if err != nil {
		sendError(w, err, http.StatusInternalServerError, "Failed to marshal object.")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(respJSON)
}

//RouteID Struct for routeID parameter
//swagger:parameters getRouteByID updateRouteByID deleteRouteByID
type RouteID struct {

	// in: path
	RouteID string `json:"routeId"`
}
