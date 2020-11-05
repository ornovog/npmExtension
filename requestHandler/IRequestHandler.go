package requestHandler

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type IRequestHandler interface {
	HandleRequest(response http.ResponseWriter, _ *http.Request, params httprouter.Params)
	HandleTreeFormatRequest(response http.ResponseWriter, _ *http.Request, params httprouter.Params)
}
