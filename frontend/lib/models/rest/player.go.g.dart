// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'player.go.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

Player _$PlayerFromJson(Map<String, dynamic> json) => Player(
      id: json['id'] as int,
      name: json['name'] as String,
      email: json['Email'] as String,
      password: json['Password'] as String,
      createdAt: DateTime.parse(json['CreatedAt'] as String),
      updatedAt: DateTime.parse(json['UpdatedAt'] as String),
      deletedAt: json['DeletedAt'] == null
          ? null
          : DateTime.parse(json['DeletedAt'] as String),
      options: (json['Options'] as Map<String, dynamic>?)?.map(
            (k, e) => MapEntry(k, e as String),
          ) ??
          {},
      tags:
          (json['Tags'] as List<dynamic>?)?.map((e) => e as String).toList() ??
              [],
    );

Map<String, dynamic> _$PlayerToJson(Player instance) => <String, dynamic>{
      'id': instance.id,
      'name': instance.name,
      'Email': instance.email,
      'Password': instance.password,
      'CreatedAt': instance.createdAt.toIso8601String(),
      'UpdatedAt': instance.updatedAt.toIso8601String(),
      'DeletedAt': instance.deletedAt?.toIso8601String(),
      'Options': instance.options,
      'Tags': instance.tags,
    };
