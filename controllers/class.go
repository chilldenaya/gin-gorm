package controllers

import (
	"gin-gorm/dto"
	m "gin-gorm/models"
)

func GetClasses() (dto.GetClassesResponse, error) {
	resultObj, err := m.GetClasses()
	if err != nil {
		return dto.GetClassesResponse{}, err
	}

	result := dto.BuildGetClassesResponse(resultObj)

	return result, nil
}

func GetClassById(id int) (dto.GetClassByIdResponse, error) {
	resultObj, err := m.GetClassById(id)
	if err != nil {
		return dto.GetClassByIdResponse{}, err
	}

	result := dto.BuildGetClassByIdResponse(resultObj)

	return result, nil
}

func CreateClass(
	req dto.CreateClassRequest,
) (dto.CreateClassResponse, error) {
	classObj := m.Class{
		Code: req.Code,
		Name: req.Name,
	}
	resultObj, err := m.CreateClass(classObj)
	if err != nil {
		return dto.CreateClassResponse{}, err
	}

	result := dto.BuildCreateClassResponse(resultObj)

	return result, nil
}
