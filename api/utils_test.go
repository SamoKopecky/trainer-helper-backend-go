package api

import (
	"testing"
	"time"
)

func TestHumanTime(t *testing.T) {
	parsedTime, _ := time.Parse(time.RFC3339, "2022-03-23T07:21:00+01:00")
	got := humanTime(parsedTime)
	want := "07:21"

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}

}

func TestHumanDate(t *testing.T) {
	parsedTime, _ := time.Parse(time.RFC3339, "2022-03-23T07:21:00+01:00")
	got := humanDate(parsedTime)
	want := "23-03"

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}

}
