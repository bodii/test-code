import 'package:flutter/material.dart';

import 'widgets/stagger_route.dart';

class HomePage extends StatelessWidget {
  const HomePage({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('stagger animation 交织动画')),
      body: const StaggerRoute(),
    );
  }
}
