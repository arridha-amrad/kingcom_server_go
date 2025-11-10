package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Role string

const (
	RoleUser  Role = "user"
	RoleAdmin Role = "admin"
)

type Provider string

const (
	ProviderCredentials Provider = "credentials"
)

type User struct {
	ID         uuid.UUID      `gorm:"type:uuid;primaryKey" json:"id"`
	CreatedAt  time.Time      `json:"createdAt"`
	UpdatedAt  time.Time      `json:"-"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
	Username   string         `gorm:"not null;unique" json:"username"`
	Name       string         `gorm:"not null" json:"name"`
	Email      string         `gorm:"not null;unique" json:"email"`
	Password   string         `gorm:"not null" json:"-"`
	Provider   Provider       `gorm:"not null" json:"-"`
	Role       Role           `gorm:"not null" json:"role"`
	JwtVersion string         `gorm:"column:jwt_version;not null" json:"-"`
	IsVerified bool           `gorm:"column:is_verified;not null" json:"-"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	return
}
