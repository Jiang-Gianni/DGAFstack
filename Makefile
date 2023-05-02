# Flutter
# json_annotation
# json_serializable (dev)
# build_runner (dev)

make go-to-dart:
	go-to-dart -i $(dir) -o frontend/lib/models/rest -m json && (cd frontend/ && dart run build_runner build)

rest:
	rm -rf frontend/lib/models/rest && find backend/rest/ -type d -exec make go-to-dart dir={} \;
