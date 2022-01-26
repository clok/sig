package main

import (
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/clok/cdocs"
	"github.com/clok/kemba"
	"github.com/clok/sig/commands"
	"github.com/urfave/cli/v2"
)

var (
	version string
	k       = kemba.New("sig")
)

func main() {
	k.Println("executing")

	im, err := cdocs.InstallManpageCommand(&cdocs.InstallManpageCommandInput{
		AppName: "sig",
		Hidden:  true,
	})
	if err != nil {
		log.Fatal(err)
	}

	app := cli.NewApp()
	app.Name = "sig"
	app.Copyright = "(c) 2022 Derek Smith"
	app.Authors = []*cli.Author{
		{
			Name:  "Derek Smith",
			Email: "derek@clokwork.net",
		},
	}
	app.Version = version
	app.Usage = "Statistics in Go - CLI tool for quick statistical analysis of data streams"
	app.Commands = []*cli.Command{
		commands.CommandHello,
		im,
		{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "Print version info",
			Action: func(c *cli.Context) error {
				fmt.Printf("%s %s (%s/%s)\n", "sig", version, runtime.GOOS, runtime.GOARCH)
				return nil
			},
		},
	}

	if os.Getenv("DOCS_MD") != "" {
		docs, err := cdocs.ToMarkdown(app)
		if err != nil {
			panic(err)
		}
		fmt.Println(docs)
		return
	}

	if os.Getenv("DOCS_MAN") != "" {
		docs, err := cdocs.ToMan(app)
		if err != nil {
			panic(err)
		}
		fmt.Println(docs)
		return
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
