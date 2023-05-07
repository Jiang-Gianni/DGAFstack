// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'mystruct.go.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

MyStruct _$MyStructFromJson(Map<String, dynamic> json) => MyStruct(
      uuid: json['uuid'] as String,
      name: json['name'] as String,
      number: json['number'] as int,
      myBoolean: json['myBoolean'] as bool,
    );

Map<String, dynamic> _$MyStructToJson(MyStruct instance) => <String, dynamic>{
      'uuid': instance.uuid,
      'name': instance.name,
      'number': instance.number,
      'myBoolean': instance.myBoolean,
    };

// **************************************************************************
// RestApiGenerator
// **************************************************************************

class MyStructRestApi {
  final http.Client client;
  final String urlMyStruct;
  MyStructRestApi({
    required this.client,
    required this.urlMyStruct,
  });

  List<MyStruct> decodeJsonResponseBody(http.Response response) {
    Iterable it = json.decode(response.body);
    return List<MyStruct>.from(
      it.map((e) => MyStruct.fromJson(e)),
    );
  }

  String encodeMyStructsToJson(List<MyStruct> listMyStruct) {
    return json.encode(
        List<Map<String, dynamic>>.from(listMyStruct.map((e) => e.toJson())));
  }

  Future<List<MyStruct>> getMyStruct() async {
    final response = await http.get(Uri.parse(urlMyStruct));
    if (response.statusCode == HttpStatus.ok) {
      return decodeJsonResponseBody(response);
    } else {
      throw Exception('Failed to get MyStruct');
    }
  }

  Future<MyStruct> getMyStructById(String id) async {
    final response = await http.get(Uri.parse("$urlMyStruct\$id"));
    if (response.statusCode == HttpStatus.ok) {
      return MyStruct.fromJson(json.decode(response.body));
    } else {
      throw Exception('Failed to get MyStruct with id: $id');
    }
  }

  Future<List<MyStruct>> createMyStructsWithReturn(
      List<MyStruct> toBeCreatedMyStructs) async {
    final response = await http.post(Uri.parse(urlMyStruct),
        body: encodeMyStructsToJson(toBeCreatedMyStructs),
        headers: {
          "Content-Type": "application/json",
        });
    if (response.statusCode == HttpStatus.ok) {
      return decodeJsonResponseBody(response);
    } else {
      throw Exception('Failed to create MyStruct ');
    }
  }

  void createMyStructs(List<MyStruct> toBeCreatedMyStructs) async {
    final response = await http.post(Uri.parse(urlMyStruct),
        body: encodeMyStructsToJson(toBeCreatedMyStructs),
        headers: {
          "Content-Type": "application/json",
        });
    if (response.statusCode == HttpStatus.created) {
      return;
    } else {
      throw Exception('Failed to create MyStruct ');
    }
  }

  Future<MyStruct> updateMyStruct(
      String id, MyStruct toBeUpdatedMyStruct) async {
    final response = await http.put(
      Uri.parse("$urlMyStruct\$id"),
      body: toBeUpdatedMyStruct.toJson(),
      headers: {
        "Content-Type": "application/json",
      },
    );
    if (response.statusCode == HttpStatus.ok) {
      return MyStruct.fromJson(json.decode(response.body));
    } else {
      throw Exception('Failed to update MyStruct with id: $id');
    }
  }

  void deleteMyStruct(String id, MyStruct toBeDeletedMyStruct) async {
    final response = await http.delete(
      Uri.parse("$urlMyStruct\$id"),
      body: toBeDeletedMyStruct.toJson(),
      headers: {
        "Content-Type": "application/json",
      },
    );
    if (response.statusCode == HttpStatus.ok) {
      return;
    } else {
      throw Exception('Failed to delete MyStruct with id: $id');
    }
  }
}
