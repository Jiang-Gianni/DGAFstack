// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'user.go.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

User _$UserFromJson(Map<String, dynamic> json) => User(
      id: json['id'] as String,
      name: json['name'] as String,
    );

Map<String, dynamic> _$UserToJson(User instance) => <String, dynamic>{
      'id': instance.id,
      'name': instance.name,
    };

// **************************************************************************
// RestApiGenerator
// **************************************************************************

class UserRestApi {
  final http.Client client;
  final String urlUser;
  UserRestApi({
    required this.client,
    required this.urlUser,
  });

  List<User> decodeJsonResponseBody(http.Response response) {
    Iterable it = json.decode(response.body);
    return List<User>.from(
      it.map((e) => User.fromJson(e)),
    );
  }

  String encodeUsersToJson(List<User> listUser) {
    return json.encode(
        List<Map<String, dynamic>>.from(listUser.map((e) => e.toJson())));
  }

  Future<List<User>> getUser() async {
    final response = await http.get(Uri.parse(urlUser));
    if (response.statusCode == HttpStatus.ok) {
      return decodeJsonResponseBody(response);
    } else {
      throw Exception('Failed to get User');
    }
  }

  Future<User> getUserById(String id) async {
    final response = await http.get(Uri.parse("$urlUser\$id"));
    if (response.statusCode == HttpStatus.ok) {
      return User.fromJson(json.decode(response.body));
    } else {
      throw Exception('Failed to get User with id: $id');
    }
  }

  Future<List<User>> createUsersWithReturn(List<User> toBeCreatedUsers) async {
    final response = await http.post(Uri.parse(urlUser),
        body: encodeUsersToJson(toBeCreatedUsers),
        headers: {
          "Content-Type": "application/json",
        });
    if (response.statusCode == HttpStatus.ok) {
      return decodeJsonResponseBody(response);
    } else {
      throw Exception('Failed to create User ');
    }
  }

  void createUsers(List<User> toBeCreatedUsers) async {
    final response = await http.post(Uri.parse(urlUser),
        body: encodeUsersToJson(toBeCreatedUsers),
        headers: {
          "Content-Type": "application/json",
        });
    if (response.statusCode == HttpStatus.created) {
      return;
    } else {
      throw Exception('Failed to create User ');
    }
  }

  Future<User> updateUser(String id, User toBeUpdatedUser) async {
    final response = await http.put(
      Uri.parse("$urlUser\$id"),
      body: toBeUpdatedUser.toJson(),
      headers: {
        "Content-Type": "application/json",
      },
    );
    if (response.statusCode == HttpStatus.ok) {
      return User.fromJson(json.decode(response.body));
    } else {
      throw Exception('Failed to update User with id: $id');
    }
  }

  void deleteUser(String id, User toBeDeletedUser) async {
    final response = await http.delete(
      Uri.parse("$urlUser\$id"),
      body: toBeDeletedUser.toJson(),
      headers: {
        "Content-Type": "application/json",
      },
    );
    if (response.statusCode == HttpStatus.ok) {
      return;
    } else {
      throw Exception('Failed to delete User with id: $id');
    }
  }
}
