package main

import (
	"context"
	"crypto/tls"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/Rocket-Pool-Rescue-Node/rescue-proxy/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	addr := flag.String("addr", "0.0.0.0:8080", "the address where the api is responding to grpc requests")
	odao := flag.Bool("odao", false, "pass this to get the list of odao nodes")
	solo := flag.Bool("solo", false, "pass this to get the list of solo validator withdrawal addresses")
	useTLS := flag.Bool("tls", false, "use TLS to connect to the api")

	flag.Parse()

	var tc credentials.TransportCredentials
	if *useTLS {
		// An empty tls.Config{} will use the system's root CA set.
		tc = credentials.NewTLS(&tls.Config{})
	} else {
		tc = insecure.NewCredentials()
	}

	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(tc))
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
		return
	}
	defer conn.Close()

	c := pb.NewApiClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var nodeIds [][]byte

	if *odao {
		r, err := c.GetOdaoNodes(ctx, &pb.OdaoNodesRequest{})
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
			return
		}

		nodeIds = r.GetNodeIds()
	} else if *solo {
		r, err := c.GetSoloValidators(ctx, &pb.SoloValidatorsRequest{})
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
			return
		}

		nodeIds = r.GetWithdrawalAddresses()
	} else {
		r, err := c.GetRocketPoolNodes(ctx, &pb.RocketPoolNodesRequest{})
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
			return
		}

		nodeIds = r.GetNodeIds()
	}
	out := make([]string, 0, len(nodeIds))

	for _, addr := range nodeIds {
		out = append(out, "0x"+hex.EncodeToString(addr))
	}

	j, err := json.Marshal(out)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
		return
	}

	fmt.Printf("%s\n", j)
}
