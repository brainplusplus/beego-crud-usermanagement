package services

import (
	"ScaleAffWeb/requests"
	"ScaleAffWeb/responses"
	"ScaleAffWeb/repositories"
	"ScaleAffWeb/models"
	"ScaleAffWeb/util"
	"ScaleAffWeb/middlewares"
	"time"
)

type UserService interface {
	Login(username string, password string) (*responses.Response)
	GetById(id string) (*responses.Response)
	GetByUsername(username string) (*responses.Response)
	GetAllByRequest(req *requests.UserRequestList) (*responses.ResponseList)
	DeleteById(id string) (*responses.Response)
	Save(req *requests.UserRequest) (*responses.Response)
	Update(req *requests.UserRequest) (*responses.Response)
}

type userService struct {

}

func GetUserService() UserService {
	return &userService{}
}

// FindProfileById search User profile by user id and returns UserProfile
func (r *userService) Login(username string, password string) (*responses.Response){
	resp := new(responses.Response)
	data, err := repositories.GetUserRepository().GetByUsername(username)
	if err != nil {
		resp.Data = nil
		resp.Message = err.Error()
		resp.Success = false
	} else {
		match := util.CheckPasswordHash(password,data.Password)
		if match {
			et := middlewares.EasyToken{
				Username: username,
				Expires:  time.Now().Unix() + (24*3600), //Segundos
			}
			resp.Data,_ = et.GetJWTToken()
		}else{
			resp.Message = "Invalid username or password"
		}
		resp.Success = match
	}
	return resp
}

func (r *userService) GetByUsername(username string) (*responses.Response) {
	resp := new(responses.Response)
	data, err := repositories.GetUserRepository().GetByUsername(username)
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

func (r *userService) GetById(id string) (*responses.Response){
	resp := new(responses.Response)
	data, err := repositories.GetUserRepository().GetById(id)
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

func (r *userService) GetAllByRequest(req *requests.UserRequestList) (*responses.ResponseList){
	resp := new(responses.ResponseList)
	list, err := repositories.GetUserRepository().GetAllByRequest(req)
	total, _ := repositories.GetUserRepository().CountByRequest(req)
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

func (r *userService) DeleteById(id string) (*responses.Response) {
	resp := new(responses.Response)
	err := repositories.GetUserRepository().DeleteById(id)
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

func (r *userService) Save(req *requests.UserRequest) (*responses.Response) {
	resp := new(responses.Response)
	reqData := models.User{}
	reqData.Id 		 = req.Id
	reqData.Username = req.Username
	reqData.Password = req.Password
	reqData.Name 	 = req.Name
	reqData.Phone	 = req.Phone
	reqData.Email	 = req.Username
	listRole := []models.Role{}
	for i := 0; i < len(req.Roles); i++ {
		roleId := req.Roles[i]
		role, _ := repositories.GetRoleRepository().GetById(roleId)
		listRole = append(listRole,role)
	}
	reqData.Roles = listRole
	data,err := repositories.GetUserRepository().Save(&reqData)
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

func (r *userService) Update(req *requests.UserRequest) (*responses.Response) {
	resp := new(responses.Response)
	reqData,_ := repositories.GetUserRepository().GetById(req.Id)
	reqData.Username = req.Username
	if req.Password != reqData.Password && req.Password != "" {
		reqData.Password,_ = util.HashPassword(req.Password)
	}
	reqData.Name 	 = req.Name
	reqData.Phone	 = req.Phone
	reqData.Email	 = req.Username
	listRole := []models.Role{}
	for i := 0; i < len(req.Roles); i++ {
		roleId := req.Roles[i]
		role, _ := repositories.GetRoleRepository().GetById(roleId)
		listRole = append(listRole,role)
	}
	reqData.Roles = listRole
	data,err := repositories.GetUserRepository().Update(&reqData)
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



