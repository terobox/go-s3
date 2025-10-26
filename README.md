# go-s3
一个 Go 封装库（SDK），让操作 S3 / MinIO / 自建存储更轻量、更方便、更可复用。

Lightweight Golang SDK for S3/MinIO-compatible storage.

## 上传功能

- 自动重命名，函数返回 s3 链接
- 名称重复，冲突策略：覆盖、重命名、返回报错
- s3 存储，区分用户可见文件，如何实现？
- 多文件上传
- public bucket but private object
- 指定文件 MIME 类型（Content-Type 头）告诉浏览器或 API 调用者文件是什么类型（影响预览、下载方式）。