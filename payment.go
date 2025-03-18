package api_client

import (
	"errors"
	"fmt"
	"time"

	types "github.com/snakehunterr/hacs_dbapi_types"
	"github.com/snakehunterr/hacs_dbapi_types/validators"
)

func (c APIClient) PaymentGetAll() ([]types.Payment, *types.APIResponse, error) {
	var ps []types.Payment
	r, err := c.resourceGet(fmt.Sprintf("%s/payment/all", c.baseAPIURL), &ps)

	switch {
	case err != nil:
		return nil, nil, err
	case r != nil:
		return nil, r, nil
	default:
		return ps, nil, nil
	}
}

func (c APIClient) PaymentGetAllByClientID(id int64) ([]types.Payment, *types.APIResponse, error) {
	var ps []types.Payment
	r, err := c.resourceGet(fmt.Sprintf("%s/payment/client/id/%d", c.baseAPIURL, id), &ps)

	switch {
	case err != nil:
		return nil, nil, err
	case r != nil:
		return nil, r, nil
	default:
		return ps, nil, nil
	}
}

func (c APIClient) PaymentGetAllByRoomID(id int64) ([]types.Payment, *types.APIResponse, error) {
	var ps []types.Payment
	r, err := c.resourceGet(fmt.Sprintf("%s/payment/room/id/%d", c.baseAPIURL, id), &ps)

	switch {
	case err != nil:
		return nil, nil, err
	case r != nil:
		return nil, r, nil
	default:
		return ps, nil, nil
	}
}

func (c APIClient) PaymentGetByID(id int64) (*types.Payment, *types.APIResponse, error) {
	var p types.Payment
	r, err := c.resourceGet(fmt.Sprintf("%s/payment/id/%d", c.baseAPIURL, id), &p)

	switch {
	case err != nil:
		return nil, nil, err
	case r != nil:
		return nil, r, nil
	default:
		return &p, nil, nil
	}
}

func (c APIClient) PaymentGetByDate(date time.Time) ([]types.Payment, *types.APIResponse, error) {
	var ps []types.Payment
	r, err := c.resourceGet(
		fmt.Sprintf("%s/payment/date/%s",
			c.baseAPIURL,
			date.Format(validators.DATE_FORMAT),
		),
		&ps,
	)

	switch {
	case err != nil:
		return nil, nil, err
	case r != nil:
		return nil, r, nil
	default:
		return ps, nil, nil
	}
}

func (c APIClient) PaymentGetByDateRange(date_start time.Time, date_end time.Time) ([]types.Payment, *types.APIResponse, error) {
	var ps []types.Payment
	r, err := c.resourcePost(
		fmt.Sprintf("%s/payment/date/range", c.baseAPIURL),
		Form{
			"date_start": date_start.Format(validators.DATE_FORMAT),
			"date_end":   date_end.Format(validators.DATE_FORMAT),
		},
		&ps,
	)

	switch {
	case err != nil:
		return nil, nil, err
	case r != nil:
		return nil, r, nil
	default:
		return ps, nil, nil
	}
}

func (c APIClient) PaymentCreate(p *types.Payment) (*types.APIResponse, error) {
	if p == nil {
		return nil, errors.New("*types.Payment is nil")
	}

	var _p types.Payment
	r, err := c.resourceCreateDecode(
		fmt.Sprintf("%s/payment/new", c.baseAPIURL),
		formatPayment(p),
		&_p,
	)

	switch {
	case err != nil:
		return nil, err
	case r != nil:
		return r, nil
	}

	p.ID = _p.ID
	return nil, nil
}

func (c APIClient) PaymentDelete(id int64) (*types.APIResponse, error) {
	return c.resourceDelete(fmt.Sprintf("%s/payment/id/%d", c.baseAPIURL, id))
}

func (c APIClient) PaymentPatch(p *types.Payment) (*types.APIResponse, error) {
	if p == nil {
		return nil, errors.New("*types.Payment is nil")
	}

	return c.resourcePatch(
		fmt.Sprintf("%s/payment/id/%d", c.baseAPIURL, p.ID),
		formatPayment(p),
	)
}
