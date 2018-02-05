package services

import (
	"ScaleAffWeb/responses"
	"ScaleAffWeb/repositories"
)

type UserRoleService interface {
	GetAllRoleByUserId(userId string) (*responses.Response)
	GetAllRoleByUsername(username string) (*responses.Response)
}

type userRoleService struct {

}

func GetUserRoleService() UserRoleService {
	return &userRoleService{}
}

// FindProfileById search User profile by user id and returns UserProfile
func (r *userRoleService) GetAllRoleByUserId(userId string) (*responses.Response){
	resp := new(responses.Response)
	data, err := repositories.GetUserRoleRepository().GetAllRoleByUserId(userId)
	if err != nil {
		resp.Data = nil
		resp.Message = err.Error()
		resp.Success = false
	} else {
		resp.Data = data
		resp.Success = true
	}
	return resp
}

func (r *userRoleService) GetAllRoleByUsername(username string) (*responses.Response){
	resp := new(responses.Response)
	data, err := repositories.GetUserRoleRepository().GetAllRoleByUsername(username)
	if err != nil {
		resp.Data = nil
		resp.Message = err.Error()
		resp.Success = false
	} else {
		resp.Data = data
		resp.Success = true
	}
	return resp
}