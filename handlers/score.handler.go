package handlers

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tienanh129902/go-rest-api/datatransfers"
	"github.com/tienanh129902/go-rest-api/models"
)

type ScoreBoardHandlerFunction interface {
	CreateScoreBoard(c *gin.Context, answerArray datatransfers.AnswerSubmitArray) (score int, err error)
	RetrieveScoreBoardByUserId(userId uint) (scoreBoard []models.ScoreBoard, err error)
}

func isContains(obj string, src []string) bool {
	for _, v := range src {
		if v == obj {
			return true
		}
	}
	return false
}

func (m *module) CreateScoreBoard(c *gin.Context, answerArray datatransfers.AnswerSubmitArray) (score int, err error) {
	var status bool
	answerSubmit := answerArray.AnswerSubmit
	for _, v := range answerSubmit {
		ques, _ := m.db.questionOrmer.GetQuestionByID(v.QuestionId)
		for _, c := range v.Choices {
			if !isContains(c, ques.Choices) {
				status = false
			}
		}
		if status {
			score++
		}
	}
	if _, err = m.db.scoreOrmer.InsertScoreBoard(models.ScoreBoard{
		UserId:    Handler.Me(c),
		Score:     uint(score),
		CreatedAt: time.Now(),
	}); err != nil {
		log.Print(err)
		return 0, fmt.Errorf("error inserting score board. %v", err)
	}
	return score, err
}

func (m *module) RetrieveScoreBoardByUserId(userId uint) (scoreBoard []models.ScoreBoard, err error) {
	if scoreBoard, err = m.db.scoreOrmer.GetScoreBoardByUserID(userId); err != nil {
		return []models.ScoreBoard{}, fmt.Errorf("cannot find score board of user with id %d", userId)
	}
	return
}
