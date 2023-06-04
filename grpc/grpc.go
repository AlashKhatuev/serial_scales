package grpc

import (
	"context"
	"fmt"
	"serial_scales/globals"
	pb "serial_scales/protobuf"
	"strconv"
)

type SerialScaleServer struct {
	pb.UnimplementedApiCallerScaleServer
}

func (s *SerialScaleServer) GetInstantWeight(ctx context.Context, in *pb.Empty) (*pb.ResponseInstantWeight, error) {
	weight, division, tare, err := globals.SerialConn.GetMassa()
	if division == 0 {
		return &pb.ResponseInstantWeight{
			Error:   err.Error(),
			Message: fmt.Sprintf("weight: %d mg, tare: %d mg", weight*100, tare*100),
		}, nil
	}
	if division == 1 {
		return &pb.ResponseInstantWeight{
			Error:   err.Error(),
			Message: fmt.Sprintf("weight: %d g, tare: %d g", weight, tare),
		}, nil
	}
	if division == 2 {
		return &pb.ResponseInstantWeight{
			Error:   err.Error(),
			Message: fmt.Sprintf("weight: %d g, tare: %d g", weight*10, tare*10),
		}, nil
	}
	return &pb.ResponseInstantWeight{
		Error:   err.Error(),
		Message: fmt.Sprintf("weight: %d kg, tare: %d kg", weight, tare),
	}, nil
}

func (s *SerialScaleServer) GetState(ctx context.Context, in *pb.Empty) (*pb.ResponseScale, error) {
	// do nothin
	return nil, nil
}

func (s *SerialScaleServer) ScalesMessageOutChannel(in pb.ApiCallerScale_ScalesMessageOutChannelServer) error {
	// Do nothing
	return nil
}

func (s *SerialScaleServer) SetTare(ctx context.Context, in *pb.Empty) (*pb.ResponseSetScale, error) {
	err := globals.SerialConn.SetTare(0)
	return &pb.ResponseSetScale{
		Error: err.Error(),
	}, nil
}

func (s *SerialScaleServer) SetTareValue(ctx context.Context, in *pb.RequestTareValue) (*pb.ResponseSetScale, error) {
	tare, err := strconv.Atoi(in.Message)
	tareErr := globals.SerialConn.SetTare(int64(tare))
	return &pb.ResponseSetScale{
		Error: err.Error(),
	}, tareErr
}

func (s *SerialScaleServer) SetZero(ctx context.Context, in *pb.Empty) (*pb.ResponseSetScale, error) {
	err := globals.SerialConn.SetZero()
	return &pb.ResponseSetScale{
		Error: err.Error(),
	}, nil
}
