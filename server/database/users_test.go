package database_test

import (
	"testing"

	"github.com/mguid65/osb-website/server/database"
)

func testUserDB(t *testing.T, db database.UserDatabase) {
	user := &database.User{
		Name:     "test",
		Email:    "test@test.com",
		Password: "supersecretpassword",
	}

	id, err := db.AddUser(user)
	if err != nil {
		t.Fatal(err)
	}

	user.ID = id
	user.Name = "updated"
	if err := db.UpdateUser(user); err != nil {
		t.Fatal(err)
	}

	gotUser, err := db.GetUser(user.ID)
	if err != nil {
		t.Fatal(err)
	}

	if got, want := gotUser.Name, user.Name; got != want {
		t.Errorf("Update user: got %q, want %q", got, want)
	}

	gotUserCred, err := db.GetUserByCredentials(user.Name, user.Password)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := gotUserCred.ID, user.ID; got != want {
		t.Errorf("Get user by credentials: got %d, want %d", got, want)
	}

	if err := db.DeleteUser(user.ID); err != nil {
		t.Fatal(err)
	}

	if _, err := db.GetUser(user.ID); err == nil {
		t.Fatal("want non-nil error")
	}
	if _, err := db.GetUserByCredentials(user.Name, user.Password); err == nil {
		t.Fatal("want non-nil error")
	}
}
