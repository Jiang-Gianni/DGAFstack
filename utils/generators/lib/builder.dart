import 'package:build/build.dart';
import 'package:source_gen/source_gen.dart';

import 'src/extension_generator.dart';
import 'src/my_generator.dart';

Builder extGen(BuilderOptions options) =>
    SharedPartBuilder([ExtensionGenerator()], 'extension_generator');

Builder myGen(BuilderOptions options) =>
    SharedPartBuilder([MyGenerator()], 'my_generator');
