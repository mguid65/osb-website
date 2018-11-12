package database

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

// SpecsDatabase provides thread-safe access to a database of specs.
type SpecsDatabase interface {
	// ListSpecs returns a list of all specs.
	ListSpecs() ([]*Specs, error)

	// ListResultsCreatedBy returns a list of specs created by a user with the given id.
	ListSpecsCreatedBy(id int64) ([]*Specs, error)

	// GetUser retrieves specs by its id.
	GetSpecs(id int64) (Specs, error)

	// AddResult saves the given specs.
	AddSpecs(specs *Specs) error

	// DeleteResult deletes the specs with the given id.
	DeleteSpecs(id int64) error

	// UpdateResult updates the given specs.
	UpdateSpecs(specs *Specs) error
}
