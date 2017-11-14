package skype

import (
	"bytes"
	"fmt"
	"net/url"
	"strconv"
)

type AuthorizeService service

const (
	authURL = "https://login.microsoftonline.com/botframework.com/oauth2/v2.0/token"
)

type Authorization struct {
	TokenType    string `json:"token_type, omitempty"`
	ExpiresIn    int    `json:"expires_in, omitempty"`
	ExtExpiresIn int    `json:"ext_expires_in, omitempty"`
	AccessToken  string `json:"access_token, omitempty"`
}

func (s *AuthorizeService) Authorize() (*Response, error) {

	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Add("client_id", s.client.ClientID)
	data.Add("client_secret", s.client.ClientSecret)
	data.Add("scope", "https://api.botframework.com/.default")
	req, err := s.client.NewRequest("POST", authURL, bytes.NewBufferString(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Host", "login.microsoftonline.com")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	if err != nil {
		return nil, err
	}

	var auth = new(Authorization)

	resp, err := s.client.Do(req, auth)
	if err != nil {
		return resp, err
	}
	s.client.Token = fmt.Sprintf("%s %s", auth.TokenType, auth.AccessToken)
	return resp, err
}
