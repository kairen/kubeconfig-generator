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

	"github.com/inwinstack/kubeconfig-generator/pkg/types"
	"github.com/inwinstack/kubeconfig-generator/pkg/util/kubeconfig"
)

const (
	timeout          = 30 * time.Second
	dialTimeout      = 10 * time.Second
	keepaliveTimeout = 30 * time.Second
	handshakeTimeout = 5 * time.Second
	responseTimeout  = 10 * time.Second
	expectTimeout    = 1 * time.Second
)

type Client struct {
	httpClient *http.Client
	url        string
	user       *types.User
}

func NewClient(url, dn, password string) *Client {
	c := &Client{url: url, user: &types.User{DN: dn, Password: password}}
	c.httpClient = &http.Client{
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
	}
	return c
}

func (c *Client) GenerateKubeconfig(path string) error {
	g, err := c.login()
	if err != nil {
		return err
	}
	if err := kubeconfig.Generate(g, path); err != nil {
		return err
	}
	fmt.Printf("Generate the Kubernetes config to `%s` path.\n", path)
	return nil
}

func (c *Client) login() (*types.Generator, error) {
	buff := new(bytes.Buffer)
	json.NewEncoder(buff).Encode(c.user)

	if _, err := url.Parse(c.url); err != nil {
		return nil, err
	}

	newURL := fmt.Sprintf("%s/login", c.url)
	req, err := http.NewRequest("POST", newURL, buff)
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return nil, err
	}

	r, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	var g types.Generator
	b, _ := ioutil.ReadAll(r.Body)
	if err := json.Unmarshal(b, &g); err != nil {
		return nil, err
	}

	if r.StatusCode != 200 {
		return nil, fmt.Errorf("the status is \"%s\"", g.Status)
	}
	return &g, nil
}
