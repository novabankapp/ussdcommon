package adapters

import (
	"github.com/novabankapp/ussdcommon/domain/models/adapters"
	"github.com/novabankapp/ussdcommon/domain/models/engine"
)

type EngineRequest interface {
	CallEngine(req adapters.Request) (response engine.Response, error error)
}

type RequestAdapter interface {
	GetRequest() *adapters.Request
}

type ResponseAdapter interface {
	SetResponse(engine.Response)
}
