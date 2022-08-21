package stores

import (
	"context"
	"github.com/novabankapp/ussdcommon/domain/models/adapters"
	"github.com/novabankapp/ussdcommon/routes"
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

func (s Session) Set(context context.Context, r routes.Route) {
	route := r.Ctrl + "." + r.Action
	err := s.store.SetValue(context, s.routeKey, route)
	if err != nil {
		utils.Panicln("Error setting Session: %s %v", s.routeKey, err)
	}
}

func (s Session) Get(context context.Context) routes.Route {
	rStr, err := s.store.GetValue(context, s.routeKey)
	if err != nil {
		utils.Panicln("Session: %s %v", s.routeKey, err)
	}
	routes_ := strings.Split(rStr, ".")
	if len(routes_) != 2 {
		utils.Panicln("Session: route not found")
	}
	return routes.Route{Ctrl: routes_[0], Action: routes_[1]}
}

func (s Session) Exists(context context.Context) bool {
	b, err := s.store.ValueExists(context, s.routeKey)
	if err != nil {
		utils.Panicln("Session: %v", err)
	}
	return b
}

func (s Session) Close(context context.Context) {
	s.store.DeleteValue(context, s.routeKey)
}
