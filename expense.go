package api_client

// package api_client

// import (
// 	"errors"
// 	"fmt"
// 	"time"

// 	types "github.com/snakehunterr/hacs_dbapi_types"
// 	"github.com/snakehunterr/hacs_dbapi_types/validators"
// )

// func (c APIClient) ExpenseGetAll() ([]types.Expense, *types.APIResponse, error) {
// 	var es []types.Expense

// 	res, err := c.resourceGet(fmt.Sprintf("%s/expense/all", c.baseAPIURL), &es)
// 	switch {
// 	case err != nil:
// 		return nil, nil, err
// 	case res != nil:
// 		return nil, res, nil
// 	default:
// 		return es, nil, nil
// 	}
// }

// func (c APIClient) ExpenseGetByID(id int64) (*types.Expense, *types.APIResponse, error) {
// 	var e types.Expense

// 	res, err := c.resourceGet(fmt.Sprintf("%s/expense/id/%d", c.baseAPIURL, id), &e)
// 	switch {
// 	case err != nil:
// 		return nil, nil, err
// 	case res != nil:
// 		return nil, res, nil
// 	default:
// 		return &e, nil, nil
// 	}
// }

// func (c APIClient) ExpenseGetByDate(date time.Time) ([]types.Expense, *types.APIResponse, error) {
// 	var es []types.Expense

// 	res, err := c.resourceGet(
// 		fmt.Sprintf("%s/expense/date/%s", c.baseAPIURL, date.Format(validators.DATE_FORMAT)),
// 		&es,
// 	)
// 	switch {
// 	case err != nil:
// 		return nil, nil, err
// 	case res != nil:
// 		return nil, res, nil
// 	default:
// 		return es, nil, nil
// 	}
// }

// func (c APIClient) ExpenseGetByDateRange(date_start time.Time, date_end time.Time) ([]types.Expense, *types.APIResponse, error) {
// 	var es []types.Expense

// 	res, err := c.resourcePost(
// 		fmt.Sprintf("%s/expense/date/range", c.baseAPIURL),
// 		Form{
// 			"date_start": date_start.Format(validators.DATE_FORMAT),
// 			"date_end":   date_end.Format(validators.DATE_FORMAT),
// 		},
// 		&es,
// 	)

// 	switch {
// 	case err != nil:
// 		return nil, nil, err
// 	case res != nil:
// 		return nil, res, nil
// 	default:
// 		return es, nil, nil
// 	}
// }

// func (c APIClient) ExpenseCreate(e *types.Expense) (*types.APIResponse, error) {
// 	if e == nil {
// 		return nil, errors.New("*types.Expense is nil")
// 	}

// 	var _e types.Expense
// 	res, err := c.resourceCreateDecode(
// 		fmt.Sprintf("%s/expense/new", c.baseAPIURL),
// 		formatExpense(e),
// 		&_e,
// 	)

// 	switch {
// 	case err != nil:
// 		return nil, err
// 	case res != nil:
// 		return res, nil
// 	}

// 	e.ID = _e.ID
// 	return nil, nil
// }

// func (c APIClient) ExpenseDelete(id int64) (*types.APIResponse, error) {
// 	return c.resourceDelete(fmt.Sprintf("%s/expense/id/%d", c.baseAPIURL, id))
// }

// func (c APIClient) ExpensePatch(e *types.Expense) (*types.APIResponse, error) {
// 	if e == nil {
// 		return nil, errors.New("*types.Expense is nil")
// 	}

// 	return c.resourcePatch(
// 		fmt.Sprintf("%s/expense/id/%d", c.baseAPIURL, e.ID),
// 		formatExpense(e),
// 	)
// }
