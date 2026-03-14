package main

import (
	"time"

	"github.com/SamiRemi/project/app/calendar"
	"github.com/SamiRemi/project/app/events"
	"github.com/SamiRemi/project/app/validation"
)

func main() {

	e, err := events.NewEvent("Встреча", "2026/03/05 20:00")
	if err != nil {
		validation.ErrorDisplay(err)
		return
	}
	e1, err1 := events.NewEvent("учеба", "2026.03.04")
	if err1 != nil {
		validation.ErrorDisplay(err1)
		return
	}
	calendar.AddEvent("event", e)
	calendar.AddEvent("event1", e1)
	calendar.AddEvent("event3", e)
	calendar.DeleteEvent("event")
	calendar.EditEvent("event3", "созвон", "2026/03/19 20:00")
	calendar.ShowEvent()
	time.Sleep(5 * time.Second)
}
