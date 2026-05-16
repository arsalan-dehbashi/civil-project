package rbac

import (
	"errors"
	"net/http"
	"strconv"

	rbacrepo "github.com/arsalan-dehbashi/civil-project.git/internal/repository/rbac"
	rbacservice "github.com/arsalan-dehbashi/civil-project.git/internal/service/rbac"
	"github.com/arsalan-dehbashi/civil-project.git/internal/utils"
	"github.com/gin-gonic/gin"
)

func parseIDParam(c *gin.Context, name string) (uint64, bool) {
	id, err := strconv.ParseUint(c.Param(name), 10, 64)
	if err != nil || id == 0 {
		utils.ValidationError(c, "invalid "+name)
		return 0, false
	}
	return id, true
}

func handleError(c *gin.Context, err error, messages errorMessages) {
	var validationErr rbacservice.ValidationError
	switch {
	case errors.As(err, &validationErr):
		utils.ValidationError(c, validationErr.Error())
	case errors.Is(err, rbacrepo.ErrNotFound):
		utils.Error(c, http.StatusNotFound, messages.notFound)
	case errors.Is(err, rbacrepo.ErrDuplicate):
		utils.Error(c, http.StatusConflict, messages.duplicate)
	default:
		utils.Error(c, http.StatusInternalServerError, messages.internal)
	}
}

type errorMessages struct {
	notFound  string
	duplicate string
	internal  string
}
