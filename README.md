# Go Blog


#### Generate proto files
#### to generate grpc files

```
protoc --go_out=./service \
--go-grpc_out=./service \
proto/*.proto
```

### Requirements
***just read the documentation*** :smiley:
- see [Document](https://grpc.io/docs/languages/go/quickstart/)  to install requirements for re-generating grpc files


## Usage

#### To Run Project
config .env file
`.env-sample => .env`

run docker-compose by:
> docker-compose up --build -d