// Copyright © 2022 AMIS Technologies
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package paillier

import (
	"crypto/elliptic"
	"math/big"

	"github.com/btcsuite/btcd/btcec"
)

const (
	LFactor          = 1
	LpaiFactor       = 5
	MINIMALCHALLENGE = 10
	epsilonFactor    = 2
)

type CurveConfig struct {
	Curve                elliptic.Curve
	TwoExpLAddepsilon    *big.Int
	TwoExpLpaiAddepsilon *big.Int
	TwoExpL              *big.Int
	LAddEpsilon          uint64
	LpaiAddEpsilon       uint64
	Lpai                 uint
}

func NewS256() *CurveConfig {
	curve := btcec.S256()
	epsilon := uint(epsilonFactor * curve.N.BitLen())
	L := uint(LFactor * curve.N.BitLen())
	Lpai := uint(LpaiFactor * curve.N.BitLen())
	return &CurveConfig{
		Curve:                curve,
		TwoExpLAddepsilon:    new(big.Int).Lsh(big2, L+epsilon),
		TwoExpLpaiAddepsilon: new(big.Int).Lsh(big2, Lpai+epsilon),
		TwoExpL:              new(big.Int).Lsh(big2, L),
		LAddEpsilon:          uint64(L + epsilon),
		LpaiAddEpsilon:       uint64(Lpai + epsilon),
		Lpai:                 Lpai,
	}
}
