package gotsrpc

type GoRPCProxy interface {
	Start() error
	Stop()
}
