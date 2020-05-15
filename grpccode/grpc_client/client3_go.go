package main

import (
	"context"
	"fmt"
	person "github.com/go-grpc-crs/grpccode/protopb"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main()  {
	conn, err := grpc.Dial("localhost:50051",grpc.WithInsecure())
	if err != nil{

	}
	defer conn.Close()
	c := person.NewPersonServiceClient(conn)
	createPerson(c)
}

func createPerson(client person.PersonServiceClient){
	persons := []person.Person{
		person.Person{
			Id:1,
			Name:"velan",
		},
		person.Person{
			Id:2,
			Name:"uday",
		},
		person.Person{
			Id:3,
			Name:"debarhsi",
		},
		person.Person{
			Id:4,
			Name:"raghav",
		},
		person.Person{
			Id:5,
			Name:"bogdan",
		},
		person.Person{
			Id:6,
			Name:"abhi",
		},
		person.Person{
			Id:7,
			Name:"priyo",
		},
		person.Person{
			Id:8,
			Name:"Darlene",
		},
		person.Person{
			Id:9,
			Name:"indu",
		},
		person.Person{
			Id:10,
			Name:"leo",
		},
	}
	stream, err := client.CreatePerson(context.Background())
	if err != nil{
		log.Fatalln("Error Occured")
	}
	for _, v :=range persons{
		personClientStream := person.PersonClientStreamingRequest{
			Person:&v,
		}
		_= stream.Send(&personClientStream)
		time.Sleep(1000*time.Millisecond)
	}
	resp, err := stream.CloseAndRecv()
	if err != nil{
		log.Fatal("Error Occured while close and recieve the message from server")
	}
	fmt.Printf("%v", resp)





}
