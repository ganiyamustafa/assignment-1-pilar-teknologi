package services

import (
	"math"

	"github.com/ganiyamustafa/assignment-1-pilar-teknologi/internal/models"
	"github.com/ganiyamustafa/assignment-1-pilar-teknologi/utils"
	"gorm.io/gorm"
)

type ChatService struct {
	Handler *utils.Handler
}

func (s ChatService) GetChatRooms(query *models.GetChatRoomsRequest) (*[]models.ChatRoom, *models.MetaResponse, error) {
	var chatRooms []models.ChatRoom
	var meta models.MetaResponse

	offset := ((query.Page - 1) * query.Limit)

	q := s.Handler.SQLite.Table("chat_rooms")
	if q.Error != nil {
		return nil, nil, q.Error
	}

	var count int64
	if err := q.Count(&count).Error; err != nil {
		return nil, nil, err
	}

	if query.Limit > 0 {
		q = q.Limit(int(query.Limit)).Offset(int(offset))
	}
	if query.OrderBy != "" && query.Sort != "" {
		q = q.Order(query.OrderBy + " " + query.Sort)
	}

	q = q.Find(&chatRooms)
	if q.Error != nil {
		return nil, nil, q.Error
	}

	meta.Page = query.Page
	meta.Total = count
	if query.Limit <= 0 {
		meta.LastPage = 1
		meta.Limit = nil
	} else if meta.LastPage = int64(math.Ceil(float64(count) / float64(query.Limit))); meta.LastPage == 0 {
		meta.LastPage = 1
		meta.Limit = &query.Limit
	}

	return &chatRooms, &meta, nil
}

func (s ChatService) GetChatsByRoomId(query *models.GetChatsRequest, param string) (*[]models.Chat, *models.MetaResponse, error) {
	var chats []models.Chat
	var meta models.MetaResponse

	offset := ((query.Page - 1) * query.Limit)

	q := s.Handler.SQLite.Table("chats").Where("chat_room_id = ?", param)
	if q.Error != nil {
		return nil, nil, q.Error
	}

	var count int64
	if err := q.Count(&count).Error; err != nil {
		return nil, nil, err
	}

	if query.Limit > 0 {
		q = q.Limit(int(query.Limit)).Offset(int(offset))
	}
	if query.OrderBy != "" && query.Sort != "" {
		q = q.Order(query.OrderBy + " " + query.Sort)
	}

	q = q.Find(&chats)
	if q.Error != nil {
		return nil, nil, q.Error
	}

	meta.Page = query.Page
	meta.Total = count
	if query.Limit <= 0 {
		meta.LastPage = 1
		meta.Limit = nil
	} else if meta.LastPage = int64(math.Ceil(float64(count) / float64(query.Limit))); meta.LastPage == 0 {
		meta.LastPage = 1
		meta.Limit = &query.Limit
	}

	return &chats, &meta, nil
}

func (s ChatService) GetAllChats(query *models.GetChatsRequest) (*[]models.Chat, *models.MetaResponse, error) {
	var chats []models.Chat
	var meta models.MetaResponse

	offset := ((query.Page - 1) * query.Limit)

	q := s.Handler.SQLite.Table("chats").Joins("LEFT JOIN chat_rooms ON chats.chat_room_id = chat_rooms.id")
	if q.Error != nil {
		return nil, nil, q.Error
	}

	if query.SenderID != "" {
		q = q.Where("chat_rooms.sender_id = ?", query.SenderID)
	}

	if query.ReceiverID != "" {
		q = q.Where("chat_rooms.sender_id = ?", query.ReceiverID)
	}

	var count int64
	if err := q.Count(&count).Error; err != nil {
		return nil, nil, err
	}

	if query.Limit > 0 {
		q = q.Limit(int(query.Limit)).Offset(int(offset))
	}
	if query.OrderBy != "" && query.Sort != "" {
		q = q.Order("chats." + query.OrderBy + " " + query.Sort)
	}

	q.Preload("ChatRoom", func(db *gorm.DB) *gorm.DB {
		return db.Select(`id, sender_id, receiver_id`)
	})

	q.Preload("ChatRoom.Sender")
	q.Preload("ChatRoom.Receiver")

	q = q.Find(&chats)
	if q.Error != nil {
		return nil, nil, q.Error
	}

	meta.Page = query.Page
	meta.Total = count
	if query.Limit <= 0 {
		meta.LastPage = 1
		meta.Limit = nil
	} else if meta.LastPage = int64(math.Ceil(float64(count) / float64(query.Limit))); meta.LastPage == 0 {
		meta.LastPage = 1
		meta.Limit = &query.Limit
	}

	return &chats, &meta, nil
}
