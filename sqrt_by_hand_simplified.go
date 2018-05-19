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

// sqrtLittleSquaresSimpler is the simplified version of the algorithm after I reaslized that
// I don't need to use a quadratic function to calculate the next step, but instead I can go
// directly into the next element, recursively until my desired precision.
// I call it "simpler" because it's the same original algorithm simplyfied with some basic
// algebra, and since there is an even better algorithm born from the same idea, I used "er"
// instead of "est".
// f(n) = L(n) + g(n, L(n))
func sqrtLittleSquaresSimpler(n float64, pow2 uint) float64 {
	// Approximation to end recursion
	if n > 1<<50 {
		closest := closestIntSimpler(n)
		return closest / float64(int64(1)<<pow2)
	}

	// g(n) = f(4n) / 2
	return sqrtLittleSquaresSimpler(4*n, pow2+1)
}

// L(n)=closest int whose square is less than or equal to n
func closestIntSimpler(n float64) float64 {
	if n <= 4 {
		return 2
	}

	low, high := 2.0, n
	for low+1 < high {
		closest := float64(int64((high + low) / 2.0))

		switch squared := closest * closest; {
		case squared == n:
			return closest
		case squared < n:
			low = closest
		case squared > n:
			high = closest
		}
	}

	return low
}
