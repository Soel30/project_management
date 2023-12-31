package domain

import (
	"gorm.io/gorm"
)

const (
	CollectionRole = "roles"
)

type Role struct {
	gorm.Model
	Name  string `json:"name" bson:"name"`
	Users []User `gorm:"foreignKey:RoleId"`
}

type PaginationRole struct {
	Limit     int    `json:"limit"`
	Page      int    `json:"page"`
	Sort      string `json:"sort"`
	NextPage  int    `json:"next_page"`
	PrevPage  int    `json:"prev_page"`
	TotalData int    `json:"total_data"`
	TotalPage int    `json:"total_page"`
	Data      []Role `json:"data"`
}
