package calendar

import (
	"errors"
	"fmt"

	"github.com/SamiRemi/project/app/events"
)

var eventsMap = make(map[string]events.Event)

func AddEvent(key string, e events.Event) error {
	if _, ok := eventsMap[key]; ok {
		return errors.New("Событие с именем " + key + " уже существует!")
	}
	if len(key) == 0 {
		return errors.New("Нельзя ввести пустое имя")
	}
	eventsMap[key] = e
	fmt.Println("Событие добавлено:", e.Title)
	return nil
}

func ShowEvent() error {
	if len(eventsMap) == 0 {
		return errors.New("Список пуст")
	}
	for _, v := range eventsMap {
		// Приводим к UTC и выводим с меткой
		utcTime := v.StartAt.UTC()
		fmt.Println(v.Title, "", utcTime.Format("02.01.2006 15:04"))
	}
	return nil
}
