package engine

type Request struct {
	SessionID        string
	ResponseRequired bool
	ScreenID         int
	Message          string
}
