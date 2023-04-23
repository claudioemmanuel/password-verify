package tests

import (
	"testing"

	"github.com/claudioemmanue/go-api/application/services"
	domain "github.com/claudioemmanue/go-api/domain/entities"
	"github.com/stretchr/testify/assert"
)

func TestValidatePassword(t *testing.T) {
	t.Run("Should return false verify when password is not valid", testShouldReturnFalseVerifyWhenPasswordIsNotValid)
	t.Run("Should return true verify when password is valid", testShouldReturnTrueVerifyWhenPasswordIsValid)
	t.Run("Should return error when rule is not valid", testShouldReturnErrorWhenRuleIsNotValid)
}

func testShouldReturnFalseVerifyWhenPasswordIsNotValid(t *testing.T) {
	failPayload := createMockPayload(
		"12345",
		[]domain.Rule{
			{Rule: "minSize", Value: 8},
			{Rule: "minSpecialChars", Value: 1},
			{Rule: "noRepeted", Value: 0},
		},
	)

	validator := &services.PasswordValidator{}
	response, err := validator.ValidatePassword(failPayload)

	assert.NotEmpty(t, response.NoMatch)
	assert.False(t, response.Verify)
	assert.Nil(t, err)
}

func testShouldReturnTrueVerifyWhenPasswordIsValid(t *testing.T) {
	successPayload := createMockPayload(
		"12345@abc",
		[]domain.Rule{
			{Rule: "minSize", Value: 8},
			{Rule: "minSpecialChars", Value: 1},
			{Rule: "noRepeted", Value: 0},
			{Rule: "minDigit", Value: 1},
		},
	)

	validator := &services.PasswordValidator{}
	response, err := validator.ValidatePassword(successPayload)

	assert.Empty(t, response.NoMatch)
	assert.True(t, response.Verify)
	assert.Nil(t, err)
}

func testShouldReturnErrorWhenRuleIsNotValid(t *testing.T) {
	successPayload := createMockPayload(
		"12345@abc",
		[]domain.Rule{
			{Rule: "minSize", Value: 8},
			{Rule: "invalidRule", Value: 1},
		},
	)

	validator := &services.PasswordValidator{}
	_, err := validator.ValidatePassword(successPayload)

	assert.NotNil(t, err)
}

func createMockPayload(password string, rules []domain.Rule) domain.UserInput {
	return domain.UserInput{
		Password: password,
		Rules:    rules,
	}
}
