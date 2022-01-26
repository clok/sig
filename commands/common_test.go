package commands

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_processLine(t *testing.T) {
	is := assert.New(t)

	t.Run("prepend hello [single]", func(t *testing.T) {
		rescueStdout := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		input := "Test"
		err := processLine(&processLineInput{
			output: []rune(input),
		})
		is.NoError(err)

		_ = w.Close()
		out, _ := ioutil.ReadAll(r)
		os.Stdout = rescueStdout
		is.Equal(fmt.Sprintf("hello: %s\n", input), string(out))
	})

	t.Run("prepend hello [carriage return]", func(t *testing.T) {
		rescueStdout := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		input := "Test\r\n"
		err := processLine(&processLineInput{
			output: []rune(input),
		})
		is.NoError(err)

		_ = w.Close()
		out, _ := ioutil.ReadAll(r)
		os.Stdout = rescueStdout
		is.Equal(fmt.Sprintf("hello: %s\n", "Test"), string(out))
	})
}

func Test_processReader(t *testing.T) {
	is := assert.New(t)

	sample := `test1
test2
test3`

	t.Run("processes many lines and outputs matches", func(t *testing.T) {
		rescueStdout := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		err := processReader(&processReaderInput{
			reader: bufio.NewReader(strings.NewReader(sample)),
		})
		is.NoError(err)

		_ = w.Close()
		out, _ := ioutil.ReadAll(r)
		os.Stdout = rescueStdout
		is.Equal("hello: test1\nhello: test2\nhello: test3\n", string(out))
	})
}
