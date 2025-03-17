package api_client

import (
	"os"
	"testing"
	"time"

	types "github.com/snakehunterr/hacs_dbapi_types"
	api_errors "github.com/snakehunterr/hacs_dbapi_types/errors"
	"github.com/snakehunterr/hacs_dbapi_types/validators"
)

func Test_expense_get_by_id(t *testing.T) {
	var (
		client = NewAPIClient(os.Getenv("HACS_DBAPI_HOST"), os.Getenv("HACS_DBAPI_PORT"))
		e      = &types.Expense{
			Date:   time.Now(),
			Amount: 9.99,
		}
	)

	res, err := client.ExpenseCreate(e)
	if err != nil {
		t.Fatal("ExpenseCreate:", err)
	}
	if res != nil {
		t.Fatal("ExpenseCreate *types.APIResponse:", res)
	}
	if e.ID == 0 {
		t.Fatal("ExpenseCreate *e.ID is 0")
	}

	_e, res, err := client.ExpenseGetByID(e.ID)
	if err != nil {
		t.Fatal("ExpenseGetByID:", err)
	}
	if res != nil {
		t.Fatal("ExpenseGetByID *types.APIResponse:", res)
	}
	if _e == nil {
		t.Fatal("ExpenseGetByID *types.Expense is nil")
	}
	if _e.ID != e.ID {
		t.Fatal("ExpenseGetByID returns wrong e.ID")
	}
	if _e.Date.Format(validators.DATE_FORMAT) != e.Date.Format(validators.DATE_FORMAT) {
		t.Fatal("ExpenseGetByID returns wrong e.Date")
	}
	if _e.Amount != e.Amount {
		t.Fatal("ExpenseGetByID returns wrong e.Amount")
	}

	_e, res, err = client.ExpenseGetByID(0)
	if err != nil {
		t.Fatal("ExpenseGetByID:", err)
	}
	if _e != nil {
		t.Fatal("ExpenseGetByID unexpected *types.Expense:", _e)
	}
	if !api_errors.IsChildErr(res.Error, api_errors.ErrSQLNoRows) {
		t.Fatal("ExpenseGetByID unexpected *APIError:", res.Error)
	}

	res, err = client.ExpenseDelete(e.ID)
	if err != nil {
		t.Fatal("ExpenseDelete:", err)
	}
	if res == nil {
		t.Fatal("ExpenseDelete *types.APIResponse is nil")
	}
	if res.Error != nil {
		t.Fatal("ExpenseDelete *APIError:", res.Error)
	}
}

func Test_expense_get_all(t *testing.T) {
	var (
		client = NewAPIClient(os.Getenv("HACS_DBAPI_HOST"), os.Getenv("HACS_DBAPI_PORT"))
	)

	es := []*types.Expense{
		{Date: time.Now(), Amount: 1.99},
		{Date: time.Now(), Amount: 2.99},
		{Date: time.Now(), Amount: 3.99},
		{Date: time.Now(), Amount: 4.99},
		{Date: time.Now(), Amount: 5.99},
	}

	for _, e := range es {
		res, err := client.ExpenseCreate(e)
		if err != nil {
			t.Fatal("ExpenseCreate:", err)
		}
		if res != nil {
			t.Fatal("ExpenseCreate *types.APIResponse:", res)
		}
		if e.ID == 0 {
			t.Fatal("ExpenseCreate e.ID is 0")
		}
	}

	_es, res, err := client.ExpenseGetAll()
	if err != nil {
		t.Fatal("ExpenseGetAll:", err)
	}
	if res != nil {
		t.Fatal("ExpenseGetAll *types.APIResponse:", res)
	}
	if len(_es) != len(es) {
		t.Fatal("ExpenseGetAll returns:", _es)
	}

	for _, e := range es {
		res, err := client.ExpenseDelete(e.ID)
		if err != nil {
			t.Fatal("ExpenseDelete:", err)
		}
		if res == nil {
			t.Fatal("ExpenseDelete *types.APIResponse is nil")
		}
		if res.Error != nil {
			t.Fatal("ExpenseDelete *APIError:", res.Error)
		}
	}
}

func Test_expense_get_by_date(t *testing.T) {
	var (
		client = NewAPIClient(os.Getenv("HACS_DBAPI_HOST"), os.Getenv("HACS_DBAPI_PORT"))
	)

	t1, _ := time.Parse("2006-01-02", "2014-01-10")
	t2, _ := time.Parse("2006-01-02", "2014-01-20")

	es := []*types.Expense{
		{Date: t1, Amount: 1},
		{Date: t2, Amount: 1},
		{Date: t1, Amount: 1},
	}

	for _, e := range es {
		res, err := client.ExpenseCreate(e)
		if err != nil {
			t.Fatal("ExpenseCreate:", err)
		}
		if res != nil {
			t.Fatal("ExpenseCreate *types.APIResponse:", res)
		}
		if e.ID == 0 {
			t.Fatal("ExpenseCreate e.ID is 0")
		}
	}

	_es, res, err := client.ExpenseGetByDate(t1)
	if err != nil {
		t.Fatal("ExpenseGetByDate:", err)
	}
	if res != nil {
		t.Fatal("ExpenseGetByDate *types.APIResponse:", res)
	}
	if len(_es) != 2 {
		t.Fatal("ExpenseGetByDate returns:", _es)
	}

	_es, res, err = client.ExpenseGetByDate(t2)
	if err != nil {
		t.Fatal("ExpenseGetByDate:", err)
	}
	if res != nil {
		t.Fatal("ExpenseGetByDate *types.APIResponse:", res)
	}
	if len(_es) != 1 {
		t.Fatal("ExpenseGetByDate returns:", _es)
	}

	for _, e := range es {
		res, err := client.ExpenseDelete(e.ID)
		if err != nil {
			t.Fatal("ExpenseDelete:", err)
		}
		if res == nil {
			t.Fatal("ExpenseDelete *types.APIResponse is nil")
		}
		if res.Error != nil {
			t.Fatal("ExpenseDelete *APIError:", res.Error)
		}
	}
}

func Test_expense_get_by_date_range(t *testing.T) {

	var (
		client = NewAPIClient(os.Getenv("HACS_DBAPI_HOST"), os.Getenv("HACS_DBAPI_PORT"))
	)

	t1, _ := time.Parse("2006-01-02", "2014-01-10")
	t2, _ := time.Parse("2006-01-02", "2014-01-20")
	t3, _ := time.Parse("2006-01-02", "2014-01-30")
	t4, _ := time.Parse("2006-01-02", "2014-02-20")

	es := []*types.Expense{
		{Date: t1, Amount: 1},
		{Date: t2, Amount: 1},
		{Date: t3, Amount: 1},
	}

	for _, e := range es {
		res, err := client.ExpenseCreate(e)
		if err != nil {
			t.Fatal("ExpenseCreate:", err)
		}
		if res != nil {
			t.Fatal("ExpenseCreate *types.APIResponse:", res)
		}
		if e.ID == 0 {
			t.Fatal("ExpenseCreate e.ID is 0")
		}
	}

	_es, res, err := client.ExpenseGetByDateRange(t1, t1)
	if err != nil {
		t.Fatal("ExpenseGetByDateRange:", err)
	}
	if res != nil {
		t.Fatal("ExpenseGetByDateRange *types.APIResponse:", res)
	}
	if len(_es) != 1 {
		t.Fatal("ExpenseGetByDateRange returns:", _es)
	}

	_es, res, err = client.ExpenseGetByDateRange(t1, t2)
	if err != nil {
		t.Fatal("ExpenseGetByDateRange:", err)
	}
	if res != nil {
		t.Fatal("ExpenseGetByDateRange *types.APIResponse:", res)
	}
	if len(_es) != 2 {
		t.Fatal("ExpenseGetByDateRange returns:", _es)
	}

	_es, res, err = client.ExpenseGetByDateRange(t1, t3)
	if err != nil {
		t.Fatal("ExpenseGetByDateRange:", err)
	}
	if res != nil {
		t.Fatal("ExpenseGetByDateRange *types.APIResponse:", res)
	}
	if len(_es) != 3 {
		t.Fatal("ExpenseGetByDateRange returns:", _es)
	}

	_es, res, err = client.ExpenseGetByDateRange(t1, t4)
	if err != nil {
		t.Fatal("ExpenseGetByDateRange:", err)
	}
	if res != nil {
		t.Fatal("ExpenseGetByDateRange *types.APIResponse:", res)
	}
	if len(_es) != 3 {
		t.Fatal("ExpenseGetByDateRange returns:", _es)
	}

	for _, e := range es {
		res, err := client.ExpenseDelete(e.ID)
		if err != nil {
			t.Fatal("ExpenseDelete:", err)
		}
		if res == nil {
			t.Fatal("ExpenseDelete *types.APIResponse is nil")
		}
		if res.Error != nil {
			t.Fatal("ExpenseDelete *APIError:", res.Error)
		}
	}
}

func Test_expense_patch(t *testing.T) {
	var (
		client = NewAPIClient(os.Getenv("HACS_DBAPI_HOST"), os.Getenv("HACS_DBAPI_PORT"))
		e      = &types.Expense{
			Date:   time.Now(),
			Amount: 22.22,
		}
	)

	res, err := client.ExpenseCreate(e)
	if err != nil {
		t.Fatal("ExpenseCreate:", err)
	}
	if res != nil {
		t.Fatal("ExpenseCreate *types.APIResponse:", res)
	}
	if e.ID == 0 {
		t.Fatal("ExpenseCreate *e.ID is 0")
	}
	defer func() {
		res, err := client.ExpenseDelete(e.ID)
		if err != nil {
			t.Fatal("ExpenseDelete:", err)
		}
		if res == nil {
			t.Fatal("ExpenseDelete *types.APIResponse is nil")
		}
		if res.Error != nil {
			t.Fatal("ExpenseDelete *APIError:", res.Error)
		}
	}()

	check := func() {
		res, err := client.ExpensePatch(e)
		if err != nil {
			t.Fatal("ExpensePatch:", err)
		}
		if res == nil {
			t.Fatal("ExpensePatch *types.APIResponse is nil")
		}
		if res.Error != nil {
			t.Fatal("ExpensePatch *APIError:", res.Error)
		}

		_e, res, err := client.ExpenseGetByID(e.ID)
		if err != nil {
			t.Fatal("ExpenseGetByID:", err)
		}
		if res != nil {
			t.Fatal("ExpenseGetByID *types.APIResponse:", res)
		}

		if _e.ID != e.ID {
			t.Fatal("ExpensePatch wrong e.ID")
		}
		if _e.Date.Format(validators.DATE_FORMAT) != e.Date.Format(validators.DATE_FORMAT) {
			t.Fatal("ExpensePatch wrong e.Date")
		}
		if _e.Amount != e.Amount {
			t.Fatal("ExpensePatch wrong e.Amount")
		}
	}

	e.Amount = 10.5
	check()

	e.Date = time.Now().Add(time.Hour*12345 + time.Minute*12345 + time.Second*12345)
	check()

	e.Date = time.Now()
	e.Amount = 0.11
	check()
}
