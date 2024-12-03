package handler

import (
	"log/slog"
	"net/http"
	"strings"
	"text/template"
	"time"

	"github.com/jiaw1/sport-matchmaking-google-calendar-service/model"
	"github.com/labstack/echo/v4"
	"google.golang.org/api/calendar/v3"
)

func (h *Handler) CreateCalendarEvent(c echo.Context) error {
	matchDetails := model.MatchDetails{}

	// Bind the request body to matchDetails
	if err := c.Bind(&matchDetails); err != nil {
		return HTTPError(err)
	}
	if err := c.Validate(matchDetails); err != nil {
		return HTTPError(err)
	}

	// Create new event
	event := &calendar.Event{
		Summary:     matchDetails.Sport,
		Location:    matchDetails.Location,
		Description: matchDetails.Description,
		Start: &calendar.EventDateTime{
			DateTime: matchDetails.StartsAt.Format(time.RFC3339),
			TimeZone: "UTC",
		},
		End: &calendar.EventDateTime{
			DateTime: matchDetails.EndsAt.Format(time.RFC3339),
			TimeZone: "UTC",
		},
	}

	// Insert event into Google Calendar
	createdEvent, err := h.calendarService.Events.Insert("e4b3894afa61d429a631de599d1f320c9a18a7cc647542abbcbdac0cb304bded@group.calendar.google.com", event).Do()

	if err != nil {
		slog.Error("failed to create calendar event", slog.String("error", err.Error()))
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create calendar event")
	}

	return c.JSON(http.StatusCreated, createdEvent)
}

func (h *Handler) UpdateCalendarEvent(c echo.Context) error {
	eventID := c.Param("id") // Fetch event ID from URL parameter
	matchDetails := model.MatchDetails{}

	// Bind request body to matchDetails
	if err := c.Bind(&matchDetails); err != nil {
		return HTTPError(err)
	}
	if err := c.Validate(matchDetails); err != nil {
		return HTTPError(err)
	}

	// Update event
	event := &calendar.Event{
		Summary:     matchDetails.Sport,
		Location:    matchDetails.Location,
		Description: matchDetails.Description,
		Start: &calendar.EventDateTime{
			DateTime: matchDetails.StartsAt.Format(time.RFC3339),
			TimeZone: "UTC",
		},
		End: &calendar.EventDateTime{
			DateTime: matchDetails.EndsAt.Format(time.RFC3339),
			TimeZone: "UTC",
		},
	}

	updatedEvent, err := h.calendarService.Events.Update("e4b3894afa61d429a631de599d1f320c9a18a7cc647542abbcbdac0cb304bded@group.calendar.google.com", eventID, event).Do()

	if err != nil {
		slog.Error("failed to update calendar event", slog.String("error", err.Error()))
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to update calendar event")
	}

	return c.JSON(http.StatusOK, updatedEvent)
}

func (h *Handler) DeleteCalendarEvent(c echo.Context) error {
	eventID := c.Param("id") // Fetch event ID from URL parameter

	// Delete event from Google Calendar
	err := h.calendarService.Events.Delete("e4b3894afa61d429a631de599d1f320c9a18a7cc647542abbcbdac0cb304bded@group.calendar.google.com", eventID).Do()
	if err != nil {
		slog.Error("failed to delete calendar event", slog.String("error", err.Error()))
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to delete calendar event")
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *Handler) GetCalendarEvents(c echo.Context) error {
	// Define time range for fetching events
	now := time.Now()
	timeMin := now.Format(time.RFC3339)
	timeMax := now.AddDate(0, 1, 0).Format(time.RFC3339) // 1 month ahead

	// Fetch events from Google Calendar
	events, err := h.calendarService.Events.List("primary").
		ShowDeleted(false).
		SingleEvents(true).
		TimeMin(timeMin).
		TimeMax(timeMax).
		OrderBy("startTime").
		Do()

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to fetch calendar events: "+err.Error())
	}

	// Define HTML template with placeholders for event data
	const tmpl = `
	<!DOCTYPE html>
	<html>
		<head><title>Google Calendar Events</title></head>
		<body>
			<h1>Google Calendar Events</h1>
			<ul>
				{{ range . }}
					<li>
						<strong>{{ .Summary }}</strong><br>
						Description: {{ .Description }}<br>
						Location: {{ .Location }}<br>
						Start: {{ .Start.DateTime }}<br>
						End: {{ .End.DateTime }}<br>
						<a href="{{ .HtmlLink }}" target="_blank">View in Google Calendar</a>
					</li>
				{{ end }}
			</ul>
		</body>
	</html>`

	// Compile the template
	t, err := template.New("events").Parse(tmpl)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to parse HTML template: "+err.Error())
	}

	// Render the template to a string (using a buffer)
	var htmlContent strings.Builder
	err = t.Execute(&htmlContent, events.Items)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to render HTML template: "+err.Error())
	}

	// Return the generated HTML
	c.Response().Header().Set("Content-Type", "text/html; charset=utf-8")
	return c.HTML(http.StatusOK, htmlContent.String())
}
