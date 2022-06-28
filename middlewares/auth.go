package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/tienanh129902/go-rest-api/config"
	"github.com/tienanh129902/go-rest-api/constants"
	"github.com/tienanh129902/go-rest-api/datatransfers"
	"github.com/tienanh129902/go-rest-api/handlers"
)

func AuthenMiddleware(c *gin.Context) {
	token := strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer ")
	if token == "" {
		c.Set(constants.IsAuthenticatedKey, false)
		c.Next()
		return
	}
	claims, err := handlers.Handler.ParseToken(token, config.AppConfig.JWTSecret)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, datatransfers.Response{Error: err.Error()})
		return
	}
	c.Set(constants.IsAuthenticatedKey, true)
	c.Set(constants.UserIDKey, claims.ID)
	c.Next()
}
