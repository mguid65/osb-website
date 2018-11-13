package database

import (
	"encoding/json"
	"strings"
)

// SpecsDatabase provides thread-safe access to a database of specs.
type SpecsDatabase interface {
	// ListSpecs returns a list of all specs.
	ListSpecs() ([]*Specs, error)

	// ListSpecsCreatedBy returns a list of specs created by a user with the given id.
	ListSpecsCreatedBy(id int64) ([]*Specs, error)

	// GetSpecs retrieves specs by its id.
	GetSpecs(id int64) (Specs, error)

	// AddSpecs saves the given specs.
	AddSpecs(specs *Specs) error

	// DeleteSpecs deletes the specs with the given id.
	DeleteSpecs(id int64) error

	// UpdateSpecs updates the given specs.
	UpdateSpecs(specs *Specs) error
}

// Specs represents the Specs MySQL table.
type Specs struct {
	ID       int64          // specs ID
	ResultID int64          // connected result ID
	SysInfo  `json:"specs"` // system information
}

// SysInfo represents the `specs` JSON object stored in the Specs table.
type SysInfo struct {
	Vendor      string `json:"vendor"`       // CPU vendor
	Model       string `json:"model"`        // CPU model
	ClockSpeed  string `json:"speed"`        // CPU clock speed
	Threads     string `json:"threads"`      // number of physical CPU cores
	Overclocked string `json:"overclocked"`  // specifies if the CPU is overclocked
	ByteOrder   string `json:"byte_order"`   // CPU byte order
	PhysicalMem string `json:"physical_mem"` // physical memory
	VirtualMem  string `json:"virtual_mem"`  // virtual memory
	SwapMem     string `json:"swap_mem"`     // swap memory
}

// scanSpecs returns specs from a database row.
func scanSpecs(s rowScanner) (*Specs, error) {
	var (
		id       int64
		resultID int64
		sysInfo  string
	)
	if err := s.Scan(&id, &resultID, &sysInfo); err != nil {
		return nil, err
	}
	specs := &Specs{
		ID:       id,
		ResultID: resultID,
	}
	err := json.NewDecoder(strings.NewReader(sysInfo)).Decode(&specs.SysInfo)
	if err != nil {
		return nil, err
	}
	return specs, nil
}
