package commands

import "bufio"

type processLineInput struct {
	line []rune
}

type processReaderInput struct {
	reader *bufio.Reader
}

type processReaderStreamInput struct {
	reader  *bufio.Reader
	refresh int64
	factor  int64
	cap     int64
}

type ResultSet struct {
	n               int
	min             float64
	max             float64
	mean            float64
	mode            []float64
	median          float64
	sum             float64
	stdev           float64
	variance        float64
	p50             float64
	p75             float64
	p90             float64
	p95             float64
	p99             float64
	q1              float64
	q2              float64
	q3              float64
	mildOutliers    []float64
	extremeOutliers []float64
}

func (r ResultSet) ListFields() []string {
	return []string{
		"n",
		"min",
		"max",
		"mean",
		"mode",
		"median",
		"sum",
		"stdev",
		"variance",
		"p50",
		"p75",
		"p90",
		"p95",
		"p99",
		"q1",
		"q2",
		"q3",
		"outliers",
		"mild",
		"extreme",
	}
}

func (r *ResultSet) Get(field string) interface{} {
	switch field {
	case "n":
		return r.n
	case "min":
		return r.min
	case "max":
		return r.max
	case "mean":
		return r.mean
	case "median":
		return r.median
	case "mode":
		if len(r.mode) > 0 {
			return r.mode[0]
		}
		return float64(0)
	case "sum":
		return r.sum
	case "stdev":
		return r.stdev
	case "variance":
		return r.variance
	case "p50":
		return r.p50
	case "p75":
		return r.p75
	case "p90":
		return r.p90
	case "p95":
		return r.p95
	case "p99":
		return r.p99
	case "q1":
		return r.q1
	case "q2":
		return r.q2
	case "q3":
		return r.q3
	case "mild":
		return len(r.mildOutliers)
	case "extreme":
		return len(r.extremeOutliers)
	case "outliers":
		return len(r.mildOutliers) + len(r.extremeOutliers)
	}
	return nil
}

func (r *ResultSet) GetHeader(field string) string {
	switch field {
	case "n":
		return "N"
	case "min":
		return "Min"
	case "max":
		return "Max"
	case "mean":
		return "Mean"
	case "median":
		return "Median"
	case "mode":
		return "Mode"
	case "sum":
		return "Sum"
	case "stdev":
		return "Std Dev"
	case "variance":
		return "Variance"
	case "p50":
		return "p50"
	case "p75":
		return "p75"
	case "p90":
		return "p90"
	case "p95":
		return "p95"
	case "p99":
		return "p99"
	case "q1":
		return "Q1"
	case "q2":
		return "Q2"
	case "q3":
		return "Q3"
	case "mild":
		return "Mild"
	case "extreme":
		return "Extreme"
	case "outliers":
		return "Outliers"
	}
	return ""
}

func (r *ResultSet) GetFormat(field string) string {
	switch field {
	// int64
	case "n":
		return "%d"
	case "mild":
		return "%d"
	case "extreme":
		return "%d"
	case "outliers":
		return "%d"
	// float64
	case "min":
		return "%g"
	case "max":
		return "%g"
	case "mean":
		return "%.2f"
	case "median":
		return "%g"
	case "sum":
		return "%g"
	case "stdev":
		return "%.4f"
	case "variance":
		return "%.4f"
	case "p50":
		return "%g"
	case "p75":
		return "%g"
	case "p90":
		return "%g"
	case "p95":
		return "%g"
	case "p99":
		return "%g"
	case "q1":
		return "%g"
	case "q2":
		return "%g"
	case "q3":
		return "%g"
	// These are arrays
	case "mode":
		return "%g"
	}
	return ""
}
