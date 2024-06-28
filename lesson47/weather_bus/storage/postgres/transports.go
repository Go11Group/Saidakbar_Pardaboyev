package postgres

import (
	"database/sql"
	pb "weather_and_bus/genproto/transport"

	"github.com/lib/pq"
)

type TransportRepo struct {
	DB *sql.DB
}

func NewTransportRepo(db *sql.DB) *TransportRepo {
	return &TransportRepo{DB: db}
}

func (t *TransportRepo) GetByBusNumber(busNumber int) (*pb.Transport, error) {
	query := `select * from transports where bus_number = $1`

	busInfo := pb.Transport{}
	err := t.DB.QueryRow(query, busNumber).Scan(&busInfo.BusNumber, pq.Array(&busInfo.Stations))
	return &busInfo, err
}
