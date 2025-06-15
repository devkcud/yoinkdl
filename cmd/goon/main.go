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
		WithCommands(
			tool.
				NewCommand("size").
				WithFlags(&cli.BoolFlag{Name: "short", Aliases: []string{"s"}, Usage: "Display only the byte count without additional text or formatting."}).
				WithUsage("Convert size strings (e.g., 1.5MiB, 72KiB, 2PB) into integer bytes. Supports scientific notation (e.g., 532e-3eib).").
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
				}),
		)

	if err := builder.Run(); err != nil {
		log.Fatal(err)
	}
}
