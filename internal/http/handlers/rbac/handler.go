package rbac

import (
	rbacservice "github.com/arsalan-dehbashi/civil-project.git/internal/service/rbac"
)

type Handler struct {
	service rbacservice.Service
}

func NewHandler(service rbacservice.Service) *Handler {
	return &Handler{service: service}
}
