package commands

import (
	"bufio"
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/yargevad/filepathx"
	"os"
)

func printRows(res *ResultSet) {
	var header string
	for _, field := range res.ListFields() {
		if header == "" {
			header = res.GetHeader(field)
		} else {
			header = fmt.Sprintf("%s\t%s", header, res.GetHeader(field))
		}
	}
	fmt.Println(header)

	var row string
	for _, field := range res.ListFields() {
		if row == "" {
			row = fmt.Sprintf(res.GetFormat(field), res.Get(field))
		} else {
			format := "%s\t" + res.GetFormat(field)
			row = fmt.Sprintf(format, row, res.Get(field))
		}
	}
	fmt.Println(row)
}

func printTranspose(res *ResultSet) {
	for _, field := range res.ListFields() {
		format := "%s\t" + res.GetFormat(field) + "\n"
		fmt.Printf(format, res.GetHeader(field), res.Get(field))
	}
}

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
			&cli.BoolFlag{
				Name:    "transpose",
				Aliases: []string{"t"},
				Usage:   "transpose table output",
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
					inner, err := processReader(&processReaderInput{
						reader: reader,
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
				sample, err = processReader(&processReaderInput{
					reader: reader,
				})
				if err != nil {
					return err
				}
			}

			res, err := processSample(sample)
			if err != nil {
				return err
			}
			// TODO: Add support for refresh rate
			// fmt.Printf("\033[2K\r")

			if c.Bool("transpose") {
				printTranspose(&res)
			} else {
				printRows(&res)
			}

			return nil
		},
	}
)
