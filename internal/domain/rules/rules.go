package rules

import "regexp"

// NOTE: `email` field in User must be a valid email-address -> handle it using gin-binding-validator
var (
	usernamePattern = regexp.MustCompile("^[a-z0-9_]{3,64}$")
	passwordPattern = regexp.MustCompile("^[A-Za-z0-9!@#$%&*]{8,64}$")
	tagPattern      = regexp.MustCompile("^[ا-یa-z0-9_]{1,32}$")
)

func CheckUsernamePattern(username string) bool {
	return usernamePattern.MatchString(username)
}

func CheckPasswordPattern(password string) bool {
	return passwordPattern.MatchString(password)
}

func CheckTagPattern(tag string) bool {
	return tagPattern.MatchString(tag)
}
