package commands

import (
	"github.com/goravel/framework/contracts/console"
	"github.com/goravel/framework/contracts/console/command"
	"github.com/goravel/framework/facades"
	"github.com/pressly/goose/v3"
)

type GooseCreate struct {
}

func (g *GooseCreate) Signature() string {
	return "goose:create"
}

func (g *GooseCreate) Description() string {
	return "Create migration via goose"
}

func (g *GooseCreate) Extend() command.Extend {
	return command.Extend{
		Category: "goose",
	}
}

func (g *GooseCreate) Handle(ctx console.Context) error {
	db, _ := facades.Orm().DB()
	goose.SetBaseFS(nil)
	dialect := goose.DialectSQLite3
	name := ctx.Argument(0)
	migrationType := ctx.Argument(1)
	migrationPath := "database/" + facades.Config().GetString("database.migrations")
	if migrationType == "" {
		migrationType = "go"
	}
	dbDialect := facades.Config().GetString("database.default")
	switch dbDialect {
	case "mysql":
		dialect = goose.DialectMySQL
	case "sqlserver":
		dialect = goose.DialectMSSQL
	case "postgresql":
		dialect = goose.DialectPostgres
	}
	if err := goose.SetDialect(string(dialect)); err != nil {
		return err
	}
	return goose.Create(db, migrationPath, name, migrationType)
}
