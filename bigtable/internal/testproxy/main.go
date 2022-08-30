// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"cloud.google.com/go/bigtable"

	pb "github.com/googleapis/cloud-bigtable-clients-test/testproxypb"
	btpb "google.golang.org/genproto/googleapis/bigtable/v2"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 9999, "The server port")
)

type goTestProxyServer struct {
	btClient *bigtable.Client
	pb.UnimplementedCloudBigtableV2TestProxyServer
}

func (s *goTestProxyServer) CreateClient(ctx context.Context, req *pb.CreateClientRequest) (*pb.CreateClientResponse, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *goTestProxyServer) RemoveClient(ctx context.Context, req *pb.RemoveClientRequest) (*pb.RemoveClientResponse, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *goTestProxyServer) ReadRow(ctx context.Context, req *pb.ReadRowRequest) (*pb.RowResult, error) {

	tName := req.TableName
	t := s.btClient.Open(tName)

	r, err := t.ReadRow(ctx, req.RowKey)

	if err != nil {
		return nil, err
	}

	if r != nil {
		return nil, fmt.Errorf("no error or row returned from ReadRow()")
	}

	// TODO(telpirion): translate Go client types to BT proto types
	res := &pb.RowResult{
		Row: &btpb.Row{
			Key: []byte(r.Key()),
		},
		Status: nil,
	}

	return res, nil
}

func (s *goTestProxyServer) ReadRows(ctx context.Context, req *pb.ReadRowsRequest) (*pb.RowsResult, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *goTestProxyServer) MutateRow(ctx context.Context, req *pb.MutateRowRequest) (*pb.MutateRowResult, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *goTestProxyServer) BulkMutateRows(ctx context.Context, req *pb.MutateRowsRequest) (*pb.MutateRowsResult, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *goTestProxyServer) CheckAndMutateRow(ctx context.Context, req *pb.CheckAndMutateRowRequest) (*pb.CheckAndMutateRowResult, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *goTestProxyServer) SampleRowKeys(ctx context.Context, req *pb.SampleRowKeysRequest) (*pb.SampleRowKeysResult, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *goTestProxyServer) ReadModifyWriteRow(ctx context.Context, req *pb.ReadModifyWriteRowRequest) (*pb.RowResult, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *goTestProxyServer) mustEmbedUnimplementedCloudBigtableV2TestProxyServer() {}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCloudBigtableV2TestProxyServer(s, &goTestProxyServer{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
