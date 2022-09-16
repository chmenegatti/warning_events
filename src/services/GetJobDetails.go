package services

import (
	"encoding/json"
	"fmt"

	"github.com/oracle_with_warning/src/models"
)

type JobsToGetDetail struct {
	Id            int32
	EventSeriesId string
}

type X struct {
	J JobsToGetDetail
}

func GetJobDetails() {

	jobsToVerify := GetJobsToVerify()

	jobsFromDatabase := []models.WarningEvents{}

	json.Unmarshal([]byte(jobsToVerify), &jobsFromDatabase)

	var jobsToGetDetail []JobsToGetDetail

	for r := range jobsFromDatabase {
		jobs := JobsToGetDetail{
			Id:            jobsFromDatabase[r].Id,
			EventSeriesId: jobsFromDatabase[r].EventData.EventSeriesId,
		}

		jobsToGetDetail = append(jobsToGetDetail, jobs)

	}

	fmt.Println(jobsToGetDetail)
}
