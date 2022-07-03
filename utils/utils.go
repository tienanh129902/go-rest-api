package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/tienanh129902/go-rest-api/constants"
	"github.com/tienanh129902/go-rest-api/datatransfers"
)

func AuthenOnly(c *gin.Context) {
	if !c.GetBool(constants.IsAuthenticatedKey) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, datatransfers.Error{Error: "user not authenticated"})
	}
}
