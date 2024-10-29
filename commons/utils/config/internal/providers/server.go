package ConfigCloud

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/url"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

const BaseUrl = "s3://kod-configurations/"
const region = "ap-south-1"

var DevEnvironmentError = errors.New("not fetching Spring Cloud configuration for dev environment")

func GetCloud(fileUrl string) (ConfigResponse, error) {
	var springResponse ConfigResponse
	u, err := url.Parse(fileUrl)
	if err != nil {
		return ConfigResponse{}, err
	}
	bucketName := u.Host
	fileKey := u.Path

	awsSession, err := session.NewSession(&aws.Config{Region: aws.String(region)})
	if err != nil {
		return ConfigResponse{}, err
	}

	objectOutput, err := s3.New(awsSession).GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(fileKey),
	})
	if err != nil {
		return ConfigResponse{}, err
	}

	jsonArray, err := io.ReadAll(objectOutput.Body)
	if err != nil {
		return ConfigResponse{}, err
	}

	jsonString := fmt.Sprintf("%s", jsonArray)
	err = json.NewDecoder(strings.NewReader(jsonString)).Decode(&springResponse)
	if err != nil {
		return ConfigResponse{}, err
	}

	return springResponse, nil
}
