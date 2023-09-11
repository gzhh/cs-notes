package main

import (
	"context"
	"fmt"
	"os"
)

const (
	TASK_NAME_EXPORT_MYSQL_TO_ES    = "export_mysql_to_es"
	TASK_NAME_EXPORT_MYSQL_TO_REDIS = "export_mysql_to_redis"
)

type Callback func(ctx context.Context)

var TaskCallback map[string]Callback

type task struct {
}

func NewTask() *task {
	return &task{}
}

func (srv *task) ExportMysqlToEs(ctx context.Context) {
	fmt.Println("export_mysql_to_es")
}

func (srv *task) ExportMysqlToRedis(ctx context.Context) {
	fmt.Println("export_mysql_to_redis")
}

func main() {
	srv := NewTask()

	TaskCallback = map[string]Callback{
		TASK_NAME_EXPORT_MYSQL_TO_ES:    srv.ExportMysqlToEs,
		TASK_NAME_EXPORT_MYSQL_TO_REDIS: srv.ExportMysqlToRedis,
	}

	tasks := []string{
		TASK_NAME_EXPORT_MYSQL_TO_ES,
		TASK_NAME_EXPORT_MYSQL_TO_REDIS,
		"not-exist",
	}

	for _, taskName := range tasks {
		callback, ok := TaskCallback[taskName]
		if !ok {
			fmt.Printf("task [%s] not found\n", taskName)
			os.Exit(1)
		}

		callback(context.Background())
	}
}
