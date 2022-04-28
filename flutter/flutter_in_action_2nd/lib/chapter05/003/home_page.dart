import 'package:flutter/material.dart';

import 'widgets/use_decorated_box_01_widget.dart';

class HomePage extends StatelessWidget {
  const HomePage({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('装饰容器DecoratedBox')),
      body: Container(
        margin: const EdgeInsets.all(5.0),
        child: Column(
          children: const [
            UseDecoratedBox01Widget(),
          ],
        ),
      ),
    );
  }
}
