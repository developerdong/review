package sql

import (
	"net/url"
	"testing"
)

func TestSqlite(t *testing.T) {
	s := Sqlite{}
	if u, err := url.Parse("https://pkg.go.dev/net/url#Parse"); err != nil {
		t.Fatal(err)
	} else if err := s.Add(u); err != nil {
		t.Error(err)
	} else if urlNeedReview, err := s.Review(); err != nil {
		t.Error(err)
	} else {
		t.Log(urlNeedReview)
		if err := s.Delete(u); err != nil {
			t.Error(err)
		}
	}
}
