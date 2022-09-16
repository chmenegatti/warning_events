package services

import (
	"encoding/json"
	"fmt"

	"github.com/oracle_with_warning/src/database"
	"github.com/oracle_with_warning/src/models"
	"github.com/oracle_with_warning/src/utils"
)

func GetJobsToVerify() string {
	db := database.GetDatabase()

	utils.Logger("Info", "GetJobsToVerify", "Resgatando resgistros com Verificar Detalhe")
	jobs := []models.WarningEvents{}

	db.Where("status", "VERIFICAR_DETALHE").Find(&jobs)

	total := fmt.Sprintf("Resgistros para verificar: %v", len(jobs))

	utils.Logger("Info", "GetJobsToVerify", total)

	jsonJobs, err := json.Marshal(jobs)

	if err != nil {
		utils.Logger("Error", "", "JSON mal formado")
	}

	return string(jsonJobs)
}
