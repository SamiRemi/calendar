package main

import (
	"fmt"
	"time"

	"github.com/SamiRemi/project/app/calendar"
	"github.com/SamiRemi/project/app/events"
)

func main() {

	e, err := events.NewEvent("Встреча", "2026.03.03 15:00")
	if err != nil {
		fmt.Println("Ошибка создания события:", err)
		return
	}
	e1, err1 := events.NewEvent("Отдых", "2026/02/01")
	if err1 != nil {
		fmt.Println("Ошибка создания события:", err)
		return
	}
	calendar.AddEvent("event1", e)
	calendar.AddEvent("event2", e1)
	calendar.ShowEvent()
	time.Sleep(5 * time.Second)
}
