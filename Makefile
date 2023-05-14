# Flutter
# json_annotation
# grpc
# json_serializable (dev)
# build_runner (dev)

dart-dep:
	(cd frontend && dart pub add json_annotation grpc && dart pub add --dev json_serializable build_runner && dart pub add --dev --path="../utils/generators/" generators)

rest:
	rm -rf frontend/lib/models/rest && find backend/rest/ -type d -exec make go-to-dart dir={} \; && make br && make rungen n=1

rungen:
	find frontend/lib/models/rest -type f -name "*.go.dart" -exec dart rungen.dart {} $(n) \;

go-to-dart:
	go-to-dart -i $(dir) -o frontend/lib/models/rest -m json && make rungen n=0

br:
	(cd frontend/ && dart run build_runner build)

igrpc:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest && dart pub global activate protoc_plugin && export PATH="$PATH:$HOME/.pub-cache/bin"

grpc:
	(rm -rf frontend/lib/models/grpc && mkdir -p frontend/lib/models/grpc && find backend/grpc/ -type f -name "*.proto" -exec make gen-grpc proto={} \;) && make tidy

tidy:
	(cd backend && go mod tidy) && (cd frontend && flutter pub get)

gen-grpc:
	protoc $(proto) --go_out=. --go-grpc_out=. --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative && protoc --dart_out=. $(proto) && find backend/grpc/ -type f -name "*.dart" -exec mv {} ./frontend/lib/models/grpc/ \;

go:
	(cd backend && go build -o bin/app && bin/app)

flutter:
	(cd frontend && flutter run)
