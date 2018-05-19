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
	"math"
)

// sqrtLittleSquaresFloats is implementing the same original algorithm of grabbing a square and trying to
// get the error, but usin floating points instead of integers because integers overflow when using larger
// precision due to the recursive/increasing behavior of the algorithm.
func sqrtLittleSquaresFloats(n float64) float64 {
	closest := closestIntFloats(n)
	x := quadraticEqFloats(-2*closest, closest*closest-n)
	return float64(closest) - x
}

func closestIntFloats(n float64) float64 {
	if n <= 4 {
		return 2
	}

	low, high := 1.0, n
	for low+1 < high {
		closest := math.Floor((high + low) / 2)

		switch squared := closest * closest; {
		case squared == n:
			return closest
		case squared < n:
			low = closest
		case squared > n:
			high = closest
		}
	}

	return high
}

// ( -b - sqrt(b²-4ac) ) / (2a)
func quadraticEqFloats(b, c float64) float64 {
	disc := b*b - 4.0*c

	if p := disc / (b * b); p >= .9999999 {
		return 0
	}
	discSqrt := sqrtLittleSquaresFloats(disc)

	return (-b - discSqrt) / 2
}
