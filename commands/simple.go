package commands

import (
	"bufio"
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
	"github.com/yargevad/filepathx"
)

var (
	CommandSimple = &cli.Command{
		Name:    "simple",
		Aliases: []string{"s"},
		Usage:   "simple statistics",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "path",
				Aliases: []string{"p"},
				Usage:   "File path to files to stream, can be a glob. If not set, a pipe is assumed.",
			},
		},
		Action: func(c *cli.Context) error {
			// Verify inputs
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
					err = processReader(&processReaderInput{
						reader: reader,
					})
					if err != nil {
						return err
					}
				}

				return nil
			}

			// run in pipe pass blocks
			info, err := os.Stdin.Stat()
			if err != nil {
				return err
			}

			if info.Mode()&os.ModeCharDevice != 0 || info.Size() <= 0 {
				// if neither, throw error
				return fmt.Errorf("please use this command with a pipe or the --path flag set")
			}

			reader := bufio.NewReader(os.Stdin)
			err = processReader(&processReaderInput{
				reader: reader,
			})
			if err != nil {
				return err
			}

			return nil
		},
	}
)
