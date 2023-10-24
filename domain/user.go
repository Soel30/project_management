package domain

import (
	"gorm.io/gorm"
)

const (
	CollectionUser = "users"
)

type User struct {
	gorm.Model
	Username   string `json:"username" bson:"username"`
	Password   string `json:"password" bson:"password"`
	Email      string `json:"email" bson:"email"`
	Name       string `json:"name" bson:"name"`
	Photo      string `json:"photo" bson:"photo"`
	RoleId     int    `json:"role_id" bson:"role_id"`
	Role       Role
	Workspaces []*Workspace `gorm:"many2many:user_workspaces;"`
}

type UserWorkspace struct {
	gorm.Model
	UserId      uint      `json:"user_id" bson:"user_id"`
	User        User      `gorm:"foreignKey:UserId"`
	WorkspaceId int       `json:"workspace_id" bson:"workspace_id"`
	Workspace   Workspace `gorm:"foreignKey:WorkspaceId"`
	RoleId      int       `json:"role_id" bson:"role_id"`
	Role        Role      `gorm:"foreignKey:RoleId"`
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
