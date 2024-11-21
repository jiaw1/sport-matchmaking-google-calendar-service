package handler

import (
	"net/http"
	"time"

	"github.com/Nerzal/gocloak/v13"
	"github.com/jiaw1/sport-matchmaking-google-calendar-service/model"
	"github.com/labstack/echo/v4"
	"google.golang.org/api/calendar/v3"
)

type Handler struct {
	calendarService *calendar.Service
	goCloakClient   *gocloak.GoCloak
}

// Initialize new handler with Google Calendar service and GoCloak client
func New(calendarService *calendar.Service, goCloakClient *gocloak.GoCloak) *Handler {
	return &Handler{
		calendarService: calendarService,
		goCloakClient:   goCloakClient,
	}
}

// RegisterRoutes sets up the routes for the application
func (h *Handler) RegisterRoutes(g *echo.Group) {
	// General routes
	g.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello world!")
	})

	g.GET("/time", func(c echo.Context) error {
		return c.JSON(http.StatusOK, model.TimeResponse{Time: time.Now().UTC()})
	})

	// Redirect to Google Calendar
	g.GET("/calendar", func(c echo.Context) error {
		return c.Redirect(http.StatusFound, "https://calendar.google.com/calendar/embed?src=e4b3894afa61d429a631de599d1f320c9a18a7cc647542abbcbdac0cb304bded%40group.calendar.google.com&ctz=Europe%2FHelsinki")
	})

	// Calendar-specific routes
	g.POST("/calendar/event", h.CreateCalendarEvent)       // Create new event
	g.PUT("/calendar/event/:id", h.UpdateCalendarEvent)    // Update event
	g.DELETE("/calendar/event/:id", h.DeleteCalendarEvent) // Delete event
	g.GET("/calendar/events", h.GetCalendarEvents)         // Fetch all events
}
