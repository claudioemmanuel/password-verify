package service

import (
	"fmt"
	"regexp"

	"github.com/claudioemmanue/go-api/pkg/types"
)

// List of possible rules to validate the password
var possibleRules = []string{"minSize", "minUppercase", "minLowercase", "minDigit", "minSpecialChars", "noRepeted"}

// checkAllowedRules is responsible for checking if the rules in the payload are valid
func checkAllowedRules(p types.Payload) error {
	for _, rule := range p.Rules {
		if !contains(possibleRules, rule.Rule) {
			return fmt.Errorf("rule '%s' is not allowed", rule.Rule)
		}
	}
	return nil
}

func contains(s []string, rule string) bool {
	for _, a := range s {
		if a == rule {
			return true
		}
	}
	return false
}

// ValidatePassword is responsible for validating the password
func ValidatePassword(p types.Payload) (types.ValidationResponse, error) {

	// Checks if the sent rules are valid even before entering the for to avoid unnecessary processing
	if err := checkAllowedRules(p); err != nil {
		return types.ValidationResponse{}, err
	}

	// Create a slice to store the rules that the password does not match
	noMatch := []string{}

	// Iterate over the rules and validate the password
	for _, rule := range p.Rules {

		// Validate the password based on the rule
		switch rule.Rule {

		// Check if the password minimum size is valid
		case "minSize":
			if len(p.Password) < rule.Value {
				noMatch = append(noMatch, "minSize")
			}

		// Check if the password has uppercase characters
		case "minUppercase":
			nUppercase := 0
			for _, r := range p.Password {
				if r >= 'A' && r <= 'Z' {
					nUppercase++
				}
			}
			if nUppercase < rule.Value {
				noMatch = append(noMatch, "minUppercase")
			}

		// Check if the password has lowercase characters
		case "minLowercase":
			nLowercase := 0
			for _, r := range p.Password {
				if r >= 'a' && r <= 'z' {
					nLowercase++
				}
			}
			if nLowercase < rule.Value {
				noMatch = append(noMatch, "minLowercase")
			}

		// Check if the password has special characters
		case "minSpecialChars":
			specialCharsRegex := regexp.MustCompile(`[!@#%^&*()\-+\/{}[\]]`)
			specialChars := 0
			for _, r := range p.Password {
				if specialCharsRegex.MatchString(string(r)) {
					specialChars++
				}
			}
			if specialChars < rule.Value {
				noMatch = append(noMatch, "minSpecialChars")
			}

		// Check if the password has no repeated characters, the p.Value in this case is not used
		case "noRepeted":
			flag := true
			last := rune(0)
			for _, r := range p.Password {
				if last == r {
					flag = false
					break
				}
				last = r
			}
			if !flag {
				noMatch = append(noMatch, "noRepeted")
			}

		// Check if the password minimum digit is valid
		case "minDigit":
			nDigits := 0
			for _, r := range p.Password {
				if r >= '0' && r <= '9' {
					nDigits++
				}
			}
			if nDigits < rule.Value {
				noMatch = append(noMatch, "minDigit")
			}

		// Default case, check if the password minimum size is valid
		default:
			if len(p.Password) < 3 {
				noMatch = append(noMatch, "minSize")
			}
		}
	}

	// Create the response
	response := types.ValidationResponse{
		Verify:  len(noMatch) == 0,
		NoMatch: noMatch,
	}

	return response, nil
}
