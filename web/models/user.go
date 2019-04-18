
package models

import (
  "ps-editor/common"
)

type User struct {
  Id     int
  Status int
}

type UserRepository struct {
}

func NewUserRepository() UserRepository {
  return UserRepository{}
}

func (m UserRepository) Create() {
  db := common.GetDB()
  user := User{Id: 1, Status: 0}
  db.Create(&user)
}

/*
func (m UserRepository) UpdateStatus {
  db := common.GetDB()
}
*/

func (m UserRepository) GetUserById(Id int) *User {
  db := common.GetDB()
  var user User
  db.Where(User{Id: Id}).Find(&user)
  return &user
}
