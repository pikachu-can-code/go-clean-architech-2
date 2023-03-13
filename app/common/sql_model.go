package common

import "time"

type SQLModel struct {
	ID        uint64     `json:"-" gorm:"column:id;primary_key;autoIncrement:true"`
	FakeId    *UID       `json:"id" gorm:"-"`
	Status    int        `json:"status" gorm:"column:status;default:1;type:tinyint;"`
	CreatedAt *time.Time `json:"created_at,omitempty" gorm:"column:created_at;type:timestamp;autoCreateTime;"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" gorm:"column:updated_at;type:timestamp;autoUpdateTime;"`
}

func (m *SQLModel) GenUID(dbType uint16) {
	uid := NewUID(m.ID, dbType, 1)
	m.FakeId = &uid
}
