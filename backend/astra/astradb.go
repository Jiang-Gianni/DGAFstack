package astra

import (
	"crypto/tls"
	"encoding/hex"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/stargate/stargate-grpc-go-client/stargate/pkg/auth"
	"github.com/stargate/stargate-grpc-go-client/stargate/pkg/client"
	pb "github.com/stargate/stargate-grpc-go-client/stargate/pkg/proto"

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

func (astraDb *AstraDB) CreatePlayer(name string, number int) (string, error) {
	start := time.Now()
	defer func(start time.Time) {
		fmt.Println("time since start: ", time.Since(start).Milliseconds())
	}(start)

	uuid := uuid.New()

	insertQuery := &pb.Query{
		Cql: fmt.Sprintf("insert into test.my_table(uuid, name, number) values (%s, '%s', %d);", uuid.String(), name, number),
	}

	_, err := astraDb.StargateClient.ExecuteQuery(insertQuery)
	if err != nil {
		fmt.Printf("error executing query %v", err)
		return "", err
	}

	fmt.Println(uuid, name, number)

	return uuid.String(), nil
}

func (astraDb *AstraDB) GetPlayersByUuids(uuids []string) ([]map[string]string, error) {
	start := time.Now()
	defer func(start time.Time) {
		fmt.Println("time since start: ", time.Since(start).Milliseconds())
	}(start)

	joinedList := strings.Join(uuids, ", ")
	selectQuery := &pb.Query{
		Cql: fmt.Sprintf("select * from test.my_table where uuid in (%s);", joinedList),
		// Cql: "select * from test.my_table;",
	}

	response, err := astraDb.StargateClient.ExecuteQuery(selectQuery)
	if err != nil {
		fmt.Printf("error executing query %v", err)
		return []map[string]string{}, err
	}

	columnMap := map[int]string{
		0: "uuid",
		1: "name",
		2: "number",
	}

	re := regexp.MustCompile(`[^:]*$`)
	rows := response.GetResultSet().Rows
	results := make([]map[string]string, len(rows))

	for _, row := range rows {
		result := make(map[string]string)
		for index, value := range row.Values {

			if index == 0 {
				result[columnMap[index]] = GetUuidFromValue(value)
			} else {
				v := string(re.Find([]byte(value.String())))
				result[columnMap[index]] = v
			}

			results[index] = result
		}
	}

	return results, nil

}

func GetUuidFromValue(value *pb.Value) string {
	uuidString := hex.EncodeToString((value.GetUuid().GetValue()))
	uuid, err := uuid.Parse(uuidString)
	if err != nil {
		fmt.Println(err)
	}
	return uuid.String()
}
