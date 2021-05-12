package pgp

import (
	"github.com/ProtonMail/gopenpgp/v2/crypto"
)

//go:generate msgp

// PrivateKey Key struct with functions
type PrivateKey struct {
	// Armored is Public because it needs to be stored in session for decryption
	Armored string `msg:"privateKey"`
}

// PublicKey Key struct with functions
type PublicKey struct {
	Armored string `msg:"publicKey"`
}

// NewKey Generates a new Curve25519 Pgp key
func NewKey(name, userStr string) (*PrivateKey, error) {
	privateKey, err := crypto.GenerateKey(name, userStr, "x25519", 0)
	if err != nil {
		return nil, err
	}

	armoredKey, err := privateKey.Armor()
	if err != nil {
		return nil, err
	}

	return &PrivateKey{Armored: armoredKey}, nil
}

// PublicKey of Private Key
func (pk *PrivateKey) PublicKey() (*PublicKey, error) {
	privateKeyObj, err := crypto.NewKeyFromArmored(pk.Armored)
	if err != nil {
		return nil, err
	}

	publicKey, err := privateKeyObj.GetArmoredPublicKey()
	if err != nil {
		return nil, err
	}

	return &PublicKey{Armored: publicKey}, nil
}

// EncryptArmored encrypt data with public key
func (pk *PublicKey) EncryptArmored(data []byte) (string, error) {
	publicKeyRing, err := newKeyRingFromArmored(pk.Armored)
	if err != nil {
		return "", err
	}

	pgpMessage, err := publicKeyRing.Encrypt(crypto.NewPlainMessage(data), nil)
	if err != nil {
		return "", err
	}
	return pgpMessage.GetArmored()
}

// DecryptArmored data with private key
func (pk *PrivateKey) DecryptArmored(armored string) ([]byte, error) {
	pgpMessage, err := crypto.NewPGPMessageFromArmored(armored)
	if err != nil {
		return nil, err
	}

	privateKeyRing, err := newKeyRingFromArmored(pk.Armored)
	if err != nil {
		return nil, err
	}

	message, err := privateKeyRing.Decrypt(pgpMessage, nil, crypto.GetUnixTime())
	if err != nil {
		return nil, err
	}

	privateKeyRing.ClearPrivateParams()

	return message.GetBinary(), nil
}

func newKeyRingFromArmored(armored string) (*crypto.KeyRing, error) {
	keyObj, err := crypto.NewKeyFromArmored(armored)
	if err != nil {
		return nil, err
	}
	keyRing, err := crypto.NewKeyRing(keyObj)
	if err != nil {
		return nil, err
	}
	return keyRing, nil
}
