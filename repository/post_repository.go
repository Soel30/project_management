package repository

import (
	"sharing_vision/config/db"
	"sharing_vision/domain"
)

type PostRepositoryImpl struct {
	DB         *db.Database
	Collection string
}

func NewPostRepository(db *db.Database) domain.PostRepository {
	return &PostRepositoryImpl{
		DB:         db,
		Collection: domain.CollectionPost,
	}
}

func (r *PostRepositoryImpl) FindAll() ([]domain.Post, error) {
	var posts []domain.Post
	result := r.DB.DB.Find(&posts)
	return posts, result.Error
}

func (r *PostRepositoryImpl) FindById(id int) (domain.Post, error) {
	var post domain.Post
	result := r.DB.DB.First(&post, id)
	return post, result.Error
}

func (r *PostRepositoryImpl) Create(post domain.Post) (domain.Post, error) {
	result := r.DB.DB.Create(&post)
	return post, result.Error
}

func (r *PostRepositoryImpl) Update(post domain.Post) (domain.Post, error) {
	result := r.DB.DB.Save(&post)
	return post, result.Error
}

func (r *PostRepositoryImpl) Delete(id int) error {
	result := r.DB.DB.Delete(&domain.Post{}, id)
	return result.Error
}

func (r *PostRepositoryImpl) FindByLimitOffset(limit int, offset int) ([]domain.Post, error) {
	var posts []domain.Post
	result := r.DB.DB.Limit(limit).Offset(offset).Find(&posts)
	return posts, result.Error
}
