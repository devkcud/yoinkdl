package app

import "github.com/devkcud/goondl/internal/command"

func New() *command.Builder {
	return command.
		New().
		WithCommands(
			sizeCommand(),
			mathCommand(),
		)
}
