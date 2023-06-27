package echo

import "github/mrbryside/pocket-service/internal/service"

type Handler struct {
	saverService service.ISaverService
}

func NewHandler(ss service.ISaverService) Handler {
	return Handler{
		saverService: ss,
	}
}
