package stores

import (
	"github.com/novabankapp/ussdcommon/domain/models/adapters"
	"github.com/novabankapp/ussdcommon/utils"
	"strings"
)

type Session struct {
	store    Store
	routeKey string
}

func NewSession(store Store, request *adapters.Request) *Session {
	return &Session{
		store:    store,
		routeKey: request.PhoneNumber + "Route",
	}
}

func (s Session) Set(r utils.Route) {
	route := r.Ctrl + "." + r.Action
	err := s.store.SetValue(s.routeKey, route)
	if err != nil {
		utils.Panicln("Session: %v", err)
	}
}

func (s Session) Get() utils.Route {
	rStr, err := s.store.GetValue(s.routeKey)
	if err != nil {
		utils.Panicln("Session: %v", err)
	}
	routes := strings.Split(rStr, ".")
	if len(routes) != 2 {
		utils.Panicln("Session: route not found")
	}
	return utils.Route{Ctrl: routes[0], Action: routes[1]}
}

func (s Session) Exists() bool {
	b, err := s.store.ValueExists(s.routeKey)
	if err != nil {
		utils.Panicln("Session: %v", err)
	}
	return b
}

func (s Session) Close() {
	s.store.DeleteValue(s.routeKey)
}
