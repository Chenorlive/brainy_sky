package userrole

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
	router.HandleFunc("/api/v1/user_roles", h.createUserRole)
	router.HandleFunc("/api/v1/user_roles/", h.getUserRole)
	router.HandleFunc("/api/v1/user_roles/list", h.getUserRoles)
	router.HandleFunc("/api/v1/user_roles/update", h.updateUserRole)
	router.HandleFunc("/api/v1/user_roles/delete", h.deleteUserRole)
}

func (h *handler) createUserRole(w http.ResponseWriter, r *http.Request) {
	newUserRole := &types.NewUserRole{}
	if err := utils.ParseJSON(r, newUserRole); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	obj, err := h.store.CreateUserRole(newUserRole)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, obj)
}

func (h *handler) getUserRole(w http.ResponseWriter, r *http.Request) {
	userRoleID, err := utils.GetParamUUID(r, "id")
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	userRole, err := h.store.GetUserRole(userRoleID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, userRole)
}

func (h *handler) getUserRoles(w http.ResponseWriter, r *http.Request) {
	userRoles, err := h.store.GetUserRoles()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, userRoles)
}

func (h *handler) updateUserRole(w http.ResponseWriter, r *http.Request) {
	updateUserRole := &types.UpdateUserRole{}
	if err := utils.ParseJSON(r, updateUserRole); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := h.store.UpdateUserRole(updateUserRole); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, "User role updated successfully")
}

func (h *handler) deleteUserRole(w http.ResponseWriter, r *http.Request) {
	userRoleID, err := utils.GetParamUUID(r, "id")
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := h.store.DeleteUserRole(userRoleID); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, "User role deleted successfully")
}
