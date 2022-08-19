package stores

import (
	"github.com/novabankapp/ussdcommon/domain/models/adapters"
	"github.com/novabankapp/ussdcommon/utils"
	"strings"
)

type session struct {
	store    Store
	routeKey string
}

func newSession(store Store, request *adapters.Request) *session {
	return &session{
		store:    store,
		routeKey: request.PhoneNumber + "Route",
	}
}

func (s session) Set(r utils.Route) {
	route := r.Ctrl + "." + r.Action
	err := s.store.SetValue(s.routeKey, route)
	if err != nil {
		utils.Panicln("session: %v", err)
	}
}

func (s session) Get() utils.Route {
	rStr, err := s.store.GetValue(s.routeKey)
	if err != nil {
		utils.Panicln("session: %v", err)
	}
	routes := strings.Split(rStr, ".")
	if len(routes) != 2 {
		utils.Panicln("session: route not found")
	}
	return utils.Route{Ctrl: routes[0], Action: routes[1]}
}

func (s session) Exists() bool {
	b, err := s.store.ValueExists(s.routeKey)
	if err != nil {
		utils.Panicln("session: %v", err)
	}
	return b
}

func (s session) Close() {
	s.store.DeleteValue(s.routeKey)
}
