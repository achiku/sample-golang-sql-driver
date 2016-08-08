package main

import (
	"testing"

	"github.com/AdrianLungu/decimal"
)

func TestD(t *testing.T) {
	testData := []struct {
		s        string
		expected decimal.Decimal
		isError  bool
	}{
		{s: "100", expected: decimal.NewFromFloat(100.0), isError: false},
		{s: "200", expected: decimal.NewFromFloat(200.0), isError: false},
		{s: "0.000001", expected: decimal.NewFromFloat(0.000001), isError: false},
		{s: "aa", expected: decimal.NewFromFloat(0), isError: true},
	}
	for _, d := range testData {
		n, err := decimal.NewFromString(d.s)
		if err != nil {
			if d.isError {
				break
			}
			t.Fatal(err)
		}
		if !n.Equals(d.expected) {
			t.Errorf("%+v", n)
		}
		t.Logf("%+v", n)
	}
}
