package models

import "database/sql"

type User struct {
	ID            int
	SlackID       string
	UserName      string
	EmailAddress  string
	RealName      string
	AvatarHash    sql.NullString
	AvatarIndex   string
	SlackUserName sql.NullString
}
