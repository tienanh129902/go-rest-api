package models

import (
	"time"

	"gorm.io/gorm"
)

type Token struct {
	ID            uint `gorm:"primaryKey" json:"-"`
	Access_token  string
	Refresh_token string
	CreatedAt     time.Time `gorm:"autoCreateTime" json:"-"`
}

type tokenOrm struct {
	db *gorm.DB
}

type TokenOrmer interface {
	InsertToken(token Token) (id uint, err error)
	DeleteToken(token Token) (err error)
	GetTokenByRefreshToken(refreshToken string) (token Token, err error)
}

func NewTokenOrmer(db *gorm.DB) TokenOrmer {
	_ = db.AutoMigrate(&Token{}) // builds table when enabled, auto-synchronize with database
	return &tokenOrm{db}
}

func (o *tokenOrm) InsertToken(token Token) (id uint, err error) {
	result := o.db.Model(&Token{}).Create(&token)
	return token.ID, result.Error
}

func (o *tokenOrm) GetTokenByRefreshToken(refreshToken string) (token Token, err error) {
	result := o.db.Model(&Token{}).Where("refresh_token = ?", refreshToken).First(&token)
	return token, result.Error
}

func (o *tokenOrm) DeleteToken(token Token) (err error) {
	result := o.db.Model(&Token{}).Model(&token).Delete(&token)
	return result.Error
}
