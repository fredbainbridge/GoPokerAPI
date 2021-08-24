package ringgame

import (
	"github.com/FredBainbridge/PokerAPI/mavens/operations"
	"github.com/FredBainbridge/PokerAPI/secrets"
)
func AdminMessage(tableName string, message string) (string, error) {
	secrets, err := secrets.GetSecrets()
	if(err != nil) {
		return "error", err
	}
	var params = make(map[string]string)
	params["Command"] = "RingGamesMessage"
	params["Name"] = tableName
	params["Message"] = message
	params["Password"] = secrets.MavenAPIPassword
	params["JSON"] = "Yes"
	return operations.EncodedHttpClientRequest("POST", secrets.MavenAPIUrl, params)
}
