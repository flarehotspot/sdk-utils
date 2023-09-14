package strings

import (
	"crypto/sha1"
	"encoding/hex"
	"strings"
)

func Sha1Hash(texts ...string) string {
	allstr := strings.Join(texts, "")
	hash := sha1.Sum([]byte(allstr))
	return hex.EncodeToString(hash[:])
}
