package service

import (
	"bytes"
	"context"
	"crud_fire/model"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"firebase.google.com/go/auth"
	"gopkg.in/gomail.v2"
)

type ServiceAuth interface {
	Login(ctx context.Context, form model.ModelAuth) (string, error)
	Register(ctx context.Context, form model.ModelAuth) (*auth.UserRecord, error)
	VerifyEmail(ctx context.Context, token string) error
}

type serviceAuth struct {
	authClient   *auth.Client
	apikeys      string
	smtpUsername string
	smtpPassword string
	smtpHost     string
	smtpPort     int
	senderEmail  string
}

func NewServiceAuth(authClient *auth.Client, apikeys, smtpUsername, smtpPassword, smtpHost, senderEmail string, smtpPort int) ServiceAuth {
	return &serviceAuth{
		authClient:   authClient,
		apikeys:      apikeys,
		smtpUsername: smtpUsername,
		smtpPassword: smtpPassword,
		smtpHost:     smtpHost,
		smtpPort:     smtpPort,
		senderEmail:  senderEmail,
	}
}

func (s *serviceAuth) Login(ctx context.Context, form model.ModelAuth) (string, error) {
	// define url
	url := fmt.Sprintf("https://identitytoolkit.googleapis.com/v1/accounts:signInWithPassword?key=%s", s.apikeys)
	payload := map[string]string{
		"email":             form.Email,
		"password":          form.Password,
		"returnSecureToken": "true",
	}

	// convert payload to json
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	// create new request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return "", err
	}
	// set header
	req.Header.Set("Content-Type", "application/json")

	// send request
	client := &http.Client{}

	// get response
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	// close response body
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var errResp map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&errResp); err != nil {
			return "", fmt.Errorf("failid to decode response body: %v", err)
		}
		return "", fmt.Errorf("failid to login: %v", errResp)
	}

	// decode response body
	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	// get id token
	idToken, ok := result["idToken"].(string)
	if !ok {
		return "", fmt.Errorf("invalid id token")
	}

	return idToken, nil
}

func (s *serviceAuth) Register(ctx context.Context, form model.ModelAuth) (*auth.UserRecord, error) {
	// create user
	params := (&auth.UserToCreate{}).Email(form.Email).Password(form.Password)
	userRecord, err := s.authClient.CreateUser(ctx, params)
	if err != nil {
		return nil, err
	}

	// send link email verification
	link, err := s.authClient.EmailVerificationLinkWithSettings(ctx, form.Email, nil)
	if err != nil {
		return nil, err
	}

	// creating message
	m := gomail.NewMessage()
	m.SetHeader("From", s.senderEmail)
	m.SetHeader("To", form.Email)
	m.SetHeader("Subject", "Email Verification")
	m.SetBody("text/plain", fmt.Sprintf("Click this link to verify your email: %v", link))
	m.AddAlternative("text/html", fmt.Sprintf("<p>Please verify your email by clicking the following link: <a href=\"%s\">Verify Email</a></p>", link))

	// sending email
	d := gomail.NewDialer(s.smtpHost, s.smtpPort, s.smtpUsername, s.smtpPassword)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		return nil, err
	}

	// validation email verification
	for {
		veriviedUser, err := s.authClient.GetUserByEmail(ctx, form.Email)
		if err != nil {
			return nil, err
		}

		if veriviedUser.EmailVerified {
			fmt.Printf("Successfully verified email for user: %v\n", veriviedUser.UID)
			return userRecord, nil
		} else {
			fmt.Printf("Waiting for user to verify email: %v\n", veriviedUser.UID)
			time.Sleep(20 * time.Second)
		}
	}
}

func (s *serviceAuth) VerifyEmail(ctx context.Context, token string) error {
	// verify email
	tokenInfo, err := s.authClient.VerifyIDTokenAndCheckRevoked(ctx, token)
	if err != nil {
		return err
	}

	fmt.Printf("Successfully verified email for user: %v\n", tokenInfo.UID)

	return nil
}
