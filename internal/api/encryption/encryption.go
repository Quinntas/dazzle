package encryption

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/crypto/pbkdf2"
	"golang.org/x/crypto/sha3"
)

type HashSalt struct {
	Result string
}

func GenerateSalt(length uint32) []byte {
	secret := make([]byte, length)
	_, _ = rand.Read(secret)
	return secret
}

type CryptoParams struct {
	value      string
	Salt       []byte
	Pepper     string
	Iterations int
	Length     int
}

func ParseEncryption(value string, hash string, pepper string) (*CryptoParams, error) {
	params := &CryptoParams{}

	hashParams := strings.Split(hash, "$")

	salt, err := hex.DecodeString(hashParams[1])
	if err != nil {
		return nil, err
	}
	iterations, err := strconv.Atoi(hashParams[2])
	if err != nil {
		return nil, err
	}
	length, err := strconv.Atoi(hashParams[3])
	if err != nil {
		return nil, err
	}

	params.Salt = salt
	params.Iterations = iterations
	params.Length = length
	params.value = value
	params.Pepper = pepper

	return params, nil
}

func ConstantTimeStringCompare(a, b string) bool {
	return subtle.ConstantTimeCompare([]byte(a), []byte(b)) == 1
}

func CompareEncryption(value string, hash string, pepper string) (bool, error) {
	hashParams, err := ParseEncryption(value, hash, pepper)
	if err != nil {
		return false, err
	}

	newHash, err := GenerateEncryptionWithParams(hashParams)
	if err != nil {
		return false, err
	}

	return ConstantTimeStringCompare(hash, newHash), nil
}

func GenerateDefaultEncryption(value string, pepper string) (string, error) {
	return GenerateEncryptionWithParams(&CryptoParams{
		value:      value,
		Salt:       GenerateSalt(32),
		Pepper:     pepper,
		Iterations: 10000,
		Length:     32,
	})
}

func GenerateEncryptionWithParams(cryptoParams *CryptoParams) (string, error) {
	hash := pbkdf2.Key([]byte(cryptoParams.Pepper+cryptoParams.value), cryptoParams.Salt, cryptoParams.Iterations, cryptoParams.Length, sha3.New256)

	hashHex := hex.EncodeToString(hash)
	saltHex := hex.EncodeToString(cryptoParams.Salt)

	result := fmt.Sprintf("sha256$%s$%d$%d$%s", saltHex, cryptoParams.Iterations, cryptoParams.Length, hashHex)

	return result, nil
}
