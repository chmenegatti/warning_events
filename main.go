package main

import (
	"fmt"

	"github.com/oracle_with_warning/src/database"
	"github.com/oracle_with_warning/src/models"
	"github.com/oracle_with_warning/src/services"
)

func main() {
	database.ConnectDb()

	res, err := services.GetOracleJobs()

	if err != nil {
		fmt.Println(err)
	}

	var jobsToAdd []models.Job

	for r := range res {
		checkIfExists := services.CheckBeforeInsertJobs(res[r].EventSeriesId)
		if !checkIfExists {
			jobsToAdd = append(jobsToAdd, res[r])
		}
	}

	total, _ := services.InsertJobs(jobsToAdd)
	fmt.Println("Total Inserido: ", total)
}
