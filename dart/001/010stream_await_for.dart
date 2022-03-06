import 'dart:io';

const String searchPath = './';
const String searchTerms = 'dart';
void main(List<String> arguments) {
  s001();
  s002();
}

void s001() {
  FileSystemEntity.isDirectory(searchPath).then((isDir) {
    if (isDir) {
      final startingDir = Directory(searchPath);
      startingDir.list().listen((entity) {
        if (entity is File) {
          searchFile(entity, searchTerms);
        }
      });
    } else {
      searchFile(File(searchPath), searchTerms);
    }
  });
}

void s002() async {
  if (await FileSystemEntity.isDirectory(searchPath)) {
    final startDir = Directory(searchPath);
    await for (final entity in startDir.list()) {
      if (entity is File) {
        searchFile(entity, searchTerms);
      }
    }
  } else {
    searchFile(File(searchPath), searchTerms);
  }
}

void searchFile(File entity, String searchTerms) {}
