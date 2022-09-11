package controllers

import (
	"gin-gorm/dto"
	m "gin-gorm/models"
)

func GetStudents() (dto.GetStudentesResponse, error) {
	resultObj, err := m.GetStudents()
	if err != nil {
		return dto.GetStudentesResponse{}, err
	}

	result := dto.BuildGetStudentesResponse(resultObj)

	return result, nil
}

func GetStudentById(id int) (dto.GetStudentByIdResponse, error) {
	resultObj, err := m.GetStudentById(id)
	if err != nil {
		return dto.GetStudentByIdResponse{}, err
	}

	result := dto.BuildGetStudentByIdResponse(resultObj)

	return result, nil
}

func CreateStudent(
	req dto.CreateStudentRequest,
) (dto.CreateStudentResponse, error) {
	obj := m.Student{
		Age:  req.Age,
		Name: req.Name,
	}
	resultObj, err := m.CreateStudent(obj)
	if err != nil {
		return dto.CreateStudentResponse{}, err
	}

	result := dto.BuildCreateStudentResponse(resultObj)

	return result, nil
}

func DeleteStudentById(id int) error {
	err := m.DeleteStudentById(id)
	if err != nil {
		return err
	}

	return nil
}
