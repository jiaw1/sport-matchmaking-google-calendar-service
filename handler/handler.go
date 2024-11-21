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

func New(calendarService *calendar.Service, goCloakClient *gocloak.GoCloak) *Handler {
	return &Handler{
		calendarService: calendarService,
		goCloakClient:   goCloakClient,
	}
}

func (h *Handler) RegisterRoutes(g *echo.Group) {
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

	matchGroup := g.Group("/calendar")
	matchGroup.POST("/event", h.CreateOrUpdateCalendarEvent)
}
