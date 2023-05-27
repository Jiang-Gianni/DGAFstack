package astra

import (
	"crypto/tls"
	"fmt"
	"os"

	"github.com/stargate/stargate-grpc-go-client/stargate/pkg/auth"
	"github.com/stargate/stargate-grpc-go-client/stargate/pkg/client"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type AstraDB struct {
	Keyspace       string
	StargateClient *client.StargateClient
}

func New() *AstraDB {

	var astraUri = os.Getenv("ASTRA_URI")
	var bearerToken = os.Getenv("ASTRA_TOKEN")
	var keyspace = os.Getenv("ASTRA_KEYSPACE")

	config := &tls.Config{
		InsecureSkipVerify: false,
	}

	conn, err := grpc.Dial(astraUri, grpc.WithTransportCredentials(credentials.NewTLS(config)),
		grpc.WithBlock(),
		grpc.WithPerRPCCredentials(
			auth.NewStaticTokenProvider(bearerToken),
		),
	)

	if err != nil {
		fmt.Printf("error %s ", err.Error())
	}

	stargateClient, err := client.NewStargateClientWithConn(conn)
	if err != nil {
		fmt.Printf("error creating client %v", err)
	}

	fmt.Println("Connected to Astra DB")

	return &AstraDB{
		Keyspace:       keyspace,
		StargateClient: stargateClient,
	}
}
