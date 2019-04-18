
package controllers

import(
  "ps-editor/models"
  "strconv"
)

type Answer struct {
}

func NewAnswer() Answer {
  return Answer{}
}

func (c Answer) GetId(n string) interface{} {
  repo := models.NewAnswerRepository()
  castedN, _ := strconv.Atoi(n)
  answer := repo.GetAnswer(castedN)
  return answer
}

// DEBUG:
func (c Answer) GetMessage(n string) string {
  repo := models.NewAnswerRepository()
  castedN, _ := strconv.Atoi(n)
  message := repo.GetMessage(castedN)
  return message
}

func (c Answer) GetBody(n string) string {
  repo := models.NewAnswerRepository()
  castedN, _ := strconv.Atoi(n)
  body := repo.GetBody(castedN)
  return body
}

func (c Answer) GetKeyWord(n string) string {
  repo := models.NewAnswerRepository()
  castedN, _ := strconv.Atoi(n)
  keyword := repo.GetKeyWord(castedN)
  return keyword
}
