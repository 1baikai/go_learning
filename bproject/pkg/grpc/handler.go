package grpc

import (
	"context"

	endpoint "bproject/pkg/endpoint"
	pb "bproject/pkg/grpc/pb"
	grpc "github.com/go-kit/kit/transport/grpc"
	context1 "golang.org/x/net/context"
)

// makeAddUserHandler creates the handler logic
func makeAddUserHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.AddUserEndpoint, decodeAddUserRequest, encodeAddUserResponse, options...)
}

// decodeAddUserResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain AddUser request.
// TODO implement the decoder
func decodeAddUserRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.AddUserRequest)
	return endpoint.AddUserRequest{
		Username: req.Username,
		Fraction: req.Fraction,
	}, nil
}

// encodeAddUserResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC req := r.(*pb.UpDataWalAndBlockSizeRequest)
func encodeAddUserResponse(_ context.Context, r interface{}) (interface{}, error) {
	reply := r.(endpoint.AddUserResponse)
	if reply.Err != nil {
		return nil, reply.Err
	}
	return &pb.AddUserReply{}, nil
}
func (g *grpcServer) AddUser(ctx context1.Context, req *pb.AddUserRequest) (*pb.AddUserReply, error) {
	_, rep, err := g.addUser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.AddUserReply), nil
}

// makeGetUserHandler creates the handler logic
func makeGetUserHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.GetUserEndpoint, decodeGetUserRequest, encodeGetUserResponse, options...)
}

// decodeGetUserResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain GetUser request.
// TODO implement the decoder
func decodeGetUserRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.GetUserRequest)
	return endpoint.GetUserRequest{
		UserID: req.UserID,
	}, nil
}

// encodeGetUserResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeGetUserResponse(_ context.Context, r interface{}) (interface{}, error) {
	reply := r.(endpoint.GetUserResponse)
	if reply.Err != nil {
		return nil, reply.Err
	}
	return &pb.GetUserReply{
		Username: reply.Res.Username,
		Fraction: reply.Res.Fraction,
	}, nil
}
func (g *grpcServer) GetUser(ctx context1.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
	_, rep, err := g.getUser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetUserReply), nil
}
