package models

import (
	"github.com/google/uuid"
	"github.com/goravel/framework/database/orm"
	"mgufrone.dev/job-tracking/app/ext"
)

type JobApplication struct {
	ext.UUID
	UserID     uuid.UUID `gorm:"index"`
	JobID      uuid.UUID `gorm:"index"`
	User       *User
	Job        *Job
	ResumeName string `gorm:"index"`
	orm.Timestamps
}
