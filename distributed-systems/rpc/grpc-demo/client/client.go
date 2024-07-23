package main

import (
	pb "api/src/go"
	"context"
	"io"
	"log"
	"math/rand"
	"strconv"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var createdTasks []string

func createTask(client pb.TaskServiceClient, createTaskRequest *pb.CreateTaskRequest) {
	log.Printf("Creating Task: { description: '%s', user_id: '%s', deadline: '%s' }", createTaskRequest.Description, createTaskRequest.UserId, createTaskRequest.Deadline)

	createdTask, err := client.CreateTask(context.Background(), createTaskRequest)
	if err != nil {
		panic(err)
	}

	if err == nil {
		log.Println(createdTask)
		createdTasks = append(createdTasks, createdTask.Id)
	}
}

func getTask(client pb.TaskServiceClient, geTaskRequest *pb.GetTaskRequest) {

	task, err := client.GetTask(context.Background(), geTaskRequest)

	if err == nil {
		log.Println(task)
	}
}

func runRecordTasks(client pb.TaskServiceClient) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	tasksCount := int(r.Int31n(10)) + 2
	var createTaskRequests []*pb.CreateTaskRequest
	for i := 0; i < tasksCount; i++ {
		deadlineDaysAfter := int(r.Int31n(100))
		createTaskRequests = append(createTaskRequests, &pb.CreateTaskRequest{
			Description: "Task " + strconv.Itoa(i),
			UserId:      "1",
			Deadline:    timestamppb.New(time.Now().AddDate(0, 0, deadlineDaysAfter)),
		})
	}

	log.Printf("Recording %d tasks", tasksCount)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stream, err := client.RecordTasks(ctx)
	if err != nil {
		log.Fatalf("error %v", err)
	}

	for _, createTaskRequest := range createTaskRequests {
		if err := stream.Send(createTaskRequest); err != nil {
			log.Fatalf("stream.Send(%v) failed: %v", createTaskRequest, err)
		}
	}
	reply, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error %v", err)
	}
	log.Printf("Tasks summary: %s tasks created", reply.NoOfTasksCreated)
}

func listTasks(client pb.TaskServiceClient, userId string, deadline string) {
	log.Printf("Listing all tasks of User %s with deadline within %s", userId, deadline)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	deadlineTime, err := time.Parse(time.RFC3339, deadline)
	if err != nil {
		log.Fatalf("error parsing time: %v", err)
	}
	stream, err := client.ListTasks(ctx, &pb.ListTasksRequest{
		UserId:   userId,
		Deadline: timestamppb.New(deadlineTime),
	})

	for {
		task, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error %v", err)
		}
		log.Printf(
			"Task: { Id: %s, Description: '%s', Status: '%s', Deadline: '%s'  }",
			task.Id, task.Description, task.Status, task.Deadline.AsTime().Format(time.RFC3339))
	}
}

func runTaskChat(client pb.TaskServiceClient) {
	comments := []*pb.TaskComment{
		{UserId: "1", TaskId: createdTasks[0], Comment: "Comment 1", CreatedAt: timestamppb.Now()},
		{UserId: "2", TaskId: createdTasks[0], Comment: "Comment 2", CreatedAt: timestamppb.Now()},
		{UserId: "1", TaskId: createdTasks[0], Comment: "Comment 3", CreatedAt: timestamppb.Now()},
		{UserId: "1", TaskId: createdTasks[0], Comment: "Comment 4", CreatedAt: timestamppb.Now()},
		{UserId: "3", TaskId: createdTasks[0], Comment: "Comment 5", CreatedAt: timestamppb.Now()},
		{UserId: "2", TaskId: createdTasks[0], Comment: "Comment 6", CreatedAt: timestamppb.Now()},
		{UserId: "3", TaskId: createdTasks[0], Comment: "Comment 7", CreatedAt: timestamppb.Now()},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stream, err := client.TaskChat(ctx)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	wc := make(chan struct{})
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				close(wc)
				return
			}
			if err != nil {
				log.Fatalf("error: %v", err)
			}
			log.Printf("Got comment at task %s by user %s", in.TaskId, in.UserId)
		}
	}()

	for _, comment := range comments {
		if err := stream.Send(comment); err != nil {
			log.Fatalf("Send failed %v", err)
		}
	}
	stream.CloseSend()
	<-wc
}

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial("localhost:9090", opts...)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	defer conn.Close()

	client := pb.NewTaskServiceClient(conn)

	log.Printf("==== Calling CreateTask ====")
	createTask(client, &pb.CreateTaskRequest{
		Description: "Task 1",
		UserId:      "1",
		Deadline:    timestamppb.New(time.Now().AddDate(0, 0, 200)),
	})

	createTask(client, &pb.CreateTaskRequest{
		Description: "Task 2",
		UserId:      "1",
		Deadline:    timestamppb.New(time.Now().AddDate(0, 0, 200)),
	})

	log.Printf("==== Calling GetTask ====")
	getTask(client, &pb.GetTaskRequest{TaskId: createdTasks[0]})

	log.Printf("==== Calling RecordTasks ====")
	runRecordTasks(client)

	log.Printf("==== Calling ListTasks ====")
	listTasks(client, "1", time.Now().AddDate(0, 0, 100).Format(time.RFC3339))

	log.Printf("==== Calling TaskChat ====")
	runTaskChat(client)
}
