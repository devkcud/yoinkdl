package command

import (
	"context"

	"github.com/urfave/cli/v3"
)

type CommandWrapper struct {
	Command *cli.Command
}

func NewCommand(name string, aliases ...string) *CommandWrapper {
	return &CommandWrapper{&cli.Command{
		Name:    name,
		Aliases: aliases,
	}}
}

func (c *CommandWrapper) WithUsage(usage string) *CommandWrapper {
	c.Command.Usage = usage
	return c
}

func (c *CommandWrapper) WithFlags(flags ...cli.Flag) *CommandWrapper {
	c.Command.Flags = append(c.Command.Flags, flags...)
	return c
}

func (c *CommandWrapper) WithAction(action func(context.Context, *cli.Command) error) *CommandWrapper {
	c.Command.Action = action
	return c
}

func (c *CommandWrapper) WithSubcommand(subcommand *CommandWrapper) *CommandWrapper {
	c.Command.Commands = append(c.Command.Commands, subcommand.Command)
	return c
}
