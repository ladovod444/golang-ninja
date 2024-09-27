package main

import (
    "context"
    "fmt"
    "log"
    "net"
    "os"
    "os/signal"
    "sync"

    "google.golang.org/grpc"
    pb "go-grpc-todo/proto3/todo"
)

type taskServer struct {
    pb.UnimplementedTaskManagerServer
    mutex sync.Mutex
    tasks []*pb.Task
}

func (s *taskServer) AddTask(ctx context.Context, task *pb.Task) (*pb.TaskResponse, error) {
    s.mutex.Lock()
    defer s.mutex.Unlock()
    s.tasks = append(s.tasks, task)
    return &pb.TaskResponse{Id: task.Id, Result: "Added"}, nil
}

func (s *taskServer) ListTasks(ctx context.Context, req *pb.ListRequest) (*pb.ListResponse, error) {
    s.mutex.Lock()
    defer s.mutex.Unlock()
    return &pb.ListResponse{Tasks: s.tasks}, nil
}

func main() {
    lis, err := net.Listen("tcp", "0.0.0.0:50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    grpcServer := grpc.NewServer()
    pb.RegisterTaskManagerServer(grpcServer, &taskServer{})

    go func() {
        if err := grpcServer.Serve(lis); err != nil {
            log.Fatalf("failed to serve: %v", err)
        }
    }()

    // Wait for Ctrl+C to exit
    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt)
    <-c

    grpcServer.GracefulStop()
    fmt.Println("Shut down server")
}