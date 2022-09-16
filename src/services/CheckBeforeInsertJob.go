package services

import (
	"fmt"

	"github.com/oracle_with_warning/src/database"
	"github.com/oracle_with_warning/src/models"
	"github.com/oracle_with_warning/src/utils"
)

func CheckBeforeInsertJobs(eventId string) bool {
	db := database.GetDatabase()

	var jobs []models.WarningEvents

	err := db.Select("event_data").Where("event_data LIKE ?", "%"+eventId+"%").Find(&jobs).Error

	if err != nil {
		utils.Logger("Info", "CheckBeforeInsertJobs", err.Error())
	}

	if len(jobs) != 0 {
		utils.Logger("Info", "CheckBeforeInsertJobs", fmt.Sprintf("Já existe um evento com EventId: %s", eventId))
		return true
	}

	return false
}
