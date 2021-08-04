package services

import (
	"net/url"
	"neopaginal-command/configs"
	"neopaginal-command/domain/entity"
	"strings"
)

type BoundService interface {
	BoundDecider(hostName *url.URL) (entity.BoundType, error)
}
type boundService struct {
	appConfig configs.AppConfig
}

func (self *boundService) BoundDecider(url *url.URL) (entity.BoundType, error) {
	var fullPath = (self.appConfig.ApplicationSettings.InitialSchema + "://" + self.appConfig.ApplicationSettings.InitialHostName)

	if url.Hostname()+url.Path == fullPath {
		return entity.InitialWithoutBound, nil
	} else if strings.Contains(fullPath, url.Hostname()) {
		return entity.InBound, nil
	} else {
		return entity.OutBound, nil
	}
}

func NewBoundService(appConfig configs.AppConfig) BoundService {

	return &boundService{
		appConfig: appConfig,
	}
}
