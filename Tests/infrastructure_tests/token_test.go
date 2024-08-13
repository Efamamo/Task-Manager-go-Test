package infrastructure

import (
	"os"
	"testing"

	"github.com/Task-Management-go/Domain/err"
	infrastructure "github.com/Task-Management-go/Infrastructure"
	"github.com/stretchr/testify/suite"
)

type TokenServiceTestSuite struct {
	suite.Suite
	TokenService infrastructure.Token
}

func (suite *TokenServiceTestSuite) SetupSuite() {
	// Setup environment variables or other configurations if needed
	os.Setenv("JwtSecret", "testsecret")
}

func (suite *TokenServiceTestSuite) TearDownSuite() {
	// Cleanup environment variables or other configurations if needed
	os.Unsetenv("JwtSecret")
}

func (suite *TokenServiceTestSuite) TestValidateToken_Success() {
	// Arrange
	username := "testuser"
	isAdmin := true
	tokenString, err := suite.TokenService.GenerateToken(username, isAdmin)
	suite.Require().NoError(err)

	// Act
	validToken, err := suite.TokenService.ValidateToken(tokenString)

	// Assert
	suite.Require().NoError(err)
	suite.Require().NotNil(validToken)
}

func (suite *TokenServiceTestSuite) TestValidateToken_InvalidToken() {
	// Arrange
	invalidToken := "invalidtoken"

	// Act
	_, e := suite.TokenService.ValidateToken(invalidToken)

	// Assert
	suite.Require().Error(e)
	suite.Require().IsType(err.NewUnauthorized("unauthorized"), e)
}

func (suite *TokenServiceTestSuite) TestValidateToken_UnexpectedSigningMethod() {
	// Arrange
	tokenString, e := suite.TokenService.GenerateToken("testuser", true)
	suite.Require().NoError(e)

	// Modify the token to create an invalid token
	modifiedToken := tokenString + "extra"

	// Act
	_, e = suite.TokenService.ValidateToken(modifiedToken)

	// Assert
	suite.Require().Error(e)
	suite.Require().IsType(err.NewUnauthorized("unauthorized"), e)
}

func (suite *TokenServiceTestSuite) TestValidateAdmin_Success() {
	// Arrange
	tokenString, err := suite.TokenService.GenerateToken("testuser", true)
	suite.Require().NoError(err)

	// Act
	parsedToken, err := suite.TokenService.ValidateToken(tokenString)
	suite.Require().NoError(err)

	// Validate admin
	isAdmin := suite.TokenService.ValidateAdmin(parsedToken)

	// Assert
	suite.Require().True(isAdmin)
}

func (suite *TokenServiceTestSuite) TestValidateAdmin_NotAdmin() {
	// Arrange
	tokenString, err := suite.TokenService.GenerateToken("testuser", false)
	suite.Require().NoError(err)

	// Act
	parsedToken, err := suite.TokenService.ValidateToken(tokenString)
	suite.Require().NoError(err)

	// Validate admin
	isAdmin := suite.TokenService.ValidateAdmin(parsedToken)

	// Assert
	suite.Require().False(isAdmin)
}

func TestTokenServiceTestSuite(t *testing.T) {
	suite.Run(t, new(TokenServiceTestSuite))
}
