package repository

import (
	"pm/domain"

	"gorm.io/gorm"
)

type RoleRepository struct {
	DB *gorm.DB
}

func GetAllRole(role *domain.Role, pagination domain.PaginationRole, db *gorm.DB, status string) (domain.PaginationRole, error) {
	var roles []domain.Role
	var totalData int64
	var NextPage int
	var PrevPage int
	var TotalPage int
	var err error

	offset := (pagination.Page - 1) * pagination.Limit
	if status == "" {
		err = db.Limit(pagination.Limit).Offset(offset).Find(&roles).Error
		db.Model(&domain.Role{}).Count(&totalData)
	} else {
		err = db.Where("status = ?", status).Limit(pagination.Limit).Offset(offset).Find(&roles).Error
		db.Model(&domain.Role{}).Where("status = ?", status).Count(&totalData)
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

	pagination.Data = roles
	pagination.TotalData = int(totalData)
	pagination.NextPage = NextPage
	pagination.PrevPage = PrevPage
	pagination.TotalPage = TotalPage

	return pagination, err
}
