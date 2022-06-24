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

type SignInWithEmailAndPasswordRequest struct {
	Email             string `json:"email"`
	Password          string `json:"password"`
	ReturnSecureToken bool   `json:"returnSecureToken"`
}

type SignInWithEmailAndPasswordResponse struct {
	IdToken      string `json:"idToken"`
	Email        string `json:"email"`
	RefreshToken string `json:"refreshToken"`
	ExpiresIn    string `json:"expiresIn"`
	LocalId      string `json:"localId"`
	Registered   bool   `json:"registered"`
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
