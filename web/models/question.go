
package models

import (
	"ps-editor/common"
)

type Question struct {
	Id       int `gorm:"AUTO_INCREMENT;primary_key"`
	Title string `gorm:"size:225;not null"`
	Body  string `gorm:"size:225;not null"`
	Hint  string `gorm:"size:225;not null"`
	LevelId  int
}

type QuestionRepository struct {
}

func NewQuestionRepository() QuestionRepository {
	return QuestionRepository{}
}

func (m QuestionRepository) GetQuestionById(id int) *Question {
	db := common.GetDB()
	var question Question
	db.Where(Question{Id: id}).Find(&question)
	return &question
}
