package s3

import (
	"fmt"
	"io"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// generateObjectKey 生成 S3 对象键名
func generateObjectKey(filePath string, options *UploadOptions) string {
	var filename string
	// 根据配置决定文件名
	if options.PreserveName {
		// 保留原文件名
		filename = filepath.Base(filePath)
	} else {
		// 使用 ULID 重命名，保留文件扩展名
		ext := filepath.Ext(filePath)
		filename = GenerateULID() + ext
	}
	// 拼接目录路径
	if options.Directory != "" {
		// 清理目录路径（移除首尾斜杠）
		directory := strings.Trim(options.Directory, "/")
		return JoinPath(directory, filename)
	}
	return filename
}

// detectContentType 尝试通过文件扩展名和内容嗅探 MIME 类型
func detectContentType(file *os.File) (string, error) {
	// 1️⃣ 优先使用扩展名
	ext := strings.ToLower(filepath.Ext(file.Name()))
	if ctype := mime.TypeByExtension(ext); ctype != "" {
		// 重置文件指针到开头，防止外部函数已读
		if _, err := file.Seek(0, io.SeekStart); err != nil {
			return "", err
		}
		return ctype, nil
	}

	// 2️⃣ 读取前 512 字节嗅探内容
	buffer := make([]byte, 512)
	n, err := file.Read(buffer)
	if err != nil && err != io.EOF {
		return "", err
	}

	ctype := http.DetectContentType(buffer[:n])

	// 3️⃣ 重置文件指针到开头，确保 S3 可完整读取
	if _, err = file.Seek(0, io.SeekStart); err != nil {
		return "", err
	}

	return ctype, nil
}

// buildObjectURL 构建完整的 S3 对象访问链接
func buildObjectURL(endpoint, bucket, key string, useSSL bool) string {
	// 格式通常是: http(s)://<endpoint>/<bucket>/<key>
	protocol := "https"
	if !useSSL {
		protocol = "http"
	}
	return fmt.Sprintf("%s://%s/%s/%s", protocol, endpoint, bucket, key)
}
