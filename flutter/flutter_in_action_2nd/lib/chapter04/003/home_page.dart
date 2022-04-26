import 'package:flutter/material.dart';

import 'widgets/use_flex_widget.dart';

class HomePage extends StatelessWidget {
  const HomePage({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('弹性布局flex、row、column')),
      body: Column(
        children: const [
          UseFlexWidget(),
        ],
      ),
    );
  }
}
