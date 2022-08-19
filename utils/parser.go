package utils

import (
	"github.com/novabankapp/ussdcommon/domain/models/adapters"
	"github.com/novabankapp/ussdcommon/domain/models/engine"
)

type Parser interface {
	ConvertResponseToXML(response engine.Response) (xml *string, error error)
	ConvertXMLToRequest(xml string) (request *adapters.Request, error error)
}
