package infrastructure_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	infrastructure "github.com/Task-Management-go/Infrastructure"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type AuthMiddlewareTestSuite struct {
	suite.Suite
	TokenService infrastructure.Token
}

func (suite *AuthMiddlewareTestSuite) SetupSuite() {
	// Setup any needed environment variables or configurations
	gin.SetMode(gin.TestMode)
	suite.TokenService = infrastructure.Token{}
}

func (suite *AuthMiddlewareTestSuite) TestAuthMiddleware_MissingAuthorizationHeader() {
	// Arrange
	r := gin.Default()
	r.Use(infrastructure.AuthMiddleware(false))
	r.GET("/test", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, gin.H{"status": "success"})
	})

	req, _ := http.NewRequest(http.MethodGet, "/test", nil)
	resp := httptest.NewRecorder()

	// Act
	r.ServeHTTP(resp, req)

	// Assert
	assert.Equal(suite.T(), http.StatusUnauthorized, resp.Code)
	assert.JSONEq(suite.T(), `{"message": "unauthorized"}`, resp.Body.String())
}

func (suite *AuthMiddlewareTestSuite) TestAuthMiddleware_InvalidTokenFormat() {
	// Arrange
	r := gin.Default()
	r.Use(infrastructure.AuthMiddleware(false))
	r.GET("/test", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, gin.H{"status": "success"})
	})

	req, _ := http.NewRequest(http.MethodGet, "/test", nil)
	req.Header.Add("Authorization", "InvalidTokenFormat")
	resp := httptest.NewRecorder()

	// Act
	r.ServeHTTP(resp, req)

	// Print the response body for debugging
	fmt.Println("Response Body:", resp.Body.String())

	// Assert
	assert.Equal(suite.T(), http.StatusUnauthorized, resp.Code)
	assert.JSONEq(suite.T(), `{"message": "unauthorized"}`, resp.Body.String())
}

func (suite *AuthMiddlewareTestSuite) TestAuthMiddleware_InvalidToken() {
	// Arrange
	r := gin.Default()
	r.Use(infrastructure.AuthMiddleware(false))
	r.GET("/test", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, gin.H{"status": "success"})
	})

	// Generate an invalid token
	invalidToken := "invalidtoken"
	req, _ := http.NewRequest(http.MethodGet, "/test", nil)
	req.Header.Add("Authorization", "Bearer "+invalidToken)
	resp := httptest.NewRecorder()

	// Act
	r.ServeHTTP(resp, req)

	// Assert
	assert.Equal(suite.T(), http.StatusUnauthorized, resp.Code)
	assert.JSONEq(suite.T(), `{"error": "unauthorized"}`, resp.Body.String())
}

func (suite *AuthMiddlewareTestSuite) TestAuthMiddleware_AuthorizedUser() {
	// Arrange
	r := gin.Default()
	r.Use(infrastructure.AuthMiddleware(false))
	r.GET("/test", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, gin.H{"status": "success"})
	})

	token, err := suite.TokenService.GenerateToken("testuser", false)
	suite.Require().NoError(err)

	req, _ := http.NewRequest(http.MethodGet, "/test", nil)
	req.Header.Add("Authorization", "Bearer "+token)
	resp := httptest.NewRecorder()

	// Act
	r.ServeHTTP(resp, req)

	// Assert
	assert.Equal(suite.T(), http.StatusOK, resp.Code)
	assert.JSONEq(suite.T(), `{"status": "success"}`, resp.Body.String())
}

func (suite *AuthMiddlewareTestSuite) TestAuthMiddleware_AuthorizedAdmin() {
	// Arrange
	r := gin.Default()
	r.Use(infrastructure.AuthMiddleware(true))
	r.GET("/test", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, gin.H{"status": "success"})
	})

	token, err := suite.TokenService.GenerateToken("testuser", true)
	suite.Require().NoError(err)

	req, _ := http.NewRequest(http.MethodGet, "/test", nil)
	req.Header.Add("Authorization", "Bearer "+token)
	resp := httptest.NewRecorder()

	// Act
	r.ServeHTTP(resp, req)

	// Assert
	assert.Equal(suite.T(), http.StatusOK, resp.Code)
	assert.JSONEq(suite.T(), `{"status": "success"}`, resp.Body.String())
}

func (suite *AuthMiddlewareTestSuite) TestAuthMiddleware_UnauthorizedAdmin() {
	// Arrange
	r := gin.Default()
	r.Use(infrastructure.AuthMiddleware(true))
	r.GET("/test", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, gin.H{"status": "success"})
	})

	token, err := suite.TokenService.GenerateToken("testuser", false)
	suite.Require().NoError(err)

	req, _ := http.NewRequest(http.MethodGet, "/test", nil)
	req.Header.Add("Authorization", "Bearer "+token)
	resp := httptest.NewRecorder()

	// Act
	r.ServeHTTP(resp, req)

	// Assert
	assert.Equal(suite.T(), http.StatusUnauthorized, resp.Code)
	assert.JSONEq(suite.T(), `{"error": "Forbidden"}`, resp.Body.String())
}

func TestAuthMiddlewareTestSuite(t *testing.T) {
	suite.Run(t, new(AuthMiddlewareTestSuite))
}
