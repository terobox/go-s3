package s3

// ConflictPolicy 文件名冲突时的处理策略
type ConflictPolicy string

const (
	ConflictOverwrite ConflictPolicy = "overwrite" // 直接覆盖
	ConflictRename    ConflictPolicy = "rename"    // 自动加后缀 (1), (2)...
	ConflictError     ConflictPolicy = "error"     // 直接返回错误
)

// UploadOptions 上传配置
type UploadOptions struct {
	Rename         bool           // 是否强制重命名（用 ULID）
	ConflictPolicy ConflictPolicy // 冲突处理策略
	ContentType    string         // 可选：MIME 类型
	PublicRead     bool           // 是否公开可读（主流公开图床使用）
}
