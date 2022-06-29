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

func GET_ScoreBoardByUserId(c *gin.Context) {
	var err error
	var scoreBoards []models.ScoreBoard
	param := c.Param("userid")
	userId, _ := strconv.ParseUint(param, 10, 32)
	if scoreBoards, err = handlers.Handler.RetrieveScoreBoardByUserId(uint(userId)); err != nil {
		c.JSON(http.StatusNotFound, datatransfers.Response{Error: "cannot find any"})
		return
	}
	var result []datatransfers.ScoreBoardInfoByUserId
	for i := 0; i < len(scoreBoards); i++ {
		result = append(result, datatransfers.ScoreBoardInfoByUserId{
			Score:  scoreBoards[i].Score,
			PlayAt: scoreBoards[i].CreatedAt,
		})
	}
	c.JSON(http.StatusOK, datatransfers.Response{Data: result})
}

func POST_UserSubmit(c *gin.Context) {
	var err error
	var score int
	var answerArray datatransfers.AnswerSubmitArray
	if err = c.ShouldBind(&answerArray); err != nil {
		c.JSON(http.StatusBadRequest, datatransfers.Response{Error: err.Error()})
		return
	}
	if score, err = handlers.Handler.CreateScoreBoard(c, answerArray); err != nil {
		c.JSON(http.StatusUnauthorized, datatransfers.Response{Error: "failed creating score board"})
		return
	}
	res := fmt.Sprintf("Your score: %d", score)
	c.JSON(http.StatusCreated, datatransfers.Response{Status: "Score board created", Data: res})
}
