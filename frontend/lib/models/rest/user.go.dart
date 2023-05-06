import 'dart:io';
import 'dart:convert';
import 'package:http/http.dart' as http;

part 'user.go.g.dart';

class User {
	final int id;
	final String name;
	
	User({
		required this.id,
		required this.name,
	});
	
	Map<String, dynamic> toJson() => _$UserToJson(this);
	
	factory User.fromJson(Map<String, dynamic> json) => _$UserFromJson(json);
}

