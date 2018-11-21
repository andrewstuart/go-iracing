package iracing

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"path"
	"strings"

	"astuart.co/clyde"
	cookiejar "astuart.co/persistent-cookiejar"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/pkg/errors"
	"golang.org/x/net/publicsuffix"
)

// LoginURL is the URL for Login
var LoginURL = "https://members.iracing.com/membersite/Login"

// A Client holds state related to iRacing session
type Client struct {
	*http.Client
}

func (c *Client) JSON(u string, out interface{}) error {
	res, err := c.Get(u)
	if err != nil {
		return errors.Wrap(err, "error performing get")
	}
	defer res.Body.Close()

	return errors.Wrap(json.NewDecoder(res.Body).Decode(&out), "error decoding JSON")
}

var home string

func init() {
	if hd, err := homedir.Dir(); err == nil {
		home = hd
	}
}

// NewClient returns a Client instance with a cookie jar
func NewClient() (*Client, error) {
	var c Client
	o := &cookiejar.Options{
		PublicSuffixList: publicsuffix.List,
		Filter:           cookiejar.AllowAllFilter,
	}
	if home != "" {
		o.Filename = path.Join(home, ".config", "iracing.cookies")
	}

	j, err := cookiejar.New(o)
	if err != nil {
		return nil, err
	}

	c.Client = &http.Client{
		Jar: j,
		Transport: clyde.RoundTripFunc(func(req *http.Request) (*http.Response, error) {
			defer func() {
				if err := j.Save(); err != nil {
					log.Println("error saving ", err)
				}
			}()

			return http.DefaultTransport.RoundTrip(req)
		}),
	}

	return &c, nil
}

// Login logs in to the iRacing API
func (c *Client) Login(user, password string) error {
	u, err := url.Parse(LoginURL)
	if err != nil {
		return errors.Wrap(err, "could not parse login URL")
	}
	cookies := c.Jar.Cookies(u)
	if len(cookies) > 0 {
		_, err := c.GetDriverCounts()
		if err == nil {
			return nil
		}
	}

	v := url.Values{
		"username":  []string{user},
		"password":  []string{password},
		"AUTOLOGIN": []string{"on"},
		"utcoffset": []string{"0"},
	}

	res, err := c.Post(LoginURL, "application/x-www-form-urlencoded", strings.NewReader(v.Encode()))
	if err != nil {
		return err
	}
	return res.Body.Close()
}
