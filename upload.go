package s3

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// Upload 上传本地文件到 S3
// 功能特性：
// - 自动检测文件 MIME 类型
// - 支持 ULID 重命名或保留原文件名
// - 支持设置公开可读权限
// - 返回完整的访问链接和元数据
func (c *Client) Upload(
	ctx context.Context,
	filePath string,
	options *UploadOptions,
) (*UploadResult, error) {
	// 参数校验
	if filePath == "" {
		return nil, fmt.Errorf("文件路径不能为空")
	}

	// 设置默认选项
	if options == nil {
		options = &UploadOptions{}
	}

	// 打开并读取本地文件
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("打开文件失败: %w", err)
	}
	defer file.Close()

	// 获取文件信息
	fileInfo, err := file.Stat()
	if err != nil {
		return nil, fmt.Errorf("获取文件信息失败: %w", err)
	}

	// 生成 S3 对象键名
	s3Key := generateObjectKey(filePath, options)

	// 检测 MIME 类型（使用文件句柄）
	contentType, err := detectContentType(file)
	if err != nil {
		return nil, fmt.Errorf("检测文件类型失败: %w", err)
	}

	// 构建上传参数
	putInput := &s3.PutObjectInput{
		Bucket:      aws.String(c.Bucket),
		Key:         aws.String(s3Key),
		Body:        file, // 直接传文件句柄
		ContentType: aws.String(contentType),
	}

	// 执行上传
	_, err = c.s3.PutObject(ctx, putInput)
	if err != nil {
		return nil, fmt.Errorf("上传文件失败: %w", err)
	}

	// 构建访问链接
	url := buildObjectURL(c.Endpoint, c.Bucket, s3Key, c.UseSSL)

	// 返回上传结果
	return &UploadResult{
		Key:         s3Key,
		URL:         url,
		Size:        fileInfo.Size(),
		ContentType: contentType,
	}, nil
}
