package services

import (
	"fmt"
	"log"

	"github.com/oracle_with_warning/src/database"
	"github.com/oracle_with_warning/src/models"
)

func CheckBeforeInsertJobs(eventId string) bool {
	db := database.GetDatabase()

	var jobs []models.WarningEvents

	err := db.Select("event_data").Where("event_data LIKE ?", "%"+eventId+"%").Find(&jobs).Error

	if err != nil {
		log.Fatal(err)
	}

	if len(jobs) != 0 {
		fmt.Println("Já existe esse job", eventId)
		return true
	}

	return false
}
