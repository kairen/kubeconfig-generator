package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/kairen/kubeconfig-generator/pkg/types"
	"github.com/kairen/kubeconfig-generator/pkg/util/kubeconfig"
)

const (
	timeout          = 30 * time.Second
	dialTimeout      = 10 * time.Second
	keepaliveTimeout = 30 * time.Second
	handshakeTimeout = 5 * time.Second
	responseTimeout  = 10 * time.Second
	expectTimeout    = 1 * time.Second
)

type Flags struct {
	URL      string
	DN       string
	Password string
}

type Client struct {
	httpClient *http.Client
	flags      *Flags
}

func NewClient(flags Flags) *Client {
	return &Client{
		flags: &flags,
		httpClient: &http.Client{
			Timeout: timeout,
			Transport: &http.Transport{
				Dial: (&net.Dialer{
					Timeout:   dialTimeout,
					KeepAlive: keepaliveTimeout,
				}).Dial,
				TLSHandshakeTimeout:   handshakeTimeout,
				ResponseHeaderTimeout: responseTimeout,
				ExpectContinueTimeout: expectTimeout,
			},
		},
	}
}

func (c *Client) GenerateKubeconfig(path string) error {
	g, err := c.login()
	if err != nil {
		return err
	}
	if err := kubeconfig.Generate(g, path); err != nil {
		return err
	}
	return nil
}

func (c *Client) login() (*types.Generator, error) {
	u := &types.User{DN: c.flags.DN, Password: c.flags.Password}
	b, err := c.post("/login", u)
	if err != nil {
		return nil, err
	}

	var g types.Generator
	if err := json.Unmarshal(b, &g); err != nil {
		return nil, err
	}
	return &g, nil
}

func (c *Client) post(path string, body interface{}) ([]byte, error) {
	buff := new(bytes.Buffer)
	json.NewEncoder(buff).Encode(body)

	if _, err := url.Parse(c.flags.URL); err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s", c.flags.URL, path), buff)
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return nil, err
	}

	r, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	if r.StatusCode == http.StatusUnauthorized || r.StatusCode == http.StatusBadRequest {
		return nil, fmt.Errorf("%s", http.StatusText(r.StatusCode))
	}
	return b, nil
}
