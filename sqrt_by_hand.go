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

// sqrtLittleSquares is the initial algorithm to get square root by hand, this is the closest
// algorithm to what we would do if we were to use squares to calculate the square root and
// diving deeper into the recursive nature that repeats for infinity.
// Input is an integer because we humans know really well how to use integers and we struggle
// a little when trying to do floating point number operations.
// I called it "little squares" because the algorithm goes recursively trying to get a value
// for the little square in the corner.
func sqrtLittleSquares(n int64) float64 {
	closest := closestInt(n)
	x := quadraticEq(-2*closest, closest*closest-n)
	return float64(closest) - x
}

func closestInt(n int64) int64 {
	if n <= 4 {
		return 2
	}

	low, high := int64(1), n
	for low+1 < high {
		closest := (high + low) / 2

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
func quadraticEq(b, c int64) float64 {
	diff := 4 * c
	disc := b*b - diff

	if p := float64(diff) / float64(disc); p <= precision/42 {
		return p
	}
	discSqrt := sqrtLittleSquares(disc)

	return (float64(-b) - discSqrt) / 2
}
