package dto

import m "gin-gorm/models"

type CreateStudentRequest struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type CreateStudentResponse struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type GetStudentesResponseItem struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type GetStudentesResponse struct {
	List []GetStudentesResponseItem `json:"list"`
}

type GetStudentByIdResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func BuildCreateStudentResponse(obj m.Student) CreateStudentResponse {
	return CreateStudentResponse{
		Age:  obj.Age,
		Name: obj.Name,
	}
}

func BuildGetStudentesResponse(obj []m.Student) GetStudentesResponse {
	items := make([]GetStudentesResponseItem, 0)

	for _, class := range obj {
		item := GetStudentesResponseItem{
			Name: class.Name,
			Age:  class.Age,
		}
		items = append(items, item)
	}

	return GetStudentesResponse{
		List: items,
	}
}

func BuildGetStudentByIdResponse(obj m.Student) GetStudentByIdResponse {
	return GetStudentByIdResponse{
		ID:   int(obj.ID),
		Age:  obj.Age,
		Name: obj.Name,
	}
}
