package pkg

import (
	"context"
	"encoding/base64"
	"io/ioutil"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func GetS3ImageEncode(key string) (string, error) {

	if key == "" {
		return "", nil
	}

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return "", err
	}

	client := s3.NewFromConfig(cfg)

	obj, err := client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String("emotion-sns"),
		Key:    aws.String("user/" + key),
	})
	if err != nil {
		return "", err
	}

	b, err := ioutil.ReadAll(obj.Body)
	if err != nil {
		return "", err
	}

	//JS側でそのまま表示できるように「data:image/png;base64,....」の形式での受け渡し
	return "data:" + *obj.ContentType + ";base64," + base64.StdEncoding.EncodeToString(b), nil
}
