# Go Blog
a simple blog
with:
- Authentication
- CRUD post

### Client Usage
the clients can fetch posts in three ways:
+ HTTP Request (REST API)
+ gRPC request
+ Graphql


#### DataBases
+ Using Redis to store JWT token
+ Using Postgres for post, user



#### Generate files
#### to regenerate grpc files

```
protoc --go_out=./service \
--go-grpc_out=./service \
proto/*.proto
```

#### to regenerate Graphql files

```
go generate cmd/graph/graph
```

### Requirements
***just read the documentation*** :smiley:
- see [Document](https://grpc.io/docs/languages/go/quickstart/)  to install requirements for re-generating grpc files


## Usage

#### To Run Project
config __.env__ file by:
`.env-sample => .env`
###### then
run docker-compose by:
> docker-compose up --build -d
