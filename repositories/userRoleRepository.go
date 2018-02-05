package repositories

import (
	"ScaleAffWeb/models"
)

type UserRoleRepository interface {
	GetAllRoleByUserId(userId string) ([]models.Role, error)
	GetAllRoleByUsername(username string) ([]models.Role, error)
}

type userRoleRepository struct {
	db *DataSource
}

func GetUserRoleRepository() UserRoleRepository {
	return &userRoleRepository{db: datasource}
}

func (repo *userRoleRepository) GetAllRoleByUserId(userId string) ([]models.Role, error) {
	var data models.User
	err := repo.db.Model(&models.User{}).Preload("Roles").Where(&models.User{Id: userId}).First(&data)

	return data.Roles, err.Error
}

func (repo *userRoleRepository) GetAllRoleByUsername(username string) ([]models.Role, error) {
	var data models.User
	err := repo.db.Model(&models.User{}).Preload("Roles").Where(&models.User{Username: username}).First(&data)

	return data.Roles, err.Error
}