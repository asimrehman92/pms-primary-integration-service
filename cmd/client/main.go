package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/asimrehman/pms-primary-integration-service/internal/config"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// PublishMsg()

	switch os.Args[1] {
	case "queue":
		go PublishMsgSqs()
	case "topic":
		go PublishMsgSns()
	}

	<-sigs
	// GetMessages()
	// ListQ()
}

func PublishMsgSqs() {
	reader := bufio.NewReader(os.Stdin)
	for {
		text, _ := reader.ReadString('\n')
		if text == "\n" {
			continue
		}
		SendMsgSQS(text[:len(text)-1])
	}
}

func PublishMsgSns() {
	reader := bufio.NewReader(os.Stdin)
	for {
		text, _ := reader.ReadString('\n')
		if text == "\n" {
			continue
		}
		SendSNS(text[:len(text)-1])
	}
}
func SendMsgSQS(message string) error {

	awsSession := config.BuildSession()
	svc := sqs.New(awsSession)
	QueueName := aws.String(os.Args[2])
	// Get URL of queue
	result, err := GetQueueURL(awsSession, QueueName)
	if err != nil {
		fmt.Println("Got an error getting the queue URL:")
		fmt.Println(err)
	}

	queueURL := result.QueueUrl

	_, err = svc.SendMessage(&sqs.SendMessageInput{
		DelaySeconds: aws.Int64(0),
		// MessageAttributes: map[string]*sqs.MessageAttributeValue{
		// 	"Title": &sqs.MessageAttributeValue{
		// 		DataType:    aws.String("String"),
		// 		StringValue: aws.String("The Whistler"),
		// 	},
		// 	"Author": &sqs.MessageAttributeValue{
		// 		DataType:    aws.String("String"),
		// 		StringValue: aws.String("John Grisham"),
		// 	},
		// 	"WeeksOn": &sqs.MessageAttributeValue{
		// 		DataType:    aws.String("Number"),
		// 		StringValue: aws.String("6"),
		// 	},
		// },
		MessageBody: aws.String(message),
		QueueUrl:    queueURL,
	})

	if err != nil {
		return err
	}
	fmt.Println("Sent message to queue ")
	return nil
}

func SendSNS(message string) {
	awsSession := config.BuildSession()
	svc := sns.New(awsSession)
	Topicarn := aws.String(os.Args[2])
	pubInput := &sns.PublishInput{

		Message:  aws.String(message),
		TopicArn: Topicarn,
	}

	_, err := svc.Publish(pubInput)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Sent message to queue ")

	//fmt.Println(output.MessageId)
}
