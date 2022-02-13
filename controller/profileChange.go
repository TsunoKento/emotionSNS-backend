package controller

import (
	"TsunoKento/emotionSNS/model"
	"bytes"
	"context"
	"encoding/base64"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
)

func ProfileChange(img, uid, name string, id uint) error {
	user := &model.User{ID: id, UserID: uid, Name: name}

	if img != "" {
		uu := uuid.NewString()
		image, err := base64.StdEncoding.DecodeString(img)
		if err != nil {
			return err
		}

		cfg, err := config.LoadDefaultConfig(context.TODO())
		if err != nil {
			return err
		}

		client := s3.NewFromConfig(cfg)

		uploader := manager.NewUploader(client)

		_, err = uploader.Upload(context.TODO(), &s3.PutObjectInput{
			Bucket:      aws.String("emotion-sns"),
			Key:         aws.String("user/" + uu),
			ContentType: aws.String("image/png"),
			Body:        bytes.NewReader(image),
		})
		if err != nil {
			return err
		}

		user.Image = uu
	}

	if err := user.UpdateUser(); err != nil {
		return err
	}

	return nil
}
