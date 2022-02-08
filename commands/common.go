package commands

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/clok/kemba"
	"github.com/montanaflynn/stats"
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

			// TODO: Add support for refresh rate
			// if lines%100 == 0 {
			//	res, err := processSample(sample)
			//	if err != nil {
			//		return nil, err
			//	}
			//	var row string
			//	for _, field := range res.ListFields() {
			//		if row == "" {
			//			row = fmt.Sprintf(res.GetFormat(field), res.Get(field))
			//		} else {
			//			format := "%s\t" + res.GetFormat(field)
			//			row = fmt.Sprintf(format, row, res.Get(field))
			//		}
			//	}
			//	fmt.Printf("%s\r", row)
			// }

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

func processSample(data []float64) (ResultSet, error) {
	var res ResultSet
	var err error
	res.n = len(data)

	var min float64
	min, err = stats.Min(data)
	if err != nil {
		return res, err
	}
	res.min = min

	var max float64
	max, err = stats.Max(data)
	if err != nil {
		return res, err
	}
	res.max = max

	var mean float64
	mean, err = stats.Mean(data)
	if err != nil {
		return res, err
	}
	res.mean = mean

	var mode []float64
	mode, err = stats.Mode(data)
	if err != nil {
		return res, err
	}
	res.mode = mode

	var median float64
	median, err = stats.Median(data)
	if err != nil {
		return res, err
	}
	res.median = median

	var sum float64
	sum, err = stats.Sum(data)
	if err != nil {
		return res, err
	}
	res.sum = sum

	var variance float64
	variance, err = stats.Variance(data)
	if err != nil {
		return res, err
	}
	res.variance = variance

	var stdev float64
	stdev, err = stats.StandardDeviation(data)
	if err != nil {
		return res, err
	}
	res.stdev = stdev

	var p50 float64
	p50, err = stats.Percentile(data, 50)
	if err != nil {
		return res, err
	}
	res.p50 = p50

	var p75 float64
	p75, err = stats.Percentile(data, 75)
	if err != nil {
		return res, err
	}
	res.p75 = p75

	var p90 float64
	p90, err = stats.Percentile(data, 90)
	if err != nil {
		return res, err
	}
	res.p90 = p90

	var p95 float64
	p95, err = stats.Percentile(data, 95)
	if err != nil {
		return res, err
	}
	res.p95 = p95

	var p99 float64
	p99, err = stats.Percentile(data, 99)
	if err != nil {
		return res, err
	}
	res.p99 = p99

	var qs stats.Quartiles
	qs, err = stats.Quartile(data)
	if err != nil {
		return res, err
	}
	res.q1 = qs.Q1
	res.q2 = qs.Q2
	res.q3 = qs.Q3

	var outliers stats.Outliers
	outliers, err = stats.QuartileOutliers(data)
	if err != nil {
		return res, err
	}
	res.mildOutliers = outliers.Mild
	res.extremeOutliers = outliers.Extreme

	return res, nil
}
