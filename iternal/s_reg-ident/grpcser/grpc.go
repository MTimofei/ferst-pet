package grpcser

import (
	"context"
	"fmt"
	"log"
	"net"
	"pet/integration_auth/api"
	"pet/iternal/s_reg-ident/jwt/ac"

	"google.golang.org/grpc"
)

type GRPCSer struct {
	api.UnimplementedGreeterServer
	Key *ac.KeyAcc
}

func (s *GRPCSer) GetKey(ctx context.Context, r *api.Request) (k *api.Key, err error) {
	k = &api.Key{PublicKey: s.Key.GetPublicKey()}
	return k, nil
}

func StartServerGRPC(addrgrcp string, keyacc *ac.KeyAcc) {
	log.Println("grpc")
	n := &net.ListenConfig{}
	lis, err := n.Listen(context.Background(), "tcp", addrgrcp)
	if err != nil {
		log.Fatal(fmt.Errorf("grpc failed to listen: %v", err))
	}
	s := grpc.NewServer()
	api.RegisterGreeterServer(s, &GRPCSer{Key: keyacc})
	go func() {
		err := s.Serve(lis)
		log.Fatal(err)
	}()

}
