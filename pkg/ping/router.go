package ping

import (
	"net/http"
)

type Hander struct {
}

func NewHander() *Hander {
	return &Hander{}
}

func (h *Hander) RegisterRoutes(router *http.ServeMux) {

	router.HandleFunc("GET /ping", h.healthCheck)

}

func (h *Hander) healthCheck(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
