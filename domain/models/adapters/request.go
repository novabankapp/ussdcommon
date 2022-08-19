package adapters

type Request struct {
	Stage       int
	SessionID   string
	Message     string
	PhoneNumber string
	ScreenID    int
	Network     string
	Tag         string
}
