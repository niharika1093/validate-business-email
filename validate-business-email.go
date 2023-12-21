package auth

import (
	"regexp"
	"strings"

	"golang.org/x/exp/slices"
)

var emailRegexp = regexp.MustCompile("^.+?@.+?\\..+?$") //nolint

// IsValidEmail checks if email given is valid
func IsValidEmail(email string) bool {
	if strings.Count(email, "@") > 1 {
		return false
	}
	return emailRegexp.MatchString(email)
}

func ValidateBusinessEmail(email string) bool {

	IsValidEmail(email) // First check if email is valid

	splitEmail := strings.SplitAfter(email, "@")
	domain := splitEmail[1]
	if slices.Contains(Free_email_provider_domains, domain) {
		return false
	}
	return true // if the free provider set doesn't have this domain, then most likely it's a company email address
}
