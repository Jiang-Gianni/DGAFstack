import 'package:build/build.dart';
import 'package:source_gen/source_gen.dart';

import 'src/extension_generator.dart';
import 'src/my_generator.dart';
import 'src/rest_generator.dart';

Builder extGen(BuilderOptions options) =>
    SharedPartBuilder([ExtensionGenerator()], 'extension_generator');

Builder myGen(BuilderOptions options) =>
    SharedPartBuilder([MyGenerator()], 'my_generator');

Builder restApiGen(BuilderOptions options) =>
    SharedPartBuilder([RestApiGenerator()], 'rest_generator');
