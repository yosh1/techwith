
package controllers

import(
  "ps-editor/models"
)

type User struct {
}

func NewUser() User {
  return User{}
}

func (c User) CreateUser() {
  repo := models.NewUserRepository()
  repo.Create()
}

func (c User) GetUser(n int) interface{} {
  repo := models.NewUserRepository()
  user := repo.GetUserById(1)
  // user := repo.GetUserById(n)
  return user
}
