package services

import (
	"time"
	"github.com/tajud99n/go-micro/src/api/utils/errors"
	"github.com/tajud99n/oauth-go/src/api/domain/oauth"
)

type oauthService struct{}

type oauthServiceInterface interface {
	CreateAccessToken(request oauth.AccessTokenRequest) (*oauth.AccessToken, errors.APIError)
	GetAccessToken(accessToken string) (*oauth.AccessToken, errors.APIError)
}

var (
	OauthService oauthServiceInterface
)

func init() {
	OauthService = &oauthService{}
}

func (s *oauthService) CreateAccessToken(request oauth.AccessTokenRequest) (*oauth.AccessToken, errors.APIError) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	user, err := oauth.GetUserByUsernameAndPassword(request.Username, request.Password)
	if err != nil {
		return nil, err
	}
	
	token := oauth.AccessToken{
		UserId: user.Id,
		Expires: time.Now().UTC().Add(24*time.Hour).Unix(),
	}

	if err := token.Save(); err != nil {
		return nil, err
	}
	return &token, nil
}

func (s *oauthService) GetAccessToken(accessToken string) (*oauth.AccessToken, errors.APIError) {
	token, err := oauth.GetAccessTokenByToken(accessToken)
	if err != nil {
		return nil, err
	}

	return token, nil
}