package requestHandler

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type IRequestHandler interface {
	HandleJsonFormatRequest(response http.ResponseWriter, _ *http.Request, params httprouter.Params)
	HandleTreeFormatRequest(response http.ResponseWriter, _ *http.Request, params httprouter.Params)
}
