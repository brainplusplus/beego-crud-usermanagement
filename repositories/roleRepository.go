package repositories

import (
	"ScaleAffWeb/requests"
	"ScaleAffWeb/models"
	"ScaleAffWeb/util"
	"strings"
)

type RoleRepository interface {
	GetById(id string) (models.Role, error)
	GetByName(name string) (models.Role, error)
	GetAll()([]models.Role,error)
	GetAllByRequest(req *requests.RoleRequestList)([]models.Role,error)
	CountByRequest(req *requests.RoleRequestList)(int,error)
	Save(data *models.Role) (*models.Role, error)
	Update(data *models.Role) (*models.Role, error)
	DeleteById(id string) (error)
}

type roleRepository struct {
	db *DataSource
}

func GetRoleRepository() RoleRepository {
	return &roleRepository{db: datasource}
}

func (repo *roleRepository) GetById(id string) (models.Role, error) {
	var data models.Role
	err := repo.db.Model(&models.Role{}).Where(&models.Role{Id: id}).First(&data)

	return data, err.Error
}

func (repo *roleRepository) GetByName(name string) (models.Role, error) {
	var data models.Role
	err := repo.db.Model(&models.Role{}).Where(&models.Role{Name: name}).First(&data)

	return data, err.Error
}


func (repo *roleRepository) GetAll()([]models.Role,error) {
	var list []models.Role
	err := repo.db.Model(&models.Role{}).Find(&list)

	return list, err.Error
}

func (repo *roleRepository) GetAllByRequest(req *requests.RoleRequestList)([]models.Role,error) {
	var list []models.Role
	search := "%"+strings.ToLower(req.Search)+"%"
	err := repo.db.Model(&models.Role{}).
		Where("lower(role.name) like ? OR lower(role.notes) like ?",search,search).
		Order(req.OrderBy).Offset(req.Offset).Limit(req.Limit).Find(&list)

	return list, err.Error
}

func (repo *roleRepository) CountByRequest(req *requests.RoleRequestList)(int,error) {
	var total int
	search := "%"+strings.ToLower(req.Search)+"%"
	err := repo.db.Model(&models.Role{}).
		Where("lower(role.name) like ? OR lower(role.notes) like ?",search,search).
		Count(&total)

	return total, err.Error
}

func (repo *roleRepository) Save(data *models.Role) (*models.Role, error){
	var err error
	data.Id = util.GenerateSnowflake()
	tx := repo.db.Begin()
	if err = tx.Model(&models.Role{}).Create(&data).Error; err != nil {
		tx.Rollback()
		return nil,err
	}

	tx.Commit()
	return data, err
}

func (repo *roleRepository) Update(data *models.Role) (*models.Role, error){
	var err error
	tx := repo.db.Begin()
	if err = tx.Model(&models.Role{}).Save(&data).Error; err != nil {
		tx.Rollback()
		return nil,err
	}

	tx.Commit()
	return data, err
}

func (repo *roleRepository) DeleteById(id string) (error){
	var err error
	var data models.Role
	repo.db.Model(&models.Role{}).Where(&models.Role{Id: id}).First(&data)
	tx := repo.db.Begin()
	if err = tx.Model(&models.Role{}).Delete(&data).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return  err
}
