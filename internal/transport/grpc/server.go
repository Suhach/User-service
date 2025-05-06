package grpc

import "net"

func RunGRPC(svc *user.Service) error {
	net.Listen("tcp", ":50051")
	grpc.NewServer()
	userpb.RegisterUserServiceServer(grpcSrv, NewHandler(svc))
	grpcSrv.Serve(listener)
}
