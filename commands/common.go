package commands

import (
	"fmt"
	"github.com/clok/kemba"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"io"
	"math"
	"os"
	"os/exec"
	"runtime"
	"sort"
	"strconv"
	"strings"
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

func clearTerminal() error {
	var cmd *exec.Cmd
	goos := runtime.GOOS
	switch goos {
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	case "linux", "darwin":
		cmd = exec.Command("clear")
	}

	if cmd == nil {
		return fmt.Errorf("%s is not supported", runtime.GOOS)
	}

	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func processReader(opts *processReaderInput) ([]float64, error) {
	var data []rune
	var lines int64
	var fails int64
	var sample []float64

	for {
		input, _, err := opts.reader.ReadRune()
		if err != nil && err == io.EOF {
			break
		}
		kfpd.Printf("%c", input)
		data = append(data, input)
		if input == '\n' {
			value, err := processLine(&processLineInput{
				line: data,
			})
			if err != nil && value == 0 {
				fails++
			} else {
				sample = append(sample, value)
			}
			lines++

			// reset
			data = []rune{}
			kfpd.Println("-- RESET OUTPUT --")
		}
	}

	if len(data) > 0 {
		value, err := processLine(&processLineInput{
			line: data,
		})
		if err != nil && value == 0 {
			fails++
		}
		lines++
		sample = append(sample, value)
	}
	kfpl.Printf("%d / %d lines processed failed to parse", fails, lines)
	return sample, nil
}

func processReaderStream(opts *processReaderStreamInput) ([]float64, error) {
	var data []rune
	var lines int64
	var fails int64
	var sample []float64

	p := message.NewPrinter(language.English)

	// Clear terminal screen
	err := clearTerminal()
	if err != nil {
		return nil, err
	}

	var steps int64
	steps = opts.refresh
	iteration := 0
	for {
		input, _, err := opts.reader.ReadRune()
		if err != nil && err == io.EOF {
			break
		}
		kfpd.Printf("%c", input)
		data = append(data, input)
		if input == '\n' {
			value, err := processLine(&processLineInput{
				line: data,
			})
			if err != nil && value == 0 {
				fails++
			} else {
				sample = append(sample, value)
			}
			lines++

			if lines%steps == 0 {
				res, err := processSample(sample)
				if err != nil {
					return nil, err
				}
				fmt.Printf("\033[0;0H")
				for _, field := range res.ListFields() {
					format := "%s\t" + res.GetFormat(field) + "\n"
					fmt.Printf(format, res.GetHeader(field), res.Get(field))
				}
				steps *= opts.factor
				if steps > opts.cap {
					steps = opts.cap
				}
				iteration++
				_, err = p.Printf("\n[%d] next refresh at N modulo %d == 0", iteration, steps)
				if err != nil {
					return nil, err
				}
			}

			// reset
			data = []rune{}
			kfpd.Println("-- RESET OUTPUT --")
		}
	}

	if len(data) > 0 {
		value, err := processLine(&processLineInput{
			line: data,
		})
		if err != nil && value == 0 {
			fails++
		}
		lines++
		sample = append(sample, value)
	}
	kfpl.Printf("%d / %d lines processed failed to parse", fails, lines)
	return sample, nil
}

func processLine(opts *processLineInput) (float64, error) {
	line := strings.ReplaceAll(strings.ReplaceAll(string(opts.line), "\r", ""), "\n", "")
	kfpld.Printf("stringify %s", line)

	// convert to float64
	f, err := strconv.ParseFloat(line, 64)
	if err != nil {
		// return error if not valid
		return 0, fmt.Errorf("could not parse '%s' to float64", line)
	}
	kfpd.Printf("parse result: %v\n", f)

	return f, nil
}

func processSample(data []float64) (res ResultSet, err error) {
	// Sort list
	sort.Float64s(data)

	// set data based on sorted slice
	res.n = len(data)
	res.min = data[0]
	res.max = data[res.n-1]

	// sum
	res.sum = calcSum(data)

	// unstable mean
	res.mean = calcMean(data)

	// median
	res.median = calcMedian(data)

	// mode
	res.mode = calcMode(data)

	// variance
	res.variance = calcVariance(data, res.mean)

	// standard deviation
	res.stdev = math.Sqrt(res.variance)

	// Percentiles
	res.p50 = calcPercentile(data, 50)
	res.p75 = calcPercentile(data, 75)
	res.p90 = calcPercentile(data, 90)
	res.p95 = calcPercentile(data, 95)
	res.p99 = calcPercentile(data, 99)

	// quartiles
	var c1 int
	var c2 int
	if res.n%2 == 0 {
		c1 = res.n / 2
		c2 = res.n / 2
	} else {
		c1 = (res.n - 1) / 2
		c2 = c1 + 1
	}
	res.q1 = calcMedian(data[:c1])
	res.q2 = res.median
	res.q3 = calcMedian(data[c2:])

	return res, nil
}

func calcSum(data []float64) (sum float64) {
	for _, v := range data {
		sum += v
	}
	return sum
}

func calcMode(data []float64) (mode []float64) {
	// Return the data if there's only one number
	l := len(data)
	if l == 1 {
		return data
	} else if l == 0 {
		return nil
	}

	// Traverse sorted array,
	// tracking the longest repeating sequence
	mode = make([]float64, 5)
	cnt, maxCnt := 1, 1
	for i := 1; i < l; i++ {
		switch {
		case data[i] == data[i-1]:
			cnt++
		case cnt == maxCnt && maxCnt != 1:
			mode = append(mode, data[i-1])
			cnt = 1
		case cnt > maxCnt:
			mode = append(mode[:0], data[i-1])
			maxCnt, cnt = cnt, 1
		default:
			cnt = 1
		}
	}
	switch {
	case cnt == maxCnt:
		mode = append(mode, data[l-1])
	case cnt > maxCnt:
		mode = append(mode[:0], data[l-1])
		maxCnt = cnt
	}

	// Since length must be greater than 1,
	// check for slices of distinct values
	if maxCnt == 1 || len(mode)*maxCnt == l && maxCnt != l {
		return []float64{}
	}

	return mode
}

func calcVariance(data []float64, mean float64) (variance float64) {
	// Sum the square of the mean subtracted from each number
	for _, n := range data {
		variance += (n - mean) * (n - mean)
	}

	// mean of the squared differences
	return variance / float64(len(data))
}

func calcMean(data []float64) float64 {
	var sum float64
	for _, v := range data {
		sum += v
	}
	return sum / float64(len(data))
}

func calcMedian(data []float64) (median float64) {
	n := len(data)
	if n%2 == 0 {
		median = calcMean(data[n/2-1 : n/2+1])
	} else {
		median = data[n/2]
	}
	return median
}

func calcPercentile(data []float64, percentile float64) (result float64) {
	// generate predicted index
	index := (percentile / 100) * float64(len(data))

	// Check if the index is a whole number
	if index == float64(int64(index)) {

		// Convert float to int
		idx := int(index)

		// Find the value at the index
		result = data[idx-1]

	} else {

		// Convert float to int via truncation
		idx := int(index)

		// Find the average of the index and following values
		result = calcMean([]float64{data[idx-1], data[idx]})

	}

	return result
}
