import 'package:analyzer/dart/element/visitor.dart';
import 'package:analyzer/dart/element/element.dart';

class ModelVisitor extends SimpleElementVisitor {
  String className;

  Map<String, dynamic> fields = {};

  @override
  visitConstructorElement(ConstructorElement element) {
    className = element.type.returnType.toString().replaceFirst("*", "");
  }

  @override
  visitFieldElement(FieldElement element) {
    fields[element.name] = element.type.toString().replaceFirst("*", "");
  }
}
