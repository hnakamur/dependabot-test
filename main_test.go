package main

import (
	"testing"

	"github.com/Masterminds/goutils"
)

func TestAbbreviate(t *testing.T) {
	str := "abcdefg"
	got, err := goutils.Abbreviate(str, 6)
	if err != nil {
		t.Errorf("got unexpected error: %v", err)
	}
	if want := "abc..."; got != want {
		t.Errorf("result mismatch, got=%q, want=%q", got, want)
	}
}
