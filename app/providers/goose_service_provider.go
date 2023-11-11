package providers

import (
	"github.com/goravel/framework/contracts/console"
	"github.com/goravel/framework/contracts/foundation"
	"github.com/goravel/framework/facades"
)

type GooseServiceProvider struct {
}

func (receiver *GooseServiceProvider) Register(app foundation.Application) {
}

func (receiver *GooseServiceProvider) Boot(app foundation.Application) {
	var commands []console.Command
	facades.Artisan().Register(commands)
}
