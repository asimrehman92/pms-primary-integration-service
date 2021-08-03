package main

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func GetQueueURL(sess *session.Session, queueName *string) (*sqs.GetQueueUrlOutput, error) {
	svc := sqs.New(sess)
	result, err := svc.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: queueName,
	})
	if err != nil {
		return nil, err
	}

	return result, nil

}
