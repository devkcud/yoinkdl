package app

import "github.com/devkcud/yoinkdl/internal/cli/command"

func New() *command.Builder {
	return command.
		New().
		WithCommands(
			sizeCommand(),
			mathCommand(),
		)
}
