package api_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	types "github.com/snakehunterr/hacs_dbapi_types"
)

type Form = map[string]string

func NewAPIClient(host string, port string) APIClient {
	if host == "" {
		panic("Empty arg: host")
	}
	if port == "" {
		panic("Empty arg: port")
	}
	return APIClient{
		HTTPClient: &http.Client{},
		apiHost:    host,
		apiPort:    port,
		baseAPIURL: fmt.Sprintf("http://%s:%s/api", host, port),
		Debug:      false,
	}
}

type APIClient struct {
	HTTPClient *http.Client
	Debug      bool
	apiHost    string
	apiPort    string
	baseAPIURL string
}

func (c *APIClient) SetDebug(mode bool) {
	c.Debug = mode
}

func (c APIClient) newFormRequest(method string, apiurl string, form Form) (*http.Request, error) {
	var vs = url.Values{}

	for k, v := range form {
		vs.Set(k, v)
	}

	req, err := http.NewRequest(
		method,
		apiurl,
		bytes.NewBufferString(vs.Encode()),
	)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	return req, nil
}

func (c APIClient) resourceDelete(url string) (*types.APIResponse, error) {
	req, err := http.NewRequest(
		http.MethodDelete,
		url,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("http.NewRequest(): %w", err)
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("c.HTTPClient.Do(): %w", err)
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	var r types.APIResponse

	if err := decoder.Decode(&r); err != nil {
		return nil, fmt.Errorf("decoder.Decode(): %w", err)
	}

	return &r, nil
}

func (c APIClient) resourcePatch(url string, form Form) (*types.APIResponse, error) {
	req, err := c.newFormRequest(http.MethodPatch, url, form)
	if err != nil {
		return nil, fmt.Errorf("c.newFormRequest(): %w", err)
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("c.HTTPClient.Do(): %w", err)
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	var r types.APIResponse

	if err := decoder.Decode(&r); err != nil {
		return nil, fmt.Errorf("c.HTTPClient.Do(): %w", err)
	}

	return &r, nil
}

func (c APIClient) resourceGet(url string, dest any) (*types.APIResponse, error) {
	res, err := c.HTTPClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("c.HTTPClient.Get(): %w", err)
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)

	if res.StatusCode != 200 {
		var r types.APIResponse

		if err := decoder.Decode(&r); err != nil {
			return nil, fmt.Errorf("decoder.Decode(): %w", err)
		}

		return &r, nil
	}

	if err := decoder.Decode(dest); err != nil {
		return nil, fmt.Errorf("decoder.Decode(): %w", err)
	}

	return nil, nil
}

func (c APIClient) resourceCreate(url string, form Form) (*types.APIResponse, error) {
	req, err := c.newFormRequest(
		http.MethodPost,
		url,
		form,
	)
	if err != nil {
		return nil, fmt.Errorf("c.NewFormRequest(): %w", err)
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("c.HTTPClient.Do(): %w", err)
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	var r types.APIResponse

	if err := decoder.Decode(&r); err != nil {
		return nil, fmt.Errorf("decoder.Decode(): %w", err)
	}

	return &r, nil
}
