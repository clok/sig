package commands

import "math"

// calcSum adds all values in a slice
func calcSum(data []float64) (sum float64) {
	if len(data) == 0 {
		return math.NaN()
	}
	for _, v := range data {
		sum += v
	}
	return sum
}

// calcMode determines the mode values for a given slice. Assumes that the slice is
// already sorted.
func calcMode(data []float64) (mode []float64) {
	// Return the data if there's only one number
	l := len(data)
	if l == 1 {
		return data
	} else if l == 0 {
		return nil
	}

	// track the longest repeating sequence
	mode = make([]float64, 5)
	cnt, maxCnt := 1, 1
	for idx := 1; idx < l; idx++ {
		switch {
		case data[idx] == data[idx-1]:
			cnt++
		case cnt == maxCnt && maxCnt != 1:
			mode = append(mode, data[idx-1])
			cnt = 1
		case cnt > maxCnt:
			mode = append(mode[:0], data[idx-1])
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

	// check for slices of distinct values to avoid returning original dataset
	// and cases where all values occur at the same rate
	if maxCnt == 1 || len(mode)*maxCnt == l && maxCnt != l {
		return []float64{}
	}

	// TODO: add count to return
	return mode
}

// calcVariance determines the population variance via mean of the squared difference
// from the population mean. For performance reasons, the mean of the population must be
// provided to the calcVariance method.
func calcVariance(data []float64, mean float64) (variance float64) {
	if len(data) == 0 {
		return math.NaN()
	}

	// Sum the square of the mean subtracted from each number
	for _, n := range data {
		variance += (n - mean) * (n - mean)
	}

	// mean of the squared differences
	return variance / float64(len(data))
}

// calcMean will perform a simple mean computation
func calcMean(data []float64) float64 {
	if len(data) == 0 {
		return math.NaN()
	}
	var sum float64
	for _, v := range data {
		sum += v
	}
	return sum / float64(len(data))
}

// calcMedian determines the median value. Assumes that the slice is
// already sorted.
func calcMedian(data []float64) (median float64) {
	if len(data) == 0 {
		return math.NaN()
	}
	n := len(data)
	if n%2 == 0 {
		median = calcMean(data[n/2-1 : n/2+1])
	} else {
		median = data[n/2]
	}
	return median
}

// calcPercentile will choose the correct index or average of two values for a given
// percentile value. Assumes that the slice is already sorted.
func calcPercentile(data []float64, percentile float64) (result float64) {
	if len(data) == 0 {
		return math.NaN()
	} else if len(data) == 1 {
		return data[0]
	}
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
