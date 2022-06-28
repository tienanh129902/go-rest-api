package models

import (
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Question struct {
	ID             uint           `gorm:"primaryKey" json:"-"`
	Content        string         `json:"-"`
	Choices        pq.StringArray `gorm:"type:text[]"`
	CorrectAnswers pq.StringArray `gorm:"type:text[]"`
	CreatedAt      time.Time      `gorm:"autoCreateTime" json:"-"`
	UpdatedAt      time.Time      `gorm:"autoUpdateTime" json:"-"`
}

type questionOrm struct {
	db *gorm.DB
}

type QuestionOrmer interface {
	GetAllQuestion() (ques []Question, err error)
	GetQuestionByID(id uint) (ques Question, err error)
	InsertQuestion(ques Question) (id uint, err error)
	UpdateQuestion(ques Question) (err error)
	DeleteQuestion(ques Question) (err error)
}

func NewQuestionOrmer(db *gorm.DB) QuestionOrmer {
	// builds table when enabled, auto-synchronize with database
	_ = db.AutoMigrate(&Question{})
	return &questionOrm{db}
}

func (o *questionOrm) GetQuestionByID(id uint) (ques Question, err error) {
	result := o.db.Model(&Question{}).Where("id = ?", id).First(&ques)
	return ques, result.Error
}

func (o *questionOrm) GetAllQuestion() (ques []Question, err error) {
	result := o.db.Limit(10).Find(&ques) // pagination with 10 question
	return ques, result.Error
}

func (o *questionOrm) InsertQuestion(ques Question) (id uint, err error) {
	result := o.db.Model(&Question{}).Create(&ques)
	return ques.ID, result.Error
}

func (o *questionOrm) UpdateQuestion(ques Question) (err error) {
	result := o.db.Model(&Question{}).Model(&ques).Updates(&ques)
	return result.Error
}

func (o *questionOrm) DeleteQuestion(ques Question) (err error) {
	result := o.db.Model(&Question{}).Model(&ques).Delete(&ques)
	return result.Error
}
