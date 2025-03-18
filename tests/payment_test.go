package api_client

// package api_client

// import (
// 	"os"
// 	"testing"
// 	"time"

// 	types "github.com/snakehunterr/hacs_dbapi_types"
// 	"github.com/snakehunterr/hacs_dbapi_types/validators"
// )

// func Test_payment_get_by_id(t *testing.T) {
// 	var (
// 		client = NewAPIClient(os.Getenv("HACS_DBAPI_HOST"), os.Getenv("HACS_DBAPI_PORT"))
// 		c      = &types.Client{
// 			ID:   1,
// 			Name: "foo",
// 		}
// 		r = &types.Room{
// 			ID:          1,
// 			ClientID:    c.ID,
// 			Area:        1,
// 			PeopleCount: 1,
// 		}
// 		p = &types.Payment{
// 			ClientID: c.ID,
// 			RoomID:   r.ID,
// 			Amount:   100.50,
// 			Date:     time.Now(),
// 		}
// 		res *types.APIResponse
// 		err error
// 	)

// 	if _, err := client.ClientCreate(c); err != nil {
// 		t.Fatal("ClientCreate:", err)
// 	}
// 	defer func() {
// 		if _, err := client.ClientDelete(c.ID); err != nil {
// 			t.Fatal("ClientDelete:", err)
// 		}
// 	}()

// 	if _, err := client.RoomCreate(r); err != nil {
// 		t.Fatal("RoomCreate:", err)
// 	}
// 	defer func() {
// 		if _, err := client.RoomDelete(r.ID); err != nil {
// 			t.Fatal("RoomDelete:", err)
// 		}
// 	}()

// 	res, err = client.PaymentCreate(p)
// 	if err != nil {
// 		t.Fatal("PaymentCreate:", err)
// 	}
// 	if res != nil && res.Error != nil {
// 		t.Fatal("PaymentCreate *APIError:", res.Error)
// 	}
// 	if p.ID == 0 {
// 		t.Fatal("PaymentCreate new payment ID is 0")
// 	}
// 	defer func() {
// 		res, err = client.PaymentDelete(p.ID)
// 		if err != nil {
// 			t.Fatal("PaymentDelete:", err)
// 		}
// 		if res == nil {
// 			t.Fatal("PaymentDelete *types.APIResponse is nil")
// 		}
// 		if res.Error != nil {
// 			t.Fatal("PaymentDelete *APIError:", res.Error)
// 		}
// 	}()

// 	_p, res, err := client.PaymentGetByID(p.ID)
// 	if err != nil {
// 		t.Fatal("PaymentGetByID:", err)
// 	}
// 	if res != nil {
// 		t.Fatal("PaymentGetByID *types.APIResponse:", err)
// 	}
// 	if _p == nil {
// 		t.Fatal("PaymentGetByID *types.Payment is nil")
// 	}
// 	if _p.ID != p.ID {
// 		t.Fatalf("PaymentGetByID wrong payment_id: %v != %v", _p.ID, p.ID)
// 	}
// 	if _p.ClientID != p.ClientID {
// 		t.Fatalf("PaymentGetByID wrong client_id: %v != %v", _p.ClientID, p.ClientID)
// 	}
// 	if _p.RoomID != p.RoomID {
// 		t.Fatalf("PaymentGetByID wrong room_id: %v != %v", _p.RoomID, p.RoomID)
// 	}
// 	if _p.Date.Format(validators.DATE_FORMAT) != p.Date.Format(validators.DATE_FORMAT) {
// 		t.Fatalf("PaymentGetByID wrong payment_date: %v != %v", _p.Date, p.Date)
// 	}
// 	if _p.Amount != p.Amount {
// 		t.Fatalf("PaymentGetByID wrong payment_amount: %v != %v", _p.Amount, p.Amount)
// 	}
// }

// func Test_payment_get_all(t *testing.T) {
// 	var (
// 		client = NewAPIClient(os.Getenv("HACS_DBAPI_HOST"), os.Getenv("HACS_DBAPI_PORT"))
// 		c      = &types.Client{
// 			ID:   1,
// 			Name: "foo",
// 		}
// 		r = &types.Room{
// 			ID:          1,
// 			ClientID:    c.ID,
// 			Area:        1,
// 			PeopleCount: 1,
// 		}
// 	)

// 	if _, err := client.ClientCreate(c); err != nil {
// 		t.Fatal("ClientCreate:", err)
// 	}
// 	defer func() {
// 		if _, err := client.ClientDelete(c.ID); err != nil {
// 			t.Fatal("ClientDelete:", err)
// 		}
// 	}()

// 	if _, err := client.RoomCreate(r); err != nil {
// 		t.Fatal("RoomCreate:", err)
// 	}
// 	defer func() {
// 		if _, err := client.RoomDelete(r.ID); err != nil {
// 			t.Fatal("RoomDelete:", err)
// 		}
// 	}()

// 	ps := []*types.Payment{
// 		{ClientID: c.ID, RoomID: r.ID, Date: time.Now(), Amount: 1},
// 		{ClientID: c.ID, RoomID: r.ID, Date: time.Now(), Amount: 2},
// 		{ClientID: c.ID, RoomID: r.ID, Date: time.Now(), Amount: 3},
// 		{ClientID: c.ID, RoomID: r.ID, Date: time.Now(), Amount: 4},
// 	}

// 	for _, p := range ps {
// 		if _, err := client.PaymentCreate(p); err != nil {
// 			t.Fatal("PaymentCreate:", err)
// 		}
// 	}

// 	_ps, res, err := client.PaymentGetAll()
// 	if err != nil {
// 		t.Fatal("PaymentGetAll:", err)
// 	}
// 	if res != nil {
// 		t.Fatal("PaymentGetAll *types.APIResponse:", res)
// 	}
// 	if len(_ps) != len(ps) {
// 		t.Fatal("PaymentGetAll returned:", _ps)
// 	}

// 	for _, p := range ps {
// 		if res, err := client.PaymentDelete(p.ID); err != nil {
// 			t.Fatal("PaymentDelete:", err)
// 		} else if res.Error != nil {
// 			t.Fatal("PaymentDelete *APIError:", res.Error)
// 		}
// 	}
// }

// func Test_payment_get_by_client_id(t *testing.T) {
// 	var (
// 		client = NewAPIClient(os.Getenv("HACS_DBAPI_HOST"), os.Getenv("HACS_DBAPI_PORT"))
// 		c1     = &types.Client{
// 			ID:   1,
// 			Name: "foo",
// 		}
// 		c2 = &types.Client{
// 			ID:   2,
// 			Name: "foo",
// 		}
// 		r = &types.Room{
// 			ID:          1,
// 			ClientID:    c1.ID,
// 			Area:        1,
// 			PeopleCount: 1,
// 		}
// 	)

// 	if _, err := client.ClientCreate(c1); err != nil {
// 		t.Fatal("ClientCreate:", err)
// 	}
// 	defer func() {
// 		if _, err := client.ClientDelete(c1.ID); err != nil {
// 			t.Fatal("ClientDelete:", err)
// 		}
// 	}()

// 	if _, err := client.ClientCreate(c2); err != nil {
// 		t.Fatal("ClientCreate:", err)
// 	}
// 	defer func() {
// 		if _, err := client.ClientDelete(c2.ID); err != nil {
// 			t.Fatal("ClientDelete:", err)
// 		}
// 	}()

// 	if _, err := client.RoomCreate(r); err != nil {
// 		t.Fatal("RoomCreate:", err)
// 	}
// 	defer func() {
// 		if _, err := client.RoomDelete(r.ID); err != nil {
// 			t.Fatal("RoomDelete:", err)
// 		}
// 	}()

// 	var ps = []*types.Payment{
// 		{ClientID: c1.ID, RoomID: r.ID, Date: time.Now(), Amount: 1},
// 		{ClientID: c2.ID, RoomID: r.ID, Date: time.Now(), Amount: 1},
// 		{ClientID: c1.ID, RoomID: r.ID, Date: time.Now(), Amount: 1},
// 		{ClientID: c2.ID, RoomID: r.ID, Date: time.Now(), Amount: 1},
// 		{ClientID: c1.ID, RoomID: r.ID, Date: time.Now(), Amount: 1},
// 	}

// 	for _, p := range ps {
// 		if _, err := client.PaymentCreate(p); err != nil {
// 			t.Fatal("PaymentCreate:", err)
// 		}
// 	}

// 	temp := []*types.Payment{}
// 	for _, p := range ps {
// 		if p.ClientID == c1.ID {
// 			temp = append(temp, p)
// 		}
// 	}
// 	_ps, res, err := client.PaymentGetAllByClientID(c1.ID)
// 	if err != nil {
// 		t.Fatal("PaymentGetAllByClientID:", err)
// 	}
// 	if res != nil {
// 		t.Fatal("PaymentGetAllByClientID *types.APIResponse:", res)
// 	}
// 	if len(_ps) != len(temp) {
// 		t.Fatal("PaymentGetAllByClientID returns:", _ps)
// 	}

// 	temp = []*types.Payment{}
// 	for _, p := range ps {
// 		if p.ClientID == c2.ID {
// 			temp = append(temp, p)
// 		}
// 	}
// 	_ps, res, err = client.PaymentGetAllByClientID(c2.ID)
// 	if err != nil {
// 		t.Fatal("PaymentGetAllByClientID:", err)
// 	}
// 	if res != nil {
// 		t.Fatal("PaymentGetAllByClientID *types.APIResponse:", res)
// 	}
// 	if len(_ps) != len(temp) {
// 		t.Fatal("PaymentGetAllByClientID returns:", _ps)
// 	}

// 	for _, p := range ps {
// 		if res, err := client.PaymentDelete(p.ID); err != nil {
// 			t.Fatal("PaymentDelete:", err)
// 		} else if res.Error != nil {
// 			t.Fatal("PaymentDelete *APIError:", res.Error)
// 		}
// 	}
// }

// func Test_payment_get_by_room_id(t *testing.T) {
// 	var (
// 		client = NewAPIClient(os.Getenv("HACS_DBAPI_HOST"), os.Getenv("HACS_DBAPI_PORT"))
// 		c      = &types.Client{
// 			ID:   1,
// 			Name: "foo",
// 		}
// 		r1 = &types.Room{
// 			ID:          1,
// 			ClientID:    c.ID,
// 			Area:        1,
// 			PeopleCount: 1,
// 		}
// 		r2 = &types.Room{
// 			ID:          2,
// 			ClientID:    c.ID,
// 			Area:        1,
// 			PeopleCount: 1,
// 		}
// 	)

// 	if _, err := client.ClientCreate(c); err != nil {
// 		t.Fatal("ClientCreate:", err)
// 	}
// 	defer func() {
// 		if _, err := client.ClientDelete(c.ID); err != nil {
// 			t.Fatal("ClientDelete:", err)
// 		}
// 	}()

// 	if _, err := client.RoomCreate(r1); err != nil {
// 		t.Fatal("RoomCreate:", err)
// 	}
// 	defer func() {
// 		if _, err := client.RoomDelete(r1.ID); err != nil {
// 			t.Fatal("RoomDelete:", err)
// 		}
// 	}()

// 	if _, err := client.RoomCreate(r2); err != nil {
// 		t.Fatal("RoomCreate:", err)
// 	}
// 	defer func() {
// 		if _, err := client.RoomDelete(r2.ID); err != nil {
// 			t.Fatal("RoomDelete:", err)
// 		}
// 	}()

// 	var ps = []*types.Payment{
// 		{ClientID: c.ID, RoomID: r1.ID, Date: time.Now(), Amount: 1},
// 		{ClientID: c.ID, RoomID: r2.ID, Date: time.Now(), Amount: 1},
// 		{ClientID: c.ID, RoomID: r1.ID, Date: time.Now(), Amount: 1},
// 		{ClientID: c.ID, RoomID: r2.ID, Date: time.Now(), Amount: 1},
// 		{ClientID: c.ID, RoomID: r1.ID, Date: time.Now(), Amount: 1},
// 	}

// 	for _, p := range ps {
// 		if _, err := client.PaymentCreate(p); err != nil {
// 			t.Fatal("PaymentCreate:", err)
// 		}
// 	}

// 	temp := []*types.Payment{}
// 	for _, p := range ps {
// 		if p.RoomID == r1.ID {
// 			temp = append(temp, p)
// 		}
// 	}
// 	_ps, res, err := client.PaymentGetAllByRoomID(r1.ID)
// 	if err != nil {
// 		t.Fatal("PaymentGetAllByRoomID:", err)
// 	}
// 	if res != nil {
// 		t.Fatal("PaymentGetAllByRoomID *types.APIResponse:", res)
// 	}
// 	if len(_ps) != len(temp) {
// 		t.Fatal("PaymentGetAllByRoomID returns:", _ps)
// 	}

// 	temp = []*types.Payment{}
// 	for _, p := range ps {
// 		if p.RoomID == r2.ID {
// 			temp = append(temp, p)
// 		}
// 	}
// 	_ps, res, err = client.PaymentGetAllByRoomID(r2.ID)
// 	if err != nil {
// 		t.Fatal("PaymentGetAllByRoomID:", err)
// 	}
// 	if res != nil {
// 		t.Fatal("PaymentGetAllByRoomID *types.APIResponse:", res)
// 	}
// 	if len(_ps) != len(temp) {
// 		t.Fatal("PaymentGetAllByRoomID returns:", _ps)
// 	}

// 	for _, p := range ps {
// 		if res, err := client.PaymentDelete(p.ID); err != nil {
// 			t.Fatal("PaymentDelete:", err)
// 		} else if res.Error != nil {
// 			t.Fatal("PaymentDelete *APIError:", res.Error)
// 		}
// 	}
// }

// func Test_payment_get_by_date(t *testing.T) {
// 	var (
// 		client = NewAPIClient(os.Getenv("HACS_DBAPI_HOST"), os.Getenv("HACS_DBAPI_PORT"))
// 		c      = &types.Client{
// 			ID:   1,
// 			Name: "foo",
// 		}
// 		r = &types.Room{
// 			ID:          1,
// 			ClientID:    c.ID,
// 			Area:        1,
// 			PeopleCount: 1,
// 		}
// 	)

// 	if _, err := client.ClientCreate(c); err != nil {
// 		t.Fatal("ClientCreate:", err)
// 	}
// 	defer func() {
// 		if _, err := client.ClientDelete(c.ID); err != nil {
// 			t.Fatal("ClientDelete:", err)
// 		}
// 	}()

// 	if _, err := client.RoomCreate(r); err != nil {
// 		t.Fatal("RoomCreate:", err)
// 	}
// 	defer func() {
// 		if _, err := client.RoomDelete(r.ID); err != nil {
// 			t.Fatal("RoomDelete:", err)
// 		}
// 	}()

// 	t1, _ := time.Parse("2006-01-02", "2024-01-10")
// 	t2, _ := time.Parse("2006-01-02", "2024-01-11")

// 	ps := []*types.Payment{
// 		{ClientID: c.ID, RoomID: r.ID, Date: t1, Amount: 1},
// 		{ClientID: c.ID, RoomID: r.ID, Date: t2, Amount: 1},
// 		{ClientID: c.ID, RoomID: r.ID, Date: t1, Amount: 1},
// 	}

// 	for _, p := range ps {
// 		if _, err := client.PaymentCreate(p); err != nil {
// 			t.Fatal("PaymentCreate:", err)
// 		}
// 	}

// 	_ps, res, err := client.PaymentGetByDate(t1)
// 	if err != nil {
// 		t.Fatal("PaymentGetByDate:", err)
// 	}
// 	if res != nil {
// 		t.Fatal("PaymentGetByDate *types.APIResponse:", res)
// 	}
// 	if len(_ps) != 2 {
// 		t.Fatal("PaymentGetByDate returns:", _ps)
// 	}

// 	_ps, res, err = client.PaymentGetByDate(t2)
// 	if err != nil {
// 		t.Fatal("PaymentGetByDate:", err)
// 	}
// 	if res != nil {
// 		t.Fatal("PaymentGetByDate *types.APIResponse:", res)
// 	}
// 	if len(_ps) != 1 {
// 		t.Fatal("PaymentGetByDate returns:", _ps)
// 	}

// 	for _, p := range ps {
// 		res, err := client.PaymentDelete(p.ID)
// 		if err != nil {
// 			t.Fatal("PaymentDelete:", err)
// 		}
// 		if res.Error != nil {
// 			t.Fatal("PaymentDelete *APIError:", res.Error)
// 		}
// 	}
// }

// func Test_payment_get_by_date_range(t *testing.T) {
// 	var (
// 		client = NewAPIClient(os.Getenv("HACS_DBAPI_HOST"), os.Getenv("HACS_DBAPI_PORT"))
// 		c      = &types.Client{
// 			ID:   1,
// 			Name: "foo",
// 		}
// 		r = &types.Room{
// 			ID:          1,
// 			ClientID:    c.ID,
// 			Area:        1,
// 			PeopleCount: 1,
// 		}
// 	)

// 	if _, err := client.ClientCreate(c); err != nil {
// 		t.Fatal("ClientCreate:", err)
// 	}
// 	defer func() {
// 		if _, err := client.ClientDelete(c.ID); err != nil {
// 			t.Fatal("ClientDelete:", err)
// 		}
// 	}()

// 	if _, err := client.RoomCreate(r); err != nil {
// 		t.Fatal("RoomCreate:", err)
// 	}
// 	defer func() {
// 		if _, err := client.RoomDelete(r.ID); err != nil {
// 			t.Fatal("RoomDelete:", err)
// 		}
// 	}()

// 	t1, _ := time.Parse("2006-01-02", "2024-01-10")
// 	t2, _ := time.Parse("2006-01-02", "2024-01-20")
// 	t3, _ := time.Parse("2006-01-02", "2024-01-30")
// 	t4, _ := time.Parse("2006-01-02", "2024-02-20")

// 	ps := []*types.Payment{
// 		{ClientID: c.ID, RoomID: r.ID, Date: t1, Amount: 1},
// 		{ClientID: c.ID, RoomID: r.ID, Date: t2, Amount: 1},
// 		{ClientID: c.ID, RoomID: r.ID, Date: t3, Amount: 1},
// 	}

// 	for _, p := range ps {
// 		if _, err := client.PaymentCreate(p); err != nil {
// 			t.Fatal("PaymentCreate:", err)
// 		}
// 	}

// 	_ps, res, err := client.PaymentGetByDateRange(t1, t1)
// 	if err != nil {
// 		t.Fatal("PaymentGetByDateRange:", err)
// 	}
// 	if res != nil {
// 		t.Fatal("PaymentGetByDateRange *types.APIResponse:", res)
// 	}
// 	if len(_ps) != 1 {
// 		t.Fatal("PaymentGetByDateRange returns:", _ps)
// 	}

// 	_ps, res, err = client.PaymentGetByDateRange(t1, t2)
// 	if err != nil {
// 		t.Fatal("PaymentGetByDateRange:", err)
// 	}
// 	if res != nil {
// 		t.Fatal("PaymentGetByDateRange *types.APIResponse:", res)
// 	}
// 	if len(_ps) != 2 {
// 		t.Fatal("PaymentGetByDateRange returns:", _ps)
// 	}

// 	_ps, res, err = client.PaymentGetByDateRange(t1, t3)
// 	if err != nil {
// 		t.Fatal("PaymentGetByDateRange:", err)
// 	}
// 	if res != nil {
// 		t.Fatal("PaymentGetByDateRange *types.APIResponse:", res)
// 	}
// 	if len(_ps) != 3 {
// 		t.Fatal("PaymentGetByDateRange returns:", _ps)
// 	}

// 	_ps, res, err = client.PaymentGetByDateRange(t1, t4)
// 	if err != nil {
// 		t.Fatal("PaymentGetByDateRange:", err)
// 	}
// 	if res != nil {
// 		t.Fatal("PaymentGetByDateRange *types.APIResponse:", res)
// 	}
// 	if len(_ps) != 3 {
// 		t.Fatal("PaymentGetByDateRange returns:", _ps)
// 	}

// 	for _, p := range ps {
// 		res, err := client.PaymentDelete(p.ID)
// 		if err != nil {
// 			t.Fatal("PaymentDelete:", err)
// 		}
// 		if res.Error != nil {
// 			t.Fatal("PaymentDelete *APIError:", res.Error)
// 		}
// 	}
// }

// func Test_payment_patch(t *testing.T) {
// 	var (
// 		client = NewAPIClient(os.Getenv("HACS_DBAPI_HOST"), os.Getenv("HACS_DBAPI_PORT"))
// 		c1     = &types.Client{
// 			ID:   1,
// 			Name: "foo",
// 		}
// 		c2 = &types.Client{
// 			ID:   2,
// 			Name: "foo",
// 		}
// 		r1 = &types.Room{
// 			ID:          1,
// 			ClientID:    c1.ID,
// 			Area:        1,
// 			PeopleCount: 1,
// 		}
// 		r2 = &types.Room{
// 			ID:          2,
// 			ClientID:    c2.ID,
// 			Area:        1,
// 			PeopleCount: 1,
// 		}
// 	)

// 	if _, err := client.ClientCreate(c1); err != nil {
// 		t.Fatal("ClientCreate:", err)
// 	}
// 	defer func() {
// 		if _, err := client.ClientDelete(c1.ID); err != nil {
// 			t.Fatal("ClientDelete:", err)
// 		}
// 	}()
// 	if _, err := client.ClientCreate(c2); err != nil {
// 		t.Fatal("ClientCreate:", err)
// 	}
// 	defer func() {
// 		if _, err := client.ClientDelete(c2.ID); err != nil {
// 			t.Fatal("ClientDelete:", err)
// 		}
// 	}()

// 	if _, err := client.RoomCreate(r1); err != nil {
// 		t.Fatal("RoomCreate:", err)
// 	}
// 	defer func() {
// 		if _, err := client.RoomDelete(r1.ID); err != nil {
// 			t.Fatal("RoomDelete:", err)
// 		}
// 	}()
// 	if _, err := client.RoomCreate(r2); err != nil {
// 		t.Fatal("RoomCreate:", err)
// 	}
// 	defer func() {
// 		if _, err := client.RoomDelete(r2.ID); err != nil {
// 			t.Fatal("RoomDelete:", err)
// 		}
// 	}()

// 	p := &types.Payment{
// 		ClientID: c1.ID,
// 		RoomID:   r1.ID,
// 		Date:     time.Now(),
// 		Amount:   10.50,
// 	}
// 	res, err := client.PaymentCreate(p)
// 	if err != nil {
// 		t.Fatal("PaymentCreate:", err)
// 	}
// 	if res != nil && res.Error != nil {
// 		t.Fatal("PaymentCreate *APIError:", res.Error)
// 	}
// 	defer func() {
// 		res, err := client.PaymentDelete(p.ID)
// 		if err != nil {
// 			t.Fatal("PaymentDelete:", err)
// 		}
// 		if res.Error != nil {
// 			t.Fatal("PaymentDelete *APIError:", res.Error)
// 		}
// 	}()

// 	check := func() {
// 		res, err := client.PaymentPatch(p)
// 		if err != nil {
// 			t.Fatal("PaymentPatch:", err)
// 		}
// 		if res == nil {
// 			t.Fatal("PaymentPatch *types.APIResponse is nil")
// 		}
// 		if res.Error != nil {
// 			t.Fatal("PaymentPatch *APIError:", res.Error)
// 		}

// 		_p, res, err := client.PaymentGetByID(p.ID)
// 		if err != nil {
// 			t.Fatal("PaymentGetByID:", err)
// 		}
// 		if res != nil {
// 			t.Fatal("PaymentGetByID *types.APIResponse:", res)
// 		}
// 		if _p == nil {
// 			t.Fatal("PaymentGetByID *types.Payment is nil")
// 		}
// 		if _p.ClientID != p.ClientID {
// 			t.Fatal("PaymentGetByID returns incorrect ClientID")
// 		}
// 		if _p.RoomID != p.RoomID {
// 			t.Fatal("PaymentGetByID returns incorrect RoomID")
// 		}
// 		if p.Date.Format(validators.DATE_FORMAT) != _p.Date.Format(validators.DATE_FORMAT) {
// 			t.Fatalf("PaymentGetByID returns incorrect Date (%v != %v)", _p.Date, p.Date)
// 		}
// 		if _p.Amount != p.Amount {
// 			t.Fatal("PaymentGetByID returns incorrect Amount")
// 		}
// 	}

// 	p.ClientID = c2.ID
// 	check()

// 	p.ClientID = c1.ID
// 	p.RoomID = r2.ID
// 	p.Amount = 3.88
// 	check()

// 	p.Date = time.Now().Add(time.Hour*24*31*23 + time.Minute*12345 + time.Second*12345)
// 	check()
// }
