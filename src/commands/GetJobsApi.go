package commands

import (
	"fmt"

	"github.com/oracle_with_warning/src/models"
	"github.com/oracle_with_warning/src/services"
	"github.com/oracle_with_warning/src/utils"
)

func GetJobsFromRubrik() {
	res, err := services.GetOracleJobs()

	if err != nil {
		utils.Logger("Error", "GetOracleJobs", err.Error())
	}

	var jobsToAdd []models.Job

	for r := range res {
		checkIfExists := services.CheckBeforeInsertJobs(res[r].EventSeriesId)
		if !checkIfExists {
			jobsToAdd = append(jobsToAdd, res[r])
		}
	}

	total, _ := services.InsertJobs(jobsToAdd)
	utils.Logger("Info", "InsertJobs", fmt.Sprintf("Total de registros inseridos: %v", total))
}
