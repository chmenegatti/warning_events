package services

import (
	"encoding/json"
	"io"
	"log"

	"github.com/oracle_with_warning/src/configs"
	"github.com/oracle_with_warning/src/models"
	"github.com/oracle_with_warning/src/utils"
)

func GetOracleJobs() (result []models.Job, err error) {

	url := configs.GetEnvKeys("RUBRIK_API_BASE_URL")
	url = url + "/v1/job_monitoring?limit=100&job_event_status=SuccessfulWithWarnings&job_type=Backup&object_type=OracleDatabase"

	res, err := utils.RubrikApi(url)

	if err != nil {
		log.Fatal(err.Error())
	}

	responseBody, err := io.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err.Error())
	}

	defer res.Body.Close()

	var jobInfo models.JobMonitoring

	json.Unmarshal(responseBody, &jobInfo)

	for r := range jobInfo.JobMonitoringInfoList {
		result = append(result, jobInfo.JobMonitoringInfoList[r])
	}

	return result, err

}
