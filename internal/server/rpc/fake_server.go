package rpc

import (
	"easyagent/internal/proto"
	"golang.org/x/net/context"
)

type fakeRpcServer struct{}

func (rpc *fakeRpcServer) RegisterSidecar(ctx context.Context, request *proto.RegisterRequest) (*proto.RegisterResponse, error) {
	return &proto.RegisterResponse{}, nil
}

func (rpc *fakeRpcServer) ReadyForControl(request *proto.ControlRequest, stream proto.EasyAgentService_ReadyForControlServer) error {
	return nil
}

func (rpc *fakeRpcServer) ReportEvent(ctx context.Context, event *proto.Event) (*proto.EmptyResponse, error) {
	return &proto.EmptyResponse{}, nil
}
