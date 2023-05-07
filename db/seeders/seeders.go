package seeders

import (
	"fmt"

	"github.com/ganiyamustafa/assignment-1-pilar-teknologi/internal/models"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

func Wipe(db *gorm.DB) {
	// Delete table that need seed including any of their related table
	db.Session(&gorm.Session{CreateBatchSize: 100, AllowGlobalUpdate: true})

	if db.Where("true").Delete(&models.User{}).Error != nil {
		log.Fatal(db.Error)
	}
}

func Seed(db *gorm.DB) {
	db.Session(&gorm.Session{CreateBatchSize: 100, AllowGlobalUpdate: true})
	var userDatas []models.User
	var chatRoomDatas []models.ChatRoom
	var chatDatas []models.Chat
	// seed admin, try not to wipe user data if possible because that is valuable

	for i := 0; i < 10; i++ {
		userData := models.User{
			Name:  fmt.Sprintf("User-%v", i),
			Email: fmt.Sprintf("Email%v@tuturu.com", i),
		}
		userDatas = append(userDatas, userData)
	}

	if db.Create(&userDatas).Error != nil {
		log.Fatal(db.Error)
	}

	for i := 0; i < 9; i++ {
		chatRoomData := models.ChatRoom{
			SenderID:   userDatas[i].ID,
			ReceiverID: userDatas[i+1].ID,
		}
		chatRoomDatas = append(chatRoomDatas, chatRoomData)
	}

	if db.Create(&chatRoomDatas).Error != nil {
		log.Fatal(db.Error)
	}

	for i, chatRoom := range chatRoomDatas {
		chatData := models.Chat{
			ChatRoomID: chatRoom.ID,
			Message:    fmt.Sprintf("Hello, %v, its %v", i, chatRoom.ID),
		}

		chatDatas = append(chatDatas, chatData)
	}

	if db.Create(&chatDatas).Error != nil {
		log.Fatal(db.Error)
	}
}
