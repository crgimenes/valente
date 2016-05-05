package valente

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
)

//StatusInfo represents info of app status
type StatusInfo struct {
	Started        time.Time
	OpenSessions   int
	ClosedSessions int
	Gotos          int
	SendedBytes    int
	ReceivedBytes  int
}

var (
	//Status is app status
	Status StatusInfo
)

func ValenteStatusHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, Status)
}
