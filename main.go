package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"serial_scales/globals"
	"serial_scales/grpc"
	"serial_scales/messaging"

	pb "serial_scales/protobuf"

	grpc_def "google.golang.org/grpc"
)

func main() {
	var port = flag.String("port", "", "COM-port")
	var baudrate = flag.Int("baudrate", 115200, "COM_port baudrate")
	var size = flag.Int("size", 8, "COM_port baudrate")
	var grpc_port = flag.String("grpc_port", "8080", "gRPC port")

	globals.SerialConn = messaging.NewSerialConnection(*port, *baudrate, 0, *size)

	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", *grpc_port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc_def.NewServer()
	pb.RegisterApiCallerScaleServer(grpcServer, &grpc.SerialScaleServer{})
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
