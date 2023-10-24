package domain

import (
	"gorm.io/gorm"
)

const (
	CollectionUser = "users"
)

type User struct {
	gorm.Model
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
	Email    string `json:"email" bson:"email"`
	Name     string `json:"name" bson:"name"`
	Photo    string `json:"photo" bson:"photo"`
	RoleId   int    `json:"role_id" bson:"role_id"`
	Role     Role
}

type UserRepository interface {
	FindAll() ([]User, error)
	FindById(id int) (User, error)
	Create(user User) (User, error)
	Update(user User) (User, error)
	Delete(id int) error
	FindByLimitOffset(limit int, offset int) ([]User, error)
}

type PaginationUser struct {
	Limit     int    `json:"limit"`
	Page      int    `json:"page"`
	Sort      string `json:"sort"`
	NextPage  int    `json:"next_page"`
	PrevPage  int    `json:"prev_page"`
	TotalData int    `json:"total_data"`
	TotalPage int    `json:"total_page"`
	Data      []User `json:"data"`
}
