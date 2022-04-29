import 'package:flutter/material.dart';

import 'widgets/use_container_01_widget.dart';
import 'widgets/use_container_02_widget.dart';

class HomePage extends StatelessWidget {
  const HomePage({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('Container')),
      body: Column(
        children: const [
          UseContainer01Widget(),
          UseContainer02Widget(),
        ],
      ),
    );
  }
}
