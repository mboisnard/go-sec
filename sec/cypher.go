/*
TODO :

- AES en mode CBC
- RSA

Utilisation de lib
- TODO Recherche de la lib la plus adaptée

Méthodes
- Cypher AES utilisant CBC
- UnCypher AES

- Génération clé RSA pour un user / server
- Chiffrement clé publique
- Déchiffrement clé privée
- Signature clé privée
- Vérification clé publique
*/

package sec

import (
	"crypto/sha512"
	"fmt"
)

// Sha512 computes the checksum of the provided message and returns its
// value as an hexadecimal string
func Sha512(message string) string {
	checksum := sha512.Sum512([]byte(message))
	return fmt.Sprintf("%x", checksum)
}
