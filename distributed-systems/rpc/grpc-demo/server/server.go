package main

import (
	pb "api/src/go"
	"context"
	"database/sql"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"time"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "mysecretpassword"
	dbname   = "go_grpc_demo"
)

var db *sql.DB

func initDB() error {
	var err error

	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", user, password, host, port, dbname)

	db, err = sql.Open("postgres", connectionString)

	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

type taskServiceServer struct {
	pb.UnimplementedTaskServiceServer
}

func createTaskFromRequest(request *pb.CreateTaskRequest) (*pb.Task, error) {
	var task = &pb.Task{
		Description: request.Description,
		UserId:      request.UserId,
		Deadline:    request.Deadline,
		CreatedAt:   timestamppb.Now(),
		Status:      pb.TaskStatus_TASK_STATUS_INCOMPLETE,
	}

	var taskId int

	insertStmt := `
        INSERT INTO tasks("description", "user_id", "status", "deadline", "created_at")
        VALUES($1, $2, $3, $4, $5) RETURNING id;
    `
	err := db.QueryRow(insertStmt, task.Description, task.UserId, pb.TaskStatus_name[int32(task.Status)], task.Deadline.AsTime(), task.CreatedAt.AsTime()).Scan(&taskId)
	if err != nil {
		return nil, err
	}

	task.Id = strconv.Itoa(taskId)
	return task, nil
}
func (s *taskServiceServer) CreateTask(ctx context.Context, request *pb.CreateTaskRequest) (*pb.Task, error) {
	return createTaskFromRequest(request)
}

func (s *taskServiceServer) GetTask(ctx context.Context, request *pb.GetTaskRequest) (*pb.Task, error) {
	var (
		id          int
		description string
		user_id     int
		status      string
		deadline    string
		created_at  string
	)
	err := db.QueryRow("SELECT * FROM tasks WHERE tasks.id = $1", request.TaskId).Scan(
		&id, &description, &user_id, &status, &deadline, &created_at)
	if err != nil {
		return nil, err
	}

	deadlineTime, err := time.Parse(time.RFC3339, deadline)
	if err != nil {
		log.Fatalf("Error: Invalid time for deadline: %v", err)
	}
	createdAtTime, err := time.Parse(time.RFC3339, created_at)
	if err != nil {
		log.Fatalf("Error: Invalid time for created_at: %v", err)
	}
	task := &pb.Task{
		Id:          strconv.Itoa(id),
		Description: description,
		UserId:      strconv.Itoa(user_id),
		Status:      pb.TaskStatus(pb.TaskStatus_value[status]),
		Deadline:    timestamppb.New(deadlineTime),
		CreatedAt:   timestamppb.New(createdAtTime),
	}
	return task, nil
}

func (s *taskServiceServer) ListTasks(request *pb.ListTasksRequest, stream pb.TaskService_ListTasksServer) error {
	query := `
        SELECT * FROM tasks WHERE user_id=$1 AND deadline < $2;
    `
	rows, err := db.Query(query, request.UserId, request.Deadline.AsTime())
	if err != nil {
		return err
	}
	for rows.Next() {
		var (
			id          int
			description string
			user_id     int
			status      string
			deadline    string
			created_at  string
		)
		err = rows.Scan(&id, &description, &user_id, &status, &deadline, &created_at)
		if err != nil {
			return err
		}

		deadlineTime, err := time.Parse(time.RFC3339, deadline)
		if err != nil {
			log.Fatalf("Error: Invalid time for deadline: %v", err)
		}
		createdAtTime, err := time.Parse(time.RFC3339, created_at)
		if err != nil {
			log.Fatalf("Error: Invalid time for created_at: %v", err)
		}
		task := &pb.Task{
			Id:          strconv.Itoa(id),
			Description: description,
			UserId:      strconv.Itoa(user_id),
			Status:      pb.TaskStatus(pb.TaskStatus_value[status]),
			Deadline:    timestamppb.New(deadlineTime),
			CreatedAt:   timestamppb.New(createdAtTime),
		}
		err = stream.Send(task)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *taskServiceServer) RecordTasks(stream pb.TaskService_RecordTasksServer) error {
	var createTaskRequest *pb.CreateTaskRequest
	var err error
	count := 0
	for {
		createTaskRequest, err = stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.TaskSummary{
				NoOfTasksCreated: strconv.Itoa(count),
			})
		}
		if err != nil {
			return err
		}
		_, err := createTaskFromRequest(createTaskRequest)
		if err != nil {
			return err
		}
		count++
	}
}

func (s *taskServiceServer) TaskChat(stream pb.TaskService_TaskChatServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		taskId := in.TaskId
		userId := in.UserId
		comment := in.Comment

		insertStmt := `INSERT INTO comments("task_id", "user_id", "comment") VALUES($1, $2, $3)`
		_, err = db.Exec(insertStmt, taskId, userId, comment)
		if err != nil {
			return err
		}

		taskComment := &pb.TaskComment{
			TaskId:    taskId,
			UserId:    userId,
			Comment:   comment,
			CreatedAt: timestamppb.Now(),
		}

		if err := stream.Send(taskComment); err != nil {
			return err
		}
	}
}

func main() {
	err := initDB()
	if err != nil {
		log.Fatalf("Error initiating database: %v", err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 9090))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterTaskServiceServer(grpcServer, &taskServiceServer{})

	// Register reflection service on gRPC server.
	reflection.Register(grpcServer)

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("Error starting gRPC server: %v", err)
	}
}
