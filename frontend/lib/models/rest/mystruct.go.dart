import 'dart:io';
import 'dart:convert';
import 'package:http/http.dart' as http;

part 'mystruct.go.g.dart';

class MyStruct {
	final String uuid;
	final String name;
	final int number;
	final bool myBoolean;
	
	MyStruct({
		required this.uuid,
		required this.name,
		required this.number,
		required this.myBoolean,
	});
	
	Map<String, dynamic> toJson() => _$MyStructToJson(this);
	
	factory MyStruct.fromJson(Map<String, dynamic> json) => _$MyStructFromJson(json);
}

