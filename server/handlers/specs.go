package handlers

import (
	"net/http"

	"github.com/mguid65/osb-website/server/database"
)

// ListSpecs returns a list of all specs.
func ListSpecs(db database.SpecsDatabase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		panic("implement me")
	}
}

// ListSpecsCreatedBy returns a list of specs created by a user with the given id.
func ListSpecsCreatedBy(db database.SpecsDatabase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		panic("implement me")
	}
}

// GetSpecs retrieves specs by its id.
func GetSpecs(db database.SpecsDatabase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		panic("implement me")
	}
}

// AddSpecs saves the given specs.
func AddSpecs(db database.SpecsDatabase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		panic("implement me")
	}
}

// DeleteSpecs deletes the specs with the given id.
func DeleteSpecs(db database.SpecsDatabase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		panic("implement me")
	}
}

// UpdateSpecs updates the given specs.
func UpdateSpecs(db database.SpecsDatabase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		panic("implement me")
	}
}
