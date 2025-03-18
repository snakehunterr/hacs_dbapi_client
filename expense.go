package api_client

import (
	"errors"
	"fmt"
	"time"

	types "github.com/snakehunterr/hacs_dbapi_types"
	"github.com/snakehunterr/hacs_dbapi_types/validators"
)

func (c APIClient) ExpenseGetAll() ([]types.Expense, error) {
	var es []types.Expense
	err := c.resourceGet(fmt.Sprintf("%s/expense/all", c.baseAPIURL), &es)

	return es, err
}

func (c APIClient) ExpenseGetByID(id int64) (*types.Expense, error) {
	var e types.Expense
	err := c.resourceGet(fmt.Sprintf("%s/expense/id/%d", c.baseAPIURL, id), &e)

	if err != nil {
		return nil, err
	} else {
		return &e, nil
	}
}

func (c APIClient) ExpenseGetByDate(date time.Time) ([]types.Expense, error) {
	var es []types.Expense
	err := c.resourceGet(
		fmt.Sprintf("%s/expense/date/%s", c.baseAPIURL, date.Format(validators.DATE_FORMAT)),
		&es,
	)

	return es, err
}

func (c APIClient) ExpenseGetByDateRange(date_start time.Time, date_end time.Time) ([]types.Expense, error) {
	var es []types.Expense
	err := c.resourcePost(
		fmt.Sprintf("%s/expense/date/range", c.baseAPIURL),
		Form{
			"date_start": date_start.Format(validators.DATE_FORMAT),
			"date_end":   date_end.Format(validators.DATE_FORMAT),
		},
		&es,
	)

	return es, err
}

func (c APIClient) ExpenseCreate(date time.Time, amount float64) (*types.Expense, error) {
	var e = types.Expense{
		Date:   date,
		Amount: amount,
	}
	err := c.resourceCreateWithDecode(
		fmt.Sprintf("%s/expense/new", c.baseAPIURL),
		formatExpense(&e),
		&e,
	)

	if err != nil {
		return nil, err
	} else {
		return &e, nil
	}
}

func (c APIClient) ExpenseDelete(id int64) error {
	return c.resourceDelete(fmt.Sprintf("%s/expense/id/%d", c.baseAPIURL, id))
}

func (c APIClient) ExpensePatch(e *types.Expense) error {
	if e == nil {
		return errors.New("*types.Expense is nil")
	}

	return c.resourcePatch(
		fmt.Sprintf("%s/expense/id/%d", c.baseAPIURL, e.ID),
		formatExpense(e),
	)
}
