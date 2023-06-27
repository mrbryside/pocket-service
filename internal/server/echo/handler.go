package echo

import "github/mrbryside/pocket-service/internal/service"

type Handler struct {
	saverService    service.ISaverService
	eventGenService service.IEventGenService
}

func NewHandler(ss service.ISaverService, eg service.IEventGenService) Handler {
	return Handler{
		saverService:    ss,
		eventGenService: eg,
	}
}
