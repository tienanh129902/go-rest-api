package handlers

import (
	"fmt"
	"log"

	"github.com/tienanh129902/go-rest-api/datatransfers"
	"github.com/tienanh129902/go-rest-api/models"
)

type TokenHandlerFunction interface {
	CreateToken(credentials datatransfers.JWTToken) (err error)
	DeleteToken(refreshToken string) (err error)
}

func (m *module) CreateToken(credentials datatransfers.JWTToken) (err error) {

	if _, err = m.db.tokenOrmer.InsertToken(models.Token{
		Access_token:  credentials.Access_token,
		Refresh_token: credentials.Refresh_token,
	}); err != nil {
		log.Print(err)
		return fmt.Errorf("error inserting token. %v", err)
	}
	return
}

func (m *module) DeleteToken(refreshToken string) (err error) {
	var token models.Token
	if token, err = m.db.tokenOrmer.GetTokenByRefreshToken(refreshToken); err != nil {
		return fmt.Errorf("cannot find token")
	}
	return m.db.tokenOrmer.DeleteToken(token)
}
