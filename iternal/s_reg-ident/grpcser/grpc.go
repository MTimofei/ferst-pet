package grpcser

import (
	"context"
	"fmt"
	"log"
	"net"
	"pet/api"

	"google.golang.org/grpc"
)

type GRPCSer struct {
	api.UnimplementedGreeterServer
	K []byte
}

func (s *GRPCSer) GetKey(ctx context.Context, r *api.Request) (k *api.Key, err error) {
	k = &api.Key{PublicKey: s.K}
	return k, nil
}

func ConnectionGRPC(addrgrcp string, kbyt []byte) {
	log.Println("grpc")
	// n := &net.ListenConfig{}
	// lis, err := n.Listen(context.TODO(), "tcp", addrgrcp)
	lis, err := net.Listen("tcp", addrgrcp)
	if err != nil {
		log.Fatal(fmt.Errorf("grpc failed to listen: %v", err))
	}
	s := grpc.NewServer()
	api.RegisterGreeterServer(s, &GRPCSer{K: kbyt})
	go func() {
		err := s.Serve(lis)
		log.Fatal(err)
	}()

}
