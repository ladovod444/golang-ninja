package main

import (
    "context"
    "fmt"
    "log"
    "time"

    "google.golang.org/grpc"
    pb "go-grpc-todo/proto3/todo"
)

func main() {
    conn, err := grpc.NewClient("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
    if err != nil {
        log.Fatalf("could not connect: %v", err)
    }
    defer conn.Close()
    c := pb.NewTaskManagerClient(conn)

    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()

    // Add a task
    //r, err := c.AddTask(ctx, &pb.Task{Id: "1", Title: "Learn gRPC", Completed: false})
    //r, err := c.AddTask(ctx, &pb.Task{Id: "2", Title: "Learn gRPC another", Completed: false})
    r, err := c.AddTask(ctx, &pb.Task{Id: "2", Title: "Learn gRPC one more", Completed: false})
    if err != nil {
        log.Fatalf("could not add task: %v", err)
    }
    fmt.Printf("AddTask: %s\n", r.Result)

    // List tasks
    l, err := c.ListTasks(ctx, &pb.ListRequest{})
    if err != nil {
        log.Fatalf("could not list tasks: %v", err)
    }
    for _, t := range l.Tasks {
        fmt.Printf("Task: %s - %s\n", t.Id, t.Title)
    }
}