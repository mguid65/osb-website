package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/mguid65/osb-website/server/database"
)

// ListUsers returns a list of all users.
func ListUsers(db database.UserDatabase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := db.ListUsers()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := sendJSONResponse(w, users); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

// GetUser retrieves a user by its id.
func GetUser(db database.UserDatabase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr, ok := mux.Vars(r)["id"]
		if !ok {
			http.Error(w, `route: no key "id"`, http.StatusInternalServerError)
		}

		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		user, err := db.GetUser(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := sendJSONResponse(w, user); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

// AddUser saves a given user.
func AddUser(db database.UserDatabase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		// TODO: get user values and hash password
		var user database.User

		id, err := db.AddUser(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		log.Println("successfully added result id", id)

		w.WriteHeader(http.StatusOK)
	}
}

// DeleteUser deletes a user with the given id.
func DeleteUser(db database.UserDatabase) http.HandlerFunc {
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

		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := db.DeleteUser(id); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

// UpdateUser updates a given user.
func UpdateUser(db database.UserDatabase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		// TODO: get result values
		var user database.User

		if err := db.UpdateUser(&user); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
