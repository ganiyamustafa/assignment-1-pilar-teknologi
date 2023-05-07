package routes

import (
	"github.com/ganiyamustafa/assignment-1-pilar-teknologi/internal/controllers"
	"github.com/ganiyamustafa/assignment-1-pilar-teknologi/internal/services"
	"github.com/ganiyamustafa/assignment-1-pilar-teknologi/utils"
	"github.com/gin-gonic/gin"
)

func ChatRoutes(router *gin.Engine, handler *utils.Handler) {
	chatService := services.ChatService{Handler: handler}
	controller := controllers.ChatController{ChatService: chatService}

	router.GET("/api/chat-rooms", controller.GetChatRooms)
	router.GET("/api/chats/:roomId", controller.GetChatsByRoomId)
	router.GET("/api/chats", controller.GetAllChats)
}
