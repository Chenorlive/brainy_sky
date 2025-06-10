package role

import (
	"net/http"

	"github.com/Chenorlive/brainy/types"
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

	router.HandleFunc("GET /role/{id}", h.getRole)
	router.HandleFunc("GET /role", h.getRoles)
	router.HandleFunc("POST /role", h.addRole)
	router.HandleFunc("PUT /role", h.updateRole)
	router.HandleFunc("DELETE /role/{id}", h.deleteRole)

}

func (h *Hander) getRole(w http.ResponseWriter, r *http.Request) {
	roleID, err := utils.GetParamUUID(r, "id")
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	role, err := h.store.GetRole(roleID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, role)
}

func (h *Hander) getRoles(w http.ResponseWriter, r *http.Request) {
	roles, err := h.store.GetRoles()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, roles)
}

func (h *Hander) addRole(w http.ResponseWriter, r *http.Request) {

	newRole := &types.NewRole{}
	if err := utils.ParseJSON(r, newRole); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	result, err := h.store.CreateRole(newRole)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, result)
}

func (h *Hander) updateRole(w http.ResponseWriter, r *http.Request) {

	updateRole := &types.UpdateRole{}

	if err := utils.ParseJSON(r, updateRole); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := h.store.UpdateRole(updateRole); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, updateRole)
}

func (h *Hander) deleteRole(w http.ResponseWriter, r *http.Request) {

	roleID, err := utils.GetParamUUID(r, "id")
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := h.store.DeleteRole(roleID); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
