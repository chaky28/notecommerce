package server

import (
	"net/http"
	"strconv"
)

type Server struct {
	Port int
}

type Handler struct {
	Fn      func(http.ResponseWriter, *http.Request)
	Pattern string
}

type BluePrint struct {
	Prefix   string
	Handlers []Handler
}

func (bp *BluePrint) AddHandler(pattern string, fn func(http.ResponseWriter, *http.Request)) {
	bp.Handlers = append(bp.Handlers, Handler{
		Fn:      fn,
		Pattern: bp.Prefix + pattern,
	})
}

func (s Server) ListenAndServe(bps []BluePrint) error {
	for _, bp := range bps {
		for _, handler := range bp.Handlers {
			http.HandleFunc(handler.Pattern, handler.Fn)
		}
	}

	return http.ListenAndServe(":"+strconv.Itoa(s.Port), nil)
}
