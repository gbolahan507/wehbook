package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {

	server := http.NewServeMux()

	profileName := os.Getenv("AWS_PROFILE")

	server.HandleFunc("/", HandleWebhook)

	sess, err := session.NewSessionWithOptions(session.Options{
		Config:            aws.Config{CredentialsChainVerboseErrors: aws.Bool(true)},
		Profile:           profileName,
		SharedConfigState: session.SharedConfigEnable,
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	client := s3.New(sess)

	err = uploadFile(client, "myfirstbucket507", "./example.txt")

	if err != nil {
		log.Println("Error while uploading the file", err)
	} else {
		log.Println("Successful uploading")
	}

	http.ListenAndServe(":8080", server)

}

func uploadFile(client *s3.S3, bucketName string, fileName string) error {

	file, err := os.Open(fileName)

	if err != nil {
		fmt.Println("Error while opening file")
	}

	putObjectInput := &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String("audit-logs"),
		Body:   file,
	}

	_, err = client.PutObject(putObjectInput)

	if err != nil {
		fmt.Println("Error while opening file")
	}

	return err

}

// func retrieveBucketData(client *s3.S3, bucketName string) {

// 	input := &s3.ListObjectsV2Input{
// 		Bucket: aws.String(bucketName),
// 	}

// 	output, err := client.ListObjectsV2(input)

// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	for _, object := range output.Contents {
// 		fmt.Println(*object.Key)
// 		log.Printf("Name=%s Size=%d", *aws.String(*object.Key), *object.Size)
// 	}

// }
