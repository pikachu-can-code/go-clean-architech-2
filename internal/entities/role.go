package entities

import (
	"time"

	"github.com/nguyen-phi-khanh-monorevo/go-clean-architech-2/common"
)

const roleEntityName = "roles"
const userRoleEntityName = "user_roles"

type Role struct {
	common.SQLModel `json:",inline"`
	Name            string       `json:"name" gorm:"column:name;"`
	Permissions     []Permission `json:"-" gorm:"preload:false;many2many:permission_roles;foreignKey:ID;joinForeignKey:RoleID;References:ID;JoinReferences:PermissionID;"`
}

type UserRoleCreate struct {
	UserId    uint64     `json:"user_id" gorm:"column:user_id;primary_key;"`
	RoleId    uint64     `json:"role_id" gorm:"column:role_id;primary_key;"`
	Status    int        `json:"status" gorm:"column:status;default:1;type:tinyint;"`
	CreatedAt *time.Time `json:"created_at,omitempty" gorm:"column:created_at;type:timestamp;autoCreateTime;"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" gorm:"column:updated_at;type:timestamp;autoUpdateTime;"`
}

func (Role) TableName() string {
	return roleEntityName
}

func (UserRoleCreate) TableName() string {
	return userRoleEntityName
}
