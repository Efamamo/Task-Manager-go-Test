package infrastructure_test

import (
	"testing"

	"github.com/Task-Management-go/Domain/err"
	infrastructure "github.com/Task-Management-go/Infrastructure"
	"github.com/stretchr/testify/suite"
)

type PassServiceTestSuite struct {
	suite.Suite
	PassService infrastructure.Pass
}

func (suite *PassServiceTestSuite) TestHashPassword_Success() {
	// Arrange
	password := "mySecurePassword"

	// Act
	hashedPassword, err := suite.PassService.HashPassword(password)

	// Assert
	suite.Require().NoError(err)
	suite.Require().NotEmpty(hashedPassword)
}

func (suite *PassServiceTestSuite) TestComparePassword_Success() {
	// Arrange
	password := "mySecurePassword"
	hashedPassword, err := suite.PassService.HashPassword(password)
	suite.Require().NoError(err)

	// Act
	isValid, err := suite.PassService.ComparePassword(hashedPassword, password)

	// Assert
	suite.Require().NoError(err)
	suite.Require().True(isValid)
}

func (suite *PassServiceTestSuite) TestComparePassword_InvalidPassword() {
	// Arrange
	password := "mySecurePassword"
	invalidPassword := "wrongPassword"
	hashedPassword, e := suite.PassService.HashPassword(password)
	suite.Require().NoError(e)

	// Act
	isValid, e := suite.PassService.ComparePassword(hashedPassword, invalidPassword)

	// Assert
	suite.Require().Error(e)
	suite.Require().False(isValid)
	suite.Require().IsType(err.NewUnauthorized("Invalid Credentials"), e)
}

func TestPassServiceTestSuite(t *testing.T) {
	suite.Run(t, new(PassServiceTestSuite))
}
