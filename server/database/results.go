package database

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
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
	ID     int64
	UserID int64
	Scores `json:"scores"`
}

// Score holds the metadata for a benchmark algorithm run.
type Score struct {
	Name  string   `json:"name"`  // algorithm name
	Time  Duration `json:"time"`  // total elapsed time
	Score float64  `json:"score"` // total score
}

// Scores implements driver.Valuer and sql.Scanner.
type Scores []Score

// Value returns a driver.Value for a slice of scores.
func (s Scores) Value() (driver.Value, error) {
	return json.Marshal(s)
}

// Scan scans a slice of scores from the database.
func (s *Scores) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), *s)
}

// scanResult returns a result from a database row.
func scanResult(s rowScanner) (*Result, error) {
	var (
		id     int64
		userID int64
		scores string
	)
	if err := s.Scan(&id, &userID, &scores); err != nil {
		return nil, err
	}
	result := &Result{
		ID:     id,
		UserID: userID,
	}
	err := json.NewDecoder(strings.NewReader(scores)).Decode(&result.Scores)
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

	switch val := v.(type) {
	case int, float64:
		duration := time.Duration(0)
		if i, ok := val.(int); ok {
			duration = time.Duration(i) * time.Nanosecond
		}
		if f, ok := val.(float64); ok {
			duration = time.Duration(f) * time.Nanosecond
		}
		d.Duration = duration
	case string:
		parsed, err := time.ParseDuration(val)
		if err != nil {
			return err
		}
		d.Duration = parsed
	default:
		return fmt.Errorf("could not unmarshal %v: unsupported type %T", v, v)
	}
	return nil
}
