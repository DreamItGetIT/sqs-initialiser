package main

import (
	"errors"
	"flag"
	"log"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type queuesNames struct {
	names []string
}

func (qn *queuesNames) String() string {
	return strings.Join(qn.names, ",")
}

func (qn *queuesNames) Set(val string) error {
	if len(val) == 0 {
		return errors.New("value cannot be empty")
	}

	qn.names = strings.Split(val, ",")
	return nil
}

func main() {
	qNames := &queuesNames{}
	endPoint := flag.String("endpoint", "", "IP:PORT of SQS service")
	region := flag.String("region", "", "SQS AWS region")
	ssl := flag.Bool("ssl", false, "Use SSL, false by default")
	flag.Var(qNames, "queues", "The queues' names to create, provide one or a comma separated list of them")

	flag.Parse()

	if len(*endPoint) == 0 {
		log.Fatal("-endpoint parameters is required")
	}

	if len(*region) == 0 {
		log.Fatal("-region parameters is required")
	}

	if qNames == nil {
		log.Fatal("-queues parameter is required")
	}

	svc := sqs.New(session.New(aws.NewConfig().WithDisableSSL(!*ssl).WithEndpoint(*endPoint).WithRegion(*region)))

	for _, qn := range qNames.names {
		log.Println("Creating queue", qn)
		createQueue(svc, qn)
	}
}

func createQueue(svc *sqs.SQS, queueName string) {
	params := &sqs.CreateQueueInput{
		QueueName: &queueName,
	}
	_, err := svc.CreateQueue(params)
	if err != nil {
		log.Fatal(err)
	}
}
