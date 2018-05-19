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

const (
	bits    = uint(10)
	divisor = float64(uint64(1) << bits)
	factor  = 1 << (bits << 1)
)

// sqrtLittleSquaresSimplest is the final version of the algorithm I started with.
// I arrived to this point using some slightly more advanced (but still simple) algebra.
// Since I couldn't arrive to a better solution, I called this "simplest" even though
// the std algorithm performs way better.
// I think I can get this algorithm to perform better by replacing the L() function for
// some other function tests bit by bit instead of using binary search, but this should still
// be doable easily by a person with enough patience.
// f(n)=L(4^(N)n)/2^N
func sqrtLittleSquaresSimplest(n float64) float64 {
	n = n * float64(factor)
	return closestIntSimplest(n) / divisor
}

// L(n)=closest int whose square is less than or equal to n
func closestIntSimplest(n float64) float64 {
	if n <= 4 {
		return 2
	}

	low, high := 2.0, n
	for low+1 < high {
		closest := float64(int64((high + low) / 2.0))

		switch comparison := n / closest / closest; {
		case comparison == 1:
			return closest
		case comparison > 1:
			low = closest
		case comparison < 1:
			high = closest
		}
	}

	return low
}
