import 'package:dio/dio.dart';
import 'dart:io';

Future downloadWithChunks(
  url,
  savePath, {
  required ProgressCallback onReceiveProgress,
}) async {
  const firstChunkSize = 102;
  const maxChunk = 3;

  int total = 0;
  var dio = Dio();
  List<int> progress = [];

  createCallback(no) {
    return (int received, _) {
      progress[no] = received;
      if (total != 0) {
        onReceiveProgress(progress.reduce((a, b) => a + b), total);
      }
    };
  }

  Future<Response> downloadChunk(url, start, end, no) async {
    progress.add(0);
    --end;
    return dio.download(
      url,
      savePath + 'temp$no',
      onReceiveProgress: createCallback(no),
      options: Options(
        headers: {'range': 'bytes=$start-$end'},
      ),
    );
  }

  Future mergeTempFiles(chunk) async {
    File f = File(savePath + 'temp0');
    IOSink ioSink = f.openWrite(mode: FileMode.writeOnlyAppend);
    for (int i = 1; i < chunk; ++i) {
      File _f = File(savePath + 'temp$i');
      await ioSink.addStream(_f.openRead());
      await _f.delete();
    }

    await ioSink.close();
    await f.rename(savePath);
  }

  Response response = await downloadChunk(url, 0, firstChunkSize, 0);
  if (response.statusCode == 206) {
    total = int.parse(response.headers
        .value(HttpHeaders.contentRangeHeader)!
        .split("/")
        .last);
    int reserved = total -
        int.parse(response.headers.value(HttpHeaders.contentLengthHeader)!);
    int chunk = (reserved / firstChunkSize).ceil() + 1;
    if (chunk > 1) {
      int chunkSize = firstChunkSize;
      if (chunk > maxChunk + 1) {
        chunk = maxChunk + 1;
        chunkSize = (reserved / maxChunk).ceil();
      }

      List<Future> futures = [];
      for (var i = 0; i < maxChunk; ++i) {
        var start = firstChunkSize + i * chunkSize;
        futures.add(downloadChunk(url, start, start + chunkSize, i + 1));
      }
      await Future.wait(futures);
    }
    await mergeTempFiles(chunk);
  }
}

main() async {
  var url =
      'https://code.visualstudio.com/sha/download?build=stable&os=linux-deb-x64';
  var savePath = './example/vscode-linux-deb-x64.deb';
  await downloadWithChunks(url, savePath, onReceiveProgress: (received, total) {
    if (total != -1) {
      print('${(received / total * 100).floor()}%');
    }
  });
}
