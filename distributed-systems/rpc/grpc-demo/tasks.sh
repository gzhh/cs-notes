# protoc
protoc --go_out=./src/go \
--go_opt=paths=source_relative --go-grpc_out=./src/go/ \
--go-grpc_opt=paths=source_relative \
--grpc-gateway_out=./src/go \
--grpc-gateway_opt=logtostderr=true,paths=source_relative,grpc_api_configuration=tasks_http_annotations.yaml \
--openapiv2_out=./src/openapi \
--openapiv2_opt=logtostderr=true,grpc_api_configuration=tasks_http_annotations.yaml,openapi_configuration=tasks_openapi_options.yaml \