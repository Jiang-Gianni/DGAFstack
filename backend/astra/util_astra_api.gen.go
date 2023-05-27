// GENERATED CODE - DO NOT MODIFY BY HAND

package astra

import (
	"fmt"

	
	"github.com/Jiang-Gianni/DGAFstack/rest/mytable"
	"github.com/Jiang-Gianni/DGAFstack/rest/user"
	"github.com/google/uuid"
	"github.com/stargate/stargate-grpc-go-client/stargate/pkg/client"
	pb "github.com/stargate/stargate-grpc-go-client/stargate/pkg/proto"
)



//MyTable REST API
var (
	MyTableColumnToField = []func(*pb.Value, *mytable.MyTable) *mytable.MyTable{
		
		func(val *pb.Value, p *mytable.MyTable) *mytable.MyTable {
			Uuid, _ := client.ToUUID(val)
			p.Uuid = Uuid.String()
			return p
		},
		
		func(val *pb.Value, p *mytable.MyTable) *mytable.MyTable {
			Name, _ := client.ToString(val)
			p.Name = Name
			return p
		},
		
		func(val *pb.Value, p *mytable.MyTable) *mytable.MyTable {
			Number, _ := client.ToInt(val)
			p.Number = int(Number)
			return p
		},
		
	}
	MyTableQueryResponse = func(response *pb.Response) ([]mytable.MyTable, error) {
		rows := response.GetResultSet().Rows
		result := make([]mytable.MyTable, len(rows))
		for rowIndex, row := range rows {
			rowResult := &mytable.MyTable{}
			for valueIndex, value := range row.Values {
				rowResult = MyTableColumnToField[valueIndex](value, rowResult)
			}
			result[rowIndex] = *rowResult
		}
		return result, nil
	}
)

type MyTable struct {
	Keyspace       string
	StargateClient *client.StargateClient
}

func (astraDb *AstraDB) MyTable() *MyTable {
	return &MyTable{
		Keyspace: astraDb.Keyspace,
		StargateClient: astraDb.StargateClient,
	}
}

func (t *MyTable) Get() ([]mytable.MyTable, error) {
	pbQuery := &pb.Query{Cql: fmt.Sprintf("select uuid, name, number from %s.my_table;", t.Keyspace)}
	response, err := t.StargateClient.ExecuteQuery(pbQuery)
	if err != nil {
		fmt.Printf("error executing query %v", err)
		return nil, err
	}
	return MyTableQueryResponse(response)
}

func (t *MyTable) GetById(id string) ([]mytable.MyTable, error) {
	pbQuery := &pb.Query{
		Cql: fmt.Sprintf("select uuid, name, number from %s.my_table where uuid = %s;", t.Keyspace, id),
	}
	response, err := t.StargateClient.ExecuteQuery(pbQuery)
	if err != nil {
		fmt.Printf("error executing query %v", err)
		return nil, err
	}
	return MyTableQueryResponse(response)
}

func (t *MyTable) Create(toBeCreatedMyTables []mytable.MyTable) ([]string, error) {
	newUuids := []string{}
	var queries []*pb.BatchQuery
	for _, row := range toBeCreatedMyTables {
		rowUuid := uuid.New().String()
		row.Uuid = rowUuid
		newUuids = append(newUuids, rowUuid)
		rowBatchQuery := pb.BatchQuery{
			Cql: fmt.Sprintf("insert into %s.my_table(uuid, name, number) values (%s ,'%s', %d);", t.Keyspace, row.Uuid, row.Name, row.Number),
		}
		queries = append(queries, &rowBatchQuery)
	}
	pbBatch := &pb.Batch{
		Type:    pb.Batch_LOGGED,
		Queries: queries,
	}
	_, err := t.StargateClient.ExecuteBatch(pbBatch)
	if err != nil {
		fmt.Printf("error executing query %v", err)
		return nil, err
	}
	return newUuids, nil
}

func (t *MyTable) Update(toBeUpdatedMyTables []mytable.MyTable) error {
	var queries []*pb.BatchQuery
	for _, row := range toBeUpdatedMyTables {
		rowBatchQuery := pb.BatchQuery{
			Cql: fmt.Sprintf("update %s.my_table set name = '%s', number = %d where uuid = %s;", t.Keyspace, row.Name, row.Number, row.Uuid),
		}
		_ = row
		queries = append(queries, &rowBatchQuery)
	}
	pbBatch := &pb.Batch{
		Type:    pb.Batch_LOGGED,
		Queries: queries,
	}
	_, err := t.StargateClient.ExecuteBatch(pbBatch)
	if err != nil {
		fmt.Printf("error executing query %v", err)
		return err
	}
	return nil
}

func (t *MyTable) Delete(id string) (error) {
	pbQuery := &pb.Query{
		Cql: fmt.Sprintf("delete from %s.my_table where uuid = %s;", t.Keyspace, id),
	}
	_, err := t.StargateClient.ExecuteQuery(pbQuery)
	if err != nil {
		fmt.Printf("error executing query %v", err)
		return err
	}
	return nil
}



//User REST API
var (
	UserColumnToField = []func(*pb.Value, *user.User) *user.User{
		
		func(val *pb.Value, p *user.User) *user.User {
			Id, _ := client.ToUUID(val)
			p.Id = Id.String()
			return p
		},
		
		func(val *pb.Value, p *user.User) *user.User {
			Name, _ := client.ToString(val)
			p.Name = Name
			return p
		},
		
	}
	UserQueryResponse = func(response *pb.Response) ([]user.User, error) {
		rows := response.GetResultSet().Rows
		result := make([]user.User, len(rows))
		for rowIndex, row := range rows {
			rowResult := &user.User{}
			for valueIndex, value := range row.Values {
				rowResult = UserColumnToField[valueIndex](value, rowResult)
			}
			result[rowIndex] = *rowResult
		}
		return result, nil
	}
)

type User struct {
	Keyspace       string
	StargateClient *client.StargateClient
}

func (astraDb *AstraDB) User() *User {
	return &User{
		Keyspace: astraDb.Keyspace,
		StargateClient: astraDb.StargateClient,
	}
}

func (t *User) Get() ([]user.User, error) {
	pbQuery := &pb.Query{Cql: fmt.Sprintf("select id, name from %s.user;", t.Keyspace)}
	response, err := t.StargateClient.ExecuteQuery(pbQuery)
	if err != nil {
		fmt.Printf("error executing query %v", err)
		return nil, err
	}
	return UserQueryResponse(response)
}

func (t *User) GetById(id string) ([]user.User, error) {
	pbQuery := &pb.Query{
		Cql: fmt.Sprintf("select id, name from %s.user where id = %s;", t.Keyspace, id),
	}
	response, err := t.StargateClient.ExecuteQuery(pbQuery)
	if err != nil {
		fmt.Printf("error executing query %v", err)
		return nil, err
	}
	return UserQueryResponse(response)
}

func (t *User) Create(toBeCreatedUsers []user.User) ([]string, error) {
	newUuids := []string{}
	var queries []*pb.BatchQuery
	for _, row := range toBeCreatedUsers {
		rowUuid := uuid.New().String()
		row.Id = rowUuid
		newUuids = append(newUuids, rowUuid)
		rowBatchQuery := pb.BatchQuery{
			Cql: fmt.Sprintf("insert into %s.user(id, name) values (%s ,'%s');", t.Keyspace, row.Id, row.Name),
		}
		queries = append(queries, &rowBatchQuery)
	}
	pbBatch := &pb.Batch{
		Type:    pb.Batch_LOGGED,
		Queries: queries,
	}
	_, err := t.StargateClient.ExecuteBatch(pbBatch)
	if err != nil {
		fmt.Printf("error executing query %v", err)
		return nil, err
	}
	return newUuids, nil
}

func (t *User) Update(toBeUpdatedUsers []user.User) error {
	var queries []*pb.BatchQuery
	for _, row := range toBeUpdatedUsers {
		rowBatchQuery := pb.BatchQuery{
			Cql: fmt.Sprintf("update %s.user set name = '%s' where id = %s;", t.Keyspace, row.Name, row.Id),
		}
		_ = row
		queries = append(queries, &rowBatchQuery)
	}
	pbBatch := &pb.Batch{
		Type:    pb.Batch_LOGGED,
		Queries: queries,
	}
	_, err := t.StargateClient.ExecuteBatch(pbBatch)
	if err != nil {
		fmt.Printf("error executing query %v", err)
		return err
	}
	return nil
}

func (t *User) Delete(id string) (error) {
	pbQuery := &pb.Query{
		Cql: fmt.Sprintf("delete from %s.user where id = %s;", t.Keyspace, id),
	}
	_, err := t.StargateClient.ExecuteQuery(pbQuery)
	if err != nil {
		fmt.Printf("error executing query %v", err)
		return err
	}
	return nil
}

