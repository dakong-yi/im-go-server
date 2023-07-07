package controllers

// import (
// 	"net/http"
// 	"strconv"

// 	"github.com/dakong-yi/im-go-server/services"
// 	"github.com/gin-gonic/gin"
// )

// type ConversationController struct {
// 	convService *services.ConversationService
// }

// func NewConversationController(convService *services.ConversationService) *ConversationController {
// 	return &ConversationController{convService}
// }

// func (c *ConversationController) CreateConversation(ctx *gin.Context) {
// 	name := ctx.PostForm("name")
// 	conversation, err := c.convService.CreateConversation(name)
// 	if err != nil {
// 		// handle error
// 	}
// 	ctx.JSON(http.StatusOK, gin.H{"conversation": conversation})
// }

// func (c *ConversationController) GetConversationByID(ctx *gin.Context) {
// 	id, err := strconv.Atoi(ctx.Param("id"))
// 	if err != nil {
// 		// handle error
// 	}
// 	conversation, err := c.convService.GetConversationByID(uint(id))
// 	if err != nil {
// 		// handle error
// 	}
// 	ctx.JSON(http.StatusOK, gin.H{"conversation": conversation})
// }

// func (c *ConversationController) UpdateConversation(ctx *gin.Context) {
// 	id, err := strconv.Atoi(ctx.Param("id"))
// 	if err != nil {
// 		// handle error
// 	}
// 	conversation, err := c.convService.GetConversationByID(uint(id))
// 	if err != nil {
// 		// handle error
// 	}
// 	name := ctx.PostForm("name")
// 	conversation.Name = name
// 	err = c.convService.UpdateConversation(conversation)
// 	if err != nil {
// 		// handle error
// 	}
// 	ctx.JSON(http.StatusOK, gin.H{"conversation": conversation})
// }

// func (c *ConversationController) DeleteConversation(ctx *gin.Context) {
// 	id, err := strconv.Atoi(ctx.Param("id"))
// 	if err != nil {
// 		// handle error
// 	}
// 	err = c.convService.DeleteConversation(uint(id))
// 	if err != nil {
// 		// handle error
// 	}
// 	ctx.JSON(http.StatusOK, gin.H{})
// }
