package hash

import (
	"encoding/base64"
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
