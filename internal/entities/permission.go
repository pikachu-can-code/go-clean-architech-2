package entities

import (
	"time"

	"github.com/nguyen-phi-khanh-monorevo/go-clean-architech-2/common"
)

const permissionEntityName = "permissions"
const permissionRoleEntityName = "permission_roles"

type Permission struct {
	common.SQLModel `json:",inline"`
	Name            string `json:"name" gorm:"column:name;"`
}

type PermissionRole struct {
	RoleId       uint64     `json:"role_id" gorm:"column:role_id;primary_key;"`
	PermissionId uint64     `json:"permission_id" gorm:"column:permission_id;primary_key;"`
	Status       int        `json:"status" gorm:"column:status;default:1;type:tinyint;"`
	CreatedAt    *time.Time `json:"created_at,omitempty" gorm:"column:created_at;type:timestamp;autoCreateTime;"`
	UpdatedAt    *time.Time `json:"updated_at,omitempty" gorm:"column:updated_at;type:timestamp;autoUpdateTime;"`
}

func (Permission) TableName() string {
	return permissionEntityName
}

func (PermissionRole) TableName() string {
	return permissionRoleEntityName
}
