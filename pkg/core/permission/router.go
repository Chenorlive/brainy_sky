package permission

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

	router.HandleFunc("GET /permission/{id}", h.getPermission)
	router.HandleFunc("GET /permission", h.getPermissions)
	router.HandleFunc("POST /permission", h.addPermission)
	router.HandleFunc("PUT /permission", h.updatePermission)
	router.HandleFunc("DELETE /permission/{id}", h.deletePermission)

}

func (h *Hander) getPermission(w http.ResponseWriter, r *http.Request) {
	permissionID, err := utils.GetParamUUID(r, "id")
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	permission, err := h.store.GetPermission(permissionID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, permission)
}

func (h *Hander) getPermissions(w http.ResponseWriter, r *http.Request) {
	permissions, err := h.store.GetPermissions()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, permissions)
}

func (h *Hander) addPermission(w http.ResponseWriter, r *http.Request) {
	newPermission := &types.NewPermission{}
	if err := utils.ParseJSON(r, newPermission); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	obj, err := h.store.CreatePermission(newPermission)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, obj)
}

func (h *Hander) updatePermission(w http.ResponseWriter, r *http.Request) {
	updatePermission := &types.UpdatePermission{}
	if err := utils.ParseJSON(r, updatePermission); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := h.store.UpdatePermission(updatePermission); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, updatePermission)
}

func (h *Hander) deletePermission(w http.ResponseWriter, r *http.Request) {
	permissionID, err := utils.GetParamUUID(r, "id")
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := h.store.DeletePermission(permissionID); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusNoContent, nil)
}
