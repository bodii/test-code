import 'package:flutter/material.dart';
import 'dart:io';
import 'dart:async';
import 'package:path_provider/path_provider.dart';

class FileOperationWidget extends StatefulWidget {
  const FileOperationWidget({Key? key}) : super(key: key);

  @override
  _FileOperationWidgetState createState() => _FileOperationWidgetState();
}

class _FileOperationWidgetState extends State<FileOperationWidget> {
  int _counter = 0;

  @override
  void initState() {
    super.initState();
    _readCounter().then((int value) {
      setState(() {
        _counter = value;
      });
    });
  }

  Future<File> _getLocalFile() async {
    String dir = (await getApplicationDocumentsDirectory()).path;

    return File('$dir/counter.text');
  }

  Future<int> _readCounter() async {
    try {
      File file = await _getLocalFile();
      String contents = await file.readAsString();

      return int.parse(contents);
    } on FileSystemException {
      return 0;
    }
  }

  void _incrementCounter() async {
    setState(() {
      _counter++;
    });

    // 将点击将数字符串类型写入文件
    await (await _getLocalFile()).writeAsString('$_counter');
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('文件操作')),
      body: Center(
        child: Text('点击了 $_counter 次'),
      ),
      floatingActionButton: FloatingActionButton(
        onPressed: _incrementCounter,
        tooltip: 'Increment',
        child: const Icon(Icons.add),
      ),
    );
  }
}
