package engine

import (
	"github.com/novabankapp/ussdcommon/routes"
)

type Response struct {
	SessionID          string
	NoResponseRequired bool
	ScreenID           int
	Message            string
	Route              routes.Route
	Redirect           bool
	Err                error
}
