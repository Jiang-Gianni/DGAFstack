import 'dart:convert';
import 'dart:io';

import 'package:http/http.dart' as http;

class className {
  static fromJson(e) {}

  // static toJson(className e) {}
  Map<String, dynamic> toJson() {
    return {};
  }
}

class classNameRestApi {
  final http.Client client;
  final String urlclassName;
  classNameRestApi({
    required this.client,
    required this.urlclassName,
  });

  List<className> decodeJsonResponseBody(http.Response response) {
    Iterable it = json.decode(response.body);
    return List<className>.from(
      it.map((e) => className.fromJson(e)),
    );
  }

  String encodeclassNamesToJson(List<className> listclassName) {
    return json.encode(
        List<Map<String, dynamic>>.from(listclassName.map((e) => e.toJson())));
  }

  Future<List<className>> getclassName() async {
    final response = await http.get(Uri.parse(urlclassName));
    if (response.statusCode == HttpStatus.ok) {
      return decodeJsonResponseBody(response);
    } else {
      throw Exception('Failed to get className');
    }
  }

  Future<className> getclassNameById(String id) async {
    final response = await http.get(Uri.parse("\$urlclassName\\\$id"));
    if (response.statusCode == HttpStatus.ok) {
      return className.fromJson(json.decode(response.body));
    } else {
      throw Exception('Failed to get className with id: \$id');
    }
  }

  Future<List<className>> createclassNamesWithReturn(
      List<className> toBeCreatedclassNames) async {
    final response = await http.post(Uri.parse(urlclassName),
        body: encodeclassNamesToJson(toBeCreatedclassNames),
        headers: {
          "Content-Type": "application/json",
        });
    if (response.statusCode == HttpStatus.ok) {
      return decodeJsonResponseBody(response);
    } else {
      throw Exception('Failed to create className ');
    }
  }

  void createclassNames(List<className> toBeCreatedclassNames) async {
    final response = await http.post(Uri.parse(urlclassName),
        body: encodeclassNamesToJson(toBeCreatedclassNames),
        headers: {
          "Content-Type": "application/json",
        });
    if (response.statusCode == HttpStatus.created) {
      return;
    } else {
      throw Exception('Failed to create className ');
    }
  }

  Future<className> updateclassName(
      String id, className toBeUpdatedclassName) async {
    final response = await http.put(
      Uri.parse("\$urlclassName\\\$id"),
      body: toBeUpdatedclassName.toJson(),
      headers: {
        "Content-Type": "application/json",
      },
    );
    if (response.statusCode == HttpStatus.ok) {
      return className.fromJson(json.decode(response.body));
    } else {
      throw Exception('Failed to update className with id: \$id');
    }
  }

  void deleteclassName(String id, className toBeDeletedclassName) async {
    final response = await http.delete(
      Uri.parse("\$urlclassName\\\$id"),
      body: toBeDeletedclassName.toJson(),
      headers: {
        "Content-Type": "application/json",
      },
    );
    if (response.statusCode == HttpStatus.ok) {
      return;
    } else {
      throw Exception('Failed to delete className with id: \$id');
    }
  }
}
