package main

import (
	"fmt"
	"time"

	"github.com/SamiRemi/project/app/calendar"
	"github.com/SamiRemi/project/app/events"
)

func main() {
	e := events.Event{
		Title:   "Встреча",
		StartAt: time.Now(),
	}
	calendar.AddEvent("event1", e)
	fmt.Println("Календарь обновлён")
}
