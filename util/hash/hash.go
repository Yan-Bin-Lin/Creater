package hash

import (
	"app/util/random"
	"encoding/base64"
	"golang.org/x/crypto/argon2"
	"golang.org/x/crypto/blake2b"
	"strings"
)

// get hash of blake2b 256
func GetHash(strs ...string) [32]byte {
	return blake2b.Sum256([]byte(strings.Join(strs, "|")))
}

// get hash string in hex of blake2b 256
func GetHashString(strs ...string) string {
	result := blake2b.Sum256([]byte(strings.Join(strs, "|")))
	return base64.RawURLEncoding.EncodeToString(result[:])
}

// get hash of argon2 for password hash
func GetPWHash(pw string, time, memory uint32, threads uint8, keyLen uint32) ([]byte, []byte, error) {
	salt, err := random.GetRandomBytes(16)
	if err != nil {
		return nil, nil, err
	}

	return argon2.IDKey([]byte(pw), salt, time, memory, threads, keyLen), salt, nil
}

// get hash string in base64 url raw encode of argon2 for password hash
func GetPWHashString(pw string, time, memory uint32, threads uint8, keyLen uint32) (string, string, error) {
	hash, salt, err := GetPWHash(pw, time, memory, threads, keyLen)
	if err != nil {
		return "", "", nil
	}

	return base64.RawURLEncoding.EncodeToString(hash), base64.RawURLEncoding.EncodeToString(salt), nil
}
