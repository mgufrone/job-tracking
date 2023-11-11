package routes

import (
	"github.com/goravel/framework/facades"

	"mgufrone.dev/job-tracking/app/http/controllers"
)

func Api() {
	userController := controllers.NewUserController()
	facades.Route().Get("/users/{id}", userController.Show)
	facades.Route().Resource("/application", controllers.NewApplicationController())
}
