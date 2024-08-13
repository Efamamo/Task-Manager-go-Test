package infrastructure_test

import (
	"os"
	"testing"

	infrastructure "github.com/Task-Management-go/Infrastructure"
	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type TokenTestSuite struct {
	suite.Suite
	TokenGenerator infrastructure.Token
}

func (suite *TokenTestSuite) SetupSuite() {
	// Set the environment variable for JwtSecret
	os.Setenv("JwtSecret", "testsecret")
}

func (suite *TokenTestSuite) TearDownSuite() {
	// Unset the environment variable after all tests
	os.Unsetenv("JwtSecret")
}

func (suite *TokenTestSuite) SetupTest() {
	suite.TokenGenerator = infrastructure.Token{}
}

func TestTokenTestSuite(t *testing.T) {
	suite.Run(t, new(TokenTestSuite))
}

func (suite *TokenTestSuite) TestToken_ValidateToken() {
	// Generate a valid token
	validToken, e := suite.TokenGenerator.GenerateToken("user", true)
	require.NoError(suite.T(), e)

	// Test valid token
	parsedToken, e := suite.TokenGenerator.ValidateToken(validToken)
	assert.NoError(suite.T(), e)
	assert.True(suite.T(), parsedToken.Valid)

	// Test invalid token
	_, e = suite.TokenGenerator.ValidateToken("invalidtoken")
	assert.Error(suite.T(), e)
}

func (suite *TokenTestSuite) TestToken_ValidateAdmin() {
	// Generate an admin token
	adminToken, err := suite.TokenGenerator.GenerateToken("admin", true)
	require.NoError(suite.T(), err)

	parsedToken, err := suite.TokenGenerator.ValidateToken(adminToken)
	require.NoError(suite.T(), err)

	// Test ValidateAdmin with an admin token
	isAdmin := suite.TokenGenerator.ValidateAdmin(parsedToken)
	assert.True(suite.T(), isAdmin)

	// Generate a non-admin token
	userToken, err := suite.TokenGenerator.GenerateToken("user", false)
	require.NoError(suite.T(), err)

	parsedToken, err = suite.TokenGenerator.ValidateToken(userToken)
	require.NoError(suite.T(), err)

	// Test ValidateAdmin with a non-admin token
	isAdmin = suite.TokenGenerator.ValidateAdmin(parsedToken)
	assert.False(suite.T(), isAdmin)
}

func (suite *TokenTestSuite) TestToken_GenerateToken() {
	// Generate a token
	token, err := suite.TokenGenerator.GenerateToken("user", true)
	require.NoError(suite.T(), err)
	assert.NotEmpty(suite.T(), token)

	// Validate the generated token
	parsedToken, err := suite.TokenGenerator.ValidateToken(token)
	require.NoError(suite.T(), err)
	assert.True(suite.T(), parsedToken.Valid)

	// Check the claims
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	require.True(suite.T(), ok)
	assert.Equal(suite.T(), "user", claims["username"])
	assert.True(suite.T(), claims["isAdmin"].(bool))
}
