package rbac

import (
	"context"
	"errors"

	"github.com/arsalan-dehbashi/civil-project.git/internal/models"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

type GormRepository struct {
	db *gorm.DB
}

func NewGormRepository(db *gorm.DB) *GormRepository {
	return &GormRepository{db: db}
}

func (r *GormRepository) ListRoles(ctx context.Context) ([]models.Role, error) {
	var roles []models.Role
	err := r.db.WithContext(ctx).Preload("Permissions").Order("level DESC, id ASC").Find(&roles).Error
	return roles, normalizeError(err)
}

func (r *GormRepository) CreateRole(ctx context.Context, role *models.Role) error {
	return normalizeError(r.db.WithContext(ctx).Create(role).Error)
}

func (r *GormRepository) GetRole(ctx context.Context, id uint64) (*models.Role, error) {
	var role models.Role
	err := r.db.WithContext(ctx).Preload("Permissions").First(&role, id).Error
	if err != nil {
		return nil, normalizeError(err)
	}
	return &role, nil
}

func (r *GormRepository) UpdateRole(ctx context.Context, id uint64, updates map[string]any) (*models.Role, error) {
	db := r.db.WithContext(ctx)

	var role models.Role
	if err := db.First(&role, id).Error; err != nil {
		return nil, normalizeError(err)
	}

	if err := db.Model(&role).Updates(updates).Error; err != nil {
		return nil, normalizeError(err)
	}

	return r.GetRole(ctx, id)
}

func (r *GormRepository) DeleteRole(ctx context.Context, id uint64) error {
	result := r.db.WithContext(ctx).Delete(&models.Role{}, id)
	if result.Error != nil {
		return normalizeError(result.Error)
	}
	if result.RowsAffected == 0 {
		return ErrNotFound
	}
	return nil
}

func (r *GormRepository) SyncRolePermissions(ctx context.Context, roleID uint64, permissionIDs []uint64) (*models.Role, error) {
	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var role models.Role
		if err := tx.First(&role, roleID).Error; err != nil {
			return err
		}

		var permissions []models.Permission
		if len(permissionIDs) > 0 {
			if err := tx.Where("id IN ?", permissionIDs).Find(&permissions).Error; err != nil {
				return err
			}
			if len(permissions) != len(permissionIDs) {
				return gorm.ErrRecordNotFound
			}
		}

		if err := tx.Where("role_id = ?", roleID).Delete(&models.RolePermission{}).Error; err != nil {
			return err
		}

		assignments := make([]models.RolePermission, 0, len(permissionIDs))
		for _, permissionID := range permissionIDs {
			assignments = append(assignments, models.RolePermission{
				RoleID:       roleID,
				PermissionID: permissionID,
			})
		}

		if len(assignments) == 0 {
			return nil
		}
		return tx.Create(&assignments).Error
	})
	if err != nil {
		return nil, normalizeError(err)
	}

	return r.GetRole(ctx, roleID)
}

func (r *GormRepository) ListPermissions(ctx context.Context) ([]models.Permission, error) {
	var permissions []models.Permission
	err := r.db.WithContext(ctx).Order("resource ASC, action ASC, id ASC").Find(&permissions).Error
	return permissions, normalizeError(err)
}

func (r *GormRepository) CreatePermission(ctx context.Context, permission *models.Permission) error {
	return normalizeError(r.db.WithContext(ctx).Create(permission).Error)
}

func (r *GormRepository) GetPermission(ctx context.Context, id uint64) (*models.Permission, error) {
	var permission models.Permission
	err := r.db.WithContext(ctx).First(&permission, id).Error
	if err != nil {
		return nil, normalizeError(err)
	}
	return &permission, nil
}

func (r *GormRepository) UpdatePermission(ctx context.Context, id uint64, updates map[string]any) (*models.Permission, error) {
	db := r.db.WithContext(ctx)

	var permission models.Permission
	if err := db.First(&permission, id).Error; err != nil {
		return nil, normalizeError(err)
	}

	if err := db.Model(&permission).Updates(updates).Error; err != nil {
		return nil, normalizeError(err)
	}

	return r.GetPermission(ctx, id)
}

func (r *GormRepository) DeletePermission(ctx context.Context, id uint64) error {
	result := r.db.WithContext(ctx).Delete(&models.Permission{}, id)
	if result.Error != nil {
		return normalizeError(result.Error)
	}
	if result.RowsAffected == 0 {
		return ErrNotFound
	}
	return nil
}

func (r *GormRepository) SyncUserRoles(ctx context.Context, userID uint64, roleIDs []uint64) (*models.User, error) {
	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var user models.User
		if err := tx.First(&user, userID).Error; err != nil {
			return err
		}

		var roles []models.Role
		if len(roleIDs) > 0 {
			if err := tx.Where("id IN ?", roleIDs).Find(&roles).Error; err != nil {
				return err
			}
			if len(roles) != len(roleIDs) {
				return gorm.ErrRecordNotFound
			}
		}

		if err := tx.Where("user_id = ?", userID).Delete(&models.UserRole{}).Error; err != nil {
			return err
		}

		assignments := make([]models.UserRole, 0, len(roleIDs))
		for _, roleID := range roleIDs {
			assignments = append(assignments, models.UserRole{
				UserID: userID,
				RoleID: roleID,
			})
		}

		if len(assignments) == 0 {
			return nil
		}
		return tx.Create(&assignments).Error
	})
	if err != nil {
		return nil, normalizeError(err)
	}

	var user models.User
	if err := r.db.WithContext(ctx).Preload("Roles").First(&user, userID).Error; err != nil {
		return nil, normalizeError(err)
	}
	return &user, nil
}

func normalizeError(err error) error {
	if err == nil {
		return nil
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrNotFound
	}
	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
		return ErrDuplicate
	}
	return err
}
