package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/tienanh129902/go-rest-api/constants"
	"github.com/tienanh129902/go-rest-api/datatransfers"
	"github.com/tienanh129902/go-rest-api/handlers"
)

// Login
// @Summary Login user
// @tags Authentication
// @Accept  json
// @Produce  json
// @Param user body datatransfers.UserLogin true "User info"
// @Success 200 {object} datatransfers.Token "Ok"
// @Failure 400 {object} datatransfers.Error "Bad request"
// @Failure 401 {object} datatransfers.Error "Unauthorized"
// @Router /auth/login [post]
func POST_Login(c *gin.Context) {
	var err error
	var user datatransfers.UserLogin
	if err = c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, datatransfers.Error{Error: err.Error()})
		return
	}
	var token datatransfers.JWTToken
	if token, err = handlers.Handler.AuthenticateUser(user); err != nil {
		c.JSON(http.StatusUnauthorized, datatransfers.Error{Error: "incorrect username or password"})
		return
	}
	c.JSON(http.StatusOK, datatransfers.Token{AccessToken: token.Access_token, RefreshToken: token.Refresh_token})
}

// Register
// @Summary user registration
// @tags Authentication
// @Accept  json
// @Produce  json
// @Param user body datatransfers.UserSignup true "User info"
// @Success 200 {object} datatransfers.Status "Ok"
// @Failure 400 {object} datatransfers.Error "Bad request"
// @Failure 401 {object} datatransfers.Error "Unauthorized"
// @Router /auth/signup [post]
func POST_Register(c *gin.Context) {
	var err error
	var user datatransfers.UserSignup
	if err = c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, datatransfers.Error{Error: err.Error()})
		return
	}
	if err = handlers.Handler.RegisterUser(user); err != nil {
		c.JSON(http.StatusUnauthorized, datatransfers.Error{Error: "failed registering user"})
		return
	}
	c.JSON(http.StatusCreated, datatransfers.Status{Status: "user created"})
}

// Logout
// @Summary Logout user
// @tags Authentication
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param user body datatransfers.UserLogout true "Refresh token"
// @Success 200 {object} datatransfers.Status "Ok"
// @Failure 400 {object} datatransfers.Error "Bad request"
// @Failure 401 {object} datatransfers.Error "Unauthorized"
// @Router /auth/logout [post]
func POST_Logout(c *gin.Context) {
	var err error
	var refresh datatransfers.UserLogout
	if err = c.ShouldBind(&refresh); err != nil {
		c.JSON(http.StatusBadRequest, datatransfers.Error{Error: err.Error()})
		return
	}
	if err = handlers.Handler.LogoutUser(refresh.RefreshToken); err != nil {
		c.JSON(http.StatusUnauthorized, datatransfers.Error{Error: "invalid token"})
		return
	}
	c.Set(constants.IsAuthenticatedKey, false)
	c.JSON(http.StatusNoContent, datatransfers.Status{Status: "Logout successful"})
}
