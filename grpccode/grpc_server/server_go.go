package main

import (
	"context"
	"fmt"
	"github.com/go-grpc-crs/grpccode/protopb"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"strconv"
	"time"
)

type Server struct {

}

func (server Server) PersonCreation(ctx context.Context, request *person.PersonRequest) (*person.PersonResponse, error){
	fmt.Println("I recieved the grpc request")
	name := request.GetPerson().GetName()
	respMessage := "Hello "+name
	res := &person.PersonResponse{
		Result:respMessage,
	}
	return res, nil

}


func (server Server) GetPerson(request *person.PersonStreamingRequest, stream person.PersonService_GetPersonServer) (error){
	fmt.Println("I recieved the grpc request for streaming")
	name := request.GetPerson().GetName()

	for i:=0; i<=20; i++ {
		respMessage := "Hello "+name+" number is: "+strconv.Itoa(i)
		res := &person.PersonStreamingResponse {
			Result:respMessage,
		}
		stream.Send(res)
		time.Sleep(1000*time.Millisecond)
	}
	return nil
}

func (server Server) CreatePerson(request person.PersonService_CreatePersonServer)error{
	fmt.Printf("Inside the CreatePerson()")
	result := ""
	for{
		stream, err := request.Recv()
		fmt.Println("I recieved something from the client", stream.GetPerson().GetName())
		if err == io.EOF{
			return request.SendAndClose(&person.PersonClientStreamingResponse{
				Result: result,
			})
		}
		if err != nil{
           return err
		}
		result +=  "Hello "+stream.GetPerson().GetName() + ", "
	}
}

func (server Server) ClientServerStream(stream person.PersonService_ClientServerStreamServer) error {
	fmt.Println("Inside the ClientServerStream")

	for {
		msg, err := stream.Recv()

		if err == io.EOF{
			return nil
		}
		if err != nil {
			return err
		}

		name := msg.GetPerson().Name
		result := "Hello "+name+" ! "
		err = stream.Send(&person.PersonClientServerStreamingResponse{
			Result:result,
		})

		if err != nil{
			return err
		}
	}

}


func main(){
	server, err := net.Listen("tcp","0.0.0.0:50051")
	if err != nil{

	}
	s := grpc.NewServer()
	person.RegisterPersonServiceServer(s, Server{})
	err = s.Serve(server)
	if err != nil{
		log.Fatal("")
	}
}



