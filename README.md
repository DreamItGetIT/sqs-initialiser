SQL Initialiser
===============

Command line tool to be able to do a few initialisation tasks on Amazon Simple Queue Service (a.k.a SQS).


## Tools

So far __only allows to create queues__, nothing more.

The tools is available as `digit/sqs-initialiser` docker image, but you can also build a binary if you have [Golang](https://golang.org/) installed and cloned this project in a directory which follow the [Golang Workspaces](https://golang.org/doc/code.html#Workspaces)

This is a command line tool build with [Golang flag package](https://golang.org/pkg/flag/) so you can execute it passing `-h` or `-help` to get the required options which so far the output is

```
Usage of ./build/sqsinit:
  -endpoint string
        IP:PORT of SQS service
  -queues value
        The queues' names to create, provide one or a comma separated list of them
  -region string
        SQS AWS region
  -ssl
        Use SSL, false by default
```

In [makefile](https://github.com/DreamItGetIT/sqs-initialiser/blob/master/makefile) we use the tool to create queues as it is, and thereafter we run the test to check that those queues exexist, so you can find there several examples, nonetheless here you can see one of those executions

```
@docker run --rm -e AWS_ACCESS_KEY_ID=DOESNOTMATTER -e AWS_SECRET_ACCESS_KEY=doesnotmatter digit/sqs-initialiser -endpoint $(DOCKER_IP):4568 -region eu-west-1 -ssl=false -queues "test1,test2"
```

## Contribution

We are open to accept pull request or suggestion to add more commands.

If you submit a pull request, please update [makefile](https://github.com/DreamItGetIT/sqs-initialiser/blob/master/makefile) and make sure that `full` target (the default target) runs all the tests without requiring to run anything else, for example development environment, etc.

When we merge an update of the tool to master, we'll make an update of the docker image.


## LICENSE

The MIT license, read LICENSE file for more information.
