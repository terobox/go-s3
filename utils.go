package s3

import (
	"crypto/rand"
	"strings"
	"time"

	"github.com/oklog/ulid/v2"
)

// JoinPath 安全拼接 S3 Key 路径
func JoinPath(parts ...string) string {
	return strings.Join(parts, "/")
}

// GenerateULID 生成一个新的 ULID（带时间顺序，可排序）
// 参考：01K8FJ3Q3WA031F7YPV0E9PRBP
func GenerateULID() string {
	t := time.Now()
	entropy := ulid.Monotonic(rand.Reader, 0)
	id := ulid.MustNew(ulid.Timestamp(t), entropy)
	return id.String()
}
