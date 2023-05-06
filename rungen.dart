import 'dart:io';

void main(List<String> args) async {
  const finalImports = '''
import 'dart:io';
import 'dart:convert';
import 'package:http/http.dart' as http;
''';

  const annotationsImports = '''
import 'package:json_annotation/json_annotation.dart';
import 'package:generators/src/rest_generator.dart';
''';

  const annotations = '''
@RestApiAnnotation()
@JsonSerializable()
''';

  StringBuffer cleanContents = StringBuffer();

  String filePath = './' + args[0];
  String zeroOne = args[1];

  await File(filePath).readAsLines().then(
    (value) {
      value.forEach(
        (element) {
          if (!element.startsWith("import") && !element.contains("@")) {
            cleanContents.writeln(element);
          }
        },
      );
    },
  );

  StringBuffer finalContents =
      StringBuffer(finalImports + cleanContents.toString());

  var classRegExpr = RegExp("(class.*{)");

  StringBuffer generationContents = StringBuffer(
      (annotationsImports + finalContents.toString()).replaceAllMapped(
    classRegExpr,
    (match) => "$annotations${match[1]}",
  ));

  if (zeroOne == "0") {
    await File(filePath).writeAsString(
      generationContents.toString(),
    );
  } else {
    await File(filePath).writeAsString(
      finalContents.toString(),
    );
  }
}
