package services

import (
	"errors"
	"regexp"

	domain "github.com/claudioemmanue/go-api/domain/entities"
)

type PasswordValidator struct{}

// CheckAllowedRules checks if the rules in the payload are valid
func checkAllowedRules(p domain.UserInput) error {
	possibleRules := []string{"minSize", "minUppercase", "minLowercase", "minDigit", "minSpecialChars", "noRepeted"}

	for _, rule := range p.Rules {
		if !contains(possibleRules, rule.Rule) {
			return errors.New("Invalid rule: " + rule.Rule)
		}
	}
	return nil
}

// Contains checks if a given slice contains the given element
func contains(slice []string, element string) bool {
	for _, item := range slice {
		if item == element {
			return true
		}
	}
	return false
}

// ValidatePassword is responsible for validating the password
func (pv *PasswordValidator) ValidatePassword(input domain.UserInput) (domain.ValidationResponse, error) {
	err := checkAllowedRules(input)
	if err != nil {
		return domain.ValidationResponse{}, err
	}

	noMatch := []string{}

	for _, rule := range input.Rules {
		switch rule.Rule {
		case "minSize":
			if !hasMinSize(input.Password, rule.Value) {
				noMatch = append(noMatch, "minSize")
			}
		case "minUppercase":
			if !hasMinUppercase(input.Password, rule.Value) {
				noMatch = append(noMatch, "minUppercase")
			}
		case "minLowercase":
			if !hasMinLowercase(input.Password, rule.Value) {
				noMatch = append(noMatch, "minLowercase")
			}
		case "minDigit":
			if !hasMinDigit(input.Password, rule.Value) {
				noMatch = append(noMatch, "minDigit")
			}
		case "minSpecialChars":
			if !hasMinSpecialChars(input.Password, rule.Value) {
				noMatch = append(noMatch, "minSpecialChars")
			}
		case "noRepeted":
			if !hasNoRepeted(input.Password) {
				noMatch = append(noMatch, "noRepeted")
			}
		}
	}

	return domain.ValidationResponse{
		Verify:  len(noMatch) == 0,
		NoMatch: noMatch,
	}, nil
}

func hasMinSize(password string, minSize int) bool {
	return len(password) >= minSize
}

func hasMinUppercase(password string, minUppercase int) bool {
	uppercaseCount := 0
	for _, r := range password {
		if r >= 'A' && r <= 'Z' {
			uppercaseCount++
		}
	}
	return uppercaseCount >= minUppercase
}

func hasMinLowercase(password string, minLowercase int) bool {
	lowercaseCount := 0
	for _, r := range password {
		if r >= 'a' && r <= 'z' {
			lowercaseCount++
		}
	}
	return lowercaseCount >= minLowercase
}

func hasMinDigit(password string, minDigit int) bool {
	digitCount := 0
	for _, r := range password {
		if r >= '0' && r <= '9' {
			digitCount++
		}
	}
	return digitCount >= minDigit
}

func hasMinSpecialChars(password string, minSpecialChars int) bool {
	specialCharsCount := 0
	for _, r := range password {
		if r >= '!' && r <= '/' || r >= ':' && r <= '@' || r >= '[' && r <= '`' || r >= '{' && r <= '~' {
			specialCharsCount++
		}
	}
	return specialCharsCount >= minSpecialChars
}

func hasNoRepeted(password string) bool {
	re := regexp.MustCompile(`(.)\\1`)
	return !re.MatchString(password)
}

