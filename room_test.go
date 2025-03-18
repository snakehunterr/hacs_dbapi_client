package api_client

import (
	"os"
	"testing"

	types "github.com/snakehunterr/hacs_dbapi_types"
	api_errors "github.com/snakehunterr/hacs_dbapi_types/errors"
)

func Test_room_get_by_id(t *testing.T) {
	var (
		client = NewAPIClient(os.Getenv("HACS_DBAPI_HOST"), os.Getenv("HACS_DBAPI_PORT"))
		r      *types.Room
		res    *types.APIResponse
		err    error
		_c     = &types.Client{
			ID:   1001,
			Name: "foo",
		}
		_r = &types.Room{
			ID:          1,
			ClientID:    _c.ID,
			Area:        10.5,
			PeopleCount: 3,
		}
	)

	if _, err := client.ClientCreate(_c); err != nil {
		t.Fatal("ClientCreate err:", err)
	}
	defer func() {
		if _, err := client.ClientDelete(_c.ID); err != nil {
			t.Fatal("ClientDelete err:", err)
		}
	}()

	if _, err := client.RoomCreate(_r); err != nil {
		t.Fatal("RoomCreate err:", err)
	}

	r, res, err = client.RoomGetByID(_r.ID)
	if err != nil {
		t.Fatal("RoomGetByID err:", err)
	}
	if res != nil {
		t.Fatal("RoomGetByID *types.APIResponse:", res)
	}
	if r.ID != _r.ID {
		t.Fatal("RoomGetByID returning wrong id:", r.ID)
	}
	if r.ClientID != _r.ClientID {
		t.Fatal("RoomGetByID returning wrong client_id:", r.ClientID)
	}
	if r.Area != _r.Area {
		t.Fatal("RoomGetByID returning wrong area:", r.Area)
	}
	if r.PeopleCount != _r.PeopleCount {
		t.Fatal("RoomGetByID returning wrong people_count:", r.PeopleCount)
	}

	_, err = client.RoomDelete(_r.ID)
	if err != nil {
		t.Fatal("RoomDelete err:", err)
	}

	r, res, err = client.RoomGetByID(_r.ID)
	if err != nil {
		t.Fatal("RoomGetByID (after delete) err:", err)
	}
	if res == nil {
		t.Fatal("RoomGetByID (after delete) *types.APIResponse is nil")
	}
	if !api_errors.IsChildErr(res.Error, api_errors.ErrSQLNoRows) {
		t.Fatal("RoomGetByID (after delete) returns wrong *api_errors.APIError:", res.Error)
	}
	if r != nil {
		t.Fatal("RoomGetByID (after delete) returns *types.Room:", r)
	}
}

func Test_room_create(t *testing.T) {
	var (
		client = NewAPIClient(os.Getenv("HACS_DBAPI_HOST"), os.Getenv("HACS_DBAPI_PORT"))
		c      = &types.Client{
			ID:   1,
			Name: "foo",
		}
		r = &types.Room{
			ID:          1,
			ClientID:    c.ID,
			Area:        10.5,
			PeopleCount: 3,
		}
		res *types.APIResponse
		err error
	)

	if _, err := client.ClientCreate(c); err != nil {
		t.Fatal("ClientCreate err:", err)
	}
	defer func() {
		if _, err := client.ClientDelete(c.ID); err != nil {
			t.Fatal("ClientDelete err:", err)
		}
	}()

	res, err = client.RoomCreate(r)
	if err != nil {
		t.Fatal("RoomCreate err:", err)
	}
	if res == nil {
		t.Fatal("RoomCreate *types.APIResponse is nil")
	}
	if res.Error != nil {
		t.Fatal("RoomCreate *APIError:", res.Error)
	}

	res, err = client.RoomCreate(r)
	if err != nil {
		t.Fatal("RoomCreate err:", err)
	}
	if res == nil {
		t.Fatal("RoomCreate *types.APIResponse is nil")
	}
	if !api_errors.IsChildErr(res.Error, api_errors.ErrSQLInternalError) {
		t.Fatal("RoomCreate wrong *APIError:", res.Error)
	}

	res, err = client.RoomDelete(r.ID)
	if err != nil {
		t.Fatal("RoomDelete err:", err)
	}
	if res == nil {
		t.Fatal("RoomDelete *types.APIResponse is nil")
	}
	if res.Error != nil {
		t.Fatal("RoomDelete *APIError:", res.Error)
	}
}

func Test_room_get_all(t *testing.T) {
	var (
		client = NewAPIClient(os.Getenv("HACS_DBAPI_HOST"), os.Getenv("HACS_DBAPI_PORT"))
		c      = &types.Client{
			ID:   1,
			Name: "foo",
		}
		rs = []types.Room{
			{ID: 1, ClientID: c.ID, Area: 10, PeopleCount: 1},
			{ID: 2, ClientID: c.ID, Area: 10, PeopleCount: 1},
			{ID: 3, ClientID: c.ID, Area: 10, PeopleCount: 1},
			{ID: 4, ClientID: c.ID, Area: 10, PeopleCount: 1},
		}
		_rs []types.Room
		res *types.APIResponse
		err error
	)

	if _, err := client.ClientCreate(c); err != nil {
		t.Fatal("ClientCreate err:", err)
	}
	defer func() {
		if _, err := client.ClientDelete(c.ID); err != nil {
			t.Fatal("ClientDelete err:", err)
		}
	}()

	for _, r := range rs {
		if _, err := client.RoomCreate(&r); err != nil {
			t.Fatal("RoomCreate err:", err)
		}
	}

	_rs, res, err = client.RoomGetAll()
	if err != nil {
		t.Fatal("RoomGetAll err:", err)
	}
	if res != nil {
		t.Fatal("RoomGetAll *types.APIResponse:", res)
	}
	if len(_rs) != len(rs) {
		t.Fatalf("RoomGetAll returns %d records, but must return %d", len(_rs), len(rs))
	}

	for _, r := range rs {
		if _, err := client.RoomDelete(r.ID); err != nil {
			t.Fatal("RoomDelete err:", err)
		}
	}
}

func Test_room_get_all_by_client_id(t *testing.T) {
	var (
		client = NewAPIClient(os.Getenv("HACS_DBAPI_HOST"), os.Getenv("HACS_DBAPI_PORT"))
		c1     = &types.Client{
			ID:   1,
			Name: "foo",
		}
		c2 = &types.Client{
			ID:   2,
			Name: "foo",
		}
		rs = []types.Room{
			{ID: 1, ClientID: c1.ID, Area: 10, PeopleCount: 1},
			{ID: 2, ClientID: c2.ID, Area: 10, PeopleCount: 1},
			{ID: 3, ClientID: c1.ID, Area: 10, PeopleCount: 1},
			{ID: 4, ClientID: c2.ID, Area: 10, PeopleCount: 1},
		}
		_rs []types.Room
		res *types.APIResponse
		err error
	)

	for _, c := range []*types.Client{c1, c2} {
		if _, err := client.ClientCreate(c); err != nil {
			t.Fatal("ClientCreate err:", err)
		}
		defer func(id int64) {
			if _, err := client.ClientDelete(id); err != nil {
				t.Fatal("ClientDelete err:", err)
			}
		}(c.ID)
	}

	for _, r := range rs {
		if _, err := client.RoomCreate(&r); err != nil {
			t.Fatal("RoomCreate err:", err)
		}
	}

	for _, c := range []*types.Client{c1, c2} {
		_rs, res, err = client.RoomGetAllByClientID(c.ID)
		if err != nil {
			t.Fatal("RoomGetAll err:", err)
		}
		if res != nil {
			t.Fatal("RoomGetAll *types.APIResponse:", res)
		}
		temp := []types.Room{}
		for _, r := range rs {
			if r.ClientID == c.ID {
				temp = append(temp, r)
			}
		}
		if len(_rs) != len(temp) {
			t.Fatalf("RoomGetAll returns %d records, but must return %d", len(_rs), len(temp))
		}
	}

	for _, r := range rs {
		if _, err := client.RoomDelete(r.ID); err != nil {
			t.Fatal("RoomDelete err:", err)
		}
	}
}

func Test_room_patch(t *testing.T) {
	var (
		client = NewAPIClient(os.Getenv("HACS_DBAPI_HOST"), os.Getenv("HACS_DBAPI_PORT"))
		c1     = &types.Client{
			ID:   1,
			Name: "foo",
		}
		c2 = &types.Client{
			ID:   2,
			Name: "foo",
		}
		r = &types.Room{
			ID:          1,
			ClientID:    c1.ID,
			Area:        10,
			PeopleCount: 1,
		}
		_r  *types.Room
		res *types.APIResponse
		err error
	)

	for _, c := range []*types.Client{c1, c2} {
		if _, err := client.ClientCreate(c); err != nil {
			t.Fatal("ClientCreate err:", err)
		}
		defer func(id int64) {
			if _, err := client.ClientDelete(id); err != nil {
				t.Fatal("ClientDelete err:", err)
			}
		}(c.ID)
	}

	if _, err = client.RoomCreate(r); err != nil {
		t.Fatal("RoomCreate err:", err)
	}
	defer func() {
		if _, err := client.RoomDelete(r.ID); err != nil {
			t.Fatal("RoomDelete err:", err)
		}
	}()

	r.Area = 15
	r.PeopleCount = 3
	r.ClientID = c2.ID

	res, err = client.RoomPatch(r)
	if err != nil {
		t.Fatal("RoomPatch err:", err)
	}
	if res == nil {
		t.Fatal("RoomPatch *types.APIResponse is nil")
	}
	if res.Error != nil {
		t.Fatal("RoomPatch *APIError:", res.Error)
	}

	_r, _, err = client.RoomGetByID(r.ID)
	if err != nil {
		t.Fatal("RoomGetByID err:", err)
	}

	if _r.ClientID != r.ClientID {
		t.Fatalf("RoomPatch wrong ClientID: %d != %d", r.ClientID, _r.ClientID)
	}
	if _r.Area != r.Area {
		t.Fatalf("RoomPatch wrong Area: %f != %f", r.Area, _r.Area)
	}
	if _r.PeopleCount != r.PeopleCount {
		t.Fatalf("RoomPatch wrong PeopleCount: %d != %d", r.PeopleCount, _r.PeopleCount)
	}

	res, err = client.RoomPatch(r)
	if err != nil {
		t.Fatal("RoomPatch err:", err)
	}
	if res == nil {
		t.Fatal("RoomPatch *types.APIResponse is nil")
	}
	if res.Error != nil {
		t.Fatal("RoomPatch *APIError:", res.Error)
	}

	_r, _, err = client.RoomGetByID(r.ID)
	if err != nil {
		t.Fatal("RoomGetByID err:", err)
	}

	if _r.ClientID != r.ClientID {
		t.Fatalf("RoomPatch wrong ClientID: %d != %d", r.ClientID, _r.ClientID)
	}
	if _r.Area != r.Area {
		t.Fatalf("RoomPatch wrong Area: %f != %f", r.Area, _r.Area)
	}
	if _r.PeopleCount != r.PeopleCount {
		t.Fatalf("RoomPatch wrong PeopleCount: %d != %d", r.PeopleCount, _r.PeopleCount)
	}
}
