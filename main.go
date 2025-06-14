package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Commands: []*cli.Command{
			{
				Name:    "download",
				Aliases: []string{"d", "down"},
				Usage:   "Manually start downloading files from URL",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					args := cmd.Args()
					if !args.Present() {
						return fmt.Errorf("need at least one URL")
					}

					for _, arg := range args.Slice() {
						if err := downloadFile(arg); err != nil {
							return err
						}
					}

					return nil
				},
			},
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

func downloadFile(argument string) error {
	fileURL, err := url.ParseRequestURI(argument)
	if err != nil {
		return err
	}

	path := fileURL.Path
	segments := strings.Split(path, "/")
	filename := segments[len(segments)-1]

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}

	resp, err := client.Get(argument)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 { // Considering the server doesn't send 400 - OK lol
		return fmt.Errorf("server returned non-OK status: %s", resp.Status)
	}

	size, err := io.Copy(file, resp.Body)

	fmt.Printf("Downloaded a file %s with size %d", filename, size)
	return nil
}
