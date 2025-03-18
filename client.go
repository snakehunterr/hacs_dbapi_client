package api_client

import (
	"errors"
	"fmt"

	types "github.com/snakehunterr/hacs_dbapi_types"
)

func (c APIClient) ClientGetByID(id int64) (*types.Client, error) {
	var client types.Client
	err := c.resourceGet(fmt.Sprintf("%s/client/id/%d", c.baseAPIURL, id), &client)

	if err != nil {
		return nil, err
	} else {
		return &client, err
	}
}

func (c APIClient) ClientGetAll() ([]types.Client, error) {
	var cs []types.Client
	err := c.resourceGet(fmt.Sprintf("%s/client/all", c.baseAPIURL), &cs)

	return cs, err
}

func (c APIClient) ClientGetAdmins() ([]types.Client, error) {
	var cs []types.Client
	err := c.resourceGet(fmt.Sprintf("%s/client/admins", c.baseAPIURL), &cs)

	return cs, err
}

func (c APIClient) ClientGetByName(name string) ([]types.Client, error) {
	var cs []types.Client
	err := c.resourceGet(fmt.Sprintf("%s/client/name/%s", c.baseAPIURL, name), &cs)

	return cs, err
}

func (c APIClient) ClientCreate(client *types.Client) error {
	if client == nil {
		return errors.New("*types.Client is nil")
	}

	return c.resourceCreate(
		fmt.Sprintf("%s/client/id/%d", c.baseAPIURL, client.ID),
		formatClient(client),
	)
}

func (c APIClient) ClientDelete(id int64) error {
	return c.resourceDelete(fmt.Sprintf("%s/client/id/%d", c.baseAPIURL, id))
}

func (c APIClient) ClientPatch(client *types.Client) error {
	if client == nil {
		return errors.New("*types.Client is nil")
	}

	return c.resourcePatch(
		fmt.Sprintf("%s/client/id/%d", c.baseAPIURL, client.ID),
		formatClient(client),
	)
}
