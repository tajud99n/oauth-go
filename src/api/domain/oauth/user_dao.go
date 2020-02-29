package oauth

import "github.com/tajud99n/go-micro/src/api/utils/errors"

const (
	queryGetUserByUsernameAndPassword = "SELECT id, username FROM users WHERE username=? AND password=?;"
)

var (
	users = map[string]*User {
		"john": {Id: 123, Username: "doe"},
	}
)

func GetUserByUsernameAndPassword(username string, password string) (*User, errors.APIError) {
	user := users[username]
	if user == nil {
		return nil, errors.NewNotFoundError("no user found with given parameters")
	}
	return user, nil
}