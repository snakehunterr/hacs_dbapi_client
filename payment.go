package api_client

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	types "github.com/snakehunterr/hacs_dbapi_types"
)

func (c APIClient) PaymentGetAll() ([]types.Payment, *types.APIResponse, error) {
	res, err := c.HTTPClient.Get(fmt.Sprintf("%s/payment/all", c.baseAPIURL))
	if err != nil {
		return nil, nil, fmt.Errorf("c.HTTPClient.Get(): %w", err)
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)

	if res.StatusCode != 200 {
		var r types.APIResponse

		if err := decoder.Decode(&r); err != nil {
			return nil, nil, fmt.Errorf("decoder.Decode(): %w", err)
		}

		return nil, &r, nil
	}

	var ps []types.Payment

	if err := decoder.Decode(&ps); err != nil {
		return nil, nil, fmt.Errorf("decoder.Decode(): %w", err)
	}

	return ps, nil, nil
}

func (c APIClient) PaymentGetAllByClientID(id uint) ([]types.Payment, *types.APIResponse, error) {
	res, err := c.HTTPClient.Get(fmt.Sprintf("%s/payment/client_id/%d", c.baseAPIURL, id))
	if err != nil {
		return nil, nil, fmt.Errorf("c.HTTPClient.Get(): %w", err)
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)

	if res.StatusCode != 200 {
		var r types.APIResponse

		if err := decoder.Decode(&r); err != nil {
			return nil, nil, fmt.Errorf("decoder.Decode(): %w", err)
		}

		return nil, &r, nil
	}

	var ps []types.Payment

	if err := decoder.Decode(&ps); err != nil {
		return nil, nil, fmt.Errorf("decoder.Decode(): %w", err)
	}

	return ps, nil, nil
}

func (c APIClient) PaymentGetAllByRoomID(id uint) ([]types.Payment, *types.APIResponse, error) {
	res, err := c.HTTPClient.Get(fmt.Sprintf("%s/payment/room_id/%d", c.baseAPIURL, id))
	if err != nil {
		return nil, nil, fmt.Errorf("c.HTTPClient.Get(): %w", err)
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)

	if res.StatusCode != 200 {
		var r types.APIResponse

		if err := decoder.Decode(&r); err != nil {
			return nil, nil, fmt.Errorf("decoder.Decode(): %w", err)
		}

		return nil, &r, nil
	}

	var ps []types.Payment

	if err := decoder.Decode(&ps); err != nil {
		return nil, nil, fmt.Errorf("decoder.Decode(): %w", err)
	}

	return ps, nil, nil
}

func (c APIClient) PaymentGetByID(id uint) (*types.Payment, *types.APIResponse, error) {
	res, err := c.HTTPClient.Get(fmt.Sprintf("%s/payment/id/%d", c.baseAPIURL, id))
	if err != nil {
		return nil, nil, fmt.Errorf("c.HTTPClient.Get(): %w", err)
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)

	if res.StatusCode != 200 {
		var r types.APIResponse

		if err := decoder.Decode(&r); err != nil {
			return nil, nil, fmt.Errorf("decoder.Decode(): %w", err)
		}

		return nil, &r, nil
	}

	var p types.Payment

	if err := decoder.Decode(&p); err != nil {
		return nil, nil, fmt.Errorf("decoder.Decode(): %w", err)
	}

	return &p, nil, nil
}

func (c APIClient) PaymentGetByDate(date time.Time) ([]types.Payment, *types.APIResponse, error) {
	datestr := date.Format("2006-01-02")
	res, err := c.HTTPClient.Get(fmt.Sprintf("%s/payment/date/%s", c.baseAPIURL, datestr))
	if err != nil {
		return nil, nil, fmt.Errorf("c.HTTPClient.Get(): %w", err)
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)

	if res.StatusCode != 200 {
		var r types.APIResponse

		if err := decoder.Decode(&r); err != nil {
			return nil, nil, fmt.Errorf("decoder.Decode(): %w", err)
		}

		return nil, &r, nil
	}

	var ps []types.Payment

	if err := decoder.Decode(&ps); err != nil {
		return nil, nil, fmt.Errorf("decoder.Decode(): %w", err)
	}

	return ps, nil, nil
}

func (c APIClient) PaymentGetByDateRange(date_start *time.Time, date_end *time.Time) ([]types.Payment, *types.APIResponse, error) {
	if date_start == nil && date_end == nil {
		return nil, nil, errors.New("both date_start and date_end is nil")
	}

	form := Form{}
	if date_start != nil {
		form["date_start"] = date_start.Format("2006-01-02")
	}
	if date_end != nil {
		form["date_end"] = date_end.Format("2006-01-02")
	}

	req, err := c.newFormRequest(
		http.MethodPost,
		fmt.Sprintf("%s/payment/date_range", c.baseAPIURL),
		form,
	)
	if err != nil {
		return nil, nil, fmt.Errorf("c.NewFormRequest(): %w", err)
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, nil, fmt.Errorf("c.HTTPClient.Do(): %w", err)
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)

	if res.StatusCode != 200 {
		var r types.APIResponse

		if err := decoder.Decode(&r); err != nil {
			return nil, nil, fmt.Errorf("decoder.Decode(): %w", err)
		}

		return nil, &r, nil
	}

	var ps []types.Payment

	if err := decoder.Decode(&ps); err != nil {
		return nil, nil, fmt.Errorf("decoder.Decode(): %w", err)
	}

	return ps, nil, nil
}

func (c APIClient) PaymentCreate(p *types.Payment) (*types.APIResponse, error) {
	if p == nil {
		return nil, errors.New("*types.Payment is nil")
	}

	req, err := c.newFormRequest(
		http.MethodPost,
		fmt.Sprintf("%s/payment/new", c.baseAPIURL),
		Form{
			"client_id":      strconv.FormatUint(uint64(p.ClientID), 10),
			"room_id":        strconv.FormatUint(uint64(p.RoomID), 10),
			"payment_date":   p.Date.Format("2006-01-02"),
			"payment_amount": fmt.Sprintf("%.2f", p.Amount),
		},
	)
	if err != nil {
		return nil, fmt.Errorf("c.NewFormRequest(): %w", err)
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("c.HTTPClient.Do(): %w", err)
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	var r types.APIResponse

	if err := decoder.Decode(&r); err != nil {
		return nil, fmt.Errorf("decoder.Decode(): %w", err)
	}

	return &r, nil
}

func (c APIClient) PaymentDelete(id uint) (*types.APIResponse, error) {
	return c.resourceDelete(fmt.Sprintf("%s/payment/id/%d", c.baseAPIURL, id))
}

func (c APIClient) PaymentPatch(p *types.Payment) (*types.APIResponse, error) {
	if p == nil {
		return nil, errors.New("*types.Payment is nil")
	}

	req, err := c.newFormRequest(
		http.MethodPatch,
		fmt.Sprintf("%s/payment/id/%d", c.baseAPIURL, p.ID),
		Form{
			"client_id":      strconv.FormatUint(uint64(p.ClientID), 10),
			"room_id":        strconv.FormatUint(uint64(p.RoomID), 10),
			"payment_date":   p.Date.Format("2006-01-02"),
			"payment_amount": fmt.Sprintf("%.2f", p.Amount),
		},
	)
	if err != nil {
		return nil, fmt.Errorf("c.NewFormRequest(): %w", err)
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("c.HTTPClient.Do(): %w", err)
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	var r types.APIResponse

	if err := decoder.Decode(&r); err != nil {
		return nil, fmt.Errorf("decoder.Decode(): %w", err)
	}

	return &r, nil
}
