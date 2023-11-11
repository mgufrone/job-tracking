package models

import (
	"github.com/goravel/framework/database/orm"
	"mgufrone.dev/job-tracking/app/ext"
)

type User struct {
	ext.UUID
	Name       string
	ExternalID string
	orm.Timestamps
}
