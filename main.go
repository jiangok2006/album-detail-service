package main

import (
	"context"
	"fmt"
	"log"
	"net"

	gpb "buf.build/gen/go/jiangok/buf-hello/grpc/go/album_detail_service/v1/album_detail_servicev1grpc"
	pb "buf.build/gen/go/jiangok/buf-hello/protocolbuffers/go/album_detail_service/v1"

	"google.golang.org/grpc"
)

type server struct {
	gpb.UnimplementedAlbumDetailServiceServer
}

func NewServer() *server {
	return &server{}
}

func (s *server) GetAlbumDetail(ctx context.Context, in *pb.GetAlbumDetailRequest) (*pb.GetAlbumDetailResponse, error) {
	if path, ok := grpc.Method(ctx); ok {
		fmt.Println("Received request for method:", path)
	}

	return &pb.GetAlbumDetailResponse{Id: "100", Title: "mytitle", Price: 10.00}, nil
}

// buf curl --schema buf.build/jiangok/buf-hello --protocol grpc --http2-prior-knowledge http://localhost:8081/album_detail_service.v1.AlbumDetailService/GetAlbumDetail
func main() {
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("failed to detailen: %v", err)
	}
	s := grpc.NewServer()
	gpb.RegisterAlbumDetailServiceServer(s, &server{})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
