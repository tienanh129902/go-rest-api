package datatransfers

type QuestionCreate struct {
	Content        string   `json:"content" binding:"required"`
	Choices        []string `json:"choices" binding:"required"`
	CorrectAnswers []string `json:"correctAnswers" binding:"required"`
}

type QuestionInfo struct {
	ID      uint     `uri:"id" json:"id"`
	Content string   `json:"content"`
	Choices []string `json:"choices"`
}

type QuestionResult struct {
	Content        string   `json:"content"`
	CorrectAnswers []string `json:"correctAnswers"`
}
