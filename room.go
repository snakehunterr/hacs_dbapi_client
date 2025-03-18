package api_client

// package api_client

// import (
// 	"errors"
// 	"fmt"

// 	types "github.com/snakehunterr/hacs_dbapi_types"
// )

// func (c APIClient) RoomGetAll() ([]types.Room, *types.APIResponse, error) {
// 	var rs []types.Room

// 	res, err := c.resourceGet(fmt.Sprintf("%s/room/all", c.baseAPIURL), &rs)
// 	return rs, res, err
// }

// func (c APIClient) RoomGetByID(id int64) (*types.Room, *types.APIResponse, error) {
// 	var r types.Room

// 	res, err := c.resourceGet(fmt.Sprintf("%s/room/id/%d", c.baseAPIURL, id), &r)
// 	switch {
// 	case err != nil:
// 		return nil, nil, err
// 	case res != nil:
// 		return nil, res, nil
// 	default:
// 		return &r, nil, nil
// 	}
// }

// func (c APIClient) RoomGetAllByClientID(id int64) ([]types.Room, *types.APIResponse, error) {
// 	var rs []types.Room

// 	res, err := c.resourceGet(fmt.Sprintf("%s/room/client/id/%d", c.baseAPIURL, id), &rs)
// 	switch {
// 	case err != nil:
// 		return nil, nil, err
// 	case res != nil:
// 		return nil, res, nil
// 	default:
// 		return rs, nil, nil
// 	}
// }

// func (c APIClient) RoomCreate(r *types.Room) (*types.APIResponse, error) {
// 	if r == nil {
// 		return nil, errors.New("*types.Room is nil")
// 	}

// 	return c.resourceCreate(
// 		fmt.Sprintf("%s/room/id/%d", c.baseAPIURL, r.ID),
// 		formatRoom(r),
// 	)
// }

// func (c APIClient) RoomDelete(id int64) (*types.APIResponse, error) {
// 	return c.resourceDelete(fmt.Sprintf("%s/room/id/%d", c.baseAPIURL, id))
// }

// func (c APIClient) RoomPatch(r *types.Room) (*types.APIResponse, error) {
// 	if r == nil {
// 		return nil, errors.New("*types.Room is nil")
// 	}

// 	return c.resourcePatch(
// 		fmt.Sprintf("%s/room/id/%d", c.baseAPIURL, r.ID),
// 		formatRoom(r),
// 	)
// }
