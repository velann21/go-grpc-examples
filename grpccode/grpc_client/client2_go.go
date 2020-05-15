package main

import (
	"context"
	"fmt"
	person "github.com/go-grpc-crs/grpccode/protopb"
	"google.golang.org/grpc"
	"io"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:50051",grpc.WithInsecure())
	if err != nil{

	}
	defer conn.Close()
	c := person.NewPersonServiceClient(conn)
	getPersons2(c)
}

func getPersons2(client person.PersonServiceClient){
	fmt.Println("Starting getPersons call ")
	request := person.PersonStreamingRequest{
		Person: &person.Person{
			Id:11,
			Name:"velan",
		},
	}

	resp, err := client.GetPerson(context.Background(), &request)
	if err != nil{
		log.Fatal("Error while calling the get person")
	}
	for{
		res, err := resp.Recv()
		if err == io.EOF{
			break
		}
		if err != nil{
			log.Fatal("Error while calling the get person")
		}
		log.Printf(res.Result)
	}
}

