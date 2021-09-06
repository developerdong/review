package fgt

import (
	"math"
	"testing"
	"time"
)

// The retrievability at time 0 should be 1.
func TestGetRetrievability(t *testing.T) {
	points := []time.Time{
		time.Unix(0, 0),
	}
	now := time.Unix(0, 0)
	if result := GetRetrievability(points, now); result != 1 {
		t.Error(result)
	} else {
		t.Log(result)
	}
}

// The retrievability should not be smaller than `r`.
func TestGetRetrievability1(t *testing.T) {
	points := []time.Time{
		time.Unix(0, 0),
	}
	now := time.Unix(math.MaxUint32, 0)
	if result := GetRetrievability(points, now); result >= r {
		t.Log(result)
	} else {
		t.Error(result)
	}
}

// The retrievability should be 0 if we have never learned it.
func TestGetRetrievability2(t *testing.T) {
	if result := GetRetrievability(nil, time.Now()); result != 0 {
		t.Error(result)
	}
}

// The retrievability computed by this program should be equal to that computed by hand.
func TestGetRetrievability3(t *testing.T) {
	points := []time.Time{
		time.Unix(0, 0),
	}
	now := time.Unix(18*24*3600, 0)
	if result := GetRetrievability(points, now); !(math.Abs(result-0.288) < 1e-5) {
		t.Error(result)
	} else {
		t.Log(result)
	}
}

// The later we review, the better the result is.
func TestGetRetrievability4(t *testing.T) {
	points1 := []time.Time{
		time.Unix(0, 0),
		time.Unix(10*24*3600, 0),
		time.Unix(20*24*3600, 0),
	}
	points2 := []time.Time{
		time.Unix(0, 0),
		time.Unix(10*24*3600, 0),
		time.Unix(20*24*3600+1, 0),
	}
	now := time.Unix(30*24*3600, 0)
	if result1, result2 := GetRetrievability(points1, now), GetRetrievability(points2, now); !(result1 < result2) {
		t.Error(result1, result2)
	} else {
		t.Log(result1, result2)
	}
}
