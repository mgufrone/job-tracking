package migrations

import (
	"context"
	"database/sql"
	facades2 "github.com/goravel/framework/facades"
	"github.com/pressly/goose/v3"
	"mgufrone.dev/job-tracking/app/facades"
	"mgufrone.dev/job-tracking/app/models"
)

func init() {
	goose.AddMigrationContext(upCreateUserTable, downCreateUserTable)
}

func upCreateUserTable(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	return facades.DBSource().WithContext(ctx).Migrator().
		AutoMigrate(&models.User{}, &models.Job{}, &models.JobApplication{})
}

func downCreateUserTable(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	facades2.Orm()
	return facades.DBSource().WithContext(ctx).Migrator().DropTable(&models.JobApplication{}, &models.User{}, &models.Job{})
}
