package events

import (
	"time"

	"github.com/SamiRemi/project/app/validation"
	"github.com/araddon/dateparse"
)

type Event struct {
	Title   string
	StartAt time.Time
}

func NewEvent(title string, dateStr string) (Event, error) {
	isValid := validation.IsValidTitle(title)
	if !isValid {
		return Event{}, validation.NewTitleError(title)
	}
	t, err := dateparse.ParseAny(dateStr)
	if err != nil {
		return Event{}, validation.NewDateError(dateStr)
	}
	return Event{
		Title:   title,
		StartAt: t,
	}, nil
}
