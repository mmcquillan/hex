package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/projectjane/jane/models"
)

// GetConnectors Returns all connectors, HTTP GET - /api/v1/connectors
// swagger:route GET /api/v1/connectors connectors getConnectors
//		 Returns all connectors
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
func GetConnectors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	connectors, err := data.GetConnectors()
	if err != nil {
		sendError(w, err, http.StatusInternalServerError, "Failed to get connectors.")
		return
	}

	connectorsJSON, err := json.Marshal(connectors)
	if err != nil {
		sendError(w, err, http.StatusInternalServerError, "Failed to marshal object.")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(connectorsJSON)
}

// GetConnectorByID Returns all contacts for an agent, HTTP GET - /api/v1/connectors/{id}
// swagger:route GET /api/v1/connectors/{connectorId} connectors getConnectorByID
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
func GetConnectorByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id := vars["id"]

	connector, err := data.GetConnectorByID(id)
	if err != nil {
		sendError(w, err, http.StatusInternalServerError, "Failed to get connector.")
		return
	}

	connectorJSON, err := json.Marshal(connector)
	if err != nil {
		sendError(w, err, http.StatusInternalServerError, "Failed to marshal object.")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(connectorJSON)
}

// CreateConnector Creates a connector, HTTP GET - /api/v1/connectors
// swagger:route POST /api/v1/connectors connectors createConnector
//		 Creates a connector
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
func CreateConnector(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var connector models.Connector
	err := json.NewDecoder(r.Body).Decode(&connector)
	if err != nil {
		sendError(w, err, http.StatusInternalServerError, "Invalid connector format.")
		return
	}

	connector, err := data.CreateConnector(connector)
	if err != nil {
		sendError(w, err, http.StatusInternalServerError, "Failed to create connector.")
		return
	}

	connectorJSON, err := json.Marshal(connector)
	if err != nil {
		sendError(w, err, http.StatusInternalServerError, "Failed to marshal object.")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(connectorJSON)
}

// UpdateConnectorByID Updates a connector, HTTP GET - /api/v1/connectors/{id}
// swagger:route PUT /api/v1/connectors/{connectorId} connectors updateConnector
//		 Updates a connector
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
func UpdateConnectorByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var connector models.Connector
	err := json.NewDecoder(r.Body).Decode(&connector)
	if err != nil {
		sendError(w, err, http.StatusInternalServerError, "Invalid connector format.")
		return
	}

	connector, err := data.UpdateConnectorByID(connector)
	if err != nil {
		sendError(w, err, http.StatusInternalServerError, "Failed to update connector.")
		return
	}

	connectorJSON, err := json.Marshal(connector)
	if err != nil {
		sendError(w, err, http.StatusInternalServerError, "Failed to marshal object.")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(connectorJSON)
}

//ConnectorID Struct for connectorID parameter
//swagger:parameters getConnectorByID
type ConnectorID struct {

	// in: path
	ConnectorID string `json:"connectorId"`
}
