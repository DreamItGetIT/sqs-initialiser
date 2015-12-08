package test

import (
	"errors"
	"fmt"
	"log"
	"net"
	"regexp"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

const sqsPort = 4568

var svc *sqs.SQS
var addrRegExp = regexp.MustCompile(`([\d.]+)/[\d]+`)

func init() {
	dockerIP, err := getDockerIP()
	if err != nil {
		log.Fatalf("Error when getting docker ip address: %s", err)
	}

	svc = sqs.New(session.New(aws.NewConfig().WithDisableSSL(true).WithEndpoint(fmt.Sprintf("%s:%d", dockerIP, sqsPort)).WithRegion("eu-west=1")))
}

func TestQueuesExist(t *testing.T) {
	expectedQueues := map[string]struct{}{
		fmt.Sprintf("http://0.0.0.0:%d/test1", sqsPort): struct{}{},
		fmt.Sprintf("http://0.0.0.0:%d/test2", sqsPort): struct{}{},
		fmt.Sprintf("http://0.0.0.0:%d/test3", sqsPort): struct{}{},
	}

	queues, err := svc.ListQueues(nil)
	if err != nil {
		t.Fatal(err)
	}
	if queues == nil {
		t.Fatal("Queues lists is nil")
	}

	for _, q := range queues.QueueUrls {
		if _, ok := expectedQueues[*q]; ok == false {
			t.Fatalf("Queue with URL %s wasn't expected", *q)
		}
	}

	if len(queues.QueueUrls) != len(expectedQueues) {
		t.Fatalf("The number of existent queues (%d) aren't the expected number (%d)", len(queues.QueueUrls), len(expectedQueues))
	}
}

func getDockerIP() (string, error) {
	iface, err := net.InterfaceByName("docker0")
	if err != nil {
		return "", err
	}

	addrs, err := iface.Addrs()
	if err != nil {
		return "", err
	}

	if len(addrs) == 0 {
		return "", errors.New("docker0 network interface doesn't have any address")
	}

	addrMatches := addrRegExp.FindAllStringSubmatch(addrs[0].String(), -1)
	if addrMatches == nil {
		return "", errors.New("docker0 network interface doesn't have any IP address")
	}

	return addrMatches[0][1], nil
}
