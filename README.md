
# go-s3

一个轻量级的 Go 封装库，用于更方便地操作 **S3 / MinIO / 自建兼容存储**。

> Lightweight Golang SDK for S3/MinIO-compatible storage.

---

## ✨ 特性

- 📦 一行代码上传文件，自动生成访问链接  
- 🔤 自动检测 MIME 类型（支持扩展名与内容嗅探）  
- 🔁 文件命名策略：保留原名 / ULID 自动命名  
- 🧭 完整的访问 URL 自动构建（支持 HTTP / HTTPS）  
- ⚙️ 兼容所有 S3 API（如 AWS、MinIO、Ceph、Wasabi）

---

## 🚀 快速开始

```go
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
		"ACCESS_KEY",
		"SECRET_KEY",
		"BUCKET_NAME", // bucket name
		true,          // use SSL
	)
	if err != nil {
		log.Fatal(err)
	}

	result, err := client.Upload(context.Background(), "local.txt", &s3.UploadOptions{
		Directory:    "test",      // 上传到子目录
		PreserveName: false,       // 自动重命名
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("✅ 上传成功!\n对象键名: %s\n文件大小: %d 字节\n访问链接: %s\n",
		result.Key, result.Size, result.URL)
}
````

---

## ⚙️ 可选上传参数

| 字段名            | 类型       | 说明                     | 默认值     |
| -------------- | -------- | ---------------------- | ------- |
| `Directory`    | `string` | 上传目录路径（例如 `"images/"`） | `""`    |
| `PreserveName` | `bool`   | 是否保留原文件名               | `false` |

---

## 📁 上传结果结构体

```go
type UploadResult struct {
    Key         string // 对象键名
    URL         string // 完整访问链接
    Size        int64  // 文件大小
    ContentType string // MIME 类型
}
```

---

## 🧩 TODO

* [ ] 批量上传（支持多文件并发）

---

## 📜 License

MIT
