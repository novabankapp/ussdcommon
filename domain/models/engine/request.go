package engine

type Response struct {
	SessionID        string
	ResponseRequired bool
	ScreenID         int
	Message          string
}
