package tool

import (
	"context"
	"os"

	"github.com/urfave/cli/v3"
)

type builder struct {
	rootCmd *cli.Command
}

func New() *builder {
	return &builder{
		rootCmd: &cli.Command{
			Name:    "goon",
			Usage:   "cli download manager",
			Authors: []any{"devkcud"},
		},
	}
}

func (b *builder) WithGlobalFlags(flags ...cli.Flag) *builder {
	b.rootCmd.Flags = append(b.rootCmd.Flags, flags...)
	return b
}

func (b *builder) WithCommands(commands ...*commandWrapper) *builder {
	for _, wrapper := range commands {
		b.rootCmd.Commands = append(b.rootCmd.Commands, wrapper.command)
	}
	return b
}

func (b *builder) Run() error {
	return b.rootCmd.Run(context.Background(), os.Args)
}
