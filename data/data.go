package data

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/FredBainbridge/PokerAPI/data/models"
	"github.com/FredBainbridge/PokerAPI/secrets"
	_ "github.com/denisenkom/go-mssqldb"
)

func Users(slackids ...string) []models.User {
	var users []models.User
	secret, nil := secrets.GetSecrets()
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d", secret.DBServer, secret.DBUserName, secret.DBPassword, secret.DBPort)
	conn, err := sql.Open("mssql", connString)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	if len(slackids) > 0 {
		for i := range slackids {
			id := strings.Trim(slackids[i], "\"")
			q := fmt.Sprintf("select * from Poker.dbo.[User] where SlackID = '%s'", id)
			rows, err := conn.Query(q)
			if err != nil {
				log.Fatal((err))
			}
			users = append(users, getUserObjects(rows)...)
		}
	} else { //all users
		rows, err := conn.Query("select * from Poker.dbo.[User]")
		if err != nil {
			log.Fatal(err)
		}

		defer rows.Close()
		users = append(users, getUserObjects(rows)...)
	}
	return users
}

func getUserObjects(rows *sql.Rows) []models.User {
	defer rows.Close()
	var users []models.User
	for rows.Next() {
		user := models.User{}
		err := rows.Scan(&user.ID, &user.SlackID, &user.UserName, &user.EmailAddress, &user.RealName, &user.AvatarHash, &user.AvatarIndex, &user.SlackUserName)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}
	return users
}
