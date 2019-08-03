package controllers

import (
	"4_gin_web_service/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	Message      = "message"
	BadRequest   = "Bad Request"
	InvalidInput = "Invalid Input"
	OK           = "OK"
	Created      = "Created"
	Conflict     = "Conflict"
	NotFound     = "Not Found"
)

func SendBadRequest(ctx *gin.Context) {
	SendBadRequestWithMessage(ctx, BadRequest)
}

func SendInvalidInput(ctx *gin.Context) {
	SendInvalidInputWithMessage(ctx, InvalidInput)
}

func SendBadRequestWithMessage(ctx *gin.Context, message string) {
	responseMap := models.GenericMsgResp{
		Message: message,
	}
	SendMessageWithStatus(ctx, http.StatusBadRequest, responseMap)
}

func SendRequestOK(ctx *gin.Context) {
	responseMap := models.GenericMsgResp{
		Message: OK,
	}
	SendMessageWithStatus(ctx, http.StatusOK, responseMap)
}

func SendConflict(ctx *gin.Context) {
	responseMap := models.GenericMsgResp{
		Message: Conflict,
	}
	SendMessageWithStatus(ctx, http.StatusConflict, responseMap)
}

func SendNotFound(ctx *gin.Context) {
	responseMap := models.GenericMsgResp{
		Message: NotFound,
	}
	SendMessageWithStatus(ctx, http.StatusNotFound, responseMap)
}

func SendRequestCreated(ctx *gin.Context) {
	responseMap := models.GenericMsgResp{
		Message: Created,
	}
	SendMessageWithStatus(ctx, http.StatusCreated, responseMap)
}

func SendInvalidInputWithMessage(ctx *gin.Context, message string) {
	responseMap := models.GenericMsgResp{
		Message: message,
	}
	SendMessageWithStatus(ctx, http.StatusMethodNotAllowed, responseMap)
}

func SendMessageWithStatus(ctx *gin.Context, httpStatusCode int, message interface{}) {
	ctx.JSON(httpStatusCode,
		message,
	)
}
