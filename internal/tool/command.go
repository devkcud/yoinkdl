package tool

import (
	"context"

	"github.com/urfave/cli/v3"
)

type commandWrapper struct {
	command *cli.Command
}

func NewCommand(name string, aliases ...string) *commandWrapper {
	return &commandWrapper{&cli.Command{
		Name:    name,
		Aliases: aliases,
	}}
}

func (c *commandWrapper) WithUsage(usage string) *commandWrapper {
	c.command.Usage = usage
	return c
}

func (c *commandWrapper) WithFlags(flags ...cli.Flag) *commandWrapper {
	c.command.Flags = append(c.command.Flags, flags...)
	return c
}

func (c *commandWrapper) WithAction(action func(context.Context, *cli.Command) error) *commandWrapper {
	c.command.Action = action
	return c
}

func (c *commandWrapper) WithSubcommand(subcommand *commandWrapper) *commandWrapper {
	c.command.Commands = append(c.command.Commands, subcommand.command)
	return c
}
