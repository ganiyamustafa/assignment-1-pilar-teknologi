package models

import (
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type ChatRoom struct {
	ID         uuid.UUID `json:"id" gorm:"type:uuid;primary_key"`
	SenderID   uuid.UUID `json:"sender_id" gorm:"type:uuid;foreign_key"`
	ReceiverID uuid.UUID `json:"receiver_id" gorm:"type:uuid;foreign_key"`
	Sender     *User     `json:"sender,omitempty" gorm:"foreignKey:SenderID"`
	Receiver   *User     `json:"receiver,omitempty" gorm:"foreignKey:ReceiverID"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (u *ChatRoom) BeforeCreate(tx *gorm.DB) error {
	newUuid, err := uuid.NewV4()
	if err != nil {
		return err
	}

	u.ID = newUuid
	return nil
}

type GetChatRoomsRequest struct {
	FilterRequest
	PaginateRequest
}
