package main

import (
	"time"

	"github.com/SamiRemi/project/app/calendar"
	"github.com/SamiRemi/project/app/events"
	"github.com/SamiRemi/project/app/validation"
)

func main() {

	e, err := events.NewEvent("Встреча", "2026.03.03 15:00")
	if err != nil {
		validation.ErrorDisplay(err)
		return
	}
	e1, err1 := events.NewEvent("Отдых", "2026/02/01")
	if err1 != nil {
		validation.ErrorDisplay(err1)
		return
	}
	e2, err2 := events.NewEvent("учеба", "20266.03.04")
	if err2 != nil {
		validation.ErrorDisplay(err2)
		return
	}
	calendar.AddEvent("event", e)
	calendar.AddEvent("event1", e1)
	calendar.AddEvent("event2", e2)
	calendar.ShowEvent()
	time.Sleep(5 * time.Second)
}
