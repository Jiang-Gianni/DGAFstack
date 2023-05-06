// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'arst.go.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

Arst _$ArstFromJson(Map<String, dynamic> json) => Arst(
      id: json['id'] as int,
      name: json['name'] as String,
    );

Map<String, dynamic> _$ArstToJson(Arst instance) => <String, dynamic>{
      'id': instance.id,
      'name': instance.name,
    };

// **************************************************************************
// RestApiGenerator
// **************************************************************************

class ArstRestApi {
  final http.Client client;
  final String urlArst;
  ArstRestApi({
    required this.client,
    required this.urlArst,
  });

  List<Arst> decodeJsonResponseBody(http.Response response) {
    Iterable it = json.decode(response.body);
    return List<Arst>.from(
      it.map((e) => Arst.fromJson(e)),
    );
  }

  String encodeArstsToJson(List<Arst> listArst) {
    return json.encode(
        List<Map<String, dynamic>>.from(listArst.map((e) => e.toJson())));
  }

  Future<List<Arst>> getArst() async {
    final response = await http.get(Uri.parse(urlArst));
    if (response.statusCode == HttpStatus.ok) {
      return decodeJsonResponseBody(response);
    } else {
      throw Exception('Failed to get Arst');
    }
  }

  Future<Arst> getArstById(String id) async {
    final response = await http.get(Uri.parse(urlArst + id));
    if (response.statusCode == HttpStatus.ok) {
      return Arst.fromJson(json.decode(response.body));
    } else {
      throw Exception('Failed to get Arst with id: $id');
    }
  }

  Future<List<Arst>> postArstsWithReturn(List<Arst> postArsts) async {
    final response = await http
        .post(Uri.parse(urlArst), body: encodeArstsToJson(postArsts), headers: {
      "Content-Type": "application/json",
    });
    if (response.statusCode == HttpStatus.ok) {
      return decodeJsonResponseBody(response);
    } else {
      throw Exception('Failed to post Arst ');
    }
  }

  void postArsts(List<Arst> postArsts) async {
    final response = await http
        .post(Uri.parse(urlArst), body: encodeArstsToJson(postArsts), headers: {
      "Content-Type": "application/json",
    });
    if (response.statusCode == HttpStatus.created) {
      return;
    } else {
      throw Exception('Failed to post Arst ');
    }
  }

  Future<Arst> putArst(String id, Arst putArst) async {
    final response = await http.put(
      Uri.parse(urlArst + id),
      body: putArst.toJson(),
      headers: {
        "Content-Type": "application/json",
      },
    );
    if (response.statusCode == HttpStatus.ok) {
      return Arst.fromJson(json.decode(response.body));
    } else {
      throw Exception('Failed to put Arst with id: $id');
    }
  }

  void deleteArst(String id, Arst putArst) async {
    final response = await http.put(
      Uri.parse(urlArst + id),
      body: putArst.toJson(),
      headers: {
        "Content-Type": "application/json",
      },
    );
    if (response.statusCode == HttpStatus.ok) {
      return;
    } else {
      throw Exception('Failed to delete Arst with id: $id');
    }
  }
}
