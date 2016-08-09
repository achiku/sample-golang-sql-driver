package main

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/AdrianLungu/decimal"
)

func TestDecimal(t *testing.T) {
	testData := []struct {
		s        string
		expected decimal.Decimal
		isError  bool
	}{
		{s: "100", expected: decimal.NewFromFloat(100.0), isError: false},
		{s: "200", expected: decimal.NewFromFloat(200.0), isError: false},
		{s: "0.000001", expected: decimal.NewFromFloat(0.000001), isError: false},
		{s: "aa", expected: decimal.NewFromFloat(0), isError: true},
		{s: "", expected: decimal.NewFromFloat(0), isError: true},
	}
	for _, d := range testData {
		n, err := decimal.NewFromString(d.s)
		if err != nil {
			if !d.isError {
				t.Fatal(err)
			}
		}
		if !n.Equals(d.expected) {
			t.Errorf("%+v", n)
		}
		t.Logf("%+v", n)
	}
}

func TestStrconv(t *testing.T) {
	testData := []struct {
		s        string
		expected int64
		isError  bool
	}{
		{s: "100", expected: int64(100), isError: false},
		{s: "200", expected: int64(200), isError: false},
		{s: "9", expected: int64(9), isError: false},
		{s: "aa", expected: int64(0), isError: true},
		{s: "", expected: int64(0), isError: true},
	}
	for _, d := range testData {
		i, err := strconv.ParseInt(d.s, 10, 64)
		if err != nil {
			if !d.isError {
				t.Fatal(err)
			}
		}
		t.Logf("%v %d", reflect.TypeOf(i), i)
	}
}
