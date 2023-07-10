package server

import (
	"fmt"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	Server      *grpc.Server
	netListener net.Listener
}

func NewServer(cfg *Config) (*Server, error) {
	listener, err := net.Listen(
		cfg.GRPCProtocol,
		fmt.Sprintf("%s:%s", cfg.GRPCAdress, cfg.GRPCPort),
	)
	if err != nil {
		return nil, err
	}

	server := grpc.NewServer()

	return &Server{
		netListener: listener,
		Server:      server,
	}, err
}
