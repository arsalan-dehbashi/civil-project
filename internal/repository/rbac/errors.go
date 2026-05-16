package rbac

import "errors"

var (
	ErrDuplicate = errors.New("duplicate record")
	ErrNotFound  = errors.New("record not found")
)
