package rbac

import (
	"net/http"

	rbacdto "github.com/arsalan-dehbashi/civil-project.git/internal/dto/rbac"
	"github.com/arsalan-dehbashi/civil-project.git/internal/utils"
	"github.com/gin-gonic/gin"
)

func (h *Handler) ListPermissions(c *gin.Context) {
	permissions, err := h.service.ListPermissions(c.Request.Context())
	if err != nil {
		handleError(c, err, permissionErrors())
		return
	}

	utils.Success(c, http.StatusOK, permissions)
}

func (h *Handler) CreatePermission(c *gin.Context) {
	var req rbacdto.CreatePermissionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationError(c, err.Error())
		return
	}

	permission, err := h.service.CreatePermission(c.Request.Context(), req)
	if err != nil {
		handleError(c, err, permissionErrors())
		return
	}

	utils.Created(c, permission)
}

func (h *Handler) GetPermission(c *gin.Context) {
	id, ok := parseIDParam(c, "id")
	if !ok {
		return
	}

	permission, err := h.service.GetPermission(c.Request.Context(), id)
	if err != nil {
		handleError(c, err, permissionErrors())
		return
	}

	utils.Success(c, http.StatusOK, permission)
}

func (h *Handler) UpdatePermission(c *gin.Context) {
	id, ok := parseIDParam(c, "id")
	if !ok {
		return
	}

	var req rbacdto.UpdatePermissionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationError(c, err.Error())
		return
	}

	permission, err := h.service.UpdatePermission(c.Request.Context(), id, req)
	if err != nil {
		handleError(c, err, permissionErrors())
		return
	}

	utils.Success(c, http.StatusOK, permission)
}

func (h *Handler) DeletePermission(c *gin.Context) {
	id, ok := parseIDParam(c, "id")
	if !ok {
		return
	}

	if err := h.service.DeletePermission(c.Request.Context(), id); err != nil {
		handleError(c, err, permissionErrors())
		return
	}

	utils.NoContent(c)
}

func permissionErrors() errorMessages {
	return errorMessages{
		notFound:  "permission not found",
		duplicate: "permission already exists",
		internal:  "failed to process permission request",
	}
}
