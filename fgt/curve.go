/*
Package `fgt` is the abbreviation for `forget`,
which is for computing the remaining rate of knowledge.
*/
package fgt

import (
	"math"
	"time"
)

// The stability of memory, which makes exp(-t/S)≈0.2 when t is 32 days.
const stability = 20 * 24 * time.Hour

// getRetrievability returns R=exp(-t/S), where `R` is the retrievability of memory, `t` is time, and `S` is the stability of memory.
// Ref: Trahan, Donald E. ; Larrabee, Glenn J. (1995). Effect of normal aging on rate of forgetting. Neuropsychology 6. 2 : 115-122.
func getRetrievability(duration time.Duration) float64 {
	return math.Exp(-float64(duration) / float64(stability))
}

// GetRetrievability returns the remaining retrievability of knowledge until now after some time points of learning.
func GetRetrievability(points []time.Time, now time.Time) float64 {
	if len(points) == 0 {
		return 0
	} else {
		return GetRetrievability(points[:len(points)-1], now) + (1-GetRetrievability(points[:len(points)-1], points[len(points)-1]))*getRetrievability(now.Sub(points[len(points)-1]))
	}
}
