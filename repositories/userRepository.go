package repositories

import (
	"ScaleAffWeb/models"
	"ScaleAffWeb/requests"
	"ScaleAffWeb/util"
	"strings"
)

type UserRepository interface {
	GetById(id string) (models.User, error)
	GetByUsername(username string) (models.User, error)
	GetAllByRole(roleId string) ([]models.User, error)
	GetAll()([]models.User,error)
	GetAllByRequest(req *requests.UserRequestList)([]models.User,error)
	CountByRequest(req *requests.UserRequestList)(int,error)
	Save(data *models.User) (*models.User, error)
	Update(data *models.User) (*models.User, error)
	DeleteById(id string) (error)
}

type userRepository struct {
	db *DataSource
}

func GetUserRepository() UserRepository {
	return &userRepository{db: datasource}
}

func (repo *userRepository) GetById(id string) (models.User, error) {
	var data models.User
	err := repo.db.Model(&models.User{}).Preload("Roles").Where(&models.User{Id: id}).First(&data)

	return data, err.Error
}

func (repo *userRepository) GetByUsername(username string) (models.User, error) {
	var data models.User
	err := repo.db.Model(&models.User{}).Preload("Roles").Where(&models.User{Username: username}).First(&data)

	return data, err.Error
}

func (repo *userRepository) GetAllByRole(roleId string) ([]models.User, error) {
	var list []models.User
	err := repo.db.Model(&models.User{}).Joins("left join user_role on user_role.user_id = user.id").Where("user_role.role_id = ?", roleId).Find(&list)

	return list, err.Error
}


func (repo *userRepository) GetAll()([]models.User,error) {
	var list []models.User
	err := repo.db.Model(&models.User{}).Find(&list)

	return list, err.Error
}

func (repo *userRepository) GetAllByRequest(req *requests.UserRequestList)([]models.User,error) {
	var list []models.User
	roleId := "%"+req.RoleId+"%"
	search := "%"+strings.ToLower(req.Search)+"%"
	err := repo.db.Model(&models.User{}).Joins("left join user_role on user_role.user_id = user.id").
			Where("user_role.role_id like ? AND ( lower(user.username) like ? OR lower(user.phone) like ? OR lower(user.name) like ? )",roleId,search,search,search).
			Group("user.id").Order(req.OrderBy).Offset(req.Offset).Limit(req.Limit).Find(&list)

	return list, err.Error
}

func (repo *userRepository) CountByRequest(req *requests.UserRequestList)(int,error) {
	var total int
	roleId := "%"+req.RoleId+"%"
	search := "%"+strings.ToLower(req.Search)+"%"
	err := repo.db.Model(&models.User{}).Joins("left join user_role on user_role.user_id = user.id").
		Where("user_role.role_id like ? AND (  lower(user.username) like ? OR lower(user.phone) like ? OR lower(user.name) like ? ) ",roleId,search,search,search).
		Group("user.id").Count(&total)

	return total, err.Error
}

func (repo *userRepository) Save(data *models.User) (*models.User, error){
	var err error
	data.Id = util.GenerateSnowflake()
	data.Password, _ = util.HashPassword(data.Password)
	tx := repo.db.Begin()
	if err = tx.Model(&models.User{}).Create(&data).Error; err != nil {
		tx.Rollback()
		return nil,err
	}

	tx.Commit()
	return data, err
}

func (repo *userRepository) Update(data *models.User) (*models.User, error){
	var err error
	tx := repo.db.Begin()
	roles := data.Roles
	repo.db.Model(&data).Association("Roles").Clear()
	repo.db.Model(&data).Association("Roles").Append(roles)
	if err = tx.Model(&models.User{}).Save(&data).Error; err != nil {
		tx.Rollback()
		return nil,err
	}

	tx.Commit()
	return data, err
}

func (repo *userRepository) DeleteById(id string) (error){
	var err error
	var data models.User
	repo.db.Model(&models.User{}).Where(&models.User{Id: id}).First(&data)
	tx := repo.db.Begin()
	repo.db.Model(&data).Association("Roles").Clear()
	if err = tx.Model(&models.User{}).Delete(&data).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return  err
}