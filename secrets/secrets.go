package secrets

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

func GetSecrets() (*Secrets, error) {
	gp := os.Getenv("GOPATH")
	ap := path.Join(gp, "secrets.json")
	data, err := ioutil.ReadFile(ap)
	if err != nil {
		fmt.Print(err)
		return nil, err
	}
	var secrets Secrets
	err = json.Unmarshal(data, &secrets)
	if err != nil {
		fmt.Println("error:", err)
		return nil, err
	}
	return &secrets, nil
}

type Secrets struct {
	DBName           string
	DBServer         string
	DBUserName       string
	DBPassword       string
	DBPort           int
	MavenAPIUrl      string
	MavenAPIPassword string
}
