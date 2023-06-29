package domain

import (
	"gorm.io/gorm"
)

const (
	CollectionPost = "posts"
)

type Post struct {
	gorm.Model
	Title    string `json:"title" bson:"title" binding:"required,min=20"`
	Content  string `json:"content" bson:"content" binding:"required,min=200"`
	Category string `json:"category" bson:"category" binding:"required,min=3"`
	Status   string `json:"status" bson:"status" binding:"required,oneof=Publish Draft Thrash"`
}

type PostRepository interface {
	FindAll() ([]Post, error)
	FindById(id int) (Post, error)
	Create(post Post) (Post, error)
	Update(post Post) (Post, error)
	Delete(id int) error
	FindByLimitOffset(limit int, offset int) ([]Post, error)
}
