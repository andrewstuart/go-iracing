package iracing

import (
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"

	"golang.org/x/net/publicsuffix"
)

// A Client holds state related to iRacing session
type Client struct {
	*http.Client
}

// NewClient returns a Client instance with a cookie jar
func NewClient() (*Client, error) {
	var c Client
	j, err := cookiejar.New(&cookiejar.Options{
		PublicSuffixList: publicsuffix.List,
	})
	if err != nil {
		return nil, err
	}
	c.Client = &http.Client{
		Jar: j,
	}

	return &c, nil
}

// Login logs in to the iRacing API
func (c *Client) Login(user, password string) error {
	v := url.Values{
		"username":  []string{user},
		"password":  []string{password},
		"utcoffset": []string{"0"},
	}

	res, err := c.Post("https://members.iracing.com/membersite/Login", "application/x-www-form-urlencoded", strings.NewReader(v.Encode()))
	if err != nil {
		return err
	}
	return res.Body.Close()
}
