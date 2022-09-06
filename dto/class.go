package dto

import m "gin-gorm/models"

type CreateClassRequest struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

type CreateClassResponse struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

type GetClassByIdResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

func BuildCreateClassResponse(obj m.Class) CreateClassResponse {
	return CreateClassResponse{
		Code: obj.Code,
		Name: obj.Name,
	}
}

func BuildGetClassByIdResponse(obj m.Class) GetClassByIdResponse {
	return GetClassByIdResponse{
		ID:   int(obj.ID),
		Code: obj.Code,
		Name: obj.Name,
	}
}
