package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/tienanh129902/go-rest-api/constants"
	"github.com/tienanh129902/go-rest-api/datatransfers"
	"github.com/tienanh129902/go-rest-api/handlers"
)

func POST_Login(c *gin.Context) {
	var err error
	var user datatransfers.UserLogin
	if err = c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, datatransfers.Response{Error: err.Error()})
		return
	}
	var token datatransfers.JWTToken
	if token, err = handlers.Handler.AuthenticateUser(user); err != nil {
		c.JSON(http.StatusUnauthorized, datatransfers.Response{Error: "incorrect username or password"})
		return
	}
	c.JSON(http.StatusOK, datatransfers.Response{AccessToken: token.Access_token, RefreshToken: token.Refresh_token})
}

func POST_Register(c *gin.Context) {
	var err error
	var user datatransfers.UserSignup
	if err = c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, datatransfers.Response{Error: err.Error()})
		return
	}
	if err = handlers.Handler.RegisterUser(user); err != nil {
		c.JSON(http.StatusUnauthorized, datatransfers.Response{Error: "failed registering user"})
		return
	}
	c.JSON(http.StatusCreated, datatransfers.Response{Status: "user created"})
}

func POST_Logout(c *gin.Context) {
	var err error
	var refresh datatransfers.UserLogout
	if err = c.ShouldBind(&refresh); err != nil {
		c.JSON(http.StatusBadRequest, datatransfers.Response{Error: err.Error()})
		return
	}
	if err = handlers.Handler.LogoutUser(refresh.RefreshToken); err != nil {
		c.JSON(http.StatusUnauthorized, datatransfers.Response{Error: "invalid token"})
		return
	}
	c.Set(constants.IsAuthenticatedKey, false)
	c.JSON(http.StatusNoContent, datatransfers.Response{Status: "Logout successful"})
}
