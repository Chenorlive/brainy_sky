package rolepermission

import (
	"net/http"

	"github.com/Chenorlive/brainy/types"
	"github.com/Chenorlive/brainy/utils"
)

type handler struct {
	store *Store
}

func NewHandler(store *Store) *handler {
	return &handler{
		store: store,
	}
}

func (h *handler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("/api/v1/role_permissions", h.createRolePermission)
	router.HandleFunc("/api/v1/role_permissions/", h.getRolePermission)
	router.HandleFunc("/api/v1/role_permissions/list", h.getRolePermissions)
	router.HandleFunc("/api/v1/role_permissions/update", h.updateRolePermission)
	router.HandleFunc("/api/v1/role_permissions/delete", h.deleteRolePermission)
}

func (h *handler) createRolePermission(w http.ResponseWriter, r *http.Request) {
	newRolePermission := &types.NewRolePermission{}
	if err := utils.ParseJSON(r, newRolePermission); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	obj, err := h.store.CreateRolePermission(newRolePermission)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, obj)
}

func (h *handler) getRolePermission(w http.ResponseWriter, r *http.Request) {
	rolePermissionID, err := utils.GetParamUUID(r, "id")
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	rolePermission, err := h.store.GetRolePermission(rolePermissionID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, rolePermission)
}

func (h *handler) getRolePermissions(w http.ResponseWriter, r *http.Request) {
	rolePermissions, err := h.store.GetRolePermissions()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, rolePermissions)
}

func (h *handler) updateRolePermission(w http.ResponseWriter, r *http.Request) {
	updateRolePermission := &types.UpdateRolePermission{}
	if err := utils.ParseJSON(r, updateRolePermission); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := h.store.UpdateRolePermission(updateRolePermission); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, updateRolePermission)
}

func (h *handler) deleteRolePermission(w http.ResponseWriter, r *http.Request) {
	rolePermissionID, err := utils.GetParamUUID(r, "id")
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := h.store.DeleteRolePermission(rolePermissionID); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, "Role permission deleted successfully")
}
