package endpoints

import (
	"fmt"
	"net/http"
)

type Client struct {
	client  http.Client
	baseURL string
	apiKey  string
	token   string
}

func New(projectReference string, apiKey string) *Client {
	baseURL := fmt.Sprintf("https://%s.supabase.co/auth/v1", projectReference)
	return &Client{
		client:  http.Client{},
		baseURL: baseURL,
		apiKey:  apiKey,
	}
}

func (c Client) WithCustomGoTrueURL(url string) *Client {
	return &Client{
		client:  c.client,
		baseURL: url,
		apiKey:  c.apiKey,
		token:   c.token,
	}
}

func (c Client) WithToken(token string) *Client {
	return &Client{
		client:  c.client,
		baseURL: c.baseURL,
		apiKey:  c.apiKey,
		token:   token,
	}
}

func (c Client) WithClient(client http.Client) *Client {
	return &Client{
		client:  client,
		baseURL: c.baseURL,
		apiKey:  c.apiKey,
		token:   c.token,
	}
}