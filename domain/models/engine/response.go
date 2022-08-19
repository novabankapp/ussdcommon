package engine

import "github.com/novabankapp/ussdcommon/utils"

type Response struct {
	SessionID        string
	ResponseRequired bool
	ScreenID         int
	Message          string
	Route            utils.Route
	Redirect         bool
	Err              error
}
