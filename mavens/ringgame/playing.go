package ringgame

import (
	"encoding/json"
	"log"

	"github.com/FredBainbridge/PokerAPI/mavens/operations"
	"github.com/FredBainbridge/PokerAPI/secrets"
)

func Playing(tableNames ...string) ([]SeatedPlayer, error) {
	var seatedPlayers []SeatedPlayer
	var params = make(map[string]string)
	secrets, err := secrets.GetSecrets()
	if(err != nil) {
		return nil, err
	}
	params["Command"] = "RingGamesPlaying"
	params["Password"] = secrets.MavenAPIPassword
	params["JSON"] = "Yes"
	if len(tableNames) == 0 {
		tables, err := Tables()
		if err != nil {
			log.Fatalf("Cannot find tables. %q", err)
			return nil, err
		}
		for i := range tables {
			params["Name"] = tables[i].Name
			seatedPlayers = append(seatedPlayers, getSeatedPlayers(tables[i].Name,params, *secrets)...)
		}
	} else {
		for i:= range tableNames {
			params["Name"] = tableNames[i]
			seatedPlayers = append(seatedPlayers, getSeatedPlayers(tableNames[i], params, *secrets)...)		
		}
	}
	return seatedPlayers, nil
}

func getSeatedPlayers (tableName string, params map[string]string, secrets secrets.Secrets) []SeatedPlayer {
	var seatedPlayers []SeatedPlayer
	var mavResp mavenGamePlayingResponse
	params["Name"] = tableName
	response, err := operations.EncodedHttpClientRequest("POST",secrets.MavenAPIUrl,params)
	if(err != nil) {
		log.Fatalf("Failed when getting %q", tableName)
	}
	json.Unmarshal([]byte(response),&mavResp)
	for i := 0; i < mavResp.Count; i++ {
		seatedPlayers = append(seatedPlayers, SeatedPlayer{Name: mavResp.Player[i], TableName: tableName})
	}
	return seatedPlayers
}

type SeatedPlayer struct {
	Name string
	TableName string
}
type mavenGamePlayingResponse struct {
	Result string
	Count int
	Player []string
	Seat []int
	Chips []int
	Net []int
	Away []int
}