package commands

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func Test_calcSum(t *testing.T) {
	is := assert.New(t)

	t.Run("should be NaN for empty slice", func(t *testing.T) {
		is.True(math.IsNaN(calcSum([]float64{})), "should return NaN")
	})

	t.Run("should be value of single element", func(t *testing.T) {
		is.Equal(float64(1), calcSum([]float64{1}))
	})

	t.Run("should sum values of slice", func(t *testing.T) {
		is.Equal(float64(10), calcSum([]float64{1, 2, 3, 4}))
	})
}

func Test_calcMean(t *testing.T) {
	is := assert.New(t)

	t.Run("should be NaN for empty slice", func(t *testing.T) {
		is.True(math.IsNaN(calcMean([]float64{})), "should return NaN")
	})

	t.Run("should be value of single element", func(t *testing.T) {
		is.Equal(float64(1), calcMean([]float64{1}))
	})

	t.Run("should generate a simple mean value of slice", func(t *testing.T) {
		is.Equal(2.5, calcMean([]float64{1, 2, 3, 4}))
	})
}

func Test_calcMedian(t *testing.T) {
	is := assert.New(t)

	t.Run("should be NaN for empty slice", func(t *testing.T) {
		is.True(math.IsNaN(calcMedian([]float64{})), "should return NaN")
	})

	t.Run("should be value of single element", func(t *testing.T) {
		is.Equal(float64(1), calcMedian([]float64{1}))
	})

	t.Run("should return the mean of two neighboring elements at the midpoint for even length slice", func(t *testing.T) {
		is.Equal(2.5, calcMedian([]float64{1, 2, 3, 10}))
	})

	t.Run("should return the median of an odd length slice", func(t *testing.T) {
		is.Equal(float64(2), calcMedian([]float64{1, 1, 2, 3, 10}))
	})
}

func Test_calcMode(t *testing.T) {
	is := assert.New(t)

	t.Run("should be nil for empty slice", func(t *testing.T) {
		is.Nil(calcMode([]float64{}))
	})

	t.Run("should be the input value of single element", func(t *testing.T) {
		is.Equal([]float64{1}, calcMode([]float64{1}))
	})

	t.Run("should be an empty slice if all values occur at the same rate", func(t *testing.T) {
		is.Equal([]float64{}, calcMode([]float64{1, 2, 3, 4}))
		is.Equal([]float64{}, calcMode([]float64{1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6}))
	})

	t.Run("should mode", func(t *testing.T) {
		is.Equal([]float64{1}, calcMode([]float64{1, 1, 3, 4}))
		is.Equal([]float64{1, 4}, calcMode([]float64{1, 1, 3, 4, 4}), "edges failed")
		is.Equal([]float64{1, 2, 3, 4, 5, 6}, calcMode([]float64{1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7}), "more than 5 failed")
	})
}

func Test_calcVariance(t *testing.T) {
	is := assert.New(t)

	t.Run("should be NaN for empty slice", func(t *testing.T) {
		is.True(math.IsNaN(calcVariance([]float64{}, 0)), "should return NaN")
	})

	t.Run("should be 0 for a single element slice", func(t *testing.T) {
		is.Equal(float64(0), calcVariance([]float64{100}, 100))
	})

	t.Run("should calculate the variance for a dataset", func(t *testing.T) {
		is.Equal(1.25, calcVariance([]float64{1, 2, 3, 4}, calcMean([]float64{1, 2, 3, 4})))
		is.Equal(12.5, calcVariance([]float64{1, 2, 3, 10}, calcMean([]float64{1, 2, 3, 10})))
	})
}

func Test_calcPercentile(t *testing.T) {
	is := assert.New(t)

	t.Run("should be NaN for empty slice", func(t *testing.T) {
		is.True(math.IsNaN(calcPercentile([]float64{}, 50)), "should return NaN")
	})

	t.Run("should be value of single element", func(t *testing.T) {
		is.Equal(float64(1), calcPercentile([]float64{1}, 50))
	})

	t.Run("should select whole number elements for even length slices", func(t *testing.T) {
		is.Equal(float64(2), calcPercentile([]float64{1, 2, 3, 4}, 50))
	})

	t.Run("should select the average of two neighboring elements for odd length slices", func(t *testing.T) {
		is.Equal(2.5, calcPercentile([]float64{1, 2, 3, 4, 5}, 50))
	})
}
