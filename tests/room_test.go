package api_client

import (
	"github.com/snakehunterr/hacs_dbapi_types"
	api_errors "github.com/snakehunterr/hacs_dbapi_types/errors"
	"testing"
)

func float_equal(a, b float64) bool {
	var min, max float64
	if a > b {
		max, min = a, b
	} else {
		max, min = b, a
	}

	return (max - min) < 0.01
}

func Test_room_create_delete(t *testing.T) {
	c := newTestClient(1)
	r := newTestRoom(c.ID, 1)
	var err error

	err = client.RoomCreate(r)
	if err == nil {
		t.Fatal("err is nil")
	}

	client.ClientCreate(c)
	defer client.ClientDelete(c.ID)

	err = client.RoomCreate(r)
	if err != nil {
		t.Fatal(err)
	}
	err = client.RoomCreate(r)
	if err == nil {
		t.Fatal("err is nil")
	}

	err = client.RoomDelete(r.ID)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_room_get_by_id(t *testing.T) {
	var (
		c   = newTestClient(1)
		r   = newTestRoom(1, 1)
		_r  *types.Room
		err error
	)

	client.ClientCreate(c)
	defer client.ClientDelete(c.ID)

	_r, err = client.RoomGetByID(r.ID)
	if err == nil {
		t.Fatal()
	}
	if !(api_errors.IsChildErr(err, api_errors.ErrSQLNoRows)) {
		t.Fatal()
	}

	client.RoomCreate(r)
	defer client.RoomDelete(r.ID)

	_r, err = client.RoomGetByID(r.ID)
	if err != nil {
		t.Fatal(err)
	}
	if _r.ID != r.ID || _r.ClientID != r.ClientID || !float_equal(_r.Area, r.Area) || _r.PeopleCount != r.PeopleCount {
		t.Fatal(_r, r)
	}
}

func Test_room_get_all(t *testing.T) {
	var (
		c  = newTestClient(1)
		rs = []*types.Room{
			newTestRoom(1, 1),
			newTestRoom(1, 2),
			newTestRoom(1, 3),
			newTestRoom(1, 4),
		}
		_rs []types.Room
		err error
	)

	_rs, err = client.RoomGetAll()
	if err == nil {
		t.Fatal(_rs)
	}

	client.ClientCreate(c)
	defer client.ClientDelete(c.ID)

	for _, r := range rs {
		client.RoomCreate(r)
		defer client.RoomDelete(r.ID)
	}

	_rs, err = client.RoomGetAll()
	if err != nil {
		t.Fatal(err)
	}
	if len(_rs) != len(rs) {
		t.Fatal(_rs)
	}
}

func Test_room_get_by_client_id(t *testing.T) {
	var (
		c1 = newTestClient(1)
		c2 = newTestClient(2)
		rs = []*types.Room{
			newTestRoom(1, 1),
			newTestRoom(1, 2),
			newTestRoom(2, 3),
			newTestRoom(2, 4),
			newTestRoom(1, 5),
		}
		_rs []types.Room
		err error
	)

	_rs, err = client.RoomGetAllByClientID(c1.ID)
	if err == nil {
		t.Fatal(_rs)
	}
	_rs, err = client.RoomGetAllByClientID(c2.ID)
	if err == nil {
		t.Fatal(_rs)
	}

	client.ClientCreate(c1)
	defer client.ClientDelete(c1.ID)

	client.ClientCreate(c2)
	defer client.ClientDelete(c2.ID)

	c1_count := 0
	c2_count := 0
	for _, r := range rs {
		switch r.ClientID {
		case 1:
			c1_count++
		case 2:
			c2_count++
		}

		client.RoomCreate(r)
		defer client.RoomDelete(r.ID)
	}

	_rs, err = client.RoomGetAllByClientID(c1.ID)
	if err != nil {
		t.Fatal(err)
	}
	if len(_rs) != c1_count {
		t.Fatal(_rs)
	}

	_rs, err = client.RoomGetAllByClientID(c2.ID)
	if err != nil {
		t.Fatal(err)
	}
	if len(_rs) != c2_count {
		t.Fatal(_rs)
	}
}

func Test_room_patch(t *testing.T) {
	var (
		c   = newTestClient(1)
		r   = newTestRoom(c.ID, 1)
		_r  *types.Room
		err error
	)

	err = client.RoomPatch(r)
	if !(api_errors.IsChildErr(err, api_errors.ErrSQLNoRows)) {
		t.Fatal(err)
	}

	client.ClientCreate(c)
	defer client.ClientDelete(c.ID)

	client.RoomCreate(r)
	defer client.RoomDelete(r.ID)

	r.ClientID = 12345
	err = client.RoomPatch(r)
	if err == nil {
		t.Fatal()
	}
	r.ClientID = c.ID

	check := func() {
		err = client.RoomPatch(r)
		if err != nil {
			t.Fatal(err)
		}

		_r, _ = client.RoomGetByID(r.ID)
		if _r.ID != r.ID || _r.ClientID != r.ClientID || !float_equal(_r.Area, r.Area) || _r.PeopleCount != r.PeopleCount {
			t.Fatal(_r, r)
		}
	}

	r.Area = 10.5
	check()

	r.PeopleCount = 5
	r.Area = 22
	check()

	r.PeopleCount = 3
	check()
}
