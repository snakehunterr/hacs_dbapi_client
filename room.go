package api_client

import (
	"errors"
	"fmt"

	types "github.com/snakehunterr/hacs_dbapi_types"
)

func (c APIClient) RoomGetAll() ([]types.Room, error) {
	var rs []types.Room
	err := c.resourceGet(fmt.Sprintf("%s/room/all", c.baseAPIURL), &rs)

	return rs, err
}

func (c APIClient) RoomGetByID(id int64) (*types.Room, error) {
	var r types.Room
	err := c.resourceGet(fmt.Sprintf("%s/room/id/%d", c.baseAPIURL, id), &r)

	if err != nil {
		return nil, err
	} else {
		return &r, nil
	}
}

func (c APIClient) RoomGetAllByClientID(id int64) ([]types.Room, error) {
	var rs []types.Room
	err := c.resourceGet(fmt.Sprintf("%s/room/client/id/%d", c.baseAPIURL, id), &rs)

	return rs, err
}

func (c APIClient) RoomCreate(r *types.Room) error {
	if r == nil {
		return errors.New("*types.Room is nil")
	}

	return c.resourceCreate(
		fmt.Sprintf("%s/room/id/%d", c.baseAPIURL, r.ID),
		formatRoom(r),
	)
}

func (c APIClient) RoomDelete(id int64) error {
	return c.resourceDelete(fmt.Sprintf("%s/room/id/%d", c.baseAPIURL, id))
}

func (c APIClient) RoomPatch(r *types.Room) error {
	if r == nil {
		return errors.New("*types.Room is nil")
	}

	return c.resourcePatch(
		fmt.Sprintf("%s/room/id/%d", c.baseAPIURL, r.ID),
		formatRoom(r),
	)
}
