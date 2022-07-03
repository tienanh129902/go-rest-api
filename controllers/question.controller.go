package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tienanh129902/go-rest-api/datatransfers"
	"github.com/tienanh129902/go-rest-api/handlers"
	"github.com/tienanh129902/go-rest-api/models"
)

// Create question
// @Summary Create question
// @tags Question
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param question body datatransfers.QuestionCreate true "Question data"
// @Success 200 {object} datatransfers.Status "Ok"
// @Failure 400 {object} datatransfers.Error "Bad request"
// @Failure 401 {object} datatransfers.Error "Unauthorized"
// @Router /question [post]
func POST_CreateQuestion(c *gin.Context) {
	var err error
	var question datatransfers.QuestionCreate
	if err = c.ShouldBind(&question); err != nil {
		c.JSON(http.StatusBadRequest, datatransfers.Error{Error: err.Error()})
		return
	}
	if err = handlers.Handler.CreateQuestion(question); err != nil {
		c.JSON(http.StatusUnauthorized, datatransfers.Error{Error: "failed creating question"})
		return
	}
	c.JSON(http.StatusCreated, datatransfers.Status{Status: "Question created"})
}

// Create question
// @Summary List all questions
// @Security ApiKeyAuth
// @tags Play
// @Accept  json
// @Produce  json
// @Success 200 {object} datatransfers.Data "Ok"
// @Failure 400 {object} datatransfers.Error "Bad request"
// @Failure 401 {object} datatransfers.Error "Unauthorized"
// @Router /play [get]
func GET_AllQuestions(c *gin.Context) {
	var err error
	var questions []models.Question
	if questions, err = handlers.Handler.ListQuestion(); err != nil {
		c.JSON(http.StatusNotFound, datatransfers.Error{Error: "cannot find any question"})
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
	c.JSON(http.StatusOK, datatransfers.Data{Data: result})
}

// Get question by ID
// @Summary Get question by ID
// @Security ApiKeyAuth
// @tags Question
// @Accept  json
// @Produce  json
// @Param id path int true "Question ID"
// @Success 200 {object} datatransfers.Data "Ok"
// @Failure 400 {object} datatransfers.Error "Bad request"
// @Failure 404 {object} datatransfers.Error "Not found"
// @Router /question/{id} [get]
func GET_QuestionById(c *gin.Context) {
	var err error
	var questionInfo datatransfers.QuestionInfo
	if err = c.ShouldBindUri(&questionInfo); err != nil {
		c.JSON(http.StatusBadRequest, datatransfers.Error{Error: err.Error()})
		return
	}
	var question models.Question
	if question, err = handlers.Handler.RetrieveQuestion(questionInfo.ID); err != nil {
		c.JSON(http.StatusNotFound, datatransfers.Error{Error: "cannot find question"})
		return
	}
	c.JSON(http.StatusOK, datatransfers.Data{Data: datatransfers.QuestionInfo{
		ID:      question.ID,
		Content: question.Content,
		Choices: question.Choices,
	}})
}

// Delete question
// @Summary Delete question
// @tags Question
// @Accept  json
// @Produce  json
// @Param id path int true "Question ID"
// @Success 200 {object} datatransfers.Status "Ok"
// @Failure 400 {object} datatransfers.Error "Bad request"
// @Failure 404 {object} datatransfers.Error "Not found"
// @Router /question/{id} [delete]
func DEL_QuestionById(c *gin.Context) {
	var err error
	var questionInfo datatransfers.QuestionInfo
	if err = c.ShouldBindUri(&questionInfo); err != nil {
		c.JSON(http.StatusBadRequest, datatransfers.Error{Error: err.Error()})
		return
	}
	if err = handlers.Handler.DeleteQuestion(questionInfo.ID); err != nil {
		c.JSON(http.StatusNotFound, datatransfers.Error{Error: "cannot delete question"})
		return
	}
	c.JSON(http.StatusOK, datatransfers.Status{Status: "Question deleted"})
}
