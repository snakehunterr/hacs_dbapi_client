package api_client

import (
	"errors"
	"fmt"

	types "github.com/snakehunterr/hacs_dbapi_types"
)

func (c APIClient) ClientGetByID(id int64) (*types.Client, *types.APIResponse, error) {
	var client types.Client
	r, err := c.resourceGet(fmt.Sprintf("%s/client/id/%d", c.baseAPIURL, id), &client)

	switch {
	case err != nil:
		return nil, nil, err
	case r != nil:
		return nil, r, nil
	default:
		return &client, nil, nil
	}
}

func (c APIClient) ClientGetAll() ([]types.Client, *types.APIResponse, error) {
	var cs []types.Client
	r, err := c.resourceGet(fmt.Sprintf("%s/client/all", c.baseAPIURL), &cs)

	switch {
	case err != nil:
		return nil, nil, err
	case r != nil:
		return nil, r, nil
	default:
		return cs, nil, nil
	}
}

func (c APIClient) ClientGetAdmins() ([]types.Client, *types.APIResponse, error) {
	var cs []types.Client
	r, err := c.resourceGet(fmt.Sprintf("%s/client/admins", c.baseAPIURL), &cs)

	switch {
	case err != nil:
		return nil, nil, err
	case r != nil:
		return nil, r, nil
	default:
		return cs, nil, nil
	}
}

func (c APIClient) ClientGetByName(name string) ([]types.Client, *types.APIResponse, error) {
	var cs []types.Client
	r, err := c.resourceGet(fmt.Sprintf("%s/client/name/%s", c.baseAPIURL, name), &cs)

	switch {
	case err != nil:
		return nil, nil, err
	case r != nil:
		return nil, r, nil
	default:
		return cs, nil, nil
	}
}

func (c APIClient) ClientCreate(client *types.Client) (*types.APIResponse, error) {
	if client == nil {
		return nil, errors.New("*types.Client is nil")
	}

	return c.resourceCreate(
		fmt.Sprintf("%s/client/id/%d", c.baseAPIURL, client.ID),
		formatClient(client),
	)
}

func (c APIClient) ClientDelete(id int64) (*types.APIResponse, error) {
	return c.resourceDelete(fmt.Sprintf("%s/client/id/%d", c.baseAPIURL, id))
}

func (c APIClient) ClientPatch(client *types.Client) (*types.APIResponse, error) {
	if client == nil {
		return nil, errors.New("*types.Client is nil")
	}

	return c.resourcePatch(
		fmt.Sprintf("%s/client/id/%d", c.baseAPIURL, client.ID),
		formatClient(client),
	)
}
