package handlers

import (
	"fmt"
	"log"

	"github.com/tienanh129902/go-rest-api/datatransfers"
	"github.com/tienanh129902/go-rest-api/models"
)

type QuestionHandlerFunction interface {
	CreateQuestion(credentials datatransfers.QuestionCreate) (err error)
	RetrieveQuestion(id uint) (question models.Question, err error)
	ListQuestion() (question []models.Question, err error)
	DeleteQuestion(id uint) (err error)
}

func (m *module) CreateQuestion(credentials datatransfers.QuestionCreate) (err error) {

	if _, err = m.db.questionOrmer.InsertQuestion(models.Question{
		Content:        credentials.Content,
		Choices:        credentials.Choices,
		CorrectAnswers: credentials.CorrectAnswers,
	}); err != nil {
		log.Print(err)
		return fmt.Errorf("error inserting question. %v", err)
	}
	return
}

func (m *module) RetrieveQuestion(id uint) (question models.Question, err error) {
	if question, err = m.db.questionOrmer.GetQuestionByID(id); err != nil {
		return models.Question{}, fmt.Errorf("cannot find question with id %d", id)
	}
	return
}

func (m *module) ListQuestion() (question []models.Question, err error) {
	if question, err = m.db.questionOrmer.GetAllQuestion(); err != nil {
		return []models.Question{}, fmt.Errorf("cannot find questions")
	}
	return
}

func (m *module) DeleteQuestion(id uint) (err error) {
	var question models.Question
	if question, err = m.db.questionOrmer.GetQuestionByID(id); err != nil {
		return fmt.Errorf("cannot find question with id %d", id)
	}
	return m.db.questionOrmer.DeleteQuestion(question)
}
