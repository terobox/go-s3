package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload" // 自动加载 .env 文件
	"github.com/terobox/go-s3"
)

func main() {
	// 从 .env 中读取 S3 配置
	endpoint := os.Getenv("S3_ENDPOINT")
	region := os.Getenv("S3_REGION")
	accessKey := os.Getenv("S3_ACCESS_KEY")
	secretKey := os.Getenv("S3_SECRET_KEY")
	bucket := os.Getenv("S3_BUCKET")
	useSSLStr := os.Getenv("S3_USE_SSL")

	// 转换字符串为布尔
	useSSL, _ := strconv.ParseBool(useSSLStr)

	// 初始化 S3 客户端
	client, err := s3.New(
		endpoint,
		region,
		accessKey,
		secretKey,
		bucket,
		useSSL,
	)
	if err != nil {
		log.Fatalf("❌ 初始化 S3 客户端失败: %v", err)
	}

	fmt.Println("🚀 开始上传测试文件...")

	// 模拟上传本地文件
	localFilePath := "demo/hello.txt"

	// 如果文件不存在则创建
	if _, err := os.Stat(localFilePath); os.IsNotExist(err) {
		if err := os.MkdirAll("demo", 0755); err != nil {
			log.Fatalf("❌ 创建 demo 目录失败: %v", err)
		}
		content := fmt.Sprintf("Hello from go-s3! %v\n", time.Now().Format(time.RFC3339))
		if err := os.WriteFile(localFilePath, []byte(content), 0644); err != nil {
			log.Fatalf("❌ 创建测试文件失败: %v", err)
		}
	}

	// 执行上传（带选项）
	result, err := client.Upload(context.Background(), localFilePath, &s3.UploadOptions{
		Directory:    "test", // 上传到 S3 的子目录，例如 mac/test/
		PreserveName: false,  // 不保留原文件名，使用 ULID
	})
	if err != nil {
		log.Fatalf("❌ 上传失败: %v", err)
	}

	// 输出结果
	fmt.Println("✅ 上传成功!")
	fmt.Printf("对象键名: %s\n", result.Key)
	fmt.Printf("文件大小: %d 字节\n", result.Size)
	fmt.Printf("文件类型: %s\n", result.ContentType)
	fmt.Printf("访问链接: %s\n", result.URL)
}
