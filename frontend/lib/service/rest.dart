import 'dart:convert';

import 'package:http/http.dart' as http;

T jsonFromResponse<T>(http.Response response) {
  print(T);
  return json.decode(
    response.body,
  );
}

abstract class FromJson {}
