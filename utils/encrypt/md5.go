package encrypt

import (
	"crypto/md5"
	"fmt"
	"strings"
)

// md5加密算法
func Md5(str ...string) string {
	md5Ptr := md5.New()
	_, _ = md5Ptr.Write([]byte(strings.Join(str, "")))
	return fmt.Sprintf("%x", md5Ptr.Sum(nil))
}
