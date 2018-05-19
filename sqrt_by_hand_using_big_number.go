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
	"math/big"
)

// sqrtLittleSquaresBigInt is the original algorithm trying to deal with integer overflows
// caused by the increasing/recursive behavior of the algorithm that requires getting the
// square root of larger and larger integers.
// Since big numbers allocate memory, this algorithm is the one that performs the worst,
// but it's way more precise while keeping the same "only integers" constraint.
func sqrtLittleSquaresBigInt(n *big.Int) *big.Float {
	ret := (&big.Float{}).SetPrec(fpPrec)

	closest := closestBigInt(n)
	b, c := &big.Int{}, &big.Int{}
	x := quadraticEqBigInt(b.Mul(closest, bigMinusTwo), c.Mul(closest, closest).Sub(c, n))

	return ret.SetInt(closest).Sub(ret, x)
}

var bigFour = big.NewInt(4)
var bigTwo = big.NewInt(2)
var bigMinusTwo = big.NewInt(-2)
var bigMinusOneFloat = big.NewFloat(-1)
var bigPoint5 = big.NewFloat(0.5)
var bigOne = big.NewInt(1)
var bigZero = big.NewInt(0)

func closestBigInt(n *big.Int) *big.Int {
	if n.Cmp(bigFour) <= 0 {
		return bigTwo
	}

	low, high, tmp := bigOne, &big.Int{}, &big.Int{}
	high.Set(n)
	for tmp.Add(low, bigOne).Cmp(high) < 0 {
		closest := &big.Int{}
		closest.Add(low, high)
		closest.Div(closest, bigTwo)

		squared := &big.Int{}
		squared.Mul(closest, closest)

		switch cmp := squared.Cmp(n); {
		case cmp == 0:
			return closest
		case cmp < 0:
			low = closest
		case cmp > 0:
			high = closest
		}
	}

	return high
}

var precisionBig = big.NewFloat(precision / 1000000000)

const fpPrec = 50

// ( -b - sqrt(b²-4ac) ) / (2a)
func quadraticEqBigInt(b, c *big.Int) *big.Float {
	diff, disc := &big.Int{}, &big.Int{}
	diff.Mul(bigFour, c)
	disc = disc.Mul(b, b).Sub(disc, diff)

	p := (&big.Float{}).SetPrec(fpPrec)
	p.Quo((&big.Float{}).SetInt(diff), (&big.Float{}).SetInt(disc))
	if cmp := p.Cmp(precisionBig); cmp < 0 {
		return p
	}

	discSqrt := sqrtLittleSquaresBigInt(disc)
	ret := (&big.Float{}).SetPrec(fpPrec)
	return ret.SetInt(b.Neg(b)).Sub(ret, discSqrt).Mul(ret, bigPoint5)
}
