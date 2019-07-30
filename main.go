package main

import (
	"context"

	pb "foo/proto"

	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	c := pb.NewStorageBenchWrapperClient(conn)
	if _, err := c.Write(ctx, &pb.ObjectWrite{
		BucketName:  "some-bucket",
		ObjectName:  "some-object",
		Destination: "some-destination",
	}); err != nil {
		panic(err)
	}
	if _, err := c.Read(ctx, &pb.ObjectRead{
		BucketName: "some-bucket",
		ObjectName: "some-object",
	}); err != nil {
		panic(err)
	}
}
