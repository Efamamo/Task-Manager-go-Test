// user_controller_test.go
package controllertest

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	controller "github.com/Task-Management-go/Delivery/controllers"
	domain "github.com/Task-Management-go/Domain"
	"github.com/Task-Management-go/Domain/err"
	"github.com/Task-Management-go/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type UserControllerTestSuite struct {
	suite.Suite
	mockUserService *mocks.MockUserService
	userController  *controller.UserController
	router          *gin.Engine
}

func (suite *UserControllerTestSuite) SetupTest() {
	suite.mockUserService = new(mocks.MockUserService)
	suite.userController = &controller.UserController{Service: suite.mockUserService}

	gin.SetMode(gin.TestMode)
	suite.router = gin.Default()
}

func TestUserControllerTestSuite(t *testing.T) {
	suite.Run(t, new(UserControllerTestSuite))
}

func (suite *UserControllerTestSuite) TestSignUp_Success() {
	newUser := domain.User{Username: "testuser", Password: "password"}
	suite.mockUserService.On("SignUp", newUser).Return(&newUser, nil)

	body, _ := json.Marshal(newUser)
	req, err := http.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(body))
	assert.NoError(suite.T(), err)

	rr := httptest.NewRecorder()
	suite.router.POST("/signup", suite.userController.SignUp)
	suite.router.ServeHTTP(rr, req)

	assert.Equal(suite.T(), http.StatusCreated, rr.Code)
	suite.mockUserService.AssertExpectations(suite.T())
}

func (suite *UserControllerTestSuite) TestSignUp_ValidationError() {
	newUser := domain.User{Username: "", Password: ""}
	suite.mockUserService.On("SignUp", newUser).Return(nil, errors.New("Validation Error"))

	body, _ := json.Marshal(newUser)
	req, err := http.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(body))
	assert.NoError(suite.T(), err)

	rr := httptest.NewRecorder()
	suite.router.POST("/signup", suite.userController.SignUp)
	suite.router.ServeHTTP(rr, req)

	assert.Equal(suite.T(), http.StatusBadRequest, rr.Code)
}

func (suite *UserControllerTestSuite) TestLogin_Success() {
	reqUser := domain.User{Username: "testuser", Password: "password"}
	suite.mockUserService.On("Login", reqUser).Return("token", nil)

	body, _ := json.Marshal(reqUser)
	req, err := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(body))
	assert.NoError(suite.T(), err)

	rr := httptest.NewRecorder()
	suite.router.POST("/login", suite.userController.Login)
	suite.router.ServeHTTP(rr, req)

	assert.Equal(suite.T(), http.StatusCreated, rr.Code)
	suite.mockUserService.AssertExpectations(suite.T())
}

func (suite *UserControllerTestSuite) TestLogin_Failure() {
	reqUser := domain.User{Username: "testuser", Password: "password"}
	suite.mockUserService.On("Login", reqUser).Return("", err.NewUnauthorized("INvalid Credentials"))

	body, _ := json.Marshal(reqUser)
	req, err := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(body))
	assert.NoError(suite.T(), err)

	rr := httptest.NewRecorder()
	suite.router.POST("/login", suite.userController.Login)
	suite.router.ServeHTTP(rr, req)

	assert.Equal(suite.T(), http.StatusUnauthorized, rr.Code)
	suite.mockUserService.AssertExpectations(suite.T())
}

func (suite *UserControllerTestSuite) TestPromote_Success() {
	username := "testuser"
	suite.mockUserService.On("Promote", username).Return(true, nil)

	req, err := http.NewRequest(http.MethodGet, "/promote?username="+username, nil)
	assert.NoError(suite.T(), err)

	rr := httptest.NewRecorder()
	suite.router.GET("/promote", suite.userController.Promote)
	suite.router.ServeHTTP(rr, req)

	assert.Equal(suite.T(), http.StatusOK, rr.Code)
	suite.mockUserService.AssertExpectations(suite.T())
}

func (suite *UserControllerTestSuite) TestPromote_Failure() {
	username := "testuser"
	suite.mockUserService.On("Promote", username).Return(false, err.NewNotFound("User Not Found"))

	req, err := http.NewRequest(http.MethodGet, "/promote?username="+username, nil)
	assert.NoError(suite.T(), err)

	rr := httptest.NewRecorder()
	suite.router.GET("/promote", suite.userController.Promote)
	suite.router.ServeHTTP(rr, req)

	assert.Equal(suite.T(), http.StatusNotFound, rr.Code)
	suite.mockUserService.AssertExpectations(suite.T())
}
