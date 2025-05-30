package user

import (
	"net/http"

	"github.com/Chenorlive/brainy/utils"
)

type Hander struct {
	store *Store
}

func NewHander(store *Store) *Hander {
	return &Hander{
		store: store,
	}
}

func (h *Hander) RegisterRoutes(router *http.ServeMux) {

	router.HandleFunc("GET /user/{id}", h.getUser)
	router.HandleFunc("GET /user", h.getUsers)
	router.HandleFunc("POST /todo", h.addUser)
}

func (h *Hander) getUser(w http.ResponseWriter, r *http.Request) {
	p, err := h.store.GetUsers()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	}

	utils.WriteJSON(w, http.StatusOK, p)
}

func (h *Hander) getUsers(w http.ResponseWriter, r *http.Request) {

}

func (h *Hander) addUser(w http.ResponseWriter, r *http.Request) {
}
