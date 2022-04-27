import 'package:flutter/material.dart';

import 'widgets/use_stack_and_positioned_widget.dart';
import 'widgets/use_stack_and_positioned_02_widget.dart';

class HomePage extends StatelessWidget {
  const HomePage({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('层叠布局Stack、Positioned')),
      body: const UseStackAndPositioned02Widget(),
    );
  }
}
