package models

import (
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Chat struct {
	ID         uuid.UUID `json:"id" gorm:"type:uuid;primary_key"`
	ChatRoomID uuid.UUID `json:"chat_room_id" gorm:"type:uuid;foreign_key"`
	Message    string    `json:"message" gorm:"type:text"`
	ChatRoom   *ChatRoom `json:"chat_room,omitempty" gorm:"foreignKey:ChatRoomID"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (u *Chat) BeforeCreate(tx *gorm.DB) error {
	newUuid, err := uuid.NewV4()
	if err != nil {
		return err
	}

	u.ID = newUuid
	return nil
}

type GetChatsRequest struct {
	FilterRequest
	PaginateRequest
	SenderID   string `form:"sender_id" binding:"omitempty,ascii"`
	ReceiverID string `form:"receiver_id" binding:"omitempty,ascii"`
}
