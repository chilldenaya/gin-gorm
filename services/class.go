package services

import (
	res "gin-gorm/internals"
	m "gin-gorm/models"
)

func GetClassesService() map[string]any {
	return res.SetOk(m.GetClasses())
}

func GetClassByIdService(id string) map[string]any {
	return res.SetOk(m.GetClassById(id))
}
