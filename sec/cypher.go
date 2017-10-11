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
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"crypto/x509"
	"encoding/asn1"
	"encoding/pem"
	"fmt"
)

// Sha512 computes the checksum of the provided message and returns its
// value as an hexadecimal string
func Sha512(message string) string {
	checksum := sha512.Sum512([]byte(message))
	return fmt.Sprintf("%x", checksum)
}

// GenerateKeypair generates a new RSA keypair of the given bits size.
func GenerateKeypair(bits int) (*rsa.PrivateKey, error) {
	if bits < 4096 {
		return nil, fmt.Errorf("Cannot generate a keypair using less than 4096 bits")
	}

	reader := rand.Reader
	key, err := rsa.GenerateKey(reader, bits)

	if err != nil {
		return nil, err
	}

	return key, nil
}

// EncodePrivateKey returns the raw bytes of the PEM-encoded private key.
func EncodePrivateKey(key *rsa.PrivateKey) ([]byte, error) {
	if key == nil {
		return nil, fmt.Errorf("Private key cannot be nil")
	}

	block := new(pem.Block)
	block.Type = "PRIVATE KEY"
	block.Bytes = x509.MarshalPKCS1PrivateKey(key)

	return pem.EncodeToMemory(block), nil
}

// EncodePublicKey returns the raw bytes of the PEM-encoded public key.
func EncodePublicKey(key *rsa.PublicKey) ([]byte, error) {
	if key == nil {
		return nil, fmt.Errorf("Public key cannot be nil")
	}

	bytes, err := asn1.Marshal(*key)
	if err != nil {
		return nil, err
	}

	block := new(pem.Block)
	block.Type = "PUBLIC KEY"
	block.Bytes = bytes

	return pem.EncodeToMemory(block), nil
}
