package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tienanh129902/go-rest-api/datatransfers"
	"github.com/tienanh129902/go-rest-api/handlers"
	"github.com/tienanh129902/go-rest-api/models"
)

// Get score board by user id
// @Summary Get score board by user ID
// @tags Scoreboard
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param userid path int true "User ID"
// @Success 200 {object} datatransfers.Data "Ok"
// @Failure 400 {object} datatransfers.Error "Bad request"
// @Failure 401 {object} datatransfers.Error "Unauthorized"
// @Router /score/{userid} [get]
func GET_ScoreBoardByUserId(c *gin.Context) {
	var err error
	var scoreBoards []models.ScoreBoard
	param := c.Param("userid")
	userId, _ := strconv.ParseUint(param, 10, 32)
	if scoreBoards, err = handlers.Handler.RetrieveScoreBoardByUserId(uint(userId)); err != nil {
		c.JSON(http.StatusNotFound, datatransfers.Error{Error: "cannot find any"})
		return
	}
	var result []datatransfers.ScoreBoardInfoByUserId
	for i := 0; i < len(scoreBoards); i++ {
		result = append(result, datatransfers.ScoreBoardInfoByUserId{
			Score:  scoreBoards[i].Score,
			PlayAt: scoreBoards[i].CreatedAt,
		})
	}
	c.JSON(http.StatusOK, datatransfers.Data{Data: result})
}

// Submit survey
// @Summary Submit survey
// @tags Play
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param answerArray body datatransfers.AnswerSubmitArray true "List of answer per question"
// @Success 201 {object} datatransfers.Data "Ok"
// @Failure 400 {object} datatransfers.Error "Bad request"
// @Failure 401 {object} datatransfers.Error "Unauthorized"
// @Router /play/submit [post]
func POST_UserSubmit(c *gin.Context) {
	var err error
	var score int
	var answerArray datatransfers.AnswerSubmitArray
	if err = c.ShouldBind(&answerArray); err != nil {
		c.JSON(http.StatusBadRequest, datatransfers.Error{Error: err.Error()})
		return
	}
	if score, err = handlers.Handler.CreateScoreBoard(c, answerArray); err != nil {
		c.JSON(http.StatusUnauthorized, datatransfers.Error{Error: "failed creating score board"})
		return
	}
	res := fmt.Sprintf("Your score: %d", score)
	c.JSON(http.StatusCreated, datatransfers.Data{Data: res})
}
