package fgt

import (
	"math"
	"testing"
	"time"
)

func TestGetRetrievability(t *testing.T) {
	points := []time.Time{
		time.Unix(0, 0),
		time.Unix(20*24*3600, 0),
	}
	now := time.Unix(32*24*3600, 0)
	if result := GetRetrievability(points, now); !(math.Abs(result-0.54881) < 1e-5) {
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
	points := []time.Time{
		time.Unix(0, 0),
	}
	now := time.Unix(32*24*3600, 0)
	if result := GetRetrievability(points, now); !(math.Abs(result-0.20190) < 1e-5) {
		t.Error(result)
	} else {
		t.Log(result)
	}
}

func TestGetRetrievability4(t *testing.T) {
	points := []time.Time{
		time.Unix(0, 0),
		time.Unix(10*24*3600, 0),
		time.Unix(20*24*3600, 0),
	}
	now := time.Unix(30*24*3600, 0)
	if result := GetRetrievability(points, now); !(math.Abs(result-0.60653) < 1e-5) {
		t.Error(result)
	} else {
		t.Log(result)
	}
}
