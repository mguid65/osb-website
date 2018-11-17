package database

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// ResultDatabase provides thread-safe access to a database of results.
type ResultDatabase interface {
	// ListResults returns a list of all results.
	ListResults() ([]*Result, error)

	// ListResultsCreatedBy returns a list of results created by a user with the given id.
	ListResultsCreatedBy(id int64) ([]*Result, error)

	// GetResult retrieves a result by its id.
	GetResult(id int64) (*Result, error)

	// AddResult saves a given result.
	AddResult(res *Result) (int64, error)

	// DeleteResult deletes a result with the given id.
	DeleteResult(id int64) error

	// UpdateResult updates a given result.
	UpdateResult(res *Result) error
}

// Result holds the metadata about a result.
type Result struct {
	ID      int64
	UserID  int64
	SpecsID int64
	Results []Score `json:"results"`
}

// Score holds the metadata for a benchmark algorithm run.
type Score struct {
	Name  string   `json:"name"`  // algorithm name
	Time  Duration `json:"time"`  // total elapsed time
	Score float64  `json:"score"` // total score
}

// scanResult returns a result from a database row.
func scanResult(s rowScanner) (*Result, error) {
	var (
		id      int64
		userID  int64
		specsID int64
		results string
	)
	if err := s.Scan(&id, &userID, &specsID, &results); err != nil {
		return nil, err
	}
	result := &Result{
		ID:      id,
		UserID:  userID,
		SpecsID: specsID,
	}
	err := json.NewDecoder(strings.NewReader(results)).Decode(&result.Results)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Duration wraps a time.Duration value for added json support.
type Duration struct {
	time.Duration
}

// MarshalJSON serializes a time.Duration value to JSON.
func (d *Duration) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.Duration.String())
}

// UnmarshalJSON deserializes a time.Duration value to JSON.
func (d *Duration) UnmarshalJSON(data []byte) error {
	var v interface{}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	var f float64
	switch val := v.(type) {
	case float64:
		f = val
	case string:
		parsed, err := strconv.ParseFloat(val, 64)
		if err != nil {
			return err
		}
		f = parsed
	default:
		return fmt.Errorf("could not unmarshal %v: unsupported type %T", v, v)
	}
	d.Duration = time.Duration(f) * time.Nanosecond
	return nil
}
