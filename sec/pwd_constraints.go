package sec

/*
 * Présence de majuscule, chiffres & caractères spéciaux
 *	-> éviter les attaques par dictionnaire
 *
 * 8 caractères minimum
 *	-> compliquer les attaques par bruteforce
 *
 * Expiration du mdp après 30 à 90j
 *	-> limiter la durée de validité des mdp
 *
 * Historique de 14 mdp
 *	-> renforcer la limitation de durée de validité des mdp
 *
 * 5 à 7 tentatives de connexion maxi
 *	-> éviter les attaques par bruteforce (attention au ddos !)
 */

// PasswordConstraints is a struct reprensenting
// the contraints given on password management by users
//
// @diversified bool :
type PasswordConstraints struct {
	// Password style rules
	diversified bool
	minSize     int
	maxSize     int
	// Password stockage rules
	expireTime        int
	history           int
	maxConnexionTries int
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
