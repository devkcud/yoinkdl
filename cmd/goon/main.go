package main

import (
	"context"
	"fmt"
	"log"

	"github.com/devkcud/goondl/internal/tool"
	"github.com/devkcud/goondl/pkg/size"
	"github.com/urfave/cli/v3"
)

func main() {
	builder := tool.
		New().
		WithGlobalFlags(&cli.BoolFlag{Name: "verbose", Aliases: []string{"v"}}).
		WithCommands(
			tool.
				NewCommand("size").
				WithUsage("Convert size strings (e.g., 1.5MiB, 72KiB, 2PB) into integer bytes. Supports scientific notation too (e.g., 532e-3eib).").
				WithAction(func(ctx context.Context, command *cli.Command) error {
					if !command.Args().Present() {
						return fmt.Errorf("need at least one size to parse")
					}

					i, err := size.ParseSizeFromString(command.Args().First())
					if err != nil {
						return err
					}

					fmt.Println(i)

					return nil
				}),
		)

	if err := builder.Run(); err != nil {
		log.Fatal(err)
	}
}
