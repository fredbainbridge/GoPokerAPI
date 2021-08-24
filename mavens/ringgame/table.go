package ringgame

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/FredBainbridge/PokerAPI/mavens/operations"
	"github.com/FredBainbridge/PokerAPI/secrets"
)
func Tables() ([]Table, error) {
	secrets, err := secrets.GetSecrets()
	if(err != nil) {
		return nil, err
	}
	var params = make(map[string]string)
	params["Command"] = "RingGamesList"
	params["Password"] = secrets.MavenAPIPassword
	params["JSON"] = "Yes"
	params["Fields"] = "Name,Game"
	response, err := operations.EncodedHttpClientRequest("POST", secrets.MavenAPIUrl, params)
	if(err != nil) {
		log.Fatal(err)
		return nil, err
	}
	fmt.Println(response)
	var mavResp mavenTableResponse
	json.Unmarshal([]byte(response), &mavResp)
	var tables []Table
	for i := 0; i < mavResp.RingGames; i++ {
		tables = append(tables, Table{Game: mavResp.Game[i], Name: mavResp.Name[i]})
	}
	return tables, nil
}

type mavenTableResponse struct {
	Result string
	RingGames int
	Name []string
	Game []string
}
type Table struct {
	Game string
	Name string
}