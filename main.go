package main

import (
	"context"
	"log"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func main() {
	config, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	client := s3.NewFromConfig(config)

	s3_list, err := client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: aws.String("aws-adk-test"),
	})
	if err != nil {
		log.Fatal(err)
	}

	for i, object := range s3_list.Contents {
		log.Printf("(%d) key=%s size=%d", i + 1, aws.ToString(object.Key), object.Size)
	}
}