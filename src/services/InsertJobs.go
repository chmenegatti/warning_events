package services

import (
	"time"

	"github.com/oracle_with_warning/src/database"
	"github.com/oracle_with_warning/src/models"
)

func InsertJobs(jobs []models.Job) (int, error) {

	db := database.GetDatabase()

	var event models.WarningEvents

	for j := range jobs {
		event.EventData = &jobs[j]
		event.Status = "VERIFICAR_DETALHE"
		event.CreatedAt = time.Now()
		event.UpdatedAt = time.Now()
		db.Create(&event)
	}

	return len(jobs), nil
}
