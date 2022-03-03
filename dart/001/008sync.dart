import 'dart:convert';
import 'dart:io';

void main() {
  final fileData = _readFileSync();
  final jsonData = jsonDecode(fileData);

  print('Number of JSON keys: ${jsonData.length}');
}

String _readFileSync() {
  var filename = 'test.file';
  final file = File(filename);
  final contents = file.readAsStringSync();
  return contents.trim();
}
