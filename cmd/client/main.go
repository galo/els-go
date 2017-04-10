package main

import (
	"flag"
	"fmt"
	"github.com/galo/els-go/pkg/api"
	"github.com/galo/els-go/pkg/elscli"
	"google.golang.org/grpc"
	"os"
	"time"
)

func main() {
	// The elscli presumes no service discovery system, and expects users to
	// provide the direct address of an elssvc.

	var (
		grpcAddr = flag.String("grpc.addr", "", "gRPC (HTTP) address of elssvc")
		method   = flag.String("method", "GetServiceInstanceByKey", "GetServiceInstanceByKey routingKey")
	)
	flag.Parse()

	if len(flag.Args()) != 1 {
		fmt.Fprintf(os.Stderr, "usage: elscli [flags] <routingKey> \n")
		os.Exit(1)
	}

	conn, err := grpc.Dial(*grpcAddr, grpc.WithInsecure(), grpc.WithTimeout(time.Second))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v", err)
		os.Exit(1)
	}
	defer conn.Close()

	client := api.NewElsClient(conn)

	//grpcclient.GetServiceInstanceByKey()

	switch *method {
	case "GetServiceInstanceByKey":
		routingKey := flag.Args()[0]

		v, err := elscli.GetServiceInstanceByKey(client, routingKey)

		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
		fmt.Fprintf(os.Stdout, "%d  %d\n", routingKey, v)

	default:
		fmt.Fprintf(os.Stderr, "error: invalid method %q\n", method)
		os.Exit(1)
	}

}
