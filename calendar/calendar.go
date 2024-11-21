package calendar

import (
	"log"

	"google.golang.org/api/calendar/v3"
)

// Create a new event in the specified calendar
func CreateEvent(srv *calendar.Service, calendarID string, event *calendar.Event) (*calendar.Event, error) {
	createdEvent, err := srv.Events.Insert(calendarID, event).Do()
	if err != nil {
		log.Printf("Unable to create event: %v", err)
		return nil, err
	}
	return createdEvent, nil
}

// Create a sample event for testing
func ExampleEvent() *calendar.Event {
	return &calendar.Event{
		Summary:     "Test Match Event",
		Location:    "Test Location",
		Description: "A sample match for testing Google Calendar integration.",
		Start: &calendar.EventDateTime{
			DateTime: "2024-11-20T10:00:00-07:00",
			TimeZone: "America/Los_Angeles",
		},
		End: &calendar.EventDateTime{
			DateTime: "2024-11-20T12:00:00-07:00",
			TimeZone: "America/Los_Angeles",
		},
	}
}
