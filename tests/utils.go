package api_client

import (
	"math/rand"
	"os"
	"time"

	"github.com/snakehunterr/hacs_dbapi_client"
	types "github.com/snakehunterr/hacs_dbapi_types"
)

var client = api_client.New(os.Getenv("DBAPI_SERVER_HOST"), os.Getenv("DBAPI_SERVER_PORT"))

func newTestClient(client_id int64) *types.Client {
	return &types.Client{
		ID:   client_id,
		Name: "Foo",
	}
}

func newTestRoom(client_id int64, room_id int64) *types.Room {
	return &types.Room{
		ID:          room_id,
		ClientID:    client_id,
		Area:        rand.Float64(),
		PeopleCount: uint8(rand.Uint32()),
	}
}

func newTestPayment(client_id int64, room_id int64) *types.Payment {
	return &types.Payment{
		ClientID: client_id,
		RoomID:   room_id,
		Date:     time.Now(),
		Amount:   rand.Float64(),
	}
}
