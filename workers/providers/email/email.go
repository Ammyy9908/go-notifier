package email

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/aws/aws-sdk-go-v2/service/ses/types"
	"github.com/aws/aws-sdk-go/aws"
)

type EmailProvider interface {
	SendEmail(recipientEmail, subject, body string) error
}

type EmailConfig struct {
	SenderEmail string // The default sender email address
	Region      string // Region, mainly for SES
	APIKey      string // API key, if required by the provider
	APISecret   string // API secret, if required
}
type BravoProvider struct {
	Config EmailConfig
}

func NewBravoProvider(config EmailConfig) *BravoProvider {
	return &BravoProvider{Config: config}
}

func (b *BravoProvider) SendEmail(recipientEmail, subject, body string) error {
	// Placeholder for actual Bravo API call
	fmt.Printf("Sending email via Bravo to %s\nSubject: %s\nBody:\n%s\n", recipientEmail, subject, body)
	// Implement the actual Bravo send logic here

	return nil // Return an error if sending fails
}

type SESProvider struct {
	client *ses.Client
	Config EmailConfig
}

func NewSESProvider(cfg EmailConfig) *SESProvider {
	// Initialize and configure SES client if necessary
	awsConfig, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(cfg.Region))
	if err != nil {
		log.Fatalf("unable to load AWS SDK config, %v", err)
	}

	client := ses.NewFromConfig(awsConfig)
	return &SESProvider{Config: cfg, client: client}
}

func (s *SESProvider) SendEmail(recipientEmail, subject, body string) error {
	input := &ses.SendEmailInput{
		Source: aws.String(s.Config.SenderEmail),
		Destination: &types.Destination{
			ToAddresses: []string{recipientEmail},
		},
		Message: &types.Message{
			Subject: &types.Content{
				Data: aws.String(subject),
			},
			Body: &types.Body{
				Text: &types.Content{
					Data: aws.String(body),
				},
			},
		},
	}
	_, err := s.client.SendEmail(context.TODO(), input)
	if err != nil {
		return fmt.Errorf("failed to send email via SES: %w", err)
	}

	fmt.Printf("Email sent via SES to %s\n", recipientEmail)
	return nil
}

const (
	SESProviderType   = "ses"
	BravoProviderType = "bravo"
)

// EmailProviderFactory returns an EmailProvider based on the given provider type.
func EmailProviderFactory(providerType string, config EmailConfig) (EmailProvider, error) {
	switch providerType {
	case SESProviderType:
		return NewSESProvider(config), nil
	case BravoProviderType:
		return NewBravoProvider(config), nil
	default:
		return nil, fmt.Errorf("unsupported email provider type: %s", providerType)
	}
}
