package adapters

type Request struct {
	Stage           int
	SessionID       string
	USSDString      string
	PhoneNumber     string
	ScreenID        int
	ServiceProvider string
	Tag             string
}
