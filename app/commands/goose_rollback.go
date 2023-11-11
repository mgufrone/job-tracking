package commands

import (
	"github.com/goravel/framework/contracts/console"
	"github.com/goravel/framework/contracts/console/command"
	"github.com/goravel/framework/facades"
	"github.com/pressly/goose/v3"
)

type GooseDown struct {
}

func (g *GooseDown) Signature() string {
	return "goose:down"
}

func (g *GooseDown) Description() string {
	return "Run goose migration up"
}

func (g *GooseDown) Extend() command.Extend {
	return command.Extend{
		Category: "goose",
	}
}

func (g *GooseDown) Handle(ctx console.Context) error {
	db, _ := facades.Orm().DB()
	goose.SetBaseFS(nil)
	dialect := goose.DialectSQLite3
	migrationPath := "database/" + facades.Config().GetString("database.migrations")
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
	return goose.Down(db, migrationPath)
}
