package main

import (
	"github.com/oracle_with_warning/src/commands"
	"github.com/oracle_with_warning/src/database"
	"github.com/oracle_with_warning/src/services"
	"github.com/oracle_with_warning/src/utils"
)

func main() {
	utils.Logger("Info", "Main", "Iniciando o cron de monitoramento")
	database.ConnectDb()

	commands.GetJobsFromRubrik()
	services.GetJobDetails()

}
