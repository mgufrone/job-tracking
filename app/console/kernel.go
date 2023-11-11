package console

import (
	"github.com/goravel/framework/contracts/console"
	"github.com/goravel/framework/contracts/schedule"
	"mgufrone.dev/job-tracking/app/commands"
)

type Kernel struct {
}

func (kernel *Kernel) Schedule() []schedule.Event {
	return []schedule.Event{}
}

func (kernel *Kernel) Commands() []console.Command {
	return []console.Command{
		&commands.GooseCreate{},
		&commands.GooseUp{},
		&commands.GooseDown{},
	}
}
