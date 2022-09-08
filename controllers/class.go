package controllers

import (
	"gin-gorm/dto"
	res "gin-gorm/helpers"
	m "gin-gorm/models"
)

func GetClasses() (map[string]any, error) {
	resultObj, err := m.GetClasses()
	if err != nil {
		return res.SetErr(err.Error()), err
	}

	result := dto.BuildGetClassesResponse(resultObj)

	return res.SetOk(result), nil
}

func GetClassById(id int) (map[string]any, error) {
	resultObj, err := m.GetClassById(id)
	if err != nil {
		return res.SetErr(err.Error()), err
	}

	result := dto.BuildGetClassByIdResponse(resultObj)

	return res.SetOk(result), nil
}

func CreateClass(
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
