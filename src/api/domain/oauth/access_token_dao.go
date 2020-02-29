package oauth

import (
	"fmt"

	"github.com/tajud99n/go-micro/src/api/utils/errors"
)

var (
	tokens = make(map[string]*AccessToken, 0)
)

func (at *AccessToken) Save() errors.APIError {
	at.AccessToken = fmt.Sprintf("USR_%d", at.UserId)
	tokens[at.AccessToken] = at
	return nil
}

func GetAccessTokenByToken(accessToken string) (*AccessToken, errors.APIError) {
	token := tokens[accessToken]
	if token == nil || token.IsExpired() {
		return nil, errors.NewNotFoundError("no access token found with given parameters")
	}
	return token, nil
}
