package commands

import (
	"bufio"
	"fmt"
	"math"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func withinTolerance(a, b, e float64) bool {
	if a == b {
		return true
	}

	d := math.Abs(a - b)

	if b == 0 {
		return d < e
	}

	return (d / math.Abs(b)) < e
}

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
		fmt.Printf("%# v\n", result)
		is.NoError(err)
		is.IsType(ResultSet{}, result)
		is.Equal(4, result.n)
		is.Equal(float64(1), result.min)
		is.Equal(float64(4), result.max)
		is.Equal(2.5, result.mean)
		is.Equal([]float64{}, result.mode)
		is.Equal(2.5, result.median)
		is.Equal(float64(10), result.sum)
		is.True(withinTolerance(1.118033988, result.stdev, 1e-9))
		is.Equal(1.25, result.variance)
		is.Equal(float64(2), result.p50)
		is.Equal(float64(3), result.p75)
		is.Equal(3.5, result.p90)
		is.Equal(3.5, result.p95)
		is.Equal(3.5, result.p99)
		is.Equal(1.5, result.q1)
		is.Equal(2.5, result.q2)
		is.Equal(3.5, result.q3)
	})
}
