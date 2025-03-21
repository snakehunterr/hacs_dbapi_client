package api_client

import (
	"strconv"

	types "github.com/snakehunterr/hacs_dbapi_types"
	"github.com/snakehunterr/hacs_dbapi_types/validators"
)

func formatFloat64(f float64) string {
	return strconv.FormatFloat(f, 'f', 2, 64)
}

func formatUint8(u uint8) string {
	return strconv.FormatUint(uint64(u), 10)
}

func formatInt64(i int64) string {
	return strconv.FormatInt(i, 10)
}

func formatBool(b bool) string {
	return strconv.FormatBool(b)
}

func formatClient(c *types.Client) Form {
	return Form{
		"client_name": c.Name,
		"is_admin":    formatBool(c.IsAdmin),
	}
}

func formatRoom(r *types.Room) Form {
	return Form{
		"client_id":         formatInt64(r.ClientID),
		"room_area":         formatFloat64(r.Area),
		"room_people_count": formatUint8(r.PeopleCount),
	}
}

func formatPayment(p *types.Payment) Form {
	return Form{
		"client_id":      formatInt64(p.ClientID),
		"room_id":        formatInt64(p.RoomID),
		"payment_date":   p.Date.Format(validators.DATE_FORMAT),
		"payment_amount": formatFloat64(p.Amount),
	}
}

func formatExpense(e *types.Expense) Form {
	return Form{
		"expense_date":   e.Date.Format(validators.DATE_FORMAT),
		"expense_amount": formatFloat64(e.Amount),
	}
}
