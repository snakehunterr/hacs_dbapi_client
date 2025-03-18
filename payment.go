package api_client

import (
	"errors"
	"fmt"
	"time"

	types "github.com/snakehunterr/hacs_dbapi_types"
	"github.com/snakehunterr/hacs_dbapi_types/validators"
)

func (c APIClient) PaymentGetAll() ([]types.Payment, error) {
	var ps []types.Payment
	err := c.resourceGet(fmt.Sprintf("%s/payment/all", c.baseAPIURL), &ps)

	return ps, err
}

func (c APIClient) PaymentGetAllByClientID(id int64) ([]types.Payment, error) {
	var ps []types.Payment
	err := c.resourceGet(fmt.Sprintf("%s/payment/client/id/%d", c.baseAPIURL, id), &ps)

	return ps, err
}

func (c APIClient) PaymentGetAllByRoomID(id int64) ([]types.Payment, error) {
	var ps []types.Payment
	err := c.resourceGet(fmt.Sprintf("%s/payment/room/id/%d", c.baseAPIURL, id), &ps)

	return ps, err
}

func (c APIClient) PaymentGetByID(id int64) (*types.Payment, error) {
	var p types.Payment
	err := c.resourceGet(fmt.Sprintf("%s/payment/id/%d", c.baseAPIURL, id), &p)

	if err != nil {
		return nil, err
	} else {
		return &p, nil
	}
}

func (c APIClient) PaymentGetByDate(date time.Time) ([]types.Payment, error) {
	var ps []types.Payment
	err := c.resourceGet(
		fmt.Sprintf("%s/payment/date/%s",
			c.baseAPIURL,
			date.Format(validators.DATE_FORMAT),
		),
		&ps,
	)

	return ps, err
}

func (c APIClient) PaymentGetByDateRange(date_start time.Time, date_end time.Time) ([]types.Payment, error) {
	var ps []types.Payment
	err := c.resourcePost(
		fmt.Sprintf("%s/payment/date/range", c.baseAPIURL),
		Form{
			"date_start": date_start.Format(validators.DATE_FORMAT),
			"date_end":   date_end.Format(validators.DATE_FORMAT),
		},
		&ps,
	)

	return ps, err
}

func (c APIClient) PaymentCreate(client_id int64, room_id int64, date time.Time, amount float64) (*types.Payment, error) {
	var (
		p = types.Payment{
			ClientID: client_id,
			RoomID:   room_id,
			Date:     date,
			Amount:   amount,
		}
	)

	err := c.resourceCreateWithDecode(
		fmt.Sprintf("%s/payment/new", c.baseAPIURL),
		formatPayment(&p),
		&p,
	)

	if err != nil {
		return nil, err
	} else {
		return &p, nil
	}
}

func (c APIClient) PaymentDelete(id int64) error {
	return c.resourceDelete(fmt.Sprintf("%s/payment/id/%d", c.baseAPIURL, id))
}

func (c APIClient) PaymentPatch(p *types.Payment) error {
	if p == nil {
		return errors.New("*types.Payment is nil")
	}

	return c.resourcePatch(
		fmt.Sprintf("%s/payment/id/%d", c.baseAPIURL, p.ID),
		formatPayment(p),
	)
}
