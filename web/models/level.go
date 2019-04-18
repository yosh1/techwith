
package models
 import (
  "ps-editor/common"
)
 type Level struct {
  Id       int
  Level string
}
 type LevelRepository struct {
}
 func NewLevelRepository() LevelRepository {
	return LevelRepository{}
}
 func (m LevelRepository) GetLevelById(level_id int) string {
  db := common.GetDB()
	var level Level
	db.Where(Level{Id: level_id}).Find(&level)
  return level.Level
}
