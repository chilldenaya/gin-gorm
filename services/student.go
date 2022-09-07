package services

import (
	res "gin-gorm/helpers"
	m "gin-gorm/models"
)

func GetStudentsService() map[string]any {
	return res.SetOk(m.GetStudents())
}

func GetStudentByIdService(id string) map[string]any {
	return res.SetOk(m.GetStudentById(id))
}
