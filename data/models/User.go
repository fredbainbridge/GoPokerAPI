package models

type User struct {
	ID            int
	SlackID       string
	UserName      string
	EmailAddress  string
	RealName      string
	AvatarHash    string
	AvatarIndex   string
	SlackUserName string
}
