package handlers

import (
	"encoding/json"
	"net/http"

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
	r.PathPrefix("/static/css/").Handler(http.StripPrefix("/static/css/", http.FileServer(http.Dir("./build/static/css/"))))
	r.PathPrefix("/static/js/").Handler(http.StripPrefix("/static/js/", http.FileServer(http.Dir("./build/static/js/"))))
	r.PathPrefix("/service-worker.js").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./build/service-worker.js")
	})
	r.PathPrefix("/manifest.json").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./build/manifest.json")
	})
	r.PathPrefix("/favicon.ico").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./build/favicon.ico")
	})
	r.PathPrefix("/asset-manifest.json").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./build/asset-manifest.json")
	})
	r.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		http.ServeFile(w, r, "./build/index.html")
	})
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
