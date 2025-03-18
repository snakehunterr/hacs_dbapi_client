package api_client

import (
	"os"
	"testing"

	types "github.com/snakehunterr/hacs_dbapi_types"
	api_errors "github.com/snakehunterr/hacs_dbapi_types/errors"
)

func Test_client_get_by_id(t *testing.T) {
	var (
		client = NewAPIClient(os.Getenv("HACS_DBAPI_HOST"), os.Getenv("HACS_DBAPI_PORT"))
		res    *types.APIResponse
		err    error
		c      = &types.Client{ID: 1, Name: "foo", IsAdmin: false}
	)

	res, err = client.ClientCreate(c)
	if err != nil {
		t.Fatal("ClientCreate err:", err)
	}
	if res == nil {
		t.Fatal("ClientCreate *types.APIResponse is nil")
	}
	if res.Error != nil {
		t.Fatal("ClientCreate *APIError:", res.Error)
	}

	_c, res, err := client.ClientGetByID(c.ID)
	if err != nil {
		t.Fatal("ClientGetByID:", err)
	}

	if res != nil {
		t.Fatal("ClientGetByID *types.APIResponse:", res)
	}

	if _c == nil {
		t.Fatal("ClientGetByID *types.Client is nil")
	}

	if c.ID != _c.ID {
		t.Fatal("ClientGetByID returns wrong ID")
	}
	if c.Name != _c.Name {
		t.Fatal("ClientGetByID returns wrong Name")
	}
	if c.IsAdmin != _c.IsAdmin {
		t.Fatal("ClientGetByID returns wrong IsAdmin")
	}

	_, err = client.ClientDelete(c.ID)
	if err != nil {
		t.Fatal("ClientDelete err:", err)
	}

	c, res, err = client.ClientGetByID(c.ID)
	if err != nil {
		t.Fatal("ClientGetByID (non existing id):", err)
	}

	if res == nil {
		t.Fatal("ClientGetByID (non existing id) *types.APIResponse is nil")
	}

	if !api_errors.IsChildErr(res.Error, api_errors.ErrSQLNoRows) {
		t.Fatal("ClientGetByID (non existing id) wrong *APIError:", res.Error)
	}
}

func Test_client_get_all(t *testing.T) {
	var (
		client = NewAPIClient(os.Getenv("HACS_DBAPI_HOST"), os.Getenv("HACS_DBAPI_PORT"))
		cs     []types.Client
		r      *types.APIResponse
		err    error

		c1 = &types.Client{ID: 1, Name: "hahaha"}
		c2 = &types.Client{ID: 2, Name: "hahaha"}
		c3 = &types.Client{ID: 3, Name: "hahaha"}
	)

	_, res, err := client.ClientGetAll()
	if err != nil {
		t.Fatal("ClientGetAll() (no records):", err)
	}
	if res == nil {
		t.Fatal("ClientGetAll() (no records) *types.APIResponse is nil")
	}
	if res.Error != nil && !(api_errors.IsChildErr(res.Error, api_errors.ErrSQLNoRows)) {
		t.Fatal("ClientGetAll() (no records) wrong *APIError:", res.Error)
	}

	for _, c := range []*types.Client{c1, c2, c3} {
		res, err := client.ClientCreate(c)
		if err != nil {
			t.Fatal("ClientCreate:", err)
		}
		if res.Error != nil {
			t.Fatal("ClientCreate *APIError:", res.Error)
		}

		defer func(c *types.Client) {
			res, err := client.ClientDelete(c.ID)
			if err != nil {
				t.Fatal("ClientDelete:", err)
			}
			if res.Error != nil {
				t.Fatal("ClientDelete *APIError:", res.Error)
			}
		}(c)
	}

	// 200
	cs, r, err = client.ClientGetAll()

	if err != nil {
		t.Fatal("ClientGetAll returns err:", err)
	}

	if r != nil {
		t.Fatal("ClientGetAll returns APIResponse:", r)
	}

	if len(cs) != 3 {
		t.Fatal("ClientGetAll should returns []types.Client with len at least 3:", cs)
	}
}

func Test_client_patch(t *testing.T) {
	var (
		client = NewAPIClient(os.Getenv("HACS_DBAPI_HOST"), os.Getenv("HACS_DBAPI_PORT"))
		c      *types.Client
		r      *types.APIResponse
		err    error
		_c     = &types.Client{ID: 1, Name: "hahaha"}
	)

	_, err = client.ClientCreate(_c)
	if err != nil {
		t.Fatal("ClientCreate err:", err)
	}

	c, _, err = client.ClientGetByID(_c.ID)
	if err != nil {
		t.Fatal("ClientGetByID err:", err)
	}

	if c.ID != _c.ID {
		t.Fatalf("ClientPatch %d != %d", _c.ID, c.ID)
	}
	if c.Name != _c.Name {
		t.Fatalf("ClientPatch %s != %s", _c.Name, c.Name)
	}

	_c.Name = "qwerty"
	_c.IsAdmin = true
	r, err = client.ClientPatch(_c)

	if err != nil {
		t.Fatal("ClientPatch err:", err)
	}
	if r == nil {
		t.Fatal("ClientPatch r is nil")
	}
	if r.Error != nil {
		t.Fatal("ClientPatch APIError:", r.Error)
	}

	c, _, err = client.ClientGetByID(_c.ID)
	if err != nil {
		t.Fatal("ClientGetByID err:", err)
	}

	if c.ID != _c.ID {
		t.Fatalf("ClientPatch %d != %d", _c.ID, c.ID)
	}
	if c.Name != _c.Name {
		t.Fatalf("ClientPatch %s != %s", _c.Name, c.Name)
	}
	if c.IsAdmin != _c.IsAdmin {
		t.Fatalf("ClientPatch is_admin %v != %v", _c.IsAdmin, c.IsAdmin)
	}

	_, err = client.ClientDelete(_c.ID)
	if err != nil {
		t.Fatal("ClientDelete err:", err)
	}
}

func Test_client_get_admins(t *testing.T) {
	var (
		client = NewAPIClient(os.Getenv("HACS_DBAPI_HOST"), os.Getenv("HACS_DBAPI_PORT"))
		cs     []types.Client
		r      *types.APIResponse
		err    error
	)

	as := []types.Client{
		{ID: 1, Name: "hahaha", IsAdmin: true},
		{ID: 2, Name: "hahaha", IsAdmin: true},
		{ID: 3, Name: "hahaha", IsAdmin: true},
	}

	for _, c := range as {
		if _, err := client.ClientCreate(&c); err != nil {
			t.Fatal("ClientCreate err:", err)
		}
	}

	cs, r, err = client.ClientGetAdmins()
	if err != nil {
		t.Fatal("ClientGetAdmins err:", err)
	}
	if r != nil {
		t.Fatal("ClientGetAdmins APIResponse:", r)
	}
	if len(cs) < 3 {
		t.Fatal("ClientGetAdmins []types.Client len less than 3")
	}

	for _, c := range as {
		if _, err := client.ClientDelete(c.ID); err != nil {
			t.Fatal("ClientDelete err:", err)
		}
	}
}

func Test_client_get_by_name(t *testing.T) {
	var (
		client = NewAPIClient(os.Getenv("HACS_DBAPI_HOST"), os.Getenv("HACS_DBAPI_PORT"))
		cs     []types.Client
		r      *types.APIResponse
		err    error
		c1     = &types.Client{ID: 10010, Name: "foo"}
		c2     = &types.Client{ID: 10011, Name: "ha"}
		c3     = &types.Client{ID: 10012, Name: "foohafoo"}
		c4     = &types.Client{ID: 10013, Name: "foo ha foo"}
	)

	for _, c := range []*types.Client{c1, c2, c3, c4} {
		_, err = client.ClientCreate(c)
		if err != nil {
			t.Fatal("ClientCreate err:", err)
		}
	}

	cs, r, err = client.ClientGetByName("ha")
	if err != nil {
		t.Fatal("ClientGetByName err:", err)
	}
	if r != nil {
		t.Fatal("ClientGetByName APIResponse:", r)
	}
	if len(cs) < 3 {
		t.Fatal("ClientGetByName []types.Client length less than 3")
	}

	for _, c := range []*types.Client{c1, c2, c3, c4} {
		_, err = client.ClientDelete(c.ID)
		if err != nil {
			t.Fatal("ClientDelete err:", err)
		}
	}
}
