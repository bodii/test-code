import 'dart:html';
import 'package:http/http.dart' as http;
// ignore: import_deferred_library_with_load_function
import 'lib/src/deferred/hello.dart' deferred as hello;

Future<void> greet() async {
  await hello.loadLibrary();
  hello.PrintGreeting();
}
