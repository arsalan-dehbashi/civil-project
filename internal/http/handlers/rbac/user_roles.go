package rbac

import (
	"net/http"

	rbacdto "github.com/arsalan-dehbashi/civil-project.git/internal/dto/rbac"
	"github.com/arsalan-dehbashi/civil-project.git/internal/utils"
	"github.com/gin-gonic/gin"
)

func (h *Handler) SyncUserRoles(c *gin.Context) {
	id, ok := parseIDParam(c, "id")
	if !ok {
		return
	}

	var req rbacdto.SyncUserRolesRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationError(c, err.Error())
		return
	}

	user, err := h.service.SyncUserRoles(c.Request.Context(), id, req.RoleIDs)
	if err != nil {
		handleError(c, err, errorMessages{
			notFound:  "user or role not found",
			duplicate: "user role assignment already exists",
			internal:  "failed to sync user roles",
		})
		return
	}

	utils.Success(c, http.StatusOK, user)
}
