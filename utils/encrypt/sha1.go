package encrypt

import (
	"crypto/sha1"
	"fmt"
	"hash"
	"strings"
)

type Sha1 struct {
	hash.Hash
}

// 创建Sha1
func NewSha1() *Sha1 {
	return &Sha1{
		sha1.New(),
	}
}

// 计算hash值
func (s *Sha1) Calculate(str ...string) string {
	_, _ = s.Write([]byte(strings.Join(str, "")))
	return fmt.Sprintf("%x", s.Sum(nil))
}