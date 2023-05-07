import 'package:build/src/builder/build_step.dart';
import 'package:analyzer/dart/element/element.dart';
import 'package:source_gen/source_gen.dart';

import '../model_visitor.dart';

class RestApiAnnotation {
  const RestApiAnnotation();
}

class RestApiGenerator extends GeneratorForAnnotation<RestApiAnnotation> {
  @override
  generateForAnnotatedElement(
      Element element, ConstantReader annotation, BuildStep buildStep) {
    return _generatedSource(element);
  }

  String _generatedSource(Element element) {
    var visitor = ModelVisitor();

    element.visitChildren(visitor);

    var className = visitor.className;

    var classBuffer = StringBuffer();

    var myString = '''

class ${className}RestApi {
  final http.Client client;
  final String url${className};
  ${className}RestApi({
    required this.client,
    required this.url${className},
  });

  List<${className}> decodeJsonResponseBody(http.Response response) {
    Iterable it = json.decode(response.body);
    return List<${className}>.from(
      it.map((e) => ${className}.fromJson(e)),
    );
  }

  String encode${className}sToJson(List<${className}> list${className}) {
    return json.encode(
        List<Map<String, dynamic>>.from(list${className}.map((e) => e.toJson())));
  }

  Future<List<${className}>> get${className}() async {
    final response = await http.get(Uri.parse(url${className}));
    if (response.statusCode == HttpStatus.ok) {
      return decodeJsonResponseBody(response);
    } else {
      throw Exception('Failed to get ${className}');
    }
  }

  Future<${className}> get${className}ById(String id) async {
    final response = await http.get(Uri.parse("\$url${className}\\\$id"));
    if (response.statusCode == HttpStatus.ok) {
      return ${className}.fromJson(json.decode(response.body));
    } else {
      throw Exception('Failed to get ${className} with id: \$id');
    }
  }

  Future<List<${className}>> create${className}sWithReturn(
      List<${className}> toBeCreated${className}s) async {
    final response = await http.post(Uri.parse(url${className}),
        body: encode${className}sToJson(toBeCreated${className}s),
        headers: {
          "Content-Type": "application/json",
        });
    if (response.statusCode == HttpStatus.ok) {
      return decodeJsonResponseBody(response);
    } else {
      throw Exception('Failed to create ${className} ');
    }
  }

  void create${className}s(List<${className}> toBeCreated${className}s) async {
    final response = await http.post(Uri.parse(url${className}),
        body: encode${className}sToJson(toBeCreated${className}s),
        headers: {
          "Content-Type": "application/json",
        });
    if (response.statusCode == HttpStatus.created) {
      return;
    } else {
      throw Exception('Failed to create ${className} ');
    }
  }

  Future<${className}> update${className}(
      String id, ${className} toBeUpdated${className}) async {
    final response = await http.put(
      Uri.parse("\$url${className}\\\$id"),
      body: toBeUpdated${className}.toJson(),
      headers: {
        "Content-Type": "application/json",
      },
    );
    if (response.statusCode == HttpStatus.ok) {
      return ${className}.fromJson(json.decode(response.body));
    } else {
      throw Exception('Failed to update ${className} with id: \$id');
    }
  }

  void delete${className}(String id, ${className} toBeDeleted${className}) async {
    final response = await http.delete(
      Uri.parse("\$url${className}\\\$id"),
      body: toBeDeleted${className}.toJson(),
      headers: {
        "Content-Type": "application/json",
      },
    );
    if (response.statusCode == HttpStatus.ok) {
      return;
    } else {
      throw Exception('Failed to delete ${className} with id: \$id');
    }
  }
}





''';
    classBuffer.writeln(myString);
    return classBuffer.toString();
  }
}
