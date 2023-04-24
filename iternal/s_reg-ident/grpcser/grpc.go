package grpcser

import (
	"context"
	"fmt"
	"log"
	"net"
	"pet/integration_auth/api"

	"google.golang.org/grpc"
)

type GRPCSer struct {
	api.UnimplementedGreeterServer
	Key []byte
}

func (s *GRPCSer) GetKey(ctx context.Context, r *api.Request) (k *api.Key, err error) {
	k = &api.Key{PublicKey: s.Key}
	return k, nil
}

func StartServerGRPC(addrgrcp string, keybyts []byte) {
	log.Println("grpc")

	lis, err := net.Listen("tcp", addrgrcp)
	if err != nil {
		log.Fatal(fmt.Errorf("grpc failed to listen: %v", err))
	}
	s := grpc.NewServer()
	api.RegisterGreeterServer(s, &GRPCSer{Key: keybyts})
	go func() {
		err := s.Serve(lis)
		log.Fatal(err)
	}()

}
