package domain

import (
	"gorm.io/gorm"
)

const (
	CollectionPost = "posts"
)

type Post struct {
	gorm.Model
	Title    string `json:"title" bson:"title"`
	Content  string `json:"content" bson:"content"`
	Category string `json:"category" bson:"category"`
	Status   string `json:"status" bson:"status" binding:"oneof=Publish Draft Thrash"`
}

type PostRepository interface {
	FindAll() ([]Post, error)
	FindById(id int) (Post, error)
	Create(post Post) (Post, error)
	Update(post Post) (Post, error)
	Delete(id int) error
	FindByLimitOffset(limit int, offset int) ([]Post, error)
}

type Pagination struct {
	Limit     int    `json:"limit"`
	Page      int    `json:"page"`
	Sort      string `json:"sort"`
	NextPage  int    `json:"next_page"`
	PrevPage  int    `json:"prev_page"`
	TotalData int    `json:"total_data"`
	TotalPage int    `json:"total_page"`
	Data      []Post `json:"data"`
}
