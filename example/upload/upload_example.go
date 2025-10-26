package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload" // 自动加载 .env 文件

	"github.com/terobox/go-s3"
)

func main() {
	// 从环境变量中读取配置
	endpoint := os.Getenv("S3_ENDPOINT")
	region := os.Getenv("S3_REGION")
	accessKey := os.Getenv("S3_ACCESS_KEY")
	secretKey := os.Getenv("S3_SECRET_KEY")
	bucket := os.Getenv("S3_BUCKET")
	useSSLStr := os.Getenv("S3_USE_SSL")

	// 将字符串 true/false 转换为 bool
	useSSL, _ := strconv.ParseBool(useSSLStr)

	client, err := s3.New(
		endpoint,
		region,
		accessKey,
		secretKey,
		bucket,
		useSSL,
	)
	if err != nil {
		log.Fatalf("❌ 初始化失败: %v", err)
	}

	fmt.Println("🚀 开始上传测试文件...")
	err = client.Upload("demo/hello.txt", []byte("Hello from go-s3!"))
	if err != nil {
		log.Fatalf("❌ 上传失败: %v", err)
	}
	fmt.Println("✅ 上传成功!")
}
