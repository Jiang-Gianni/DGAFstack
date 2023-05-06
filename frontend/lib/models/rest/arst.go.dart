import 'dart:io';
import 'dart:convert';
import 'package:http/http.dart' as http;

part 'arst.go.g.dart';

class Arst {
	final int id;
	final String name;
	
	Arst({
		required this.id,
		required this.name,
	});
	
	Map<String, dynamic> toJson() => _$ArstToJson(this);
	
	factory Arst.fromJson(Map<String, dynamic> json) => _$ArstFromJson(json);
}

