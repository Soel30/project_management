package repository

import (
	"pm/domain"

	"gorm.io/gorm"
)

type WorkspaceRepository struct {
	DB *gorm.DB
}

func GetAllWorkspace(workspace *domain.Workspace, pagination domain.PaginationWorkspace, db *gorm.DB, status string) (domain.PaginationWorkspace, error) {
	var workspaces []domain.Workspace
	var totalData int64
	var NextPage int
	var PrevPage int
	var TotalPage int
	var err error

	offset := (pagination.Page - 1) * pagination.Limit
	if status == "" {
		err = db.Limit(pagination.Limit).Offset(offset).Find(&workspaces).Preload("Users", func(db *gorm.DB) *gorm.DB {
			return db.Select("users.*, user_workspaces.role_id as role_id")
		}, "status = ?", "active").Preload("Categories").Error

		db.Model(&domain.Workspace{}).Count(&totalData)
	} else {
		err = db.Where("status = ?", status).Limit(pagination.Limit).Offset(offset).Find(&workspaces).Preload("Users", func(db *gorm.DB) *gorm.DB {
			return db.Select("users.*, user_workspaces.role_id as role_id")
		}, "status = ?", "active").Preload("Categories").Error

		db.Model(&domain.Workspace{}).Where("status = ?", status).Count(&totalData)
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

	pagination.Data = workspaces
	pagination.TotalData = int(totalData)
	pagination.NextPage = NextPage
	pagination.PrevPage = PrevPage
	pagination.TotalPage = TotalPage

	return pagination, err
}
