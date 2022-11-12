package repositories

import (
	"example/src/config"
	"example/src/domains"
	"example/src/errors"
)

func SaveEvent(event *domains.Event) error {
	e := config.Database.Create(event)
	if e.Error != nil {
		return errors.DataAccessLayerError(e.Error.Error())
	}
	return nil
}

func FindOneEventById(id int) *domains.Event {
	var event domains.Event
	config.Database.First(&event, id)
	return &event
}

func FindAllEvents() []domains.Event {
	var events []domains.Event
	config.Database.Find(&events)
	return events
}
