package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tienanh129902/go-rest-api/datatransfers"
	"github.com/tienanh129902/go-rest-api/handlers"
	"github.com/tienanh129902/go-rest-api/models"
)

func POST_CreateQuestion(c *gin.Context) {
	var err error
	var question datatransfers.QuestionCreate
	if err = c.ShouldBind(&question); err != nil {
		c.JSON(http.StatusBadRequest, datatransfers.Response{Error: err.Error()})
		return
	}
	if err = handlers.Handler.CreateQuestion(question); err != nil {
		c.JSON(http.StatusUnauthorized, datatransfers.Response{Error: "failed creating question"})
		return
	}
	c.JSON(http.StatusCreated, datatransfers.Response{Data: "Question created"})
}

func GET_AllQuestions(c *gin.Context) {
	var err error
	var questions []models.Question
	if questions, err = handlers.Handler.ListQuestion(); err != nil {
		c.JSON(http.StatusNotFound, datatransfers.Response{Error: "cannot find any question"})
		return
	}
	var result []datatransfers.QuestionInfo
	for i := 0; i < len(questions); i++ {
		result = append(result, datatransfers.QuestionInfo{
			ID:      questions[i].ID,
			Content: questions[i].Content,
			Choices: questions[i].Choices,
		})
	}
	c.JSON(http.StatusOK, datatransfers.Response{Data: result})
}

func GET_QuestionById(c *gin.Context) {
	var err error
	var questionInfo datatransfers.QuestionInfo
	if err = c.ShouldBindUri(&questionInfo); err != nil {
		c.JSON(http.StatusBadRequest, datatransfers.Response{Error: err.Error()})
		return
	}
	var question models.Question
	if question, err = handlers.Handler.RetrieveQuestion(questionInfo.ID); err != nil {
		c.JSON(http.StatusNotFound, datatransfers.Response{Error: "cannot find question"})
		return
	}
	c.JSON(http.StatusOK, datatransfers.Response{Data: datatransfers.QuestionInfo{
		ID:      question.ID,
		Content: question.Content,
		Choices: question.Choices,
	}})
}

func DEL_QuestionById(c *gin.Context) {
	var err error
	var questionInfo datatransfers.QuestionInfo
	if err = c.ShouldBindUri(&questionInfo); err != nil {
		c.JSON(http.StatusBadRequest, datatransfers.Response{Error: err.Error()})
		return
	}
	if err = handlers.Handler.DeleteQuestion(questionInfo.ID); err != nil {
		c.JSON(http.StatusNotFound, datatransfers.Response{Error: "cannot delete question"})
		return
	}
	c.JSON(http.StatusOK, datatransfers.Response{Status: "Question deleted"})
}
