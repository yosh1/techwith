
package models

 import (
  "ps-editor/common"
)
 type Unit struct {
  Id       int
  Title string
}
 type UnitRepository struct {
}
 func NewUnitRepository() UnitRepository {
	return UnitRepository{}
}
 func (m UnitRepository) GetTitleById(unit_id int) string {
  db := common.GetDB()
	var unit Unit
	db.Where(Unit{Id: unit_id}).Find(&unit)
  return unit.Title
}
