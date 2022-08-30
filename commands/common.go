package commands

import (
	"QLProject/actionTask"
	"github.com/urfave/cli/v2"
)

func Commons() []*cli.Command {
	return []*cli.Command{
		{
			Name:   "get",
			Usage:  "Get quanliang message",
			Action: actionTask.GetQLMessage,
		},
	}
}
