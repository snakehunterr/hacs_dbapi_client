package api_client

import (
	"encoding/json"
	"testing"
	"time"

	types "github.com/snakehunterr/hacs_dbapi_types"
	api_errors "github.com/snakehunterr/hacs_dbapi_types/errors"
)

func Test_expense_create_delete(t *testing.T) {
	var (
		e   = newTestExpense()
		_e  *types.Expense
		err error
	)

	_e, err = client.ExpenseCreate(e.Date, e.Amount)
	if err != nil {
		t.Fatal(err)
	}
	if _e == nil {
		t.Fatal()
	}

	err = client.ExpenseDelete(_e.ID)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_expense_get_by_id(t *testing.T) {
	var (
		e   = newTestExpense()
		_e  *types.Expense
		__e *types.Expense
		err error
	)

	_e, err = client.ExpenseGetByID(1234445)
	if err == nil {
		t.Fatal(_e)
	}

	_e, _ = client.ExpenseCreate(e.Date, e.Amount)
	defer client.ExpenseDelete(_e.ID)

	__e, err = client.ExpenseGetByID(_e.ID)
	if err != nil {
		t.Fatal()
	}
	if !float_equal(_e.Amount, __e.Amount) ||
		!date_equal(_e.Date, __e.Date) {
		_ejson, _ := json.MarshalIndent(_e, "", "\t")
		__ejson, _ := json.MarshalIndent(__e, "", "\t")
		t.Fatalf("_e:\n%s\n__e:\n%s", _ejson, __ejson)
	}
}

func Test_expense_get_all(t *testing.T) {
	var (
		es = []*types.Expense{
			newTestExpense(),
			newTestExpense(),
			newTestExpense(),
			newTestExpense(),
			newTestExpense(),
		}
		_es []types.Expense
		err error
	)

	_es, err = client.ExpenseGetAll()
	if err == nil {
		t.Fatal(_es)
	}

	for _, e := range es {
		_e, _ := client.ExpenseCreate(e.Date, e.Amount)
		defer client.ExpenseDelete(_e.ID)
	}

	_es, err = client.ExpenseGetAll()
	if err != nil {
		t.Fatal(err)
	}
	if len(_es) != len(es) {
		t.Fatal()
	}
}

func Test_expense_get_by_date(t *testing.T) {
	var (
		d1 = newTestDate("2000-01-10")
		d2 = newTestDate("2014-05-13")
		es = []types.Expense{
			{Date: d1, Amount: 1},
			{Date: d2, Amount: 1},
			{Date: d2, Amount: 1},
			{Date: d1, Amount: 1},
			{Date: d1, Amount: 1},
			{Date: d2, Amount: 1},
			{Date: d2, Amount: 1},
			{Date: d1, Amount: 1},
		}
		_es []types.Expense
		err error
	)

	_es, err = client.ExpenseGetByDate(d1)
	if err == nil {
		t.Fatal(_es)
	}
	_es, err = client.ExpenseGetByDate(d2)
	if err == nil {
		t.Fatal(_es)
	}

	d1_count := 0
	d2_count := 0
	for _, e := range es {
		switch {
		case date_equal(e.Date, d1):
			d1_count++
		case date_equal(e.Date, d2):
			d2_count++
		}

		_e, _ := client.ExpenseCreate(e.Date, e.Amount)
		defer client.ExpenseDelete(_e.ID)
	}

	_es, err = client.ExpenseGetByDate(d1)
	if err != nil {
		t.Fatal(err)
	}
	if len(_es) != d1_count {
		t.Fatal(_es, len(_es), d1_count)
	}

	_es, err = client.ExpenseGetByDate(d2)
	if err != nil {
		t.Fatal(err)
	}
	if len(_es) != d2_count {
		t.Fatal(_es, len(_es), d2_count)
	}
}

func Test_expense_get_by_date_range(t *testing.T) {
	var (
		d1 = newTestDate("2000-01-10")
		d2 = newTestDate("2000-01-20")
		d3 = newTestDate("2000-01-30")
		d4 = newTestDate("2000-02-10")
		es = []types.Expense{
			{Date: d2, Amount: 1}, {Date: d1, Amount: 1},
			{Date: d1, Amount: 1}, {Date: d2, Amount: 1},
			{Date: d3, Amount: 1}, {Date: d2, Amount: 1},
			{Date: d1, Amount: 1}, {Date: d1, Amount: 1},
			{Date: d1, Amount: 1}, {Date: d2, Amount: 1},
			{Date: d3, Amount: 1}, {Date: d1, Amount: 1},
			{Date: d2, Amount: 1}, {Date: d1, Amount: 1},
			{Date: d1, Amount: 1}, {Date: d2, Amount: 1},
		}
		_es []types.Expense
		err error
	)

	var check func(time.Time, time.Time)

	check = func(d1, d2 time.Time) {
		_es, err = client.ExpenseGetByDateRange(d1, d2)
		if err == nil {
			t.Fatal(d1, d2, _es)
		}
	}

	dates := [][]time.Time{
		{d1, d1}, {d2, d1}, {d3, d1}, {d4, d1},
		{d1, d2}, {d2, d2}, {d3, d2}, {d4, d2},
		{d1, d3}, {d2, d3}, {d3, d3}, {d4, d3},
		{d1, d4}, {d2, d4}, {d3, d4}, {d4, d4},
	}

	for _, grp := range dates {
		check(grp[0], grp[1])
	}

	check = func(d1, d2 time.Time) {
		count := 0
		for _, e := range es {
			if !(e.Date.After(d1) || e.Date.Equal(d1)) {
				continue
			}
			if !(d2.After(e.Date) || d2.Equal(e.Date)) {
				continue
			}
			count++
		}

		_es, err = client.ExpenseGetByDateRange(d1, d2)
		if err != nil {
			if api_errors.IsChildErr(err, api_errors.ErrSQLNoRows) && count == 0 {
				return
			}
			t.Fatal(err)
		}
		if len(_es) != count {
			js, _ := json.MarshalIndent(_es, "", "\t")
			t.Fatalf("len(%d); count(%d)\n%s", len(_es), count, js)
		}
	}

	for _, e := range es {
		_e, _ := client.ExpenseCreate(e.Date, e.Amount)
		defer client.ExpenseDelete(_e.ID)
	}

	for _, grp := range dates {
		check(grp[0], grp[1])
	}
}

func Test_expense_patch(t *testing.T) {
	var (
		e   = newTestExpense()
		_e  *types.Expense
		err error
	)

	e.ID = 12345
	err = client.ExpensePatch(e)
	if err == nil {
		t.Fatal()
	}

	_e, _ = client.ExpenseCreate(e.Date, e.Amount)
	defer client.ExpenseDelete(_e.ID)
	e.ID = _e.ID

	check := func() {
		err = client.ExpensePatch(e)
		if err != nil {
			t.Fatal(err)
		}

		_e, _ = client.ExpenseGetByID(e.ID)

		if !float_equal(e.Amount, _e.Amount) ||
			!date_equal(e.Date, _e.Date) {
			ejson, _ := json.MarshalIndent(e, "", "\t")
			_ejson, _ := json.MarshalIndent(_e, "", "\t")
			t.Fatalf("e:\n%s\n_e:\n%s\n", ejson, _ejson)
		}
	}

	e.Date = time.Now().Add(time.Hour*123 + time.Second*12345)
	check()

	e.Amount = 0.88
	check()

	e.Date = time.Now()
	e.Amount = 1234.1234
	check()
}
