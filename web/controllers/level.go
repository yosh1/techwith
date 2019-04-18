
package controllers
 import(
  "ps-editor/models"
  "strconv"
)
 type Level struct {
}
 func NewLevel() Level {
  return Level{}
}
 func (c Level) GetId(n string) interface{} {
  repo := models.NewLevelRepository()
  castedN, _ := strconv.Atoi(n)
  level := repo.GetLevelById(castedN)
  return level
}
