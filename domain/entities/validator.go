package domain

type Validator interface {
	ValidatePassword(input UserInput) (ValidationResponse, error)
}
