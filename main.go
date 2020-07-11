package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecr"
)


func main() {
	mySession := session.Must(session.NewSession())
	myRegion := "eu-west-2"

	svc := ecr.New(mySession, aws.NewConfig().WithRegion(myRegion))
	input := &ecr.DescribeRepositoriesInput{}

	result, err := svc.DescribeRepositories(input)
	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			switch awsErr.Code() {
			case ecr.ErrCodeServerException:
				fmt.Println(ecr.ErrCodeServerException, awsErr.Error())
			case ecr.ErrCodeInvalidParameterException:
				fmt.Println(ecr.ErrCodeInvalidParameterException, awsErr.Error())
			case ecr.ErrCodeRepositoryNotFoundException:
				fmt.Println(ecr.ErrCodeRepositoryNotFoundException, awsErr.Error())
			default:
				fmt.Println(awsErr.Error())
			}
		} else {
			// Print the error, cast err to err.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	totalImages := len(result.Repositories)
	fmt.Printf("total images: %d\n", totalImages)
	for _, r := range result.Repositories {
		fmt.Println(*r.RepositoryArn)
	}

}

