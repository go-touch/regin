package encrypt

import (
	"crypto/sha1"
	"fmt"
	"strings"
)

// hash加密算法
func Sha1(str ...string) string {
	Sha1Ptr := sha1.New()
	_, _ = Sha1Ptr.Write([]byte(strings.Join(str, "")))
	return fmt.Sprintf("%x", Sha1Ptr.Sum(nil))
}
