package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	reqHandle "npmExtension/requestHandler"
)

func main() {
	router := newRouter()
	http.ListenAndServe(":8080", router)
}

func newRouter() *httprouter.Router{
	handler := reqHandle.NewRequestHandler()

	router := httprouter.New()
	router.GET("/package/:name/:version", handler.HandleJsonFormatRequest)
	router.GET("/package/:name/:version/tree-format", handler.HandleTreeFormatRequest)

	return router
}



