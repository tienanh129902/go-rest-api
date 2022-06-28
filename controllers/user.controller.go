package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/tienanh129902/go-rest-api/datatransfers"
	"github.com/tienanh129902/go-rest-api/handlers"
	"github.com/tienanh129902/go-rest-api/models"
)

func GET_User(c *gin.Context) {
	var err error
	var userInfo datatransfers.UserInfo
	if err = c.ShouldBindUri(&userInfo); err != nil {
		c.JSON(http.StatusBadRequest, datatransfers.Response{Error: err.Error()})
		return
	}
	var user models.User
	if user, err = handlers.Handler.RetrieveUser(userInfo.ID); err != nil {
		c.JSON(http.StatusNotFound, datatransfers.Response{Error: "cannot find user"})
		return
	}
	c.JSON(http.StatusOK, datatransfers.Response{Data: datatransfers.UserInfo{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Bio:       user.Bio,
		CreatedAt: user.CreatedAt,
	}})
}

func PATCH_User(c *gin.Context) {
	var err error
	var user datatransfers.UserUpdate
	if err = c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, datatransfers.Response{Error: err.Error()})
		return
	}
	if _, err = handlers.Handler.UpdateUser(handlers.Handler.Me(c), user); err != nil {
		c.JSON(http.StatusNotModified, datatransfers.Response{Error: "failed updating user"})
		return
	}
	c.JSON(http.StatusOK, datatransfers.Response{Status: "User data updated", Data: user})
}

func GET_Me(c *gin.Context) {
	var err error
	var user models.User
	if user, err = handlers.Handler.RetrieveUser(handlers.Handler.Me(c)); err != nil {
		c.JSON(http.StatusNotFound, datatransfers.Response{Error: "cannot find user"})
		return
	}
	c.JSON(http.StatusOK, datatransfers.Response{Data: datatransfers.UserInfo{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Bio:       user.Bio,
		CreatedAt: user.CreatedAt,
	}})
}
