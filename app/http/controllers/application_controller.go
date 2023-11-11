package controllers

import (
	errors2 "errors"
	"github.com/goravel/framework/contracts/http"
	facades2 "github.com/goravel/framework/facades"
	"mgufrone.dev/job-tracking/app/facades"
	"mgufrone.dev/job-tracking/app/http/requests/application"
	"mgufrone.dev/job-tracking/app/models"
	"mgufrone.dev/job-tracking/app/service"
	"strings"
)

type ApplicationController struct {
	//Dependent services
}

func NewApplicationController() *ApplicationController {
	return &ApplicationController{
		//Inject services
	}
}
func (r *ApplicationController) svc() service.JobApplication {
	return facades.JobApplicationService()
}

func (r *ApplicationController) Index(ctx http.Context) http.Response {
	return nil
}

func (r *ApplicationController) Show(ctx http.Context) http.Response {
	return nil
}

func (r *ApplicationController) Store(ctx http.Context) http.Response {
	var createReq application.Create
	errors, err := ctx.Request().ValidateRequest(&createReq)
	if err != nil {
		return ctx.Response().Json(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}
	if errors != nil && len(errors.All()) > 0 {
		return ctx.Response().Json(http.StatusBadRequest, map[string]interface{}{
			"errors": errors.All(),
		})
	}
	token := strings.ReplaceAll(ctx.Request().Header("Authentication"), "Bearer ", "")
	authInstance := facades2.Auth()
	_, err = authInstance.Parse(ctx, token)
	if err != nil {
		return ctx.Response().Json(http.StatusUnauthorized, map[string]string{
			"error":   "invalid authentication",
			"message": err.Error(),
			"token":   token,
		})
	}
	var usr models.User
	err = authInstance.User(ctx, &usr)
	if err != nil {
		return ctx.Response().Json(http.StatusUnauthorized, map[string]string{
			"error": "invalid authentication",
		})
	}
	app := &models.JobApplication{
		Job: &models.Job{
			Role:    createReq.Job.Role,
			Company: createReq.Job.Company,
			Link:    createReq.Job.Link,
		},
		ResumeName: createReq.ResumeName,
	}
	err = r.svc().CreateUserApplication(ctx, &usr, app)

	if errors2.Is(err, service.ErrExists) {
		return ctx.Response().Json(http.StatusSeeOther, map[string]string{
			"error": err.Error(),
		})
	}
	if err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return ctx.Response().Json(http.StatusCreated, map[string]string{
		"application": app.ID.String(),
	})
}

func (r *ApplicationController) Update(ctx http.Context) http.Response {
	return nil
}

func (r *ApplicationController) Destroy(ctx http.Context) http.Response {
	return nil
}
