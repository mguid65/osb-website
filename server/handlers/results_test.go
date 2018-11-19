package handlers_test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gorilla/mux"

	"github.com/mguid65/osb-website/server/database"
	"github.com/mguid65/osb-website/server/handlers"
)

type resultHandlerTest struct {
	Name       string
	DBName     string
	Body       string
	StatusCode int
}

type mockResultsDB struct{}

func (db *mockResultsDB) ListResults() ([]*database.Result, error) {
	return []*database.Result{
		{
			ID:     1,
			UserID: 1,
			Scores: database.Scores{
				{Name: "Total", Time: database.Duration{Duration: 123456789 * time.Nanosecond}, Score: 1000},
			},
		},
	}, nil
}

func (db *mockResultsDB) ListResultsCreatedBy(id int64) ([]*database.Result, error) {
	return nil, nil
}

func (db *mockResultsDB) GetResult(id int64) (*database.Result, error) {
	return nil, nil
}

func (db *mockResultsDB) AddResult(res *database.Result) (int64, error) {
	return 0, nil
}

func (db *mockResultsDB) DeleteResult(id int64) error {
	return nil
}

func (db *mockResultsDB) UpdateResult(res *database.Result) error {
	return nil
}

func TestListResults(t *testing.T) {
	tt := []resultHandlerTest{
		{
			Name:       "List Existing results",
			Body:       `[{"ID":1,"UserID":1,"scores":[{"name":"Total","time":"123.456789ms","score":1000}]}]`,
			StatusCode: http.StatusOK,
		},
	}

	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "/results", nil)
			if err != nil {
				t.Fatal(err)
			}

			rec := httptest.NewRecorder()

			r := mux.NewRouter()
			r.HandleFunc("/results", handlers.ListResults(&mockResultsDB{})).Methods("GET")
			r.ServeHTTP(rec, req)

			resp := rec.Result()
			if got, want := resp.StatusCode, tc.StatusCode; got != want {
				t.Errorf("status code: want %d, got %d", want, got)
			}
			defer resp.Body.Close()

			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				t.Fatal(err)
			}
			b = bytes.TrimSpace(b)

			if got, want := string(b), tc.Body; got != want {
				t.Errorf("contents: got %q, want %q", got, want)
			}
		})
	}
}

func TestListResultsCreatedBy(t *testing.T) {

}

func GetResult(t *testing.T) {

}

func TestAddResult(t *testing.T) {

}

func TestDeleteResult(t *testing.T) {

}

func TestUpdateResult(t *testing.T) {

}
