package service

import (
	"context"
	"math/rand"
	pb "weather_and_bus/genproto/transport"
	"weather_and_bus/storage/postgres"
)

type Transport struct {
	pb.UnimplementedTransportServerServer
	TransportRepo *postgres.TransportRepo
}

func NewTransport(transportRepo *postgres.TransportRepo) *Transport {
	return &Transport{TransportRepo: transportRepo}
}

func (t *Transport) GetBusSchedule(ctx context.Context, req *pb.RequestBusSchedule) (*pb.ResponceBusSchedule, error) {
	bus, err := t.TransportRepo.GetByBusNumber(int(req.BusNumber))
	if err != nil {
		return &pb.ResponceBusSchedule{}, err
	}
	return &pb.ResponceBusSchedule{Schedules: bus.Stations}, nil
}

func (t *Transport) TrackBusLocation(ctx context.Context, req *pb.RequestBusLocation) (*pb.ResponceBusLocation, error) {
	bus, err := t.TransportRepo.GetByBusNumber(int(req.BusNumber))
	if err != nil {
		return &pb.ResponceBusLocation{}, err
	}
	ind := rand.Intn(len(bus.Stations))
	curStation := bus.Stations[ind]
	return &pb.ResponceBusLocation{StationName: curStation}, err
}

func (t *Transport) ReportTrafficJam(ctx context.Context, req *pb.RequestTrafficJam) (*pb.ResponceTrafficJam, error) {
	return &pb.ResponceTrafficJam{Success: true}, nil
}
