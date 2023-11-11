package service

import (
	"context"
	"errors"
	"mgufrone.dev/job-tracking/app/models"
)

const JobApplicationService = "service.JobApplicationService"

var (
	ErrExists = errors.New("application exists")
)

type JobApplicationFilter struct {
	ResumeName string
	User       *models.User
	Company    string
	Role       string
	Status     int
}

//go:generate mockery --name=JobApplication
type JobApplication interface {
	CreateUserApplication(ctx context.Context, user *models.User, application *models.JobApplication) error
	UpdateUserApplication(ctx context.Context, user *models.User, application *models.JobApplication) error
	DeleteUserApplication(ctx context.Context, user *models.User, application *models.JobApplication) error

	Applications(ctx context.Context, filter *JobApplicationFilter) ([]*models.JobApplication, int64, error)
}
