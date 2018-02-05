package services

import (
	"ScaleAffWeb/responses"
	"ScaleAffWeb/requests"
	"ScaleAffWeb/models"
	"ScaleAffWeb/repositories"
)

type RoleService interface {
	GetByRolenameAndPassword(Rolename string, password string) (*responses.Response)
	GetById(id string) (*responses.Response)
	GetAllByRequest(req *requests.RoleRequestList) (*responses.ResponseList)
	DeleteById(id string) (*responses.Response)
	Save(req *requests.RoleRequest) (*responses.Response)
	Update(req *requests.RoleRequest) (*responses.Response)
}

type roleService struct {

}

func GetRoleService() RoleService {
	return &roleService{}
}

// FindProfileById search Role profile by Role id and returns RoleProfile
func (r *roleService) GetByRolenameAndPassword(Rolename string, password string) (*responses.Response){
	resp := new(responses.Response)

	return resp
}

func (r *roleService) GetById(id string) (*responses.Response){
	resp := new(responses.Response)
	data, err := repositories.GetRoleRepository().GetById(id)
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

func (r *roleService) GetAllByRequest(req *requests.RoleRequestList) (*responses.ResponseList){
	resp := new(responses.ResponseList)
	list, err := repositories.GetRoleRepository().GetAllByRequest(req)
	total, _ := repositories.GetRoleRepository().CountByRequest(req)
	if err != nil {
		resp.Rows = []int{}
		resp.RecordsFiltered = 0
		resp.RecordsTotal = 0
		resp.Message = err.Error()
		resp.Success = false
	} else {
		if len(list) == 0{
			resp.Rows = []int{}
		}else {
			resp.Rows = list
		}
		resp.Success = true
		resp.RecordsTotal = total
		resp.RecordsFiltered = len(list)
	}
	return resp
}

func (r *roleService) DeleteById(id string) (*responses.Response) {
	resp := new(responses.Response)
	err := repositories.GetRoleRepository().DeleteById(id)
	if err != nil {
		resp.Data = nil
		resp.Message = err.Error()
		resp.Success = false
	} else {
		resp.Data = nil
		resp.Success = true
		resp.Message = "Success Delete"
	}
	return resp
}

func (r *roleService) Save(req *requests.RoleRequest) (*responses.Response) {
	resp := new(responses.Response)
	reqData := models.Role{}
	reqData.Id 		 = req.Id
	reqData.Name 	 = req.Name
	reqData.Notes	 = req.Notes
	data,err := repositories.GetRoleRepository().Save(&reqData)
	if err != nil {
		resp.Data = nil
		resp.Message = err.Error()
		resp.Success = false
	} else {
		resp.Data = data
		resp.Success = true
		resp.Message = "Success Save"
	}
	return resp
}

func (r *roleService) Update(req *requests.RoleRequest) (*responses.Response) {
	resp := new(responses.Response)
	reqData,_ := repositories.GetRoleRepository().GetById(req.Id)
	reqData.Id 		 = req.Id
	reqData.Name 	 = req.Name
	reqData.Notes	 = req.Notes
	data,err := repositories.GetRoleRepository().Update(&reqData)
	if err != nil {
		resp.Data = nil
		resp.Message = err.Error()
		resp.Success = false
	} else {
		resp.Data = data
		resp.Success = true
		resp.Message = "Success Update"
	}
	return resp
}



