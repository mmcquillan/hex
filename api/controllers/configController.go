package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/projectjane/jane/data"
)

// GetConfig Returns the config, HTTP GET - /api/v1/config
// swagger:route GET /api/v1/config config getConfig
//		 Returns the config
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
func GetConfig(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	config, err := data.GetConfig()
	if err != nil {
		sendError(w, err, http.StatusInternalServerError, "Failed to get config.")
		return
	}

	configJSON, err := json.Marshal(config)
	if err != nil {
		sendError(w, err, http.StatusInternalServerError, "Failed to marshal object.")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(configJSON)
}

// CreateConfig Creates a config, HTTP GET - /api/v1/config
// swagger:route POST /api/v1/config config createConfig
//		 Creates a config
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
func CreateConfig(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var config data.Config
	err := json.NewDecoder(r.Body).Decode(&config)
	if err != nil {
		sendError(w, err, http.StatusInternalServerError, "Invalid config format.")
		return
	}

	config, err = data.CreateConfig(config)
	if err != nil {
		sendError(w, err, http.StatusInternalServerError, "Failed to create config.")
		return
	}

	configJSON, err := json.Marshal(config)
	if err != nil {
		sendError(w, err, http.StatusInternalServerError, "Failed to marshal object.")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(configJSON)
}

// UpdateConfig Updates an alias, HTTP GET - /api/v1/config
// swagger:route PUT /api/v1/config aliases updateConfig
//		 Updates a config
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
//       500:
func UpdateConfig(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var config data.Config
	err := json.NewDecoder(r.Body).Decode(&config)
	if err != nil {
		sendError(w, err, http.StatusInternalServerError, "Invalid config format.")
		return
	}

	config, err = data.UpdateConfig(config)
	if err != nil {
		sendError(w, err, http.StatusInternalServerError, "Failed to update config.")
		return
	}

	configJSON, err := json.Marshal(config)
	if err != nil {
		sendError(w, err, http.StatusInternalServerError, "Failed to marshal object.")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(configJSON)
}
