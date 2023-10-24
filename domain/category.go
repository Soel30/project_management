package domain

import (
	"gorm.io/gorm"
)

const (
	CollectionCategory = "categories"
)

type Category struct {
	gorm.Model
	Name        string    `json:"name" bson:"name"`
	Color       string    `json:"color" bson:"color"`
	Description string    `json:"description" bson:"description"`
	WorkspaceId uint      `json:"workspace_id" bson:"workspace_id"`
	Workspace   Workspace `gorm:"foreignKey:WorkspaceId"`
}

type PaginationCategory struct {
	Limit     int        `json:"limit"`
	Page      int        `json:"page"`
	Sort      string     `json:"sort"`
	NextPage  int        `json:"next_page"`
	PrevPage  int        `json:"prev_page"`
	TotalData int        `json:"total_data"`
	TotalPage int        `json:"total_page"`
	Data      []Category `json:"data"`
}
