package ServiceAWS


import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)


func AWSSession(config ...aws.Config) (*session.Session, error) {
	DefaultConfig := aws.Config{
		Region:aws.String("ap-northeast-1"),
	}

	if len(config) == 0 {
		return session.NewSession(&DefaultConfig)
	} else {
		return session.NewSession(&config[0])
	}
}

