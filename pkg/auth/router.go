package auth

import (
	"net/http"
)

type Hander struct {
}

func NewHander() *Hander {
	return &Hander{}
}

func (h *Hander) RegisterRoutes(router *http.ServeMux) {

	router.HandleFunc("POST /login", h.login)

}

func (h *Hander) login(w http.ResponseWriter, r *http.Request) {
}
