package controllers

import (
	"github.com/gin-gonic/gin/binding"
)

type Context interface {
	ShouldBindBodyWith(obj interface{}, bb binding.BindingBody) (err error)
	ShouldBindWith(obj interface{}, b binding.Binding) error
	JSON(code int, obj interface{})
	Param(key string) string
	AbortWithStatusJSON(code int, jsonObj interface{})
	GetHeader(key string) string
	Header(key, value string)
	Set(key string, value interface{})
	Get(key string) (value interface{}, exists bool)
	MustGet(key string) interface{}
}

type AppContextHandlerFunc func(Context)
