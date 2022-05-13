import 'package:flutter/material.dart';

import 'widgets/use_animated_list_widget.dart';

class HomePage extends StatelessWidget {
  const HomePage({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('AnimatedList')),
      body: const SizedBox(
        child: UseAnimatedListWidget(),
      ),
    );
  }
}
