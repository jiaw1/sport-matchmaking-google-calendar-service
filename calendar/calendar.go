package calendar

import (
	"context"
	"log"

	"google.golang.org/api/calendar/v3"
)

// Insert or update event in calendar
func CreateOrUpdate(ctx context.Context, srv *calendar.Service, calendarID string, event *calendar.Event) (*calendar.Event, error) {
	createdEvent, err := srv.Events.Insert(calendarID, event).Context(ctx).Do()
	if err != nil {
		log.Printf("Unable to create or update event: %v", err)
		return nil, err
	}
	return createdEvent, nil
}

// Remove event from calendar
func Delete(ctx context.Context, srv *calendar.Service, calendarID, eventID string) error {
	return srv.Events.Delete(calendarID, eventID).Context(ctx).Do()
}

// Retrieve all events in calendar
func List(ctx context.Context, srv *calendar.Service, calendarID, timeMin, timeMax string) ([]*calendar.Event, error) {
	events, err := srv.Events.List(calendarID).
		TimeMin(timeMin).
		TimeMax(timeMax).
		SingleEvents(true).
		OrderBy("startTime").
		Do()
	if err != nil {
		log.Printf("Unable to list events: %v", err)
		return nil, err
	}
	return events.Items, nil
}
