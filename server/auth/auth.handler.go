package auth

import (
	"context"

	"github.com/clerk/clerk-sdk-go/v2"
)

var secrets struct {
	ClientSecretKey string
}

type Service struct {
	client clerk.Client
}

func initServices() (*Service, error) {
	return nil, nil
}

type UserData struct {
	ID                    string               `json:"id"`
	Username              *string              `json:"username"`
	ProfileImageURL       string               `json:"profile_image_url"`
	PrimaryEmailAddressID *string              `json:"primary_email_address_id"`
	EmailAddresses        []clerk.EmailAddress `json:"email_addresses"`
}

func (s *Service) AuthHandler(ctx context.Context, token string) {
}
