/*
Package `fgt` is the abbreviation for `forget`,
which is for computing the remaining rate of knowledge.
*/
package fgt

import (
	"math"
	"time"
)

const (
	// The self-defined forgetting factor.
	k = 1 / (3 * math.E * float64(24*time.Hour))
	// The final retrievability of memory when t➡+∞
	r = 0.2
)

// According to the Newton's Law of Cooling, T=H+(T0-H)*exp(-k(t-t0)), where
// T: the temperature of thing at `t`
// H: the ambient temperature
// T0: the temperature of thing at `t0`
// k: the cooling factor
// t: the current time
// t0: the initial time
// Similarly, the forgetting curve follows this formula.
// We can use `T` as the retrievability of memory at `t`,
// `H` as the final retrievability of memory,
// `T0` as the retrievability of memory at `t0`,
// and `k` as the forgetting factor.
// According to the data published by Ebbinghaus, let H=0.2, and t0=0,
// then of course T0 is 1, so the forgetting formula is:
// T=0.2+0.8*exp(-kt)
func getRetrievability(duration time.Duration) float64 {
	return r + (1-r)*math.Exp(-k*float64(duration))
}

// GetRetrievability returns the remaining retrievability of knowledge until now after some time points of learning.
func GetRetrievability(points []time.Time, now time.Time) float64 {
	if len(points) == 0 {
		return 0
	} else {
		return GetRetrievability(points[:len(points)-1], now) + (1-GetRetrievability(points[:len(points)-1], points[len(points)-1]))*getRetrievability(now.Sub(points[len(points)-1]))
	}
}
