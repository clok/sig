package commands

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/clok/kemba"
)

var (
	k     = kemba.New("sig:commands")
	kf    = k.Extend("filter")
	kc    = k.Extend("common")
	kfp   = kc.Extend("processReader")
	kfpl  = kfp.Extend("lines")
	kfpd  = kfp.Extend("debug")
	kfpld = kc.Extend("processLine:debug")
)

type processReaderInput struct {
	reader *bufio.Reader
}

func processReader(opts *processReaderInput) error {
	var output []rune
	var lines int64

	for {
		input, _, err := opts.reader.ReadRune()
		if err != nil && err == io.EOF {
			break
		}
		kfpd.Printf("%c", input)
		output = append(output, input)
		if input == '\n' {
			err := processLine(&processLineInput{
				output: output,
			})
			if err != nil {
				return err
			}
			lines++
			output = []rune{}
			kfpd.Println("-- RESET OUTPUT --")
		}
	}

	if len(output) > 0 {
		err := processLine(&processLineInput{
			output: output,
		})
		if err != nil {
			return err
		}
		lines++
	}
	kfpl.Printf("%d lines processed", lines)
	return nil
}

type processLineInput struct {
	output []rune
}

func processLine(opts *processLineInput) error {
	line := string(opts.output)
	kfpld.Printf("stringify %s", line)
	fmt.Printf("hello: %s\n", strings.ReplaceAll(strings.ReplaceAll(line, "\r\n", ""), "\n", ""))
	return nil
}
