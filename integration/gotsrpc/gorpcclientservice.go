package gotsrpc

import (
	"context"
)

type GoRPCClientService struct {
	client GoRPCClient
	cancel context.CancelFunc
}

func NewGoRPCClientService(client GoRPCClient) *GoRPCClientService {
	return &GoRPCClientService{
		client: client,
	}
}

func (s *GoRPCClientService) Start(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	s.cancel = cancel
	s.client.Start()
	<-ctx.Done()
	s.client.Stop()

	return ctx.Err()
}

func (s *GoRPCClientService) Close(ctx context.Context) error {
	if s.cancel != nil {
		s.cancel()
		s.cancel = nil
	}

	return nil
}
