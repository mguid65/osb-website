package database_test

import (
	"testing"

	"github.com/mguid65/osb-website/server/database"
)

func testSpecsDB(t *testing.T, db database.SpecsDatabase) {
	specs := &database.Specs{
		ResultID: 6,
		SysInfo:  database.SysInfo{},
	}

	id, err := db.AddSpecs(specs)
	if err != nil {
		t.Fatal(err)
	}

	specs.ID = id
	specs.SysInfo.Vendor = "GenuineIntel"
	if err := db.UpdateSpecs(specs); err != nil {
		t.Error(err)
	}

	gotSpecs, err := db.GetSpecs(specs.ID)
	if err != nil {
		t.Error(err)
	}
	if got, want := gotSpecs.Vendor, specs.Vendor; got != want {
		t.Errorf("Update specs: got %q, want %q", got, want)
	}

	if err := db.DeleteSpecs(specs.ID); err != nil {
		t.Error(err)
	}

	if _, err := db.GetSpecs(specs.ID); err == nil {
		t.Error("want non-nil error")
	}
}
