package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"unicode"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Commands: []*cli.Command{
			{
				Name:    "download",
				Aliases: []string{"d", "down"},
				Usage:   "Manually start downloading files from URL",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "from",
						Aliases: []string{"f"},
						Usage:   "Starting byte point",
					},
					&cli.StringFlag{
						Name:    "to",
						Aliases: []string{"t"},
						Usage:   "Ending byte point",
					},
				},
				Action: func(ctx context.Context, cmd *cli.Command) error {
					args := cmd.Args()
					if !args.Present() {
						return fmt.Errorf("need at least one URL")
					}

					var from *int64
					var to *int64

					if cmd.String("from") != "" {
						value, err := parseSize(cmd.String("from"))
						if err != nil {
							return err
						}
						from = &value
					}

					if cmd.String("to") != "" {
						value, err := parseSize(cmd.String("to"))
						if err != nil {
							return err
						}
						to = &value
					}

					for _, arg := range args.Slice() {
						// TODO: Add download threads
						if err := downloadFile(arg, from, to); err != nil {
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

func downloadFile(argument string, from, to *int64) error {
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

	// TODO: Send HEAD request first to gather information like: Accept-Ranges, Content-Length, Content-Type, Strict-Transport-Security, etc
	// INFO: With the Content-Length, we split the file into chunks then join them later on
	req, err := http.NewRequest("GET", argument, nil)
	if err != nil {
		return nil
	}

	fileRange := "bytes="

	if from != nil {
		fileRange += fmt.Sprintf("%d", *from)
	}
	fileRange += fmt.Sprint("-")
	if to != nil {
		fileRange += fmt.Sprintf("%d", *to)
	}

	req.Header.Set("Range", fileRange)

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 { // Considering the server doesn't send 400 - OK lol
		return fmt.Errorf("server returned non-OK status: %s", resp.Status)
	}

	size, err := io.Copy(file, resp.Body)

	fmt.Printf("Downloaded a file %s with size %d\n", filename, size)
	return nil
}

func parseSize(s string) (int64, error) {
	str := strings.TrimSpace(s)
	if str == "" {
		return 0, fmt.Errorf("empty size string")
	}
	str = strings.ToLower(str)

	i := 0
	for i < len(str) && (unicode.IsDigit(rune(str[i])) || str[i] == '.') {
		i++
	}
	numStr := str[:i]
	unitStr := strings.TrimSpace(str[i:])

	value, err := strconv.ParseFloat(numStr, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid number: %q", numStr)
	}

	var factor float64
	switch unitStr {
	case "", "b":
		factor = 1
	case "kb":
		factor = 1e3
	case "mb":
		factor = 1e6
	case "gb":
		factor = 1e9
	case "pb":
		factor = 1e15
	case "eb":
		factor = 1e18
	case "kib":
		factor = 1 << 10
	case "mib":
		factor = 1 << 20
	case "gib":
		factor = 1 << 30
	case "pib":
		factor = 1 << 50
	case "eib":
		factor = 1 << 60
	default:
		return 0, fmt.Errorf("unknown unit: %q", unitStr)
	}

	result := int64(value * factor)
	return result, nil
}
