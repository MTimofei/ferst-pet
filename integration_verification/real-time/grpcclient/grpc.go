package grpcclient

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"fmt"
	"log"
	"pet/integration_verification/api"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func RealTimeGetKyeМiaGRPC(addrgrpc *string, keytransfer chan *rsa.PublicKey, wg *sync.WaitGroup) {
	defer wg.Done()
	var timeSlip time.Duration

	for {
		key, err := ResponsGRPC(addrgrpc)
		if err != nil {
			log.Println("gorut grpc:", err)
			timeSlip = 5 * time.Second
		} else {
			timeSlip = time.Minute
			log.Println("RealTimeUpdateKyeМiaGRPC")
		}
		keytransfer <- key
		time.Sleep(timeSlip)
	}
}

func ResponsGRPC(addrgrpc *string) (key *rsa.PublicKey, err error) {
	log.Println("grpc")

	connect, err := grpc.Dial(*addrgrpc, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("grpc wrr con: %v", err)
	}

	defer connect.Close()

	cli := api.NewGreeterClient(connect)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	req := &api.Request{}
	r, err := cli.GetKey(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("grpc err GetKey: %v", err)
	}
	key, err = x509.ParsePKCS1PublicKey(r.PublicKey)
	if err != nil {
		return nil, fmt.Errorf("grpc err key extraction: %v", err)
	}
	return key, nil
}
