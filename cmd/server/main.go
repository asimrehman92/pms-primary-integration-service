package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

<<<<<<< HEAD
	"github.com/asimrehman/pms-primary-integration-service/internal/config"
=======
	"github.com/amalikh/pms-primary-integration-service-main/internal/config"
>>>>>>> 547fb98f6c68e4091bdcbb151f74c4b8c9a65b6a
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/aws/aws-sdk-go/service/sqs"
<<<<<<< HEAD
)

=======
	"github.com/borlinp/amazon-sns-sqs/common"
)

// func init() {
// 	config, err := config.LoadConfig("C:/Users/HP/go/src/github.com/amalikh/pms-primary-integration-service-main/internal/config")
// 	if err != nil {
// 		log.Fatal("cannot load config: ", err)
// 	}
// }

>>>>>>> 547fb98f6c68e4091bdcbb151f74c4b8c9a65b6a
func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	subscribe(sigs)
	// SubscribeSNS(sigs)
}

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

func subscribe(cancel <-chan os.Signal) {

	awsSession := config.BuildSession()
	svc := sqs.New(awsSession, nil)
	fmt.Println("publisher subscribe func")
	QueueName := aws.String(os.Args[2])
	qUrl, _ := GetQueueURL(awsSession, QueueName)
	queueURL := qUrl.QueueUrl
	for {
		messages := GetMessages()

		for _, msg := range messages {
			if msg == nil {
				continue
			}
			fmt.Println(*msg.Body)
			go DeleteMessage(svc, queueURL, msg.ReceiptHandle)
		}

		select {
		case <-cancel:
			return
		case <-time.After(1 * time.Millisecond):
			// return
		}
	}
}

func DeleteMessage(svc *sqs.SQS, queueUrl *string, handle *string) {
	delInput := &sqs.DeleteMessageInput{
		QueueUrl:      queueUrl,
		ReceiptHandle: handle,
	}
	_, err := svc.DeleteMessage(delInput)

	if err != nil {
		fmt.Println("Delete Error", err)
		return
	}
}

func GetMessages() []*sqs.Message {

	awsSession := config.BuildSession()
	svc := sqs.New(awsSession)
	QueueName := aws.String(os.Args[2])
	// Get URL of queue
	urlResult, err := GetQueueURL(awsSession, QueueName)
	if err != nil {
		fmt.Println("Got an error getting the queue URL:")
		fmt.Println(err)
	}
	queueURL := urlResult.QueueUrl
	receiveMessagesInput := &sqs.ReceiveMessageInput{
		AttributeNames: []*string{
			aws.String(sqs.MessageSystemAttributeNameSentTimestamp),
		},
		MessageAttributeNames: []*string{
			aws.String(sqs.QueueAttributeNameAll),
		},
		QueueUrl:            queueURL,
		MaxNumberOfMessages: aws.Int64(10),
		WaitTimeSeconds:     aws.Int64(3),  // max 20
		VisibilityTimeout:   aws.Int64(20), // max 20
	}
	msgResult, err := svc.ReceiveMessage(receiveMessagesInput)

	if err != nil {
		fmt.Println("Error: ", err)
		return nil
	}

	if msgResult == nil || len(msgResult.Messages) == 0 {
		return nil
	}

	// fmt.Println("Message Body:     ", *msgResult.Messages[0].Body)
	// fmt.Println("Message Handle: " + *msgResult.Messages[0].ReceiptHandle)
	// fmt.Println("msgResult: ", msgResult.Messages)
	return msgResult.Messages
}

func SubscribeSNS(cancel <-chan os.Signal) {
<<<<<<< HEAD
	awsSession := config.BuildSession()
=======
	awsSession := common.BuildSession()
>>>>>>> 547fb98f6c68e4091bdcbb151f74c4b8c9a65b6a
	svcc := sqs.New(awsSession)
	QueueName := aws.String(os.Args[2])
	qUrl, _ := GetQueueURL(awsSession, QueueName)
	queueURL := qUrl.QueueUrl

<<<<<<< HEAD
	Sess := config.BuildSession()
=======
	Sess := common.BuildSession()
>>>>>>> 547fb98f6c68e4091bdcbb151f74c4b8c9a65b6a
	svc := sns.New(Sess)
	// TopicName := aws.String(os.Args[2])
	_, err := svc.Subscribe(&sns.SubscribeInput{
		// Attributes:            nil,
		Endpoint: aws.String("arn:aws:sqs:us-east-2:421122548895:MySecondQ"),
		Protocol: aws.String("sqs"),
		// ReturnSubscriptionArn: nil,
		TopicArn: aws.String("arn:aws:sns:us-east-2:421122548895:MyFirstTopic"),
	})
	if err != nil {
		fmt.Println(err)
	}

	for {
		messages := GetMessages()

		for _, msg := range messages {
			if msg == nil {
				continue
			}
			fmt.Println(*msg.Body)
			go DeleteMessage(svcc, queueURL, msg.ReceiptHandle)
		}

		select {
		case <-cancel:
			return
		case <-time.After(1 * time.Millisecond):
			// return
		}
	}
}
<<<<<<< HEAD
=======

// func BuildSession() *session.Session {
// 	// . means current folder
// 	config, err := config.LoadConfig("C:/Users/dell/go/src/github.com/amalikh/pms-primary-integration-service")

// 	if err != nil {
// 		log.Fatal("cannot load config: ", err)
// 	}
// 	sessionConfig := &aws.Config{
// 		Region:      aws.String(config.Region),
// 		Credentials: credentials.NewStaticCredentials(config.AccessKey, config.SecretKey, ""),
// 	}

// 	sess, err := session.NewSession(sessionConfig)
// 	if err != nil {
// 		fmt.Println("error", err)

// 	}
// 	return sess
// }
>>>>>>> 547fb98f6c68e4091bdcbb151f74c4b8c9a65b6a
