package controllers

import (
	"4_gin_web_service/generated_mocks"
	"4_gin_web_service/models"
	"bytes"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type PetControllerTestSuite struct {
	suite.Suite
	mockCtrl       *gomock.Controller
	recorder       *httptest.ResponseRecorder
	context        *gin.Context
	mockPetService *generated_mocks.MockPetService
	petController  PetController
}

func TestPetControllerTestSuite(t *testing.T) {
	suite.Run(t, new(PetControllerTestSuite))
}

func (suite *PetControllerTestSuite) SetupTest() {
	suite.mockCtrl = gomock.NewController(suite.T())
	suite.recorder = httptest.NewRecorder()
	suite.context, _ = gin.CreateTestContext(suite.recorder)
	suite.mockPetService = generated_mocks.NewMockPetService(suite.mockCtrl)
	suite.petController = NewPetController(suite.mockPetService)
}

func (suite *PetControllerTestSuite) TearDownTest() {
	suite.mockCtrl.Finish()
}

func (suite *PetControllerTestSuite) TestCreatePet_WhenFailureInParsingRequestBody() {
	requestBody := `{"someKey" : "someValue"}`
	suite.context.Request, _ = http.NewRequest("POST", "/", bytes.NewBufferString(requestBody))

	suite.petController.CreatePet(suite.context)

	suite.Require().Equal(http.StatusBadRequest, suite.recorder.Code)
}

func (suite *PetControllerTestSuite) TestCreatePet_WhenFailureInRequestBodyValidation() {
	requestBody := `{"id":123123,"name":"doggie","photoUrls":["string"],"status":"unavailable"}`
	suite.context.Request, _ = http.NewRequest("POST", "/", bytes.NewBufferString(requestBody))

	suite.petController.CreatePet(suite.context)

	suite.Require().Equal(http.StatusMethodNotAllowed, suite.recorder.Code)
}

//func (suite *PetControllerTestSuite) TestCreatePet_WhenFailureInService() {
//	requestBody := `{"id":123123,"name":"doggie","photoUrls":["string"],"status":"available"}`
//	suite.context.Request, _ = http.NewRequest("POST", "/", bytes.NewBufferString(requestBody))
//	pet := models.Pet{
//		Id:        123123,
//		Name:      "doggie",
//		PhotoUrls: []string{"string"},
//		Status:    "available",
//	}
//
//	suite.mockPetService.EXPECT().CreatePet(pet).Return(errors.New("some error")).Times(1)
//
//	suite.petController.CreatePet(suite.context)
//
//	suite.Require().Equal(http.StatusConflict, suite.recorder.Code)
//}

func (suite *PetControllerTestSuite) TestCreatePet() {
	requestBody := `{"id":123123,"name":"doggie","photoUrls":["string"],"status":"available"}`
	suite.context.Request, _ = http.NewRequest("POST", "/", bytes.NewBufferString(requestBody))
	pet := models.Pet{
		Id:        123123,
		Name:      "doggie",
		PhotoUrls: []string{"string"},
		Status:    "available",
	}

	suite.mockPetService.EXPECT().CreatePet(pet).Return(nil).Times(1)

	suite.petController.CreatePet(suite.context)

	suite.Require().Equal(http.StatusCreated, suite.recorder.Code)
}

func (suite *PetControllerTestSuite) TestGetPet_WhenFailureInMissingPetID() {
	suite.context.Params = gin.Params{}

	suite.petController.GetPet(suite.context)

	suite.Require().Equal(http.StatusBadRequest, suite.recorder.Code)
}

func (suite *PetControllerTestSuite) TestGetPet_WhenFailureInParsingPetID() {
	suite.context.Params = gin.Params{gin.Param{
		Key:   "petID",
		Value: "someInvalidPetID",
	}}

	suite.petController.GetPet(suite.context)

	suite.Require().Equal(http.StatusMethodNotAllowed, suite.recorder.Code)
}

func (suite *PetControllerTestSuite) TestGetPet_WhenServiceFails() {
	suite.context.Params = gin.Params{gin.Param{
		Key:   "petID",
		Value: "123456",
	}}
	petID := 123456

	suite.mockPetService.EXPECT().GetPet(petID).Return(models.Pet{}, errors.New("some error")).Times(1)

	suite.petController.GetPet(suite.context)

	suite.Require().Equal(http.StatusNotFound, suite.recorder.Code)
}

func (suite *PetControllerTestSuite) TestGetPet() {
	suite.context.Params = gin.Params{gin.Param{
		Key:   "petID",
		Value: "123456",
	}}
	petID := 123456
	pet := models.Pet{
		Id:        123123,
		Name:      "doggie",
		PhotoUrls: []string{"string"},
		Status:    "available",
	}
	expectedResponse := `{"id":123123,"name":"doggie","photoUrls":["string"],"status":"available"}`

	suite.mockPetService.EXPECT().GetPet(petID).Return(pet, nil).Times(1)

	suite.petController.GetPet(suite.context)

	suite.Require().Equal(http.StatusOK, suite.recorder.Code)
	suite.Require().Equal(expectedResponse, suite.recorder.Body.String())
}
