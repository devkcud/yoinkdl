package command

import (
	"context"
	"os"

	"github.com/urfave/cli/v3"
)

type Builder struct {
	RootCommand *cli.Command
}

func New() *Builder {
	return &Builder{
		RootCommand: &cli.Command{
			Name:    "goon",
			Usage:   "cli download manager",
			Authors: []any{"devkcud"},
		},
	}
}

func (b *Builder) WithGlobalFlags(flags ...cli.Flag) *Builder {
	b.RootCommand.Flags = append(b.RootCommand.Flags, flags...)
	return b
}

func (b *Builder) WithCommands(commands ...*CommandWrapper) *Builder {
	for _, wrapper := range commands {
		b.RootCommand.Commands = append(b.RootCommand.Commands, wrapper.Command)
	}
	return b
}

func (b *Builder) Run() error {
	return b.RootCommand.Run(context.Background(), os.Args)
}
