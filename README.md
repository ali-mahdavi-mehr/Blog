#### to generate grpc files
`protoc --go_out=./service \
--go-grpc_out=./service \
proto/*.proto
`

see [Document](https://grpc.io/docs/languages/go/quickstart/)  to install requirements for re-generating grpc files


### To Run Project
config .env file
`.env-sample => .env`

run docker-compose by:
> docker-compose up --build -d