package app

import (
	"github.com/tajud99n/go-micro/src/api/controllers/ping"
	"github.com/tajud99n/oauth-go/src/api/controllers/oauth"
)

func mapURLS() {
	router.GET("/ping", ping.Ping)
	router.POST("/oauth/access_token", oauth.CreateAccessToken)
	router.GET("/oauth/access_token/:token_id", oauth.GetAccessToken)
}
