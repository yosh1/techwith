
package controllers

import (
  "ps-editor/models"
  "strconv"
  // DEBUG:
  "reflect"
  // DEBUG:
  "fmt"
)

type Question struct {
}

func NewQuestion() Question {
  return Question{}
}

func (c Question) GetId(n string) interface{} {
  castedN, _ := strconv.Atoi(n)
  repo := models.NewQuestionRepository()
  question := repo.GetQuestionById(castedN)
  // DEBUG:
  fmt.Println(reflect.TypeOf(question))
  return question
}
