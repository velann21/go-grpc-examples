package main

import (
	"context"
	"fmt"
	person "github.com/go-grpc-crs/grpccode/protopb"
	"google.golang.org/grpc"
	"io"
	"time"
)

func main(){
	conn, err := grpc.Dial("localhost:50051",grpc.WithInsecure())
	if err != nil{

	}
	defer conn.Close()
	c := person.NewPersonServiceClient(conn)
	ClientServerStreamingClient(c)
}

func ClientServerStreamingClient(client person.PersonServiceClient){
	persons := []person.PersonClientserverStreamingRequest{
		person.PersonClientserverStreamingRequest{Person:&person.Person{Id:1, Name:"velan"}},
		person.PersonClientserverStreamingRequest{Person:&person.Person{Id:2, Name:"uday"}},
		person.PersonClientserverStreamingRequest{Person:&person.Person{Id:3, Name:"debarhsi"}},
		person.PersonClientserverStreamingRequest{Person:&person.Person{Id:4, Name:"raghav"}},
		person.PersonClientserverStreamingRequest{Person:&person.Person{Id:5, Name:"bogdan"}},
		person.PersonClientserverStreamingRequest{Person:&person.Person{Id:6, Name:"abhi"}},
		person.PersonClientserverStreamingRequest{Person:&person.Person{Id:7, Name:"priyo"}},
		person.PersonClientserverStreamingRequest{Person:&person.Person{Id:8, Name:"Darlene"}},
		person.PersonClientserverStreamingRequest{Person:&person.Person{Id:9, Name:"indu"}},
		person.PersonClientserverStreamingRequest{Person:&person.Person{Id:10, Name:"leo"}},
	}

	stream, err := client.ClientServerStream(context.Background())
	if err != nil{

	}

	waitC := make(chan struct{})
	go func(){
		for _,v := range persons{
           err := stream.Send(&v)
           time.Sleep(1000*time.Millisecond)
           if err != nil{

		   }

		}
		stream.CloseSend()
	}()

	go func(){
		for {
			resp, err := stream.Recv()
			if err == io.EOF{
               close(waitC)
				return
			}
			if err != nil{
				close(waitC)
				return
			}

			fmt.Println(resp.GetResult())

		}
		close(waitC)

	}()

	<- waitC
}
