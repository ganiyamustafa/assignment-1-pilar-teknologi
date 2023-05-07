package controllers

import (
	"net/http"

	"github.com/ganiyamustafa/assignment-1-pilar-teknologi/internal/models"
	"github.com/ganiyamustafa/assignment-1-pilar-teknologi/internal/services"
	"github.com/gin-gonic/gin"
)

type ChatController struct {
	ChatService services.ChatService
}

func (c ChatController) GetChatRooms(ctx *gin.Context) {
	var query models.GetChatRoomsRequest
	ctx.Bind(&query)

	if chatRooms, meta, err := c.ChatService.GetChatRooms(&query); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"data": gin.H{"chat_rooms": chatRooms, "meta": meta}})
	}
}

func (c ChatController) GetChatsByRoomId(ctx *gin.Context) {
	var query models.GetChatsRequest
	param := ctx.Param("roomId")
	ctx.Bind(&query)

	if chats, meta, err := c.ChatService.GetChatsByRoomId(&query, param); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"data": gin.H{"chats": chats, "meta": meta}})
	}
}

func (c ChatController) GetAllChats(ctx *gin.Context) {
	var query models.GetChatsRequest
	ctx.Bind(&query)

	if chats, meta, err := c.ChatService.GetAllChats(&query); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"data": gin.H{"chats": chats, "meta": meta}})
	}
}
