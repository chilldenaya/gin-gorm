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

type GetClassesResponseItem struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

type GetClassesResponse struct {
	List []GetClassesResponseItem `json:"list"`
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

func BuildGetClassesResponse(obj []m.Class) GetClassesResponse {
	items := make([]GetClassesResponseItem, 0)

	for _, class := range obj {
		item := GetClassesResponseItem{
			Name: class.Name,
			Code: class.Code,
		}
		items = append(items, item)
	}

	return GetClassesResponse{
		List: items,
	}
}

func BuildGetClassByIdResponse(obj m.Class) GetClassByIdResponse {
	return GetClassByIdResponse{
		ID:   int(obj.ID),
		Code: obj.Code,
		Name: obj.Name,
	}
}
