package handlers

import (
	"net/http"
        "strconv"
//	"log"
//	"encoding/json"
        "github.com/gorilla/mux"

	"github.com/mguid65/osb-website/server/database"
)

// ListSpecs returns a list of all specs.
func ListSpecs(db database.SpecsDatabase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		specs, err := db.ListSpecs()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := sendJSONResponse(w, specs); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

// ListSpecsWithResultID returns a spec related to a result.
func ListSpecsWithResultID(db database.SpecsDatabase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr, ok := mux.Vars(r)["id"]
		if !ok {
			http.Error(w, `router: no "id" key in router`, http.StatusInternalServerError)
			return
		}

		resultID, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		specs, err := db.GetSpecs(resultID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := sendJSONResponse(w, specs); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

// GetSpecs retrieves specs by its id.
func GetSpecs(db database.SpecsDatabase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
                idStr, ok := mux.Vars(r)["id"]
                if !ok {
                        http.Error(w, `router: no "id" key`, http.StatusInternalServerError)
                        return
                }

                specsID, err := strconv.ParseInt(idStr, 10, 64)
                if err != nil {
                        http.Error(w, err.Error(), http.StatusInternalServerError)
                        return
                }

                specs, err := db.GetSpecs(specsID)
                if err != nil {
                        http.Error(w, err.Error(), http.StatusInternalServerError)
                }

                if err := sendJSONResponse(w, specs); err != nil {
                        http.Error(w, err.Error(), http.StatusInternalServerError)
                }
	}
}

// AddSpecs saves the given specs. maybe unecessary why does this add a whole entry
func AddSpecs(db database.SpecsDatabase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

// DeleteSpecs deletes the specs with the given id. maybe unecessary
func DeleteSpecs(db database.SpecsDatabase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		idStr, ok := mux.Vars(r)["id"]
		if !ok {
			http.Error(w, `router: no "id" key`, http.StatusInternalServerError)
			return
		}

		specsID, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := db.DeleteSpecs(specsID); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

// UpdateSpecs updates the given specs. maybe unecessary
func UpdateSpecs(db database.SpecsDatabase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		// TODO: get specs values
		var specs database.Specs

		if err := db.UpdateSpecs(&specs); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
