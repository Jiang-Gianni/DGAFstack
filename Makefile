# Flutter
# json_annotation
# grpc
# json_serializable (dev)
# build_runner (dev)

rest:
	rm -rf frontend/lib/models/rest && find backend/rest/ -type d -exec make go-to-dart dir={} \;

make go-to-dart:
	go-to-dart -i $(dir) -o frontend/lib/models/rest -m json && (cd frontend/ && dart run build_runner build)

igrpc:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest && dart pub global activate protoc_plugin && export PATH="$PATH:$HOME/.pub-cache/bin"

grpc:
	(rm -rf frontend/lib/models/grpc && mkdir -p frontend/lib/models/grpc && find backend/grpc/ -type f -name "*.proto" -exec make gen-grpc proto={} \;) && make tidy

tidy:
	(cd backend && go mod tidy)

gen-grpc:
	protoc $(proto) --go_out=. --go-grpc_out=. --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative && protoc --dart_out=. $(proto) && find backend/grpc/ -type f -name "*.dart" -exec mv {} ./frontend/lib/models/grpc \;
