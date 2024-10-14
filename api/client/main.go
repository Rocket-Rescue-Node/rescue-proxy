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

	"github.com/Rocket-Rescue-Node/rescue-proxy/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	addr := flag.String("addr", "0.0.0.0:8080", "the address where the api is responding to grpc requests")
	odao := flag.Bool("odao", false, "pass this to get the list of odao nodes")
	solo := flag.Bool("solo", false, "pass this to get the list of solo validator withdrawal addresses")
	validateEIP1271 := flag.Bool("validate-eip1271", false, "pass this to validate an EIP-1271 signature")
	dataHash := flag.String("data-hash", "", "data hash for EIP-1271 validation (32 bytes in hex)")
	signature := flag.String("signature", "", "signature for EIP-1271 validation (hex)")
	signerAddress := flag.String("signer-address", "", "signer address for EIP-1271 validation (20 bytes in hex)")
	useTLS := flag.Bool("tls", false, "use TLS to connect to the api")

	flag.Parse()

	var tc credentials.TransportCredentials
	if *useTLS {
		// An empty tls.Config{} will use the system's root CA set.
		tc = credentials.NewTLS(&tls.Config{})
	} else {
		tc = insecure.NewCredentials()
	}

	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(tc))
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
		return
	}
	defer conn.Close()

	c := pb.NewApiClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if *validateEIP1271 {
		if *dataHash == "" || *signature == "" || *signerAddress == "" {
			fmt.Fprintf(os.Stderr, "For EIP-1271 validation, data-hash, signature, and signer-address are required\n")
			os.Exit(1)
		}

		dataHashBytes, err := hex.DecodeString(*dataHash)
		if err != nil || len(dataHashBytes) != 32 {
			fmt.Fprintf(os.Stderr, "Invalid data hash: must be 32 bytes in hex\n")
			os.Exit(1)
		}

		signatureBytes, err := hex.DecodeString(*signature)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Invalid signature: must be in hex\n")
			os.Exit(1)
		}

		signerAddressBytes, err := hex.DecodeString(*signerAddress)
		if err != nil || len(signerAddressBytes) != 20 {
			fmt.Fprintf(os.Stderr, "Invalid signer address: must be 20 bytes in hex\n")
			os.Exit(1)
		}

		r, err := c.ValidateEIP1271(ctx, &pb.ValidateEIP1271Request{
			DataHash:  dataHashBytes,
			Signature: signatureBytes,
			Address:   signerAddressBytes,
		})
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
		if r.Error != "" {
			fmt.Printf("Validation error: %s\n", r.Error)
		}
		fmt.Printf("EIP-1271 Validation Result: %v\n", r.Valid)
		return
	}

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
