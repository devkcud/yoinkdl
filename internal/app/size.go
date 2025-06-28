package app

import (
	"context"
	"fmt"

	"github.com/devkcud/goondl/internal/command"
	"github.com/devkcud/goondl/internal/modules/size"
	"github.com/urfave/cli/v3"
)

func sizeCommand() *command.CommandWrapper {
	return command.
		NewCommand("size").
		WithFlags(&cli.BoolFlag{Name: "short", Aliases: []string{"s"}, Usage: "Show raw byte count."}).
		WithUsage("Parse human-readable size strings (e.g. 1.5MiB, 72KiB, 2PB, 532e-3EiB) and output an integer byte count.").
		WithAction(func(ctx context.Context, command *cli.Command) error {
			if !command.Args().Present() {
				return fmt.Errorf("need at least one size to parse")
			}

			for _, arg := range command.Args().Slice() {
				sizeObject, err := size.ParseSizeFromString(arg)
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
		}).
		WithSubcommand(
			command.
				NewCommand("normalize", "normal").
				WithUsage("Normalize a size into human-readable string.").
				WithAction(func(ctx context.Context, command *cli.Command) error {
					if !command.Args().Present() {
						return fmt.Errorf("need at least one size to parse")
					}

					for _, arg := range command.Args().Slice() {
						sizeObject, err := size.ParseSizeFromString(arg)
						if err != nil {
							fmt.Printf("%s: %s\n", err, arg)
							continue
						}

						fmt.Printf("%s\n", sizeObject.String())
					}

					return nil
				}),
		)
}
