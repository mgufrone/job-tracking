package facades

import (
	"github.com/goravel/framework/facades"
	"mgufrone.dev/job-tracking/app/service"
)

func JobApplicationService() service.JobApplication {
	app := facades.App()
	svc, err := app.Make(service.JobApplicationService)
	if err != nil {
		return nil
	}
	return svc.(service.JobApplication)
}
