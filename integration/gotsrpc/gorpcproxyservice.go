package gotsrpc

import (
	"context"
)

type GoRPCProxyService struct {
	proxy  GoRPCProxy
	cancel context.CancelFunc
}

func NewGoRPCProxyService(proxy GoRPCProxy) *GoRPCProxyService {
	return &GoRPCProxyService{
		proxy: proxy,
	}
}

func (s *GoRPCProxyService) Start(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)

	s.cancel = cancel
	if err := s.proxy.Start(); err != nil {
		return err
	}

	<-ctx.Done()
	s.proxy.Stop()

	return ctx.Err()
}

func (s *GoRPCProxyService) Close(ctx context.Context) error {
	if s.cancel != nil {
		s.cancel()
		s.cancel = nil
	}

	return nil
}
