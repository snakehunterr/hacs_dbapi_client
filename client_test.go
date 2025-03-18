package api_client

import (
	"os"
	"strings"
	"testing"

	types "github.com/snakehunterr/hacs_dbapi_types"
	api_errors "github.com/snakehunterr/hacs_dbapi_types/errors"
)

func newTestAPIClient() APIClient {
	return New(os.Getenv("DBAPI_SERVER_HOST"), os.Getenv("DBAPI_SERVER_PORT"))
}

func Test_client_create_delete(t *testing.T) {
	var (
		client = newTestAPIClient()
		c      = &types.Client{
			ID:   1,
			Name: "foo",
		}
		err error
	)

	err = client.ClientCreate(c)
	if err != nil {
		t.Fatal("ClientCreate err:", err)
	}

	err = client.ClientCreate(c)
	if err == nil {
		t.Fatal("ClientCreate (dublicating) err is nil")
	}

	err = client.ClientDelete(c.ID)
	if err != nil {
		t.Fatal("ClientDelete err:", err)
	}

	c.Name = ""
	err = client.ClientCreate(c)
	if err == nil {
		t.Fatal("ClientCreate (missing name) err is nil")
	}
	if !api_errors.IsChildErr(err, api_errors.ErrEmptyParam) {
		t.Fatal("ClientCreate (missing name) return wrong err:", err)
	}
}

func Test_client_get_by_id(t *testing.T) {
	var (
		client = newTestAPIClient()
		c      = &types.Client{
			ID:   1,
			Name: "foo",
		}
		_c  *types.Client
		err error
	)

	_c, err = client.ClientGetByID(c.ID)
	if err == nil {
		t.Fatal("ClientGetByID (non existing) err is nil")
	}
	if !api_errors.IsChildErr(err, api_errors.ErrSQLNoRows) {
		t.Fatal("ClientGetByID (non existing) returns wrong err:", err)
	}

	client.ClientCreate(c)
	defer client.ClientDelete(c.ID)

	_c, err = client.ClientGetByID(c.ID)
	if err != nil {
		t.Fatal("ClientGetByID err:", err)
	}

	if _c.ID != c.ID || _c.Name != c.Name || _c.IsAdmin != c.IsAdmin {
		t.Fatalf("ClientGetByID returns wrong client: %#v", _c)
	}
}

func Test_client_get_all(t *testing.T) {
	client := newTestAPIClient()
	cs := []types.Client{
		{ID: 1, Name: "foo"},
		{ID: 2, Name: "foo"},
		{ID: 3, Name: "foo"},
		{ID: 4, Name: "foo"},
	}

	for _, c := range cs {
		client.ClientCreate(&c)
	}

	_cs, err := client.ClientGetAll()
	if err != nil {
		t.Fatal("ClientGetAll err:", err)
	}
	if len(_cs) != len(cs) {
		t.Fatalf("ClientGetAll len(_cs) [%d] != len(cs) [%d]", len(_cs), len(cs))
	}

	for _, c := range cs {
		client.ClientDelete(c.ID)
	}

	_cs, err = client.ClientGetAll()
	if !(api_errors.IsChildErr(err, api_errors.ErrSQLNoRows)) {
		t.Fatal("ClientGetAll (no records) wrong err:", err)
	}
}

func Test_client_get_admins(t *testing.T) {
	client := newTestAPIClient()
	cs := []types.Client{
		{ID: 1, Name: "foo", IsAdmin: true},
		{ID: 2, Name: "foo", IsAdmin: true},
		{ID: 3, Name: "foo", IsAdmin: false},
		{ID: 4, Name: "foo", IsAdmin: false},
	}

	admin_count := 0
	for _, c := range cs {
		if c.IsAdmin {
			admin_count++
		}
		client.ClientCreate(&c)
	}

	_cs, err := client.ClientGetAdmins()
	if err != nil {
		t.Fatal("ClientGetAdmins err:", err)
	}
	if len(_cs) != admin_count {
		t.Fatalf("ClientGetAdmins len(_cs) [%d] != admin_count [%d]", len(_cs), admin_count)
	}

	for _, c := range cs {
		client.ClientDelete(c.ID)
	}

	_cs, err = client.ClientGetAdmins()
	if !(api_errors.IsChildErr(err, api_errors.ErrSQLNoRows)) {
		t.Fatal("ClientGetAdmins (no records) wrong err:", err)
	}
}

func Test_client_get_by_name(t *testing.T) {
	client := newTestAPIClient()
	cs := []types.Client{
		{ID: 1, Name: "ha"},
		{ID: 2, Name: "foo"},
		{ID: 3, Name: "fohao"},
		{ID: 4, Name: "foo ha"},
	}

	pattern := "ha"
	count := 0
	for _, c := range cs {
		if strings.Contains(c.Name, pattern) {
			count++
		}
		client.ClientCreate(&c)
	}

	_cs, err := client.ClientGetByName(pattern)
	if err != nil {
		t.Fatal(err)
	}
	if len(_cs) != count {
		t.Fatalf("len(_cs) [%d] != count [%d]", len(_cs), count)
	}

	for _, c := range cs {
		client.ClientDelete(c.ID)
	}

	_cs, err = client.ClientGetByName(pattern)
	if !(api_errors.IsChildErr(err, api_errors.ErrSQLNoRows)) {
		t.Fatal(err)
	}
}

func Test_client_patch(t *testing.T) {
	client := newTestAPIClient()
	c := &types.Client{
		ID:      1,
		Name:    "foo",
		IsAdmin: false,
	}

	client.ClientCreate(c)
	defer client.ClientDelete(c.ID)

	check := func() {
		var (
			_c  *types.Client
			err error
		)

		err = client.ClientPatch(c)
		if err != nil {
			t.Fatal(err)
		}

		_c, _ = client.ClientGetByID(c.ID)

		if _c.ID != c.ID || _c.Name != c.Name || _c.IsAdmin != c.IsAdmin {
			t.Fatalf("Updated [[ %#v ]] != Local [[ %#v ]]", _c, c)
		}
	}

	c.Name = "qwerty"
	check()

	c.IsAdmin = true
	check()

	c.Name = "zxcvbnm"
	c.IsAdmin = false
	check()

	c.Name = ""
	client.ClientPatch(c)
}
