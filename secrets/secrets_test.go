package secrets_test

import (
	"testing"

	"github.com/FredBainbridge/PokerAPI/secrets"
)

func TestGetSecrets(t *testing.T) {
	secrets, err := secrets.GetSecrets()
	if(err != nil) {
		t.Errorf("Failed to get secrets. %q", err)
	}
	if(len(secrets.DBName) == 0) {
		t.Error("DBName is blank or nil")
	}
	if(len(secrets.DBServer) == 0) {
		t.Error("DBServer is blank or nil")
	}
	if(len(secrets.MavenAPIPassword) == 0) {
		t.Error("Password is blank or nil")
	}
	if(len(secrets.MavenAPIUrl) == 0) {
		t.Error("MavensAPIURL is blank or nil")
	}
}