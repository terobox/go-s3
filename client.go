package s3

import (
	"context"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/smithy-go"
)

// Client 封装 S3 客户端
type Client struct {
	s3       *s3.Client
	Bucket   string
	Endpoint string // 新增：存储 endpoint 用于构建 URL
	Region   string // 新增：存储 region 用于构建 URL
	UseSSL   bool   // 新增：存储 SSL 配置用于构建 URL
}

// New 创建新的 S3 客户端
func New(endpoint, region, key, secret, bucket string, useSSL bool) (*Client, error) {
	ctx := context.Background()

	protocol := "https"
	if !useSSL {
		protocol = "http"
	}

	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(key, secret, "")),
		config.WithRegion(region),
		config.WithEndpointResolver(aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
			return aws.Endpoint{
				URL:               fmt.Sprintf("%s://%s", protocol, endpoint),
				SigningRegion:     region,
				HostnameImmutable: true,
			}, nil
		})),
	)
	if err != nil {
		return nil, err
	}

	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.UsePathStyle = true
	})

	return &Client{
		s3:       client,
		Bucket:   bucket,
		Endpoint: endpoint,
		Region:   region,
		UseSSL:   useSSL,
	}, nil
}

// Exists 检查对象是否存在于 S3
func (c *Client) Exists(key string) (bool, error) {
	_, err := c.s3.HeadObject(context.TODO(), &s3.HeadObjectInput{
		Bucket: aws.String(c.Bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		var apiErr smithy.APIError
		if errors.As(err, &apiErr) {
			code := apiErr.ErrorCode()
			if code == "NotFound" || code == "NoSuchKey" {
				return false, nil
			}
		}
		return false, err
	}
	return true, nil
}
