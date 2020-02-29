package oauth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tajud99n/go-micro/src/api/utils/errors"
	"github.com/tajud99n/oauth-go/src/api/domain/oauth"
	"github.com/tajud99n/oauth-go/src/api/services"
)

func CreateAccessToken(c *gin.Context) {
	var request oauth.AccessTokenRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		apiErr := errors.NewBadRequestError("invalid json body")
		c.JSON(apiErr.Status(), apiErr)
		return
	}
	token, err := services.OauthService.CreateAccessToken(request)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, token)
}

func GetAccessToken(c *gin.Context) {
	tokenId := c.Param("token_id")

	token, err := services.OauthService.GetAccessToken(tokenId)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, token)
}
