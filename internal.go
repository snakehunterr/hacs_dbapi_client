package api_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	types "github.com/snakehunterr/hacs_dbapi_types"
)

type Form = map[string]string

func New(host string, port string) APIClient {
	if host == "" {
		panic("Empty arg: host")
	}
	if port == "" {
		panic("Empty arg: port")
	}
	return APIClient{
		HTTPClient: &http.Client{},
		baseAPIURL: fmt.Sprintf("http://%s:%s/api", host, port),
	}
}

type APIClient struct {
	HTTPClient *http.Client
	baseAPIURL string
}

func (c APIClient) Foo() {}

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

func (c APIClient) request(req *http.Request) error {
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("c.HTTPClient.Do(): %w", err)
	}
	defer res.Body.Close()

	bs, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("io.ReadAll() err: %w", err)
	}

	var r types.APIResponse

	if err := json.Unmarshal(bs, &r); err != nil {
		return fmt.Errorf("json.Unmarshal() err: %w, res.Body(): %s", err, bs)
	}

	if r.Error != nil {
		return r.Error
	}

	return nil
}

func (c APIClient) requestWithDecode(req *http.Request, dest any) error {
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("c.HTTPClient.Do(): %w", err)
	}
	defer res.Body.Close()

	bs, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("io.ReadAll() err: %w", err)
	}

	if res.StatusCode >= 400 {
		var r types.APIResponse

		if err := json.Unmarshal(bs, &r); err != nil {
			return fmt.Errorf("json.Unmarshal() err: %w, res.Body(): %s", err, bs)
		}

		if r.Error != nil {
			return r.Error
		}

		return nil
	}

	if err := json.Unmarshal(bs, dest); err != nil {
		return fmt.Errorf("json.Unmarshal() err: %w, res.Body(): %s", err, bs)
	}

	return nil
}

func (c APIClient) resourcePatch(url string, form Form) error {
	req, err := c.newFormRequest(http.MethodPatch, url, form)
	if err != nil {
		return fmt.Errorf("c.newFormRequest(): %w", err)
	}

	return c.request(req)
}

func (c APIClient) resourceGet(url string, dest any) error {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("http.NewRequest(): %w", err)
	}

	return c.requestWithDecode(req, dest)
}

func (c APIClient) resourcePost(url string, form Form, dest any) error {
	req, err := c.newFormRequest(http.MethodPost, url, form)
	if err != nil {
		return fmt.Errorf("c.newFormRequest(): %w", err)
	}

	return c.requestWithDecode(req, dest)
}

func (c APIClient) resourceCreate(url string, form Form) error {
	req, err := c.newFormRequest(http.MethodPost, url, form)
	if err != nil {
		return fmt.Errorf("c.NewFormRequest(): %w", err)
	}

	return c.request(req)
}

func (c APIClient) resourceCreateWithDecode(url string, form Form, dest any) error {
	req, err := c.newFormRequest(http.MethodPost, url, form)
	if err != nil {
		return fmt.Errorf("c.NewFormRequest(): %w", err)
	}

	return c.requestWithDecode(req, dest)
}

func (c APIClient) resourceDelete(url string) error {
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return fmt.Errorf("http.NewRequest(): %w", err)
	}

	return c.request(req)
}
