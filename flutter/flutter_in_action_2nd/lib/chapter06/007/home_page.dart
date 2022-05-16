import 'package:flutter/material.dart';

import 'widgets/use_page_view_01_widget.dart';
import 'widgets/keep_alive_test.dart';

class HomePage extends StatelessWidget {
  const HomePage({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('PageView KeepAlive test')),
      body: const KeepAliveTest(),
    );
  }
}
