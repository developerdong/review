package fgt

import (
	"math"
	"testing"
	"time"
)

func TestGetRetrievability(t *testing.T) {
	point0 := time.Unix(0, 0)
	point1 := time.Unix(20*24*3600, 0)
	now := time.Unix(32*24*3600, 0)
	if result := GetRetrievability([]time.Time{point0, point1}, now); !(math.Abs(result-0.54881) < 1e-5) {
		t.Error(result)
	} else {
		t.Log(result)
	}
}

func TestGetRetrievability2(t *testing.T) {
	if result := GetRetrievability(nil, time.Now()); result != 0 {
		t.Error(result)
	}
}

func TestGetRetrievability3(t *testing.T) {
	point := time.Unix(0, 0)
	now := time.Unix(32*24*3600, 0)
	if result := GetRetrievability([]time.Time{point}, now); !(math.Abs(result-0.20190) < 1e-5) {
		t.Error(result)
	} else {
		t.Log(result)
	}
}
