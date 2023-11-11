package models

import (
	"github.com/goravel/framework/database/orm"
	"github.com/goravel/framework/support/carbon"
	"mgufrone.dev/job-tracking/app/ext"
)

type Job struct {
	ext.UUID
	Role        string
	SubmittedAt carbon.DateTime
	Link        string
	Company     string
	orm.Timestamps
}
