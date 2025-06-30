package app

import "github.com/devkcud/goondl/internal/cli/command"

func New() *command.Builder {
	return command.
		New().
		WithCommands(
			sizeCommand(),
			mathCommand(),
		)
}
