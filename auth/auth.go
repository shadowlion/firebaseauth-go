package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) url(endpoint string) string {
	return fmt.Sprintf(
		"https://identitytoolkit.googleapis.com/v1/%s?key=%s",
		endpoint,
		c.ApiKey,
	)
}

type SignUpRequest struct {
	Email             string `json:"email"`
	Password          string `json:"password"`
	ReturnSecureToken bool   `json:"returnSecuretoken"`
}

type SignUpResponse struct {
	IdToken      string `json:"idToken"`
	Email        string `json:"email"`
	RefreshToken string `json:"refreshToken"`
	ExpiresIn    string `json:"expiresIn"`
	LocalId      string `json:"localId"`
	Kind         string `json:"kind"`
}

func (c *Client) SignUp(
	email string,
	password string,
	returnSecureToken bool,
) (*SignUpResponse, error) {
	payload := SignUpRequest{
		Email:             email,
		Password:          password,
		ReturnSecureToken: returnSecureToken,
	}

	jsonData, err := json.Marshal(payload)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(
		http.MethodPost,
		c.url("accounts:signUp"),
		bytes.NewBuffer(jsonData),
	)

	if err != nil {
		return nil, err
	}

	var fullResponse SignUpResponse

	if err := c.sendRequest(req, &fullResponse); err != nil {
		return nil, err
	}

	return &fullResponse, nil
}

type SignInWithEmailAndPasswordRequest struct {
	Email             string `json:"email"`
	Password          string `json:"password"`
	ReturnSecureToken bool   `json:"returnSecureToken"`
}

type SignInWithEmailAndPasswordResponse struct {
	IdToken     string `json:"idToken"`
	Email       string `json:"email"`
	LocalId     string `json:"localId"`
	DisplayName string `json:"displayName"`
	Registered  bool   `json:"registered"`
	Kind        string `json:"kind"`
}

func (c *Client) SignInWithPassword(
	email string,
	password string,
	returnSecureToken bool,
) (*SignInWithEmailAndPasswordResponse, error) {
	payload := SignInWithEmailAndPasswordRequest{
		Email:             email,
		Password:          password,
		ReturnSecureToken: returnSecureToken,
	}

	jsonData, err := json.Marshal(payload)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(
		http.MethodPost,
		c.url("accounts:signInWithPassword"),
		bytes.NewBuffer(jsonData),
	)

	if err != nil {
		return nil, err
	}

	var fullResponse SignInWithEmailAndPasswordResponse

	if err := c.sendRequest(req, &fullResponse); err != nil {
		return nil, err
	}

	return &fullResponse, nil
}

type DeleteAccountRequest struct {
	IdToken string `json:"idToken"`
}

func (c *Client) DeleteAccount(idToken string) error {
	payload := DeleteAccountRequest{
		IdToken: idToken,
	}

	jsonData, err := json.Marshal(payload)

	if err != nil {
		return err
	}

	req, err := http.NewRequest(
		http.MethodPost,
		c.url("accounts:delete"),
		bytes.NewBuffer(jsonData),
	)

	if err != nil {
		return err
	}

	if err := c.sendRequest(req, nil); err != nil {
		return err
	}

	return nil
}
