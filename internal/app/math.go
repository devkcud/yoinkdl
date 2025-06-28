package app

import (
	"context"
	"fmt"

	"github.com/devkcud/goondl/internal/command"
	"github.com/devkcud/goondl/internal/modules/size"
	"github.com/urfave/cli/v3"
)

func mathCommand() *command.CommandWrapper {
	return command.
		NewCommand("math").
		WithFlags(&cli.BoolFlag{Name: "short", Aliases: []string{"s"}, Usage: "Show raw byte count."}).
		WithUsage("Evaluate arithmetic expressions combining size literals (e.g. 1e+3MB-5MiB+10GiB*8B) and output the result in bytes.").
		WithAction(func(ctx context.Context, command *cli.Command) error {
			if !command.Args().Present() {
				return fmt.Errorf("need at least one expression to parse")
			}

			for _, arg := range command.Args().Slice() {
				sizeObject, err := size.Evaluate(arg)
				if err != nil {
					fmt.Printf("%s: %s\n", err, arg)
					continue
				}

				if command.Bool("short") {
					fmt.Println(sizeObject.Int())
				} else {
					fmt.Printf("%s = %d bytes\n", sizeObject.String(), sizeObject.Int())
				}
			}

			return nil
		})
}
