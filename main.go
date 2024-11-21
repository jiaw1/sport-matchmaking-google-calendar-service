package main

import (
	"context"
	"log"
	"os"

	"github.com/jiaw1/sport-matchmaking-google-calendar-service/auth"
	"github.com/jiaw1/sport-matchmaking-google-calendar-service/handler"
	"github.com/jiaw1/sport-matchmaking-google-calendar-service/router"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

func main() {
	// Initialize the Echo router
	r := router.New()

	// Initialize the Google Calendar client
	credentialsPath := "config/service-account-key.json"
	calendarService, err := calendar.NewService(context.Background(), option.WithCredentialsFile(credentialsPath))
	if err != nil {
		log.Fatalf("Failed to create Google Calendar service client: %s", err.Error())
	}

	// Create GoCloak client
	goCloakClient := auth.NewGoCloakClient()

	// Initialize the handler with both the Google Calendar and GoCloak clients
	h := handler.New(calendarService, goCloakClient)

	// Register routes
	apiGroup := r.Group("")
	h.RegisterRoutes(apiGroup)

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port
	}
	r.Logger.Fatal(r.Start(":" + port))
}
