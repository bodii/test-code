import 'dart:convert';
import 'dart:io';

void main() async {
  final fileData = await _readFileAsync();
  final jsonData = jsonDecode(fileData);

  print('Number or JSON keys: ${jsonData.length}');
}

Future<String> _readFileAsync() async {
  var filename = 'test.file';
  final file = File(filename);
  final contents = await file.readAsString();
  return contents.trim();
}
