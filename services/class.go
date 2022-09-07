package services

import (
	"gin-gorm/dto"

	res "gin-gorm/helpers"
	m "gin-gorm/models"
)

func GetClassesService() (map[string]any, error) {
	result, err := m.GetClasses()
	if err != nil {
		return res.SetErr(err.Error()), err
	}

	return res.SetOk(result), nil
}

func GetClassByIdService(id int) (map[string]any, error) {
	resultObj, err := m.GetClassById(id)
	if err != nil {
		return res.SetErr(err.Error()), err
	}

	result := dto.BuildGetClassByIdResponse(resultObj)

	return res.SetOk(result), nil
}

func CreateClassService(
	req dto.CreateClassRequest,
) (map[string]any, error) {
	classObj := m.Class{
		Code: req.Code,
		Name: req.Name,
	}
	resultObj, err := m.CreateClass(classObj)
	if err != nil {
		return res.SetErr(err.Error()), err
	}

	result := dto.BuildCreateClassResponse(resultObj)

	return res.SetOk(result), nil
}
