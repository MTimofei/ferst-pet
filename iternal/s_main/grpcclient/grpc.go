package grpcclient

import (
	"context"
	"crypto"
	"crypto/x509"
	"fmt"
	"log"
	"pet/api"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ConnectionGRPC(addrgrpc string) (crypto.PublicKey, error) {
	log.Println("grpc")

	conn, err := grpc.Dial(addrgrpc, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return &api.Key{}, fmt.Errorf("grpc wrr con: %v", err)
	}

	defer conn.Close()

	cli := api.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	req := &api.Request{}
	r, err := cli.GetKey(ctx, req)
	if err != nil {
		return &api.Key{}, fmt.Errorf("grpc err GetKey: %v", err)
	}
	key, err := x509.ParsePKCS1PublicKey(r.PublicKey)
	if err != nil {
		return &api.Key{}, fmt.Errorf("grpc err key extraction: %v", err)
	}
	return key, nil
}
