package commands

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ResultSet(t *testing.T) {
	is := assert.New(t)

	result := ResultSet{
		n:               100,
		min:             1,
		max:             100,
		mean:            50.5,
		mode:            []float64{},
		median:          50.5,
		sum:             5050,
		stdev:           28.86607004772212,
		variance:        833.25,
		p50:             50,
		p75:             75,
		p90:             90,
		p95:             95,
		p99:             99,
		q1:              25.5,
		q2:              50.5,
		q3:              75.5,
		mildOutliers:    []float64{0},
		extremeOutliers: []float64{0},
	}

	t.Run("Get [complete]", func(t *testing.T) {
		for _, field := range result.ListFields() {
			is.NotNil(result.Get(field))
		}
	})

	t.Run("Get [nil]", func(t *testing.T) {
		result := ResultSet{}
		is.Nil(result.Get("test"))
	})

	t.Run("GetHeader [complete]", func(t *testing.T) {
		for _, field := range result.ListFields() {
			is.NotEqual("", result.GetHeader(field))
		}
	})

	t.Run("GetHeader [nil]", func(t *testing.T) {
		result := ResultSet{}
		is.Equal("", result.GetHeader("test"))
	})

	t.Run("GetFormat [complete]", func(t *testing.T) {
		for _, field := range result.ListFields() {
			is.NotEqual("", result.GetFormat(field))
		}
	})

	t.Run("GetFormat [nil]", func(t *testing.T) {
		result := ResultSet{}
		is.Equal("", result.GetFormat("test"))
	})
}
