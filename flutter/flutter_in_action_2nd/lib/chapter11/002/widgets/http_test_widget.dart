import 'package:flutter/material.dart';
import 'dart:io';
import 'dart:convert';

class HttpTestWidget extends StatefulWidget {
  const HttpTestWidget({Key? key}) : super(key: key);

  @override
  _HttpTestWidgetState createState() => _HttpTestWidgetState();
}

class _HttpTestWidgetState extends State<HttpTestWidget> {
  bool _loading = false;
  String _text = '';

  @override
  Widget build(BuildContext context) {
    return SingleChildScrollView(
      child: Column(
        children: [
          ElevatedButton(
            child: const Text('获取必应首页内容'),
            onPressed: _loading ? null : request,
          ),
          SizedBox(
            width: MediaQuery.of(context).size.width - 50,
            child: Text(_text.replaceAll(r'\s', '')),
          ),
        ],
      ),
    );
  }

  void request() async {
    setState(() {
      _loading = true;
      _text = '正在请求...';
    });
    try {
      // 创建一个HttpClient
      HttpClient httpClient = HttpClient();
      // 打开http连接
      HttpClientRequest request =
          await httpClient.getUrl(Uri.parse('https://cn.bing.com'));
      // 使用iphone的UA
      request.headers.add(
        'user-agent',
        'Mozilla/5.0 (iPhone; CPU iPhone OS 10_3_1 like Mac OS X) AppleWebKit/603.1.30 (KHTML, like Gecko) Version/10.0 Mobile/14E304 Safari/602.1',
      );
      // 等待连接服务器
      HttpClientResponse response = await request.close();
      // 读取响应内容
      _text = await response.transform(utf8.decoder).join();
      // 输出响应头
      debugPrint(response.headers.toString());

      // 关闭client后，通过该client发起的所有请求都会中止。
      httpClient.close();
    } catch (e) {
      _text = '请求失败: $e';
    } finally {
      setState(() {
        _loading = false;
      });
    }
  }
}
