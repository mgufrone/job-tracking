package providers

import (
	"fmt"
	"github.com/goravel/framework/contracts/database/seeder"
	"github.com/goravel/framework/contracts/foundation"
	"github.com/goravel/framework/facades"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"mgufrone.dev/job-tracking/database/seeders"
)

type DatabaseServiceProvider struct {
}

func (receiver *DatabaseServiceProvider) Register(app foundation.Application) {

}

func (receiver *DatabaseServiceProvider) Boot(app foundation.Application) {
	facades.App().Bind("db.source", func(app foundation.Application) (any, error) {
		return gorm.Open(
			postgres.New(postgres.Config{
				DSN: fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
					facades.Config().GetString("database.connections.postgresql.username"),
					facades.Config().GetString("database.connections.postgresql.password"),
					facades.Config().GetString("database.connections.postgresql.host"),
					facades.Config().GetInt("database.connections.postgresql.port"),
					facades.Config().GetString("database.connections.postgresql.database"),
				),
			}),
			&gorm.Config{},
		)
	})
	facades.Seeder().Register([]seeder.Seeder{
		&seeders.DatabaseSeeder{},
	})
}
