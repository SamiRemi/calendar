package calendar

import (
	"errors"
	"fmt"

	"github.com/SamiRemi/project/app/events"
	"github.com/SamiRemi/project/app/validation"
	"github.com/araddon/dateparse"
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
		utcTime := v.StartAt.UTC()
		fmt.Println(v.Title, "", utcTime.Format("02.01.2006 15:04"))
	}
	return nil
}

func DeleteEvent(title string) error {
	if _, ok := eventsMap[title]; !ok {
		return errors.New("Событие с именем " + title + " не существует")
	}
	delete(eventsMap, title)
	return nil
}

func EditEvent(name, newTitle, dateStr string) error {
	date, dateErr := dateparse.ParseAny(dateStr)
	if dateErr != nil {
		return dateErr
	}
	err := fullValidation(name, newTitle)
	if err != nil {
		return err
	}
	eventsMap[name] = events.Event{
		Title:   newTitle,
		StartAt: date,
	}
	fmt.Println("Событие", name, "успешно изменено!")
	return nil
}

func fullValidation(name, title string) error {
	if name == "" {
		return errors.New("Имя события не может быть пустым")
	}
	if _, ok := eventsMap[name]; !ok {
		return errors.New("Событие не найдено ")
	}
	if ok := validation.IsValidTitle(title); !ok {
		return errors.New("Заголовок введён некорректно")
	}
	if eventsMap[name].Title == title {
		return errors.New("Новый заголовок совпадает с текущим")
	}
	return nil
}
