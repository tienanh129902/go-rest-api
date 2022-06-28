package datatransfers

import "time"

type AnswerSubmit struct {
	QuestionId uint      `json:"questionId"`
	UserChoice []string  `json:"userChoices"`
	SubmitAt   time.Time `json:"submitAt"`
}
