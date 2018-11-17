package database_test

import (
	"testing"
	"time"

	"github.com/mguid65/osb-website/server/database"
)

func testResultsDB(t *testing.T, db database.ResultDatabase) {
	result := &database.Result{
		UserID:  2,
		Results: make([]database.Score, 0),
	}

	id, err := db.AddResult(result)
	if err != nil {
		t.Fatal(err)
	}

	result.ID = id
	result.Results = append(result.Results, database.Score{
		Name:  "Total",
		Time:  database.Duration{Duration: 25 * time.Millisecond},
		Score: 1000,
	})
	if err := db.UpdateResult(result); err != nil {
		t.Error(err)
	}

	gotResult, err := db.GetResult(result.ID)
	if err != nil {
		t.Error(err)
	}
	for i, got := range gotResult.Results {
		want := result.Results[i]
		if got.Name != want.Name || got.Time.String() != want.Time.String() || got.Score != want.Score {
			t.Errorf("Update scores: got %v, want %v", got, want)
		}
	}

	if err := db.DeleteResult(result.ID); err != nil {
		t.Error(err)
	}

	if _, err := db.GetResult(result.ID); err != nil {
		t.Error("want non-nil error")
	}
}
