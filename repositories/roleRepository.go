package repositories

import "github.com/maulanadityaa/bank-merchant-api/models/entity"

type RoleRepository interface {
	GetRoleByID(roleID string) (entity.Role, error)
	GetRoleByName(roleName string) (entity.Role, error)
}
