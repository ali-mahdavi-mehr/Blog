# Go Blog
a simple blog
with:
- Authentication
- CRUD post

###
the user can get posts by:
+ REST API Request
+ gRPC request
+ Graphql request


#### DataBases
+ Using Redis to store JWT token
+ Using Postgres for post, user



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
config __.env__ file by:
`.env-sample => .env`

run docker-compose by:
> docker-compose up --build -d