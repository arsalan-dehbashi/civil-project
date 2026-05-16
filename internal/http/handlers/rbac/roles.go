package rbac

import (
	"net/http"

	rbacdto "github.com/arsalan-dehbashi/civil-project.git/internal/dto/rbac"
	"github.com/arsalan-dehbashi/civil-project.git/internal/utils"
	"github.com/gin-gonic/gin"
)

func (h *Handler) ListRoles(c *gin.Context) {
	roles, err := h.service.ListRoles(c.Request.Context())
	if err != nil {
		handleError(c, err, roleErrors())
		return
	}

	utils.Success(c, http.StatusOK, roles)
}

func (h *Handler) CreateRole(c *gin.Context) {
	var req rbacdto.CreateRoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationError(c, err.Error())
		return
	}

	role, err := h.service.CreateRole(c.Request.Context(), req)
	if err != nil {
		handleError(c, err, roleErrors())
		return
	}

	utils.Created(c, role)
}

func (h *Handler) GetRole(c *gin.Context) {
	id, ok := parseIDParam(c, "id")
	if !ok {
		return
	}

	role, err := h.service.GetRole(c.Request.Context(), id)
	if err != nil {
		handleError(c, err, roleErrors())
		return
	}

	utils.Success(c, http.StatusOK, role)
}

func (h *Handler) UpdateRole(c *gin.Context) {
	id, ok := parseIDParam(c, "id")
	if !ok {
		return
	}

	var req rbacdto.UpdateRoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationError(c, err.Error())
		return
	}

	role, err := h.service.UpdateRole(c.Request.Context(), id, req)
	if err != nil {
		handleError(c, err, roleErrors())
		return
	}

	utils.Success(c, http.StatusOK, role)
}

func (h *Handler) DeleteRole(c *gin.Context) {
	id, ok := parseIDParam(c, "id")
	if !ok {
		return
	}

	if err := h.service.DeleteRole(c.Request.Context(), id); err != nil {
		handleError(c, err, roleErrors())
		return
	}

	utils.NoContent(c)
}

func (h *Handler) SyncRolePermissions(c *gin.Context) {
	id, ok := parseIDParam(c, "id")
	if !ok {
		return
	}

	var req rbacdto.SyncRolePermissionsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationError(c, err.Error())
		return
	}

	role, err := h.service.SyncRolePermissions(c.Request.Context(), id, req.PermissionIDs)
	if err != nil {
		handleError(c, err, errorMessages{
			notFound:  "role or permission not found",
			duplicate: "role permission assignment already exists",
			internal:  "failed to sync role permissions",
		})
		return
	}

	utils.Success(c, http.StatusOK, role)
}

func roleErrors() errorMessages {
	return errorMessages{
		notFound:  "role not found",
		duplicate: "role already exists",
		internal:  "failed to process role request",
	}
}
