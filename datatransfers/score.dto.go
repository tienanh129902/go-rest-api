package datatransfers

import "time"

type AnswerSubmit struct {
	QuestionId uint     `json:"questionId"`
	Choices    []string `json:"choices"`
}

type AnswerSubmitArray struct {
	AnswerSubmit []AnswerSubmit `json:"answerSubmit"`
	SubmitAt     time.Time      `json:"submitAt"`
}

type ScoreBoardInfoByUserId struct {
	UserId uint      `uri:"userid" json:"userId"`
	Score  uint      `json:"score"`
	PlayAt time.Time `json:"playAt"`
}
