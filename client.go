package releaser

import (
	"net/http"
	"net/url"
)

type Client struct {
	*http.Client

	baseURL *url.URL

	debug bool
}

type optionFn func(*Client)

// WithDebug enables debug output while interacting with MIPS
func WithDebug() optionFn {
	return func(c *Client) {
		c.debug = true
	}
}

// New returns a Github api client
func newClient(options ...optionFn) (*Client, error) {
	u, _ := url.Parse("https://api.github.com/")

	c := &Client{
		Client:  http.DefaultClient,
		baseURL: u,
	}

	for _, optionFn := range options {
		optionFn(c)
	}

	return c, nil
}
