import 'package:flutter/material.dart';

import 'widgets/http_test_widget.dart';

class HomePage extends StatelessWidget {
  const HomePage({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('通过HttpClient发起HTTP请求')),
      body: const HttpTestWidget(),
    );
  }
}
