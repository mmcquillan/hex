package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/projectjane/jane/models"
)

// GetAliases Returns all connectors, HTTP GET - /api/v1/aliases
// swagger:route GET /api/v1/aliases aliases getAliases
//		 Returns all aliases
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
func GetAliases(w http.ResponseWriter, r *http.Request) {
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

// GetAliasByID Returns an alias by id, HTTP GET - /api/v1/connectors/{id}
// swagger:route GET /api/v1/aliases/{aliasId} aliases getAliasByID
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
func GetAliasByID(w http.ResponseWriter, r *http.Request) {
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

// CreateAlias Creates a command alias, HTTP GET - /api/v1/aliases
// swagger:route POST /api/v1/aliases aliases createAlias
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
func CreateAlias(w http.ResponseWriter, r *http.Request) {
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

// UpdateAliasByID Updates an alias, HTTP GET - /api/v1/aliases/{id}
// swagger:route PUT /api/v1/aliases/{aliasId} aliases updateAliasByID
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
func UpdateAliasByID(w http.ResponseWriter, r *http.Request) {
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

//AliasID Struct for aliasID parameter
//swagger:parameters getAliasByID updateAliasByID
type AliasID struct {

	// in: path
	AliasID string `json:"aliasId"`
}
