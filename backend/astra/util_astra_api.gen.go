// GENERATED CODE - DO NOT MODIFY BY HAND

package astra

import (
	"fmt"

	
	"github.com/Jiang-Gianni/DGAFstack/rest/mystruct"
	"github.com/Jiang-Gianni/DGAFstack/rest/user"
	"github.com/google/uuid"
	"github.com/stargate/stargate-grpc-go-client/stargate/pkg/client"
	pb "github.com/stargate/stargate-grpc-go-client/stargate/pkg/proto"
)



//MyStruct REST API
var (
	MyStructColumnToField = []func(*pb.Value, *mystruct.MyStruct) *mystruct.MyStruct{
		
		func(val *pb.Value, p *mystruct.MyStruct) *mystruct.MyStruct {
			Uuid, _ := client.ToUUID(val)
			p.Uuid = Uuid.String()
			return p
		},
		
		func(val *pb.Value, p *mystruct.MyStruct) *mystruct.MyStruct {
			Name, _ := client.ToString(val)
			p.Name = Name
			return p
		},
		
		func(val *pb.Value, p *mystruct.MyStruct) *mystruct.MyStruct {
			Number, _ := client.ToInt(val)
			p.Number = int(Number)
			return p
		},
		
		func(val *pb.Value, p *mystruct.MyStruct) *mystruct.MyStruct {
			MyBoolean, _ := client.ToBoolean(val)
			p.MyBoolean = MyBoolean
			return p
		},
		
	}
	MyStructQueryResponse = func(response *pb.Response) ([]mystruct.MyStruct, error) {
		rows := response.GetResultSet().Rows
		result := make([]mystruct.MyStruct, len(rows))
		for rowIndex, row := range rows {
			rowResult := &mystruct.MyStruct{}
			for valueIndex, value := range row.Values {
				rowResult = MyStructColumnToField[valueIndex](value, rowResult)
			}
			result[rowIndex] = *rowResult
		}
		return result, nil
	}
)

func (astraDb *AstraDB) GetMyStructs() ([]mystruct.MyStruct, error) {
	pbQuery := &pb.Query{Cql: fmt.Sprintf("select * from %s.my_struct;", astraDb.Keyspace)}
	response, err := astraDb.StargateClient.ExecuteQuery(pbQuery)
	if err != nil {
		fmt.Printf("error executing query %v", err)
		return nil, err
	}
	return MyStructQueryResponse(response)
}

func (astraDb *AstraDB) GetMyStructById(id string) ([]mystruct.MyStruct, error) {
	pbQuery := &pb.Query{
		Cql: fmt.Sprintf("select * from %s.my_struct where uuid = %s;", astraDb.Keyspace, id),
	}
	response, err := astraDb.StargateClient.ExecuteQuery(pbQuery)
	if err != nil {
		fmt.Printf("error executing query %v", err)
		return nil, err
	}
	return MyStructQueryResponse(response)
}

func (astraDb *AstraDB) CreateMyStructs(toBeCreatedMyStructs []mystruct.MyStruct) ([]string, error) {
	newUuids := []string{}
	var queries []*pb.BatchQuery
	for _, row := range toBeCreatedMyStructs {
		rowUuid := uuid.New().String()
		row.Uuid = rowUuid
		newUuids = append(newUuids, rowUuid)
		rowBatchQuery := pb.BatchQuery{
			Cql: fmt.Sprintf("insert into %s.my_struct(uuid, name, number, my_boolean) values (%s, '%s', %d, %t);", astraDb.Keyspace, row.Uuid, row.Name, row.Number, row.MyBoolean),
		}
		queries = append(queries, &rowBatchQuery)
	}
	pbBatch := &pb.Batch{
		Type:    pb.Batch_LOGGED,
		Queries: queries,
	}
	_, err := astraDb.StargateClient.ExecuteBatch(pbBatch)
	if err != nil {
		fmt.Printf("error executing query %v", err)
		return nil, err
	}
	return newUuids, nil
}

func (astraDb *AstraDB) UpdateMyStructs(toBeUpdatedMyStructs []mystruct.MyStruct) error {
	var queries []*pb.BatchQuery
	for _, row := range toBeUpdatedMyStructs {
		rowBatchQuery := pb.BatchQuery{
			Cql: fmt.Sprintf("update %s.my_struct set uuid = %s name = '%s' number = %d my_boolean = %t ;", astraDb.Keyspace, row.Uuid, row.Name, row.Number, row.MyBoolean),
		}
		_ = row
		queries = append(queries, &rowBatchQuery)
	}
	pbBatch := &pb.Batch{
		Type:    pb.Batch_LOGGED,
		Queries: queries,
	}
	_, err := astraDb.StargateClient.ExecuteBatch(pbBatch)
	if err != nil {
		fmt.Printf("error executing query %v", err)
		return err
	}
	return nil
}

func (astraDb *AstraDB) DeleteMyStructById(id string) (error) {
	pbQuery := &pb.Query{
		Cql: fmt.Sprintf("delete * from %s.my_struct where uuid = %s;", astraDb.Keyspace, id),
	}
	_, err := astraDb.StargateClient.ExecuteQuery(pbQuery)
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

func (astraDb *AstraDB) GetUsers() ([]user.User, error) {
	pbQuery := &pb.Query{Cql: fmt.Sprintf("select * from %s.user;", astraDb.Keyspace)}
	response, err := astraDb.StargateClient.ExecuteQuery(pbQuery)
	if err != nil {
		fmt.Printf("error executing query %v", err)
		return nil, err
	}
	return UserQueryResponse(response)
}

func (astraDb *AstraDB) GetUserById(id string) ([]user.User, error) {
	pbQuery := &pb.Query{
		Cql: fmt.Sprintf("select * from %s.user where id = %s;", astraDb.Keyspace, id),
	}
	response, err := astraDb.StargateClient.ExecuteQuery(pbQuery)
	if err != nil {
		fmt.Printf("error executing query %v", err)
		return nil, err
	}
	return UserQueryResponse(response)
}

func (astraDb *AstraDB) CreateUsers(toBeCreatedUsers []user.User) ([]string, error) {
	newUuids := []string{}
	var queries []*pb.BatchQuery
	for _, row := range toBeCreatedUsers {
		rowUuid := uuid.New().String()
		row.Id = rowUuid
		newUuids = append(newUuids, rowUuid)
		rowBatchQuery := pb.BatchQuery{
			Cql: fmt.Sprintf("insert into %s.user(id, name) values (%s, '%s');", astraDb.Keyspace, row.Id, row.Name),
		}
		queries = append(queries, &rowBatchQuery)
	}
	pbBatch := &pb.Batch{
		Type:    pb.Batch_LOGGED,
		Queries: queries,
	}
	_, err := astraDb.StargateClient.ExecuteBatch(pbBatch)
	if err != nil {
		fmt.Printf("error executing query %v", err)
		return nil, err
	}
	return newUuids, nil
}

func (astraDb *AstraDB) UpdateUsers(toBeUpdatedUsers []user.User) error {
	var queries []*pb.BatchQuery
	for _, row := range toBeUpdatedUsers {
		rowBatchQuery := pb.BatchQuery{
			Cql: fmt.Sprintf("update %s.user set id = %s name = '%s' ;", astraDb.Keyspace, row.Id, row.Name),
		}
		_ = row
		queries = append(queries, &rowBatchQuery)
	}
	pbBatch := &pb.Batch{
		Type:    pb.Batch_LOGGED,
		Queries: queries,
	}
	_, err := astraDb.StargateClient.ExecuteBatch(pbBatch)
	if err != nil {
		fmt.Printf("error executing query %v", err)
		return err
	}
	return nil
}

func (astraDb *AstraDB) DeleteUserById(id string) (error) {
	pbQuery := &pb.Query{
		Cql: fmt.Sprintf("delete * from %s.user where id = %s;", astraDb.Keyspace, id),
	}
	_, err := astraDb.StargateClient.ExecuteQuery(pbQuery)
	if err != nil {
		fmt.Printf("error executing query %v", err)
		return err
	}
	return nil
}

