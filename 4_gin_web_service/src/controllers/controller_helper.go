package controllers

import (
	"4_gin_web_service/src/models"
	"net/http"
)

const (
	Message    = "message"
	BadRequest = "Bad Request"
	InvalidInput = "Invalid Input"
	OK         = "OK"
)

func SendBadRequest(ctx Context) {
	SendBadRequestWithMessage(ctx, BadRequest)
}

func SendInvalidInput(ctx Context) {
	SendInvalidInputWithMessage(ctx, InvalidInput)
}

func SendBadRequestWithMessage(ctx Context, message string) {
	responseMap := models.GenericMsgResp{
		Message: message,
	}
	SendMessageWithStatus(ctx, http.StatusBadRequest, responseMap)
}

func SendRequestOK(ctx Context) {
	responseMap := models.GenericMsgResp{
		Message: "OK",
	}
	SendMessageWithStatus(ctx, http.StatusOK, responseMap)
}

func SendInvalidInputWithMessage(ctx Context, message string) {
	responseMap := models.GenericMsgResp{
		Message: message,
	}
	SendMessageWithStatus(ctx, http.StatusMethodNotAllowed, responseMap)
}

func SendMessageWithStatus(ctx Context, httpStatusCode int, message models.GenericMsgResp) {
	ctx.JSON(httpStatusCode,
		message,
	)
}
