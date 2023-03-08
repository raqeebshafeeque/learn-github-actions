package maths_test

import (
	"learn-actions/maths"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddNumbers(t *testing.T) {
	testTable := []struct {
		Num1, Num2, Want int64
	}{
		{Num1: 2, Num2: 2, Want: 4},
		{Num1: 0, Num2: 0, Want: 0},
		{Num1: -1, Num2: -1, Want: -2},
		{Num1: -1, Num2: 1, Want: 0},
		{Num1: 1, Num2: -1, Want: 0},
	}
	for _, test := range testTable {
		t.Run("", func(t *testing.T) {
			assert := assert.New(t)
			got := maths.AddNumbers(test.Num1, test.Num2)
			assert.Equal(got, test.Want)
		})
	}
}

func TestSubtractNumbers(t *testing.T) {
	testTable := []struct {
		Num1, Num2, Want int64
	}{
		{Num1: 2, Num2: 2, Want: 0},
		{Num1: 0, Num2: 0, Want: 0},
		{Num1: -1, Num2: -1, Want: 0},
		{Num1: -1, Num2: 1, Want: -2},
		{Num1: 1, Num2: -1, Want: 2},
	}
	for _, test := range testTable {
		t.Run("", func(t *testing.T) {
			assert := assert.New(t)
			got := maths.SubtractNumbers(test.Num1, test.Num2)
			assert.Equal(got, test.Want)
		})
	}
}
