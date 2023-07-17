package repository

import (
	"sharing_vision/config/db"
	"sharing_vision/domain"

	"gorm.io/gorm"
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

func GetAllPost(post *domain.Post, pagination domain.Pagination, db *gorm.DB, status string) (domain.Pagination, error) {
	var posts []domain.Post
	var totalData int64
	var NextPage int
	var PrevPage int
	var TotalPage int
	var err error

	offset := (pagination.Page - 1) * pagination.Limit
	if status == "" {
		err = db.Limit(pagination.Limit).Offset(offset).Find(&posts).Error
		db.Model(&domain.Post{}).Count(&totalData)
	} else {
		err = db.Where("status = ?", status).Limit(pagination.Limit).Offset(offset).Find(&posts).Error
		db.Model(&domain.Post{}).Where("status = ?", status).Count(&totalData)
	}

	if err != nil {
		return pagination, err
	}

	if int(totalData) <= (pagination.Page * pagination.Limit) {
		NextPage = 0
	} else {
		NextPage = pagination.Page + 1
	}

	if pagination.Page == 1 {
		PrevPage = 0
	} else {
		PrevPage = pagination.Page - 1
	}

	if int(totalData)%pagination.Limit == 0 {
		TotalPage = int(totalData) / pagination.Limit
	} else {
		TotalPage = (int(totalData) / pagination.Limit) + 1
	}

	pagination.Data = posts
	pagination.TotalData = int(totalData)
	pagination.NextPage = NextPage
	pagination.PrevPage = PrevPage
	pagination.TotalPage = TotalPage

	return pagination, err
}
