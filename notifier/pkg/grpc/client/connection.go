package client

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Connection struct {
	client *grpc.ClientConn
}

func NewConnection(conf *Config) (*grpc.ClientConn, error) {
	logrus.Infof("Creating gRPC connection to %s...", conf.Address)
	cl, err := grpc.Dial(
		// ctx,
		fmt.Sprintf("%s:%s", conf.Address, conf.Port),
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	logrus.Infof("Created gRPC connection to %s", conf.Address)

	return cl, nil
}

func (c *Connection) GetClient() *grpc.ClientConn {
	return c.client
}

func (c *Connection) Close() error {
	return c.client.Close()
}
