package services

import (
	"context"
	"encoding/base64"
	"fmt"
	"kingcom_api/internal/lib"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

type mailService struct {
	env    *lib.Env
	logger *lib.Logger
}

type MailService interface {
	createGoogleOauth2Config() (*oauth2.Config, error)
	getGoogleAccessToken(ctx context.Context, config *oauth2.Config) (*oauth2.Token, error)
	sendEmail(ctx context.Context, to, subject, body string) error
	SendVerificationEmail(ctx context.Context, params VerificationParams) error
	SendPasswordResetEmail(ctx context.Context, params PasswordResetParams) error
}

func NewMailService(
	env *lib.Env,
	logger *lib.Logger,
) MailService {
	return &mailService{
		env:    env,
		logger: logger,
	}
}

func (s *mailService) createGoogleOauth2Config() (*oauth2.Config, error) {
	credentials := fmt.Sprintf(`{
		"installed": {
			"client_id": "%s",
			"project_id": "%s",
			"auth_uri": "https://accounts.google.com/o/oauth2/auth",
			"token_uri": "https://oauth2.googleapis.com/token",
			"client_secret": "%s",
			"redirect_uris": ["%s"]
		}
	}`,
		s.env.GoogleOAuth2.ClientID,
		s.env.GoogleOAuth2.ProjectID,
		s.env.GoogleOAuth2.ClientSecret,
		s.env.AppUrl,
	)
	cfg, err := google.ConfigFromJSON([]byte(credentials), gmail.GmailSendScope)
	if err != nil {
		s.logger.Error("Error parsing Oauth config : %v", err)
		return nil, err
	}
	return cfg, err
}

func (s *mailService) getGoogleAccessToken(ctx context.Context, config *oauth2.Config) (*oauth2.Token, error) {
	token := &oauth2.Token{RefreshToken: s.env.GoogleOAuth2.RefreshToken}
	tokenSource := config.TokenSource(ctx, token)
	newToken, err := tokenSource.Token()
	return newToken, err
}

func (s *mailService) sendEmail(ctx context.Context, to, subject, body string) error {
	config, err := s.createGoogleOauth2Config()
	if err != nil {
		return err
	}
	token, err := s.getGoogleAccessToken(ctx, config)
	if err != nil {
		return err
	}
	client := config.Client(ctx, token)
	service, err := gmail.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return err
	}
	_, err = service.Users.Messages.Send("me", buildMessage(to, subject, body)).Do()
	return err
}

func (s *mailService) SendVerificationEmail(ctx context.Context, params VerificationParams) error {
	subject := "Email verification"
	body := fmt.Sprintf("Hello %s.\nThis is your verification code: %s", params.Name, params.Code)
	return s.sendEmail(ctx, params.Email, subject, body)
}

func (s *mailService) SendPasswordResetEmail(ctx context.Context, params PasswordResetParams) error {
	subject := "Password reset"
	body := fmt.Sprintf("Hello %s.\nClick here to reset your password: %s", params.Name, params.Token)
	return s.sendEmail(ctx, params.Email, subject, body)
}

func buildMessage(to, subject, body string) *gmail.Message {
	raw := fmt.Sprintf("To: %s\r\nSubject: %s\r\n\r\n%s", to, subject, body)
	return &gmail.Message{Raw: base64.URLEncoding.EncodeToString([]byte(raw))}
}

type SendEmailParams struct {
	Name  string
	Email string
}
type VerificationParams struct {
	SendEmailParams
	Code string
}
type PasswordResetParams struct {
	SendEmailParams
	Token string
}
