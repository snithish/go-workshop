package controllers

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

type CreatePetRequest struct {
	Id        int      `json:"id" binding:"required"`
	Name      string   `json:"name" binding:"required"`
	PhotoUrls []string `json:"photoUrls" binding:"dive,required"`
	Status    string   `json:"status" binding:"required"`
}

func badResponseValidator(recorder *httptest.ResponseRecorder, t *testing.T) {
	assert.Equal(t, 400, recorder.Code)
}

func Test_petController_CreatePet(t *testing.T) {
	type args struct {
		requestBody       string
		responseValidator func(recorder *httptest.ResponseRecorder, t *testing.T)
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "fail when ID is not present",
			args: args{
				requestBody:       `{"name":"Tom","status":"available","photoUrls":[]}`,
				responseValidator: badResponseValidator,
			},
		},
		{
			name: "fail when ID is an integer",
			args: args{
				requestBody: `{"id":"123","name":"Tom","status":"available","photoUrls":[]}`,
				responseValidator: func(recorder *httptest.ResponseRecorder, t *testing.T) {
					assert.Equal(t, 400, recorder.Code)
				},
			},
		},
	}
	for _, tt := range tests {
		recorder := httptest.NewRecorder()
		context, _ := gin.CreateTestContext(recorder)
		context.Request, _ = http.NewRequest("POST", "/", bytes.NewBufferString(tt.args.requestBody))
		t.Run(tt.name, func(t *testing.T) {
			controller := NewPetController()
			controller.CreatePet(context)
			tt.args.responseValidator(recorder, t)
		})
	}
}
