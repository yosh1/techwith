
package models

import (
  "ps-editor/common"
  // DEBUG:
  "fmt"
)

type Answer struct {
  QuestionId int
	Body    string `gorm:"size:225;not null"`
  KeyWord string
  Message string
}

type AnswerRepository struct {
}

func NewAnswerRepository() AnswerRepository {
	return AnswerRepository{}
}

func (m AnswerRepository) GetAnswer (questionId int) *Answer {
  db := common.GetDB()
  var answer Answer
  db.Where(Answer{QuestionId: questionId}).Find(&answer)
  fmt.Println("***")
  fmt.Println(answer.Message)
  return &answer
}

func (m AnswerRepository) GetMessage(questionId int) string {
  db := common.GetDB()
  var answer Answer
  db.Where(Answer{QuestionId: questionId}).Find(&answer)
  return answer.Message
}

func (m AnswerRepository) GetBody (questionId int) string {
  db := common.GetDB()
  var answer Answer
  db.Where(Answer{QuestionId: questionId}).Find(&answer)
  return answer.Body
}

func (m AnswerRepository) GetKeyWord (questionId int) string {
  db := common.GetDB()
  var answer Answer
  db.Where(Answer{QuestionId: questionId}).Find(&answer)
  return answer.KeyWord
}

// func (m AnswerRepository) GetAnswer (questionId int) Answer {
//   db := common.GetDB()
//   var answer Answer
//   db.Where(Answer{QuestionId: questionId}).Find(&answer)
//   return answer
// }
