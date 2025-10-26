package s3

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// Client 封装 S3 客户端
type Client struct {
	s3     *s3.Client
	Bucket string
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

	return &Client{s3: client, Bucket: bucket}, nil
}
