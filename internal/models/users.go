package models

import (
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primary_key"`
	Name      string    `json:"name" binding:"required" gorm:"type:varchar(100);not null"`
	Email     string    `json:"email" binding:"required" gorm:"type:varchar(100);not null;unique"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	newUuid, err := uuid.NewV4()
	if err != nil {
		return err
	}

	u.ID = newUuid
	return nil
}
