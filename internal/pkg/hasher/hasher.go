/* Package implements lets-go-chat functions
 * Unlicensed.
 */
package hasher

import (
	"crypto/sha256"
	"fmt"
	"strings"
)

// Hash the pasword
func HashPassword(password string) (string, error) {
	hash := sha256.New()
	numBytes, err := hash.Write([]byte(password))
	if numBytes > 0 {
		return fmt.Sprintf("%x", hash.Sum(nil)), err
	}
	return "", err
}

// Check pasword hash
func CheckPasswordHash(password, hash string) bool {
	ret := false
	calculatedHash, _ := HashPassword(password)
	if result := strings.Compare(hash, calculatedHash); result == 0 {
		ret = true
	}
	return ret
}
