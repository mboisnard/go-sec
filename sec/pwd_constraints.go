package sec

import "strings"

const lowerAZ = "abcdefghijklmnopqrstuvwxyz"
const upperAZ = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const specialChars = " !\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"

// PasswordConstraints is a struct reprensenting
// the contraints given on password management by users
//
// @diversified : contains upper and lower case letters, and at least one special character
//							---> avoid dictionnary attacks
// @minSize			: minimum password size (inclusive)
// @maxSize			: maximum password size (inclusive)
//							---> make bruteforce attacks more difficult
// @expireTime	: time after which a change password request is sent
// @history			: number of old passwords kept in database to avoid a change to same password
//							---> limitate passwords validity time
// @maxCnxTries	: maximum connexion tries allowed
//							---> avoid bruteforce attacks
type PasswordConstraints struct {
	diversified       bool
	minSize           int
	maxSize           int
	expireTime        int
	history           int
	maxConnexionTries int
}

func isDiversified(password string, diversified bool) bool {
	if diversified {
		return strings.ContainsAny(password, lowerAZ) && strings.ContainsAny(password, upperAZ) && strings.ContainsAny(password, specialChars)
	}
	return true
}

func hasMinSize(password string, minSize int) bool {
	return len(password) >= minSize
}

func hasMaxSize(password string, maxSize int) bool {
	return len(password) <= maxSize
}

// NewDefaultPasswordConstraints returns a new pointer to
// an instance of default password constraints
func NewDefaultPasswordConstraints() *PasswordConstraints {
	constraints := new(PasswordConstraints)
	constraints.diversified = true
	constraints.minSize = 8
	constraints.maxSize = 15
	constraints.expireTime = 42
	constraints.history = 13
	constraints.maxConnexionTries = 7
	return constraints
}
