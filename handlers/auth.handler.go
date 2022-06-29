package handlers

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
	"golang.org/x/crypto/bcrypt"

	"github.com/tienanh129902/go-rest-api/config"
	"github.com/tienanh129902/go-rest-api/datatransfers"
	"github.com/tienanh129902/go-rest-api/models"
)

type AuthHandlerFunction interface {
	AuthenticateUser(credentials datatransfers.UserLogin) (token datatransfers.JWTToken, err error)
	RegisterUser(credentials datatransfers.UserSignup) (err error)
	LogoutUser(refreshToken string) (err error)
	ParseToken(tokenString, secret string) (claims datatransfers.JWTClaims, err error)
}

func generateAccessToken(user models.User) string {
	now := time.Now()
	expiredTime, _ := time.ParseDuration(config.AppConfig.AccessTokenExpired)
	expiry := time.Now().Add(expiredTime)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, datatransfers.JWTClaims{
		Sub:       user.ID,
		ExpiresAt: expiry.Unix(),
		IssuedAt:  now.Unix(),
	})
	t, err := token.SignedString([]byte(config.AppConfig.JWTSecret))
	if err != nil {
		panic(err)
	}
	return t
}

func generateRefreshToken(user models.User) string {
	now := time.Now()
	expiredTime, _ := time.ParseDuration(config.AppConfig.RefreshTokenExpired)
	expiry := time.Now().Add(expiredTime)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, datatransfers.JWTClaims{
		Sub:       user.ID,
		ExpiresAt: expiry.Unix(),
		IssuedAt:  now.Unix(),
	})
	t, err := token.SignedString([]byte(config.AppConfig.JWTSecret))
	if err != nil {
		panic(err)
	}
	return t
}

func generateTokenPair(user models.User) (datatransfers.JWTToken, error) {
	var tokenPair datatransfers.JWTToken
	tokenPair.Access_token = generateAccessToken(user)
	tokenPair.Refresh_token = generateRefreshToken(user)
	Handler.CreateToken(tokenPair)
	return tokenPair, nil
}

func (m *module) ParseToken(tokenString, secret string) (claims datatransfers.JWTClaims, err error) {
	if token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	}); err != nil || !token.Valid {
		return datatransfers.JWTClaims{}, fmt.Errorf("invalid token. %s", err)
	}
	return
}

func (m *module) AuthenticateUser(credentials datatransfers.UserLogin) (token datatransfers.JWTToken, err error) {
	var user models.User
	if user, err = m.db.userOrmer.GetOneByUsername(credentials.Username); err != nil {
		return token, errors.New("incorrect credentials")
	}
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password)); err != nil {
		return token, errors.New("incorrect credentials")
	}
	return generateTokenPair(user)
}

func (m *module) RegisterUser(credentials datatransfers.UserSignup) (err error) {
	var hashedPassword []byte
	if hashedPassword, err = bcrypt.GenerateFromPassword([]byte(credentials.Password), bcrypt.DefaultCost); err != nil {
		return errors.New("failed hashing password")
	}
	if _, err = m.db.userOrmer.InsertUser(models.User{
		Username: credentials.Username,
		Email:    credentials.Email,
		Password: string(hashedPassword),
		Bio:      credentials.Bio,
	}); err != nil {
		log.Print(err)
		return fmt.Errorf("error inserting user. %v", err)
	}
	return
}

func (m *module) LogoutUser(refreshToken string) (err error) {
	if _, err := m.db.tokenOrmer.GetTokenByRefreshToken(refreshToken); err != nil {
		return errors.New("invalid refresh token")
	}
	return m.DeleteToken(refreshToken)
}

// func (m *module) RefreshAuth(refreshToken string) (err error) {
// 	var user models.User

// 	if _, err := m.db.tokenOrmer.GetTokenByRefreshToken(refreshToken); err != nil {
// 		return errors.New("invalid refresh token")
// 	}
// 	m.DeleteToken(refreshToken)
// 	return m.CreateToken(user)
// }
