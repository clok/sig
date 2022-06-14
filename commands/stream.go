package commands

import (
	"bufio"
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/yargevad/filepathx"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"os"
)

var (
	CommandStream = &cli.Command{
		Name:  "stream",
		Usage: "stream process ",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "path",
				Aliases: []string{"p"},
				Usage:   "File path to files to stream, can be a glob. If not set, a pipe is assumed.",
			},
			&cli.Int64Flag{
				Name:    "refresh",
				Aliases: []string{"r"},
				Usage:   "how many rows of data between updates",
				Value:   100,
			},
			&cli.Int64Flag{
				Name:    "factor",
				Aliases: []string{"f"},
				Usage:   "rate of growth of refresh value",
				Value:   10,
			},
			&cli.Int64Flag{
				Name:    "cap",
				Aliases: []string{"c"},
				Usage:   "max value of refresh rate for updates",
				Value:   100_000_000,
			},
		},
		Action: func(c *cli.Context) error {
			var sample []float64
			if c.String("path") != "" {
				kf.Printf("globbing files with pattern: %s", c.String("path"))
				files, err := filepathx.Glob(c.String("path"))
				if err != nil {
					return err
				}
				kf.Printf("found %d files", len(files))
				kf.Log(files)

				// filter files
				// For each file, create reader, pass in reader
				for _, fPath := range files {
					kf.Printf("processing file: %s", fPath)
					file, err := os.Open(fPath)
					if err != nil {
						return err
					}

					reader := bufio.NewReader(file)
					inner, err := processReaderStream(&processReaderStreamInput{
						reader:  reader,
						refresh: c.Int64("refresh"),
						factor:  c.Int64("factor"),
						cap:     c.Int64("cap"),
					})
					if err != nil {
						return err
					}
					sample = append(sample, inner...)
				}
			} else {
				// run in pipe pass blocks
				info, err := os.Stdin.Stat()
				if err != nil {
					return err
				}

				noNamedPipe := info.Mode()&os.ModeNamedPipe == 0
				noUnixPipe := info.Mode()&os.ModeCharDevice != 0 || info.Size() <= 0
				k.Printf("noNamedPipe: %t noUnixPipe: %t\n", noNamedPipe, noUnixPipe)

				if noNamedPipe && noUnixPipe {
					// if neither, throw error
					_ = cli.ShowSubcommandHelp(c)
					return fmt.Errorf("please use this command with a pipe or the --path flag set")
				}

				reader := bufio.NewReader(os.Stdin)
				sample, err = processReaderStream(&processReaderStreamInput{
					reader:  reader,
					refresh: c.Int64("refresh"),
					factor:  c.Int64("factor"),
					cap:     c.Int64("cap"),
				})
				if err != nil {
					return err
				}
			}

			res, err := processSample(sample)
			if err != nil {
				return err
			}

			// Clear terminal screen
			err = clearTerminal()
			if err != nil {
				return err
			}

			printTranspose(&res)
			fmt.Println("")

			p := message.NewPrinter(language.English)
			_, err = p.Printf("Done. Processed %d rows\n", res.n)
			if err != nil {
				return err
			}

			return nil
		},
	}
)
