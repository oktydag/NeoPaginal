package main

import (
	"fmt"
	"neopaginal-command/application_services"
	"neopaginal-command/configs"
	"neopaginal-command/domain"
)

func main() {

	var appConfig = configs.InitializeConfigs()

	var crawlRepository = domain.InitializeRepository(appConfig)

	var applicationCompleteResult = application_services.Crawl(appConfig, crawlRepository)

	fmt.Println(applicationCompleteResult)
}
