package handlers

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"

	"github.com/mguid65/osb-website/server/database"
)

// ListUsers returns a list of all users.
func ListUsers(db database.UserDatabase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		users, err := db.ListUsers()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := sendJSONResponse(w, users); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

// GetUser retrieves a user by its id.
func GetUser(db database.UserDatabase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		idStr, ok := mux.Vars(r)["id"]
		if !ok {
			http.Error(w, `route: no key "id"`, http.StatusInternalServerError)
			return
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

		if err := r.ParseForm(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var user database.User
		for k, v := range r.Form {
			fmt.Println(k, v)
			switch k {
			case "email":
				user.Email = v[0]
			case "username":
				user.Name = v[0]
			case "password":
				hash := sha256.New()
				if _, err := io.Copy(hash, strings.NewReader(v[0])); err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				user.Password = hex.EncodeToString(hash.Sum(nil))
			}
		}

		if _, err := db.AddUser(&user); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "successfully added user %s\n", user.Name)

		w.WriteHeader(http.StatusOK)
		http.Redirect(w, r, "/", http.StatusOK)
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

		// TODO: get user values
		var user database.User

		if err := db.UpdateUser(&user); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
