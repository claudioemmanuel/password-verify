package tests

import (
	"testing"

	"github.com/claudioemmanue/go-api/pkg/service"
	"github.com/claudioemmanue/go-api/pkg/types"
	"github.com/stretchr/testify/assert"
)

// TestValidatePassword is responsible for testing the password validation
func TestValidatePassword(t *testing.T) {
	t.Run("Should return false verify when password is not valid", testShouldReturnFalseVerifyWhenPasswordIsNotValid)
	t.Run("Should return true verify when password is valid", testShouldReturnTrueVerifyWhenPasswordIsValid)
	t.Run("Should return error when rule is not valid", testShouldReturnErrorWhenRuleIsNotValid)
}

func testShouldReturnFalseVerifyWhenPasswordIsNotValid(t *testing.T) {

	// Create not valid payload to test
	failPayload := createMockPayload(
		"12345",
		[]types.Rule{
			{Rule: "minSize", Value: 8},
			{Rule: "minSpecialChars", Value: 1},
			{Rule: "noRepeted", Value: 0},
		},
	)

	// Call the service to validate the password
	response, err := service.ValidatePassword(failPayload)

	// Assert if the password is not valid
	assert.NotEmpty(t, response.NoMatch)
	assert.False(t, response.Verify)
	assert.Nil(t, err)
}

func testShouldReturnTrueVerifyWhenPasswordIsValid(t *testing.T) {

	// Create a valid payload to test
	successPayload := createMockPayload(
		"12345@abc",
		[]types.Rule{
			{Rule: "minSize", Value: 8},
			{Rule: "minSpecialChars", Value: 1},
			{Rule: "noRepeted", Value: 0},
			{Rule: "minDigit", Value: 1},
		},
	)

	// Call the service to validate the password
	response, err := service.ValidatePassword(successPayload)

	// Assert if the password is valid
	assert.Empty(t, response.NoMatch)
	assert.True(t, response.Verify)
	assert.Nil(t, err)
}

func testShouldReturnErrorWhenRuleIsNotValid(t *testing.T) {

	// Create a valid payload to test
	successPayload := createMockPayload(
		"12345@abc",
		[]types.Rule{
			{Rule: "minSize", Value: 8},
			{Rule: "invalidRule", Value: 1},
		},
	)

	// Call the service to validate the password
	_, err := service.ValidatePassword(successPayload)

	assert.NotNil(t, err)
}

func createMockPayload(password string, rules []types.Rule) types.Payload {

	// Create a mock payload
	var mockPayload types.Payload

	mockPayload.Password = password
	mockPayload.Rules = rules

	return mockPayload
}
