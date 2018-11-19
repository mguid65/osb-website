package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/mguid65/osb-website/server/database"
)

// ListResults lists all results.
func ListResults(db database.ResultDatabase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		results, err := db.ListResults()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := sendJSONResponse(w, results); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

// ListResultsCreatedBy returns all results created by the user with the given user ID.
func ListResultsCreatedBy(db database.ResultDatabase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		idStr, ok := mux.Vars(r)["id"]
		if !ok {
			http.Error(w, `router: no "id" key in router`, http.StatusInternalServerError)
			return
		}

		userID, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		results, err := db.ListResultsCreatedBy(userID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := sendJSONResponse(w, results); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

// GetResult returns the result row with the matching result id.
func GetResult(db database.ResultDatabase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		idStr, ok := mux.Vars(r)["id"]
		if !ok {
			http.Error(w, `router: no "id" key`, http.StatusInternalServerError)
			return
		}

		resultID, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		result, err := db.GetResult(resultID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		if err := sendJSONResponse(w, result); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

// AddResult inserts a new result row.
func AddResult(db database.OSBDatabase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		username, password, ok := r.BasicAuth()
		if !ok {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		user, err := db.GetUserByCredentials(username, password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}

		submission := struct {
			Scores  database.Scores  `json:"scores"`
			SysInfo database.SysInfo `json:"specs"`
		}{}
		if err := json.NewDecoder(r.Body).Decode(&submission); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		result := &database.Result{
			UserID: user.ID,
			Scores: submission.Scores,
		}
		id, err := db.AddResult(result)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		log.Println("successfully added result id", id)

		specs := &database.Specs{
			ResultID: id,
			SysInfo:  submission.SysInfo,
		}
		id, err = db.AddSpecs(specs)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		log.Println("successfully added specs id", id)

		w.WriteHeader(http.StatusOK)
	}
}

// DeleteResult deletes the result row with the matching result id.
func DeleteResult(db database.ResultDatabase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		idStr, ok := mux.Vars(r)["id"]
		if !ok {
			http.Error(w, `router: no "id" key`, http.StatusInternalServerError)
			return
		}

		resultID, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := db.DeleteResult(resultID); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

// UpdateResult updates the result with the given values.
func UpdateResult(db database.ResultDatabase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		// TODO: get result values
		var result database.Result

		if err := db.UpdateResult(&result); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
