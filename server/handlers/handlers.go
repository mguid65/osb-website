package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"

	"github.com/mguid65/osb-website/server/database"
)

// Handler returns the OSB website route handler.
func Handler(db database.OSBDatabase) *mux.Router {
	r := mux.NewRouter()
	api := r.PathPrefix("/api/").Subrouter()
	addRootHandler(r)
	addUserHandlers(api, db)
	addResultHandlers(api, db)
	addSpecsHandlers(api, db)
	return r
}

func addRootHandler(r *mux.Router) {
	build := filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "mguid65", "osb-website", "build")
	err := filepath.Walk(build, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		prefix := fmt.Sprintf("/%s", info.Name())
		r.HandleFunc(prefix, func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, path)
		})
		return nil
	})
	if err != nil {
		panic(err)
	}
}

func addUserHandlers(r *mux.Router, db database.OSBDatabase) {
	r.HandleFunc("/users", ListUsers(db)).Methods(http.MethodGet)
	r.HandleFunc("/users/{id:[0-9]+}", GetUser(db)).Methods(http.MethodGet)
	r.HandleFunc("/users/register", AddUser(db)).Methods(http.MethodPost)
	r.HandleFunc("/users/delete/{id:[0-9]+}", DeleteUser(db)).Methods(http.MethodPost)
	r.HandleFunc("/users/update/{id:[0-9]+}", UpdateUser(db)).Methods(http.MethodPost)
}

func addResultHandlers(r *mux.Router, db database.OSBDatabase) {
	r.HandleFunc("/results", ListResults(db)).Methods(http.MethodGet)
	r.HandleFunc("/results/user/{id:[0-9]+}", ListResultsCreatedBy(db)).Methods(http.MethodGet)
	r.HandleFunc("/results/{id:[0-9]+}", GetResult(db)).Methods(http.MethodGet)
	r.HandleFunc("/results/submit", AddResult(db)).Methods(http.MethodPost)
	r.HandleFunc("/results/delete/{id:[0-9]+}", DeleteResult(db)).Methods(http.MethodPost)
	r.HandleFunc("/results/update/{id:[0-9]+}", UpdateResult(db)).Methods(http.MethodPost)
}

func addSpecsHandlers(r *mux.Router, db database.OSBDatabase) {
	r.HandleFunc("/specs", ListSpecs(db)).Methods(http.MethodGet)
	r.HandleFunc("/specs/result/{id:[0-9]+}", ListSpecsWithResultID(db)).Methods(http.MethodGet)
	r.HandleFunc("/specs/{id:[0-9]+}", GetSpecs(db)).Methods(http.MethodGet)
	r.HandleFunc("/specs/add/result/{id:[0-9]+}", AddSpecs(db)).Methods(http.MethodPost)
	r.HandleFunc("/specs/delete/{id:[0-9]+}", DeleteSpecs(db)).Methods(http.MethodPost)
	r.HandleFunc("/specs/update/{id:[0-9]+}", UpdateSpecs(db)).Methods(http.MethodPost)
}

func sendJSONResponse(w http.ResponseWriter, data interface{}) error {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	enc := json.NewEncoder(w)
	return enc.Encode(data)
}
