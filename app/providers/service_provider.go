package providers

import (
	"github.com/goravel/framework/contracts/foundation"
	"mgufrone.dev/job-tracking/app/service"
)

type ServiceProvider struct {
}

func (s *ServiceProvider) Register(app foundation.Application) {
	services := map[string]func(app foundation.Application) (any, error){
		service.JobApplicationService: func(app foundation.Application) (any, error) {
			return nil, nil
		},
	}
	for k, v := range services {
		app.Bind(k, v)
	}
}

func (s *ServiceProvider) Boot(app foundation.Application) {
}
