package ServiceAWS

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/aws/aws-sdk-go/aws/awserr"
)

const (
	CharSet = "UTF-8"
)

type EmailData struct {
	Sender string
	Subject string
	HtmlBody string
	TextBody string
	Recipient []string
}

func SendWithSES( data EmailData) error {
	// Get the session.
	sess, err := AWSSession()
	if err != nil {
		return err
	}
	svc := ses.New(sess)

	ToAddresses := []*string{}
	for _, r := range data.Recipient {
		ToAddresses = append(ToAddresses, aws.String(r))
	}
	
	// Assemble the email.
	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			CcAddresses: []*string{
			},
			ToAddresses: ToAddresses,
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: aws.String(CharSet),
					Data:	aws.String(data.HtmlBody),
				},
				Text: &ses.Content{
					Charset: aws.String(CharSet),
					Data:	aws.String(data.TextBody),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String(CharSet),
				Data:	aws.String(data.Subject),
			},
		},
		Source: aws.String(data.Sender),
	}

	// Attempt to send the email.
	_, err = svc.SendEmail(input)
	
	// Display error messages if they occur.
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ses.ErrCodeMessageRejected:
				return errors.New(ses.ErrCodeMessageRejected+aerr.Error())
			case ses.ErrCodeMailFromDomainNotVerifiedException:
				return errors.New(ses.ErrCodeMailFromDomainNotVerifiedException, aerr.Error())
			case ses.ErrCodeConfigurationSetDoesNotExistException:
				return errors.New(ses.ErrCodeConfigurationSetDoesNotExistException, aerr.Error())
			default:
				return errors.New(aerr.Error())
			}
		} else {
			return err
		}
	} else {
		return nil
	}
}


