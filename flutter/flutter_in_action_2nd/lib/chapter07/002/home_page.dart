import 'package:flutter/material.dart';

import 'widgets/inherited_widget_test.dart';

class HomePage extends StatelessWidget {
  const HomePage({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('数据共享InheritedWidget')),
      body: const InheritedWidgetTest(),
    );
  }
}
