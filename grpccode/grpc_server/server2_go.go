package main

import (
	person "github.com/go-grpc-crs/grpccode/protopb"
	"google.golang.org/grpc"
	"log"
	"net"
)
import "context"

type Server1 struct {

}

func (ser Server1) CreateUser(ctx context.Context, req *person.UserReq) (*person.UserRes, error) {

	return nil, nil
}

func main(){
	server, err := net.Listen("tcp","0.0.0.0:50051")
	if err != nil{

	}
	s := grpc.NewServer()
	person.RegisterUserServiceServer(s, Server1{})
	err = s.Serve(server)
	if err != nil{
		log.Fatal("")
	}
}


