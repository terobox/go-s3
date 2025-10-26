package s3

// UploadOptions 上传配置
type UploadOptions struct {
	Directory    string // Directory S3 存储目录路径（可选）
	PreserveName bool   // PreserveName 是否保留原始文件名（默认 false）
}

// UploadResult 上传结果
type UploadResult struct {
	Key         string // Key S3 对象键名
	URL         string // URL 完整的 S3 访问链接
	Size        int64  // Size 文件大小（字节）
	ContentType string // ContentType 文件 MIME 类型
}
