package repository

import (
	"pm/config/db"
	"pm/domain"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	DB         *db.Database
	Collection string
}

func GetAllUser(user *domain.User, pagination domain.PaginationUser, db *gorm.DB, status string) (domain.PaginationUser, error) {
	// get all user with pagination and preload workspace with additional column role in pivot table user_workspace
	var users []domain.User
	var totalData int64
	var NextPage int
	var PrevPage int
	var TotalPage int
	var err error

	offset := (pagination.Page - 1) * pagination.Limit
	if status == "" {
		err = db.Limit(pagination.Limit).Offset(offset).Find(&users).Preload("Workspaces", func(db *gorm.DB) *gorm.DB {
			return db.Select("workspaces.*, user_workspaces.role_id as role_id")
		}).Error
		db.Model(&domain.User{}).Count(&totalData)
	} else {
		err = db.Where("status = ?", status).Limit(pagination.Limit).Offset(offset).Find(&users).Preload("Workspaces", func(db *gorm.DB) *gorm.DB {
			return db.Select("workspaces.*, user_workspaces.role_id as role_id")
		}).Error
		db.Model(&domain.User{}).Where("status = ?", status).Count(&totalData)
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

	pagination.Data = users
	pagination.TotalData = int(totalData)
	pagination.NextPage = NextPage
	pagination.PrevPage = PrevPage
	pagination.TotalPage = TotalPage

	return pagination, err
}
func FindUserByUsername(db *gorm.DB, user domain.User, username string) (result *gorm.DB) {
	result = db.Where("username = ?", username).First(&user)

	return result
}
