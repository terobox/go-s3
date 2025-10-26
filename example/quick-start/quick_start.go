package main

import (
	"context"
	"fmt"
	"log"

	"github.com/terobox/go-s3"
)

func main() {
	client, err := s3.New(
		"s3.xbqx.com", // endpoint
		"us-east-1",   // region
		"BUbZ7afHpTTuh1DNFhKL",
		"Nry0ZO6JSbijxqALdMQqgVKw3Bs0s1HQiPHPQC1v",
		"mac", // bucket name
		true,  // use SSL
	)
	if err != nil {
		log.Fatal(err)
	}

	result, err := client.Upload(context.Background(), "local.txt", &s3.UploadOptions{
		Directory:    "son", // 上传到子目录
		PreserveName: false, // 自动重命名
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("✅ 上传成功!\n对象键名: %s\n文件大小: %d 字节\n访问链接: %s\n",
		result.Key, result.Size, result.URL)
}
