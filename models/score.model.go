package models

import (
	"time"

	"gorm.io/gorm"
)

type ScoreBoard struct {
	ID        uint `gorm:"primaryKey" json:"-"`
	UserId    uint
	Score     uint
	CreatedAt time.Time `gorm:"autoCreateTime" json:"-"`
}

type scoreOrm struct {
	db *gorm.DB
}

type ScoreOrmer interface {
	GetScoreBoardByUserID(userId uint) (board []ScoreBoard, err error)
	InsertScoreBoard(board ScoreBoard) (id uint, err error)
}

func NewScoreBoardOrmer(db *gorm.DB) ScoreOrmer {
	_ = db.AutoMigrate(&ScoreBoard{}) // builds table when enabled, auto-synchronize with database
	return &scoreOrm{db}
}

func (o *scoreOrm) GetScoreBoardByUserID(userId uint) (board []ScoreBoard, err error) {
	result := o.db.Model(&ScoreBoard{}).Where("userid = ?", userId).Limit(10).Find(&board)
	return board, result.Error
}

func (o *scoreOrm) InsertScoreBoard(board ScoreBoard) (id uint, err error) {
	result := o.db.Model(&User{}).Create(&board)
	return board.ID, result.Error
}
