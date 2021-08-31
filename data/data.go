package data

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/FredBainbridge/PokerAPI/data/models"
	"github.com/FredBainbridge/PokerAPI/secrets"
	_ "github.com/denisenkom/go-mssqldb"
)

func Players() {
	secret, nil := secrets.GetSecrets()
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d", secret.DBServer, secret.DBUserName, secret.DBPassword, secret.DBPort)
	conn, err := sql.Open("mssql", connString)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	//stmt, err := conn.Prepare()
	rows, err := conn.Query("select * from Poker.dbo.[User]")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	user := models.User{}

	for rows.Next() {
		err := rows.Scan(&user.ID, &user.SlackID, &user.UserName, &user.EmailAddress, &user.RealName, &user.AvatarHash, &user.AvatarIndex, &user.SlackUserName)

		if err != nil {
			log.Fatal(err)
		}
		log.Println(user)
	}

}
