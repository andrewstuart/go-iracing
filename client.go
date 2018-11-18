package iracing

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
	"time"

	"golang.org/x/net/publicsuffix"
)

// A Client holds state related to iRacing session
type Client struct {
	*http.Client
}

// Login logs in to the iRacing API
func (c *Client) Login(user, password string) error {
	j, err := cookiejar.New(&cookiejar.Options{
		PublicSuffixList: publicsuffix.List,
	})
	if err != nil {
		return err
	}
	c.Client = &http.Client{
		Jar: j,
	}

	v := url.Values{
		"username":  []string{user},
		"password":  []string{password},
		"utcoffset": []string{fmt.Sprintf("%.0f", time.Now().In(time.UTC).Sub(time.Now()).Hours())},
	}

	res, err := c.Post("https://members.iracing.com/membersite/Login", "application/x-www-form-urlencoded", strings.NewReader(v.Encode()))
	if err != nil {
		return err
	}
	return res.Body.Close()
}
