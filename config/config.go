package config

import (
	"context"
	"log"

	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

// Initialize the Google Calendar client
func CalendarService(credentialsPath string) (*calendar.Service, error) {
	ctx := context.Background()
	srv, err := calendar.NewService(ctx, option.WithCredentialsFile(credentialsPath))
	if err != nil {
		log.Fatalf("Unable to create calendar client: %v", err)
		return nil, err
	}
	return srv, nil
}
