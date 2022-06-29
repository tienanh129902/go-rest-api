package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/tienanh129902/go-rest-api/config"
	"github.com/tienanh129902/go-rest-api/datatransfers"
	"github.com/tienanh129902/go-rest-api/models"
)

type UserHandlerFunction interface {
	RetrieveUser(id uint) (user models.User, err error)
	UpdateUser(id uint, userUpdate datatransfers.UserUpdate) (user models.User, err error)
	Me(c *gin.Context) (id uint)
}

func (m *module) RetrieveUser(id uint) (user models.User, err error) {
	if user, err = m.db.userOrmer.GetOneByID(id); err != nil {
		return models.User{}, fmt.Errorf("cannot find user with id %d", id)
	}
	return
}

func (m *module) UpdateUser(id uint, userUpdate datatransfers.UserUpdate) (user models.User, err error) {
	if err = m.db.userOrmer.UpdateUser(models.User{
		Email: userUpdate.Email,
		Bio:   userUpdate.Bio,
	}); err != nil {
		return models.User{}, errors.New("cannot update user")
	}
	return
}

func (m *module) Me(c *gin.Context) (id uint) {
	var user models.User
	token := strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer ")
	fmt.Println(token)
	claims, err := Handler.ParseToken(token, config.AppConfig.JWTSecret)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, datatransfers.Response{Error: err.Error()})
		return
	}
	if user, err = m.db.userOrmer.GetOneByID(claims.Sub); err != nil {
		return user.ID
	}
	return
}
