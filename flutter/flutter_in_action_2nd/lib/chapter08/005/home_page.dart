import 'package:flutter/material.dart';

import 'a.dart';

class HomePage extends StatelessWidget {
  const HomePage({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('事件总线')),
      body: const A(),
    );
  }
}
