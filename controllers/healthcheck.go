package controllers

import m "gin-gorm/models"

func GetDBHeartbeatController() error {
	return m.GetDBHeartbeat()
}
