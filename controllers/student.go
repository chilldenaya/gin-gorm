package controllers

import (
	res "gin-gorm/helpers"
	m "gin-gorm/models"
)

func GetStudentsController() map[string]any {
	return res.SetOk(m.GetStudents())
}

func GetStudentByIdController(id string) map[string]any {
	return res.SetOk(m.GetStudentById(id))
}
