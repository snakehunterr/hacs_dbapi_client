package api_client

import (
	"encoding/json"
	"testing"
	"time"

	types "github.com/snakehunterr/hacs_dbapi_types"
	api_errors "github.com/snakehunterr/hacs_dbapi_types/errors"
)

func Test_payment_create_delete(t *testing.T) {
	var (
		c   = newTestClient(1)
		r   = newTestRoom(c.ID, 1)
		p   = newTestPayment(c.ID, r.ID)
		_p  *types.Payment
		err error
	)

	_, err = client.PaymentCreate(p.ClientID, p.RoomID, p.Date, p.Amount)
	if err == nil {
		t.Fatal()
	}

	client.ClientCreate(c)
	defer client.ClientDelete(c.ID)
	client.RoomCreate(r)
	defer client.RoomDelete(r.ID)

	_p, err = client.PaymentCreate(p.ClientID, p.RoomID, p.Date, p.Amount)
	if err != nil {
		t.Fatal(err)
	}
	if _p.ClientID != p.ClientID ||
		_p.RoomID != p.RoomID ||
		!float_equal(_p.Amount, p.Amount) ||
		!date_equal(_p.Date, p.Date) {
		t.Fatal(_p, p)
	}
	err = client.PaymentDelete(_p.ID)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_payment_get_by_id(t *testing.T) {
	var (
		c   = newTestClient(1)
		r   = newTestRoom(c.ID, 1)
		p   = newTestPayment(c.ID, r.ID)
		_p  *types.Payment
		__p *types.Payment
		err error
	)

	_p, err = client.PaymentGetByID(1123456)
	if err == nil {
		t.Fatal(_p)
	}

	client.ClientCreate(c)
	defer client.ClientDelete(c.ID)
	client.RoomCreate(r)
	defer client.RoomDelete(r.ID)

	_p, _ = client.PaymentCreate(p.ClientID, p.RoomID, p.Date, p.Amount)
	defer client.PaymentDelete(_p.ID)

	__p, err = client.PaymentGetByID(_p.ID)
	if _p.ClientID != __p.ClientID ||
		_p.RoomID != __p.RoomID ||
		!date_equal(_p.Date, __p.Date) ||
		!float_equal(_p.Amount, __p.Amount) {
		t.Fatal(__p, _p)
	}
}

func Test_payment_get_all(t *testing.T) {
	var (
		c1 = newTestClient(1)
		c2 = newTestClient(2)
		r1 = newTestRoom(c1.ID, 1)
		r2 = newTestRoom(c1.ID, 2)
		ps = []*types.Payment{
			newTestPayment(c1.ID, r1.ID),
			newTestPayment(c2.ID, r2.ID),
			newTestPayment(c1.ID, r2.ID),
			newTestPayment(c2.ID, r1.ID),
			newTestPayment(c1.ID, r2.ID),
			newTestPayment(c1.ID, r1.ID),
			newTestPayment(c2.ID, r2.ID),
			newTestPayment(c1.ID, r2.ID),
			newTestPayment(c2.ID, r1.ID),
			newTestPayment(c1.ID, r2.ID),
		}
		_ps []types.Payment
		err error
	)

	_ps, err = client.PaymentGetAll()
	if err == nil {
		t.Fatal(_ps)
	}

	client.ClientCreate(c1)
	defer client.ClientDelete(c1.ID)
	client.ClientCreate(c2)
	defer client.ClientDelete(c2.ID)
	client.RoomCreate(r1)
	defer client.RoomDelete(r1.ID)
	client.RoomCreate(r2)
	defer client.RoomDelete(r2.ID)

	for _, p := range ps {
		_p, _ := client.PaymentCreate(p.ClientID, p.RoomID, p.Date, p.Amount)
		defer client.PaymentDelete(_p.ID)
	}

	_ps, err = client.PaymentGetAll()
	if err != nil {
		t.Fatal(err)
	}
	if len(_ps) != len(ps) {
		t.Fatal()
	}
}

func Test_payment_get_by_client_id(t *testing.T) {
	var (
		c1 = newTestClient(1)
		c2 = newTestClient(2)
		r  = newTestRoom(c1.ID, 1)
		ps = []*types.Payment{
			newTestPayment(c1.ID, r.ID),
			newTestPayment(c2.ID, r.ID),
			newTestPayment(c1.ID, r.ID),
			newTestPayment(c2.ID, r.ID),
			newTestPayment(c1.ID, r.ID),
		}
		_ps []types.Payment
		err error
	)

	_ps, err = client.PaymentGetAllByClientID(c1.ID)
	if err == nil {
		t.Fatal(_ps)
	}
	_ps, err = client.PaymentGetAllByClientID(c2.ID)
	if err == nil {
		t.Fatal(_ps)
	}

	client.ClientCreate(c1)
	defer client.ClientDelete(c1.ID)
	client.ClientCreate(c2)
	defer client.ClientDelete(c2.ID)
	client.RoomCreate(r)
	defer client.RoomDelete(r.ID)

	c1_count := 0
	c2_count := 0
	for _, p := range ps {
		switch p.ClientID {
		case c1.ID:
			c1_count++
		case c2.ID:
			c2_count++
		}

		_p, _ := client.PaymentCreate(p.ClientID, p.RoomID, p.Date, p.Amount)
		defer client.PaymentDelete(_p.ID)
	}

	_ps, err = client.PaymentGetAllByClientID(c1.ID)
	if err != nil {
		t.Fatal()
	}
	if len(_ps) != c1_count {
		t.Fatal(_ps)
	}

	_ps, err = client.PaymentGetAllByClientID(c2.ID)
	if err != nil {
		t.Fatal()
	}
	if len(_ps) != c2_count {
		t.Fatal(_ps)
	}
}

func Test_payment_get_by_room_id(t *testing.T) {
	var (
		c  = newTestClient(1)
		r1 = newTestRoom(c.ID, 1)
		r2 = newTestRoom(c.ID, 2)
		ps = []*types.Payment{
			newTestPayment(c.ID, r1.ID),
			newTestPayment(c.ID, r2.ID),
			newTestPayment(c.ID, r1.ID),
			newTestPayment(c.ID, r2.ID),
			newTestPayment(c.ID, r1.ID),
		}
		_ps []types.Payment
		err error
	)

	_ps, err = client.PaymentGetAllByRoomID(r1.ID)
	if err == nil {
		t.Fatal(_ps)
	}
	_ps, err = client.PaymentGetAllByRoomID(r2.ID)
	if err == nil {
		t.Fatal(_ps)
	}

	client.ClientCreate(c)
	defer client.ClientDelete(c.ID)
	client.RoomCreate(r1)
	defer client.RoomDelete(r1.ID)
	client.RoomCreate(r2)
	defer client.RoomDelete(r2.ID)

	r1_count := 0
	r2_count := 0
	for _, p := range ps {
		switch p.RoomID {
		case r1.ID:
			r1_count++
		case r2.ID:
			r2_count++
		}

		_p, _ := client.PaymentCreate(p.ClientID, p.RoomID, p.Date, p.Amount)
		defer client.PaymentDelete(_p.ID)
	}

	_ps, err = client.PaymentGetAllByRoomID(r1.ID)
	if err != nil {
		t.Fatal()
	}
	if len(_ps) != r1_count {
		t.Fatal(_ps)
	}

	_ps, err = client.PaymentGetAllByRoomID(r2.ID)
	if err != nil {
		t.Fatal()
	}
	if len(_ps) != r2_count {
		t.Fatal(_ps)
	}
}

func Test_payment_get_by_date(t *testing.T) {
	var (
		c  = newTestClient(1)
		r  = newTestRoom(c.ID, 1)
		d1 = newTestDate("2000-01-10")
		d2 = newTestDate("2004-05-05")
		ps = []*types.Payment{
			{ClientID: c.ID, RoomID: r.ID, Date: d1, Amount: 1},
			{ClientID: c.ID, RoomID: r.ID, Date: d2, Amount: 1},
			{ClientID: c.ID, RoomID: r.ID, Date: d1, Amount: 1},
			{ClientID: c.ID, RoomID: r.ID, Date: d2, Amount: 1},
			{ClientID: c.ID, RoomID: r.ID, Date: d1, Amount: 1},
		}
		_ps []types.Payment
		err error
	)

	_ps, err = client.PaymentGetByDate(d1)
	if err == nil {
		t.Fatal(_ps)
	}
	_ps, err = client.PaymentGetByDate(d2)
	if err == nil {
		t.Fatal(_ps)
	}

	client.ClientCreate(c)
	defer client.ClientDelete(c.ID)
	client.RoomCreate(r)
	defer client.RoomDelete(r.ID)

	d1_count := 0
	d2_count := 0
	for _, p := range ps {
		switch {
		case date_equal(d1, p.Date):
			d1_count++
		case date_equal(d2, p.Date):
			d2_count++
		}

		_p, _ := client.PaymentCreate(p.ClientID, p.RoomID, p.Date, p.Amount)
		defer client.PaymentDelete(_p.ID)
	}

	_ps, err = client.PaymentGetByDate(d1)
	if err != nil {
		t.Fatal(err)
	}
	if len(_ps) != d1_count {
		t.Fatal(_ps)
	}

	_ps, err = client.PaymentGetByDate(d2)
	if err != nil {
		t.Fatal(err)
	}
	if len(_ps) != d2_count {
		t.Fatal(_ps)
	}
}

func Test_payment_get_by_date_range(t *testing.T) {
	var (
		c  = newTestClient(1)
		r  = newTestRoom(c.ID, 1)
		d1 = newTestDate("2000-01-10")
		d2 = newTestDate("2000-01-20")
		d3 = newTestDate("2000-01-30")
		d4 = newTestDate("2000-02-10")
		ps = []types.Payment{
			{ClientID: c.ID, RoomID: r.ID, Date: d3, Amount: 1},
			{ClientID: c.ID, RoomID: r.ID, Date: d1, Amount: 1},
			{ClientID: c.ID, RoomID: r.ID, Date: d2, Amount: 1},
			{ClientID: c.ID, RoomID: r.ID, Date: d1, Amount: 1},
			{ClientID: c.ID, RoomID: r.ID, Date: d1, Amount: 1},
			{ClientID: c.ID, RoomID: r.ID, Date: d3, Amount: 1},
			{ClientID: c.ID, RoomID: r.ID, Date: d2, Amount: 1},
			{ClientID: c.ID, RoomID: r.ID, Date: d1, Amount: 1},
			{ClientID: c.ID, RoomID: r.ID, Date: d2, Amount: 1},
			{ClientID: c.ID, RoomID: r.ID, Date: d2, Amount: 1},
			{ClientID: c.ID, RoomID: r.ID, Date: d2, Amount: 1},
			{ClientID: c.ID, RoomID: r.ID, Date: d3, Amount: 1},
			{ClientID: c.ID, RoomID: r.ID, Date: d3, Amount: 1},
			{ClientID: c.ID, RoomID: r.ID, Date: d2, Amount: 1},
			{ClientID: c.ID, RoomID: r.ID, Date: d1, Amount: 1},
		}
		dates = [][]time.Time{
			{d1, d1}, {d2, d1}, {d3, d1}, {d4, d1},
			{d1, d2}, {d2, d2}, {d3, d2}, {d4, d2},
			{d1, d3}, {d2, d3}, {d3, d3}, {d4, d3},
			{d1, d4}, {d2, d4}, {d3, d4}, {d4, d4},
		}
		_ps   []types.Payment
		err   error
		check func(d1, d2 time.Time)
	)

	check = func(d1, d2 time.Time) {
		_ps, err = client.PaymentGetByDateRange(d1, d2)
		if err == nil {
			t.Fatal(_ps)
		}
	}

	for _, grp := range dates {
		check(grp[0], grp[1])
	}

	_ps, err = client.PaymentGetByDateRange(d1, d4)
	if err == nil {
		t.Fatal(_ps)
	}

	client.ClientCreate(c)
	defer client.ClientDelete(c.ID)
	client.RoomCreate(r)
	defer client.RoomDelete(r.ID)

	for _, p := range ps {
		_p, _ := client.PaymentCreate(p.ClientID, p.RoomID, p.Date, p.Amount)
		defer client.PaymentDelete(_p.ID)
	}

	check = func(d1, d2 time.Time) {
		count := 0
		for _, p := range ps {
			if (p.Date.After(d1) || date_equal(d1, p.Date)) &&
				(d2.After(p.Date) || date_equal(d2, p.Date)) {
				count++
			}
		}

		_ps, err = client.PaymentGetByDateRange(d1, d2)
		if err != nil {
			if api_errors.IsChildErr(err, api_errors.ErrSQLNoRows) && count == 0 && len(_ps) == 0 {
				return
			}
			t.Fatal(err)
		}
		if len(_ps) != count {
			bs, _ := json.MarshalIndent(_ps, "", "\t")
			t.Fatalf("d1(%v); d2(%v)\n_ps len (%d) != count (%d):\n%s", d1, d2, len(_ps), count, bs)
		}
	}

	for _, grp := range dates {
		check(grp[0], grp[1])
	}
}

func Test_payment_patch(t *testing.T) {
	var (
		c   = newTestClient(1)
		r   = newTestRoom(c.ID, 1)
		p   = newTestPayment(c.ID, r.ID)
		err error
	)

	p.ID = 12345667
	err = client.PaymentPatch(p)
	if err == nil {
		t.Fatal()
	}

	client.ClientCreate(c)
	defer client.ClientDelete(c.ID)
	client.RoomCreate(r)
	defer client.RoomDelete(r.ID)

	_p, _ := client.PaymentCreate(p.ClientID, p.RoomID, p.Date, p.Amount)
	p.ID = _p.ID
	defer client.PaymentDelete(p.ID)

	p.ClientID = 12345
	err = client.PaymentPatch(p)
	if err == nil {
		t.Fatal()
	}
	p.ClientID = _p.ClientID

	p.RoomID = 12345
	err = client.PaymentPatch(p)
	if err == nil {
		t.Fatal()
	}
	p.RoomID = _p.RoomID

	check := func() {
		err = client.PaymentPatch(p)
		if err != nil {
			t.Fatal(err)
		}
		_p, _ = client.PaymentGetByID(p.ID)

		if _p.ClientID != p.ClientID ||
			_p.RoomID != p.RoomID ||
			!float_equal(_p.Amount, p.Amount) ||
			!date_equal(_p.Date, p.Date) {
			pjson, _ := json.MarshalIndent(p, "", "\t")
			_pjson, _ := json.MarshalIndent(_p, "", "\t")
			t.Fatalf("p:\n%s\n_p:\n%s", pjson, _pjson)
		}
	}

	p.Amount = 1234.1234
	check()

	p.Date = time.Now().Add(time.Hour*12345 + time.Second*12345)
	check()

	p.Amount = 0.9999
	p.Date = time.Now()
	check()
}
