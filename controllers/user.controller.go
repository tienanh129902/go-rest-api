package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/tienanh129902/go-rest-api/datatransfers"
	"github.com/tienanh129902/go-rest-api/handlers"
	"github.com/tienanh129902/go-rest-api/models"
)

// GetUserById godoc
// @Summary Show user info
// @tags User
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param id path int true "User ID"
// @Success 200 {object} datatransfers.UserInfo "ok"
// @Failure 400 {object} datatransfers.Error "We need ID!!"
// @Failure 404 {object} datatransfers.Error "Can not find user"
// @Router /user/{id} [get]
func GET_User(c *gin.Context) {
	var err error
	var userInfo datatransfers.UserInfo
	if err = c.ShouldBindUri(&userInfo); err != nil {
		c.JSON(http.StatusBadRequest, datatransfers.Error{Error: err.Error()})
		return
	}
	var user models.User
	if user, err = handlers.Handler.RetrieveUser(userInfo.ID); err != nil {
		c.JSON(http.StatusNotFound, datatransfers.Error{Error: "cannot find user"})
		return
	}
	c.JSON(http.StatusOK, datatransfers.Data{Data: datatransfers.UserInfo{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Bio:       user.Bio,
		CreatedAt: user.CreatedAt,
	}})
}

// UpdateUserData
// @tags User
// @Summary Update user data
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param id path int true "User ID"
// @Success 200 {object} datatransfers.UserInfo "ok"
// @Failure 400 {object} datatransfers.Error "We need ID!!"
// @Failure 404 {object} datatransfers.Error "Can not find user"
// @Router /user/ [patch]
func PATCH_User(c *gin.Context) {
	var err error
	var user datatransfers.UserUpdate
	if err = c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, datatransfers.Error{Error: err.Error()})
		return
	}
	if _, err = handlers.Handler.UpdateUser(handlers.Handler.Me(c), user); err != nil {
		c.JSON(http.StatusNotModified, datatransfers.Error{Error: "failed updating user"})
		return
	}
	c.JSON(http.StatusOK, datatransfers.Data{Data: user})
}

// GetCurrentUser godoc
// @Summary Show current user
// @Accept  json
// @tags User
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {object} datatransfers.UserInfo "ok"
// @Failure 404 {object} datatransfers.Error "Not found"
// @Router /user/me [get]
func GET_Me(c *gin.Context) {
	var err error
	var user models.User
	if user, err = handlers.Handler.RetrieveUser(handlers.Handler.Me(c)); err != nil {
		c.JSON(http.StatusNotFound, datatransfers.Error{Error: "cannot find user"})
		return
	}
	c.JSON(http.StatusOK, datatransfers.Data{Data: datatransfers.UserInfo{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Bio:       user.Bio,
		CreatedAt: user.CreatedAt,
	}})
}
