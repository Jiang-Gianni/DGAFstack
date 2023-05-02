# DGAF stack
## The **Don't Give A F\*ck** stack
- [Docker](https://www.docker.com/): creates and runs a container for the backend written in Go. I like to use [render](https://render.com/) to host it: they also offer a Go environment but I needed the D for this stack;
- [Go](https://go.dev/): just an amazing language to write the backend with;
- [AstraDB](https://www.datastax.com/products/datastax-astra/pricing): serverless DBaaS based on Cassandra, they offer you 25$ of use for free each month (which is a lot in terms of read/write/storage) and you can use different types of API through the [Stargate](https://www.datastax.com/products/datastax-astra/apis) gateway (I was surprised to see gRPC available);
- [Flutter](https://flutter.dev/): IMHO the best mobile development framework.

## Flutter Project Name Change
I originally ran "flutter create frontend".\
This command replaces all instances of the string "frontend":
```bash
find frontend/ -type f -exec sed -i -e 's/frontend/{YOUR_FLUTTER_PROJECT_NAME}/g' {} \;
```


## Go Module Name Change
I originally ran "go mod init github.com/Jiang-Gianni/DGAFstack".\
This command replaces all instances of the above repo's string:
```bash
find backend/ -type f -exec sed -i -e 's/github.com\/Jiang-Gianni\/DGAFstack/{YOUR_GO_MODULE_NAME}/g' {} \;
```

## REST
Once you define your Go's data structs in `backend\rest` it is possible to generate Dart's equivalent json serializable classes for the REST API.
### Install [go-to-dart](https://github.com/11wizards/go-to-dart)
```bash
go install github.com/11wizards/go-to-dart
```

### Generate Dart files
```bash
make rest
```
This will **remove** the `frontend/lib/models/rest` directory and recreate it with a {go_package}.go.dart file and a {go_package}.go.g.dart file for each distinct package defined in the `backend/rest` directory (if there are multiple .go files of the same package then only one is considered). \
Refer to the `Makefile` to see what this last command does.

## gRPC
Once you define your proto files in `backend\grpc`