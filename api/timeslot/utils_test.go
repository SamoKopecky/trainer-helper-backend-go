package timeslot

import (
	"testing"
	"time"
	"trainer-helper/model"
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

func TestToModelNil(t *testing.T) {
	params := timeslotPutParams{
		Id:   1,
		Name: nil,
	}
	got := params.ToModel()
	want := model.Timeslot{
		IdModel: model.IdModel{
			Id: 1,
		},
		Name: "",
	}

	if got != want {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func TestToModelNotNil(t *testing.T) {
	name := "name"
	params := timeslotPutParams{
		Id:   1,
		Name: &name,
	}
	got := params.ToModel()
	want := model.Timeslot{
		IdModel: model.IdModel{
			Id: 1,
		},
		Name: "name",
	}

	if got != want {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
