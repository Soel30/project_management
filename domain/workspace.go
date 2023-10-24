package domain

import (
	"gorm.io/gorm"
)

const (
	CollectionWorkspace = "workspaces"
)

type Workspace struct {
	gorm.Model
	Name       string      `json:"name" bson:"name"`
	Color      string      `json:"color" bson:"color"`
	Users      []*User     `gorm:"many2many:user_workspaces;"`
	Categories []*Category `gorm:"foreignKey:WorkspaceId"`
}

type PaginationWorkspace struct {
	Limit     int         `json:"limit"`
	Page      int         `json:"page"`
	Sort      string      `json:"sort"`
	NextPage  int         `json:"next_page"`
	PrevPage  int         `json:"prev_page"`
	TotalData int         `json:"total_data"`
	TotalPage int         `json:"total_page"`
	Data      []Workspace `json:"data"`
}
