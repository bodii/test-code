import 'dart:convert';
import "dart:io";

Future<void> streamApi() async {
  var config = File('config.txt');
  Stream<List<int>> inputStream = config.openRead();

  inputStream.transform(utf8.decoder).transform(const LineSplitter()).listen(
      (String line) {
    print('Got ${line.length} characters from stream');
  }, onDone: () {
    print('file is now closed');
  }, onError: (e) {
    print(e);
  });
}
