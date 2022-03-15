set PROTO_PATH=.\auth\api
set GO_OUT_PATH=.\auth\api\gen\v1
md %GO_OUT_PATH%

protoc --proto_path=%PROTO_PATH% --go_out=plugins=grpc,paths=source_relative:%GO_OUT_PATH% auth.proto
protoc --proto_path=%PROTO_PATH% --grpc-gateway_out=paths=source_relative,grpc_api_configuration=%PROTO_PATH%\auth.yaml:%GO_OUT_PATH% auth.proto