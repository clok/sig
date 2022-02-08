package commands

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_processLine(t *testing.T) {
	is := assert.New(t)

	t.Run("parses number [simple]", func(t *testing.T) {
		input := "1.234"
		f, err := processLine(&processLineInput{
			line: []rune(input),
		})
		is.NoError(err)
		is.Equal(1.234, f)
	})

	t.Run("parses number [scientific]", func(t *testing.T) {
		input := "1.234560e+02"
		f, err := processLine(&processLineInput{
			line: []rune(input),
		})
		is.NoError(err)
		is.Equal(123.456, f)
	})

	t.Run("parses number [carriage return]", func(t *testing.T) {
		input := "1.234\r\n"
		f, err := processLine(&processLineInput{
			line: []rune(input),
		})
		is.NoError(err)
		is.Equal(1.234, f)
	})

	t.Run("parses number [new line]", func(t *testing.T) {
		input := "1.234\n"
		f, err := processLine(&processLineInput{
			line: []rune(input),
		})
		is.NoError(err)
		is.Equal(1.234, f)
	})
}

func Test_processReader(t *testing.T) {
	is := assert.New(t)

	t.Run("processes lines and generate sample", func(t *testing.T) {
		data := `1.234
test bad line
1.234560e+02`
		sample, err := processReader(&processReaderInput{
			reader: bufio.NewReader(strings.NewReader(data)),
		})
		is.NoError(err)
		is.Len(sample, 2)
		is.Equal([]float64{1.234, 123.456}, sample)
	})

	t.Run("processes lines and generate sample [carriage return]", func(t *testing.T) {
		data := "1.234\r\ntest bad line\r\n1.234560e+02\r\n"
		sample, err := processReader(&processReaderInput{
			reader: bufio.NewReader(strings.NewReader(data)),
		})
		is.NoError(err)
		is.Len(sample, 2)
		is.Equal([]float64{1.234, 123.456}, sample)
	})
}

func Test_processSample(t *testing.T) {
	is := assert.New(t)

	t.Run("processes lines and generate sample", func(t *testing.T) {
		data := []float64{1, 2, 3, 4}
		result, err := processSample(data)
		is.NoError(err)
		is.IsType(ResultSet{}, result)
		is.Equal(float64(1), result.min)
		is.Equal(float64(4), result.max)
	})
}
