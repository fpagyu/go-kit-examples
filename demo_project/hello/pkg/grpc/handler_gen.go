// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package grpc

import (
	grpc "github.com/go-kit/kit/transport/grpc"
	pb "hello/pb/gen-go/pb"
	endpoint "hello/pkg/endpoint"
)

// NewGRPCServer makes a set of endpoints available as a gRPC AddServer
type grpcServer struct {
	sayHi          grpc.Handler
	makeADate      grpc.Handler
	updateUserInfo grpc.Handler
}

func NewGRPCServer(endpoints endpoint.Endpoints, options map[string][]grpc.ServerOption) pb.HelloServer {
	return &grpcServer{
		makeADate:      makeMakeADateHandler(endpoints, options["MakeADate"]),
		sayHi:          makeSayHiHandler(endpoints, options["SayHi"]),
		updateUserInfo: makeUpdateUserInfoHandler(endpoints, options["UpdateUserInfo"]),
	}
}
