package models

type Job struct {
	JobMonitoringState    string `json:"jobMonitoringState"`
	JobStatus             string `json:"jobStatus"`
	JobType               string `json:"jobType"`
	ObjectId              string `json:"objectId"`
	ObjectType            string `json:"objectType"`
	ObjectName            string `json:"objectName"`
	LocationId            string `json:"locationId"`
	LocationName          string `json:"locationName"`
	SlaDomainId           string `json:"slaDomainId"`
	SlaDomainName         string `json:"slaDomainName"`
	StartTime             string `json:"startTime"`
	EndTime               string `json:"endTime"`
	LastSuccessfulJobTime string `json:"lastSuccessfulJobTime"`
	NextJobTime           string `json:"nextJobTime"`
	IsFirstFullSnapshot   string `json:"isFirstFullSnapshot"`
	SourceClusterName     string `json:"sourceClusterName"`
	RetryCount            string `json:"retryCount"`
	MaximumAttemptsForJob string `json:"maximumAttemptsForJob"`
	DataTransferred       string `json:"dataTransferred"`
	ObjectLogicalSize     string `json:"objectLogicalSize"`
	Throughput            string `json:"throughput"`
	EventSeriesId         string `json:"eventSeriesId"`
	Duration              string `json:"duration"`
	NodeId                string `json:"nodeId"`
	WarningCount          string `json:"warningCount"`
	LastUpdatedTime       string `json:"lastUpdatedTime"`
	IsLogTask             string `json:"isLogTask"`
	IsOnDemand            string `json:"isOnDemand"`
	RetryStatus           string `json:"retryStatus"`
}

type JobMonitoring struct {
	JobMonitoringInfoList []Job `json:"jobMonitoringInfoList"`
}
