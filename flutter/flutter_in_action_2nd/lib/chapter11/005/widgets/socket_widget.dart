import 'dart:convert';

import 'package:flutter/material.dart';

import 'dart:io';

class SocketWidget extends StatelessWidget {
  const SocketWidget({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    String _text = '';

    return FutureBuilder(
      future: _request(),
      builder: (context, snapshot) {
        if (snapshot.hasError) {
          _text = 'error: ${snapshot.error}';
        } else if (snapshot.hasData) {
          _text = '${snapshot.data}';
        }

        return Text(_text);
      },
    );
  }
}

Future<String> _request() async {
  // 建立连接
  Socket socket = await Socket.connect('cn.bing.com', 80);
  // 根据http协议， 发送请求头
  socket.writeln('GET / HTTP/1.1');
  socket.writeln('Host:cn.bing.com');
  socket.writeln('Connection:close');
  socket.writeln();

  // 发送请求
  await socket.flush();

  // 读取返回内容，按utf8解码为字符串
  String _response = await utf8.decoder.bind(socket).join();

  // 关闭连接
  await socket.close();

  return _response;
}
