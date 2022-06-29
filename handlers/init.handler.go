package handlers

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/tienanh129902/go-rest-api/config"
	"github.com/tienanh129902/go-rest-api/models"
)

type HandlerFunction interface {
	AuthHandlerFunction
	UserHandlerFunction
	QuestionHandlerFunction
	TokenHandlerFunction
	ScoreBoardHandlerFunction
}

var Handler HandlerFunction

type module struct {
	db *dbEntity
}

type dbEntity struct {
	conn          *gorm.DB
	userOrmer     models.UserOrmer
	questionOrmer models.QuestionOrmer
	tokenOrmer    models.TokenOrmer
	scoreOrmer    models.ScoreOrmer
}

func InitializeHandler() (err error) {
	// Initialize DB
	var db *gorm.DB
	db, err = gorm.Open(postgres.Open(
		fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=disable",
			config.AppConfig.DBHost, config.AppConfig.DBPort, config.AppConfig.DBDatabase,
			config.AppConfig.DBUsername, config.AppConfig.DBPassword),
	), &gorm.Config{})
	if err != nil {
		log.Println("[INIT] failed connecting to PostgreSQL")
		return
	}
	log.Println("[INIT] connected to PostgreSQL")

	// Compose handler modules
	Handler = &module{
		db: &dbEntity{
			conn:          db,
			userOrmer:     models.NewUserOrmer(db),
			questionOrmer: models.NewQuestionOrmer(db),
			tokenOrmer:    models.NewTokenOrmer(db),
			scoreOrmer:    models.NewScoreBoardOrmer(db),
		},
	}
	return
}
