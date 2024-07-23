# gRPC demo

Guide
- https://www.jetbrains.com/guide/go/tutorials/grpc_part_one/
- https://www.jetbrains.com/guide/go/tutorials/grpc_part_two/
- https://www.jetbrains.com/guide/go/tutorials/grpc_part_three/

Code
- https://github.com/heraldofsolace/go-grpc-demo/tree/part4

DB
```
# docker run postgre server
# user=postgres password=mysecretpassword
docker run --name some-postgres -p 5432:5432 -e POSTGRES_PASSWORD=mysecretpassword -d postgres

# connect postgre server
psql -h localhost -p 5432 -U postgres -W
\list
\dt

# create database
CREATE DATABASE go_grpc_demo;

# change db
\c go_grpc_demo

# create table
CREATE TYPE TaskStatus AS ENUM ('TASK_STATUS_COMPLETED', 'TASK_STATUS_INCOMPLETE', 'TASK_STATUS_STARTED');
CREATE TABLE tasks(
    id SERIAL PRIMARY KEY ,
    description TEXT,
    user_id INT,
    status TaskStatus,
    deadline timestamptz,
    created_at timestamptz
);

CREATE TABLE comments(
    id SERIAL PRIMARY KEY,
    task_id INT,
    user_id INT,
    comment TEXT,
    created_at timestamptz,
    CONSTRAINT fk_task
                     FOREIGN KEY (task_id)
                     REFERENCES tasks(id)
);

# select
select * from tasks;
select * from comments;
```

Test 
```
# run gRPC server
go run server/server.go

# run gRPC proxy
go run proxy/proxy.go

# test for gRPC request for Go
go run client/client.go  

# test for http request
curl http://localhost:8081/api/v1/tasks/1

# test for gRPC request
grpcurl -plaintext -d '{"task_id": "1"}' localhost:9090 tasks.TaskService.GetTask
```
