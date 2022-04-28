import 'package:flutter/material.dart';

import 'widgets/use_constrained_box_widget.dart';

class HomePage extends StatelessWidget {
  const HomePage({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('尺寸限制类容器')),
      body: Column(
        children: const [
          UseConstrainedBoxWidget(),
        ],
      ),
    );
  }
}
