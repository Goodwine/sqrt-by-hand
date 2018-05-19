// Copyright 2018 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sqrt

import (
	"fmt"
	"math"
	"math/big"
	"testing"
	sqrt_go "github.com/Goodwine/sqrt-by-hand/third_party/go/sqrt"
)

type SqrtFn = func(n float64) float64

var tests = []struct {
	desc string
	fn   SqrtFn
}{
	{"sqrtLittleSquares", func(n float64) float64 { return sqrtLittleSquares(int64(n)) }},
	{"sqrtBinSearch", func(n float64) float64 { return sqrtBinSearch(int64(n)) }},
	{"sqrtNewton", func(n float64) float64 { return sqrtNewton(int64(n)) }},
	{"sqrtLittleSquaresSimpler", func(n float64) float64 { return sqrtLittleSquaresSimpler(n, 0) }},
	{"sqrtLittleSquaresSimplest", sqrtLittleSquaresSimplest},
	{"sqrtLittleSquaresFloats", sqrtLittleSquaresFloats},
	{"sqrtLittleSquaresBigInt", func(n float64) float64 {
		f := sqrtLittleSquaresBigInt(big.NewInt(int64(n)))
		r, _ := f.Float64()
		return r
	}},
	// Go implementations:
	{"std", math.Sqrt},
	{"stdCopy", sqrt_go.StdSqrt},
}

func TestSqrt(t *testing.T) {
	for _, test := range tests {
		t.Run("perfect_"+test.desc, func(t *testing.T) {
			testSqrtPerfect(t, test.fn)
		})
		t.Run("imperfect_"+test.desc, func(t *testing.T) {
			testSqrtImperfect(t, test.fn)
		})
	}
}

func testSqrtPerfect(t *testing.T, fn SqrtFn) {
	tests := []float64{2, 3, 4, 5, 1000, 1337, 42, 21398}

	for _, expected := range tests {
		t.Run(fmt.Sprintf("%.0f", expected), func(t *testing.T) {
			testSqrt(t, fn, expected*expected, float64(expected))
		})
	}
}

func testSqrtImperfect(t *testing.T, fn SqrtFn) {
	tests := []struct {
		input    float64
		expected float64
	}{
		{2, 1.41421356237},
		{3, 1.73205081},
		{105, 10.246950766},
		{1337, 36.5650106},
		{1337*1337*1337 + 1337, 48887.4328432},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%.3f", test.input), func(t *testing.T) {
			testSqrt(t, fn, test.input, test.expected)
		})
	}
}

func testSqrt(t *testing.T, fn SqrtFn, input, expected float64) {
	actual := fn(input)
	if diff := abs(actual - expected); diff > precision {
		t.Errorf("sqrt(%v)=%f, want: %f", input, actual, expected)
	}
}

func BenchmarkSqrt(b *testing.B) {
	for _, test := range tests {
		inputs := []float64{2, 7, 1337, 1337 * 1337, 1337*1337*1337 + 1337}

		b.Run(test.desc, func(b *testing.B) {
			for _, in := range inputs {
				b.Run(fmt.Sprintf("%.3f", in), func(b *testing.B) {
					for n := 0; n < b.N; n++ {
						test.fn(in)
					}
				})
			}
		})
	}
}
