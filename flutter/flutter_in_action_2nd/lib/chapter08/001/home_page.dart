import 'package:flutter/material.dart';

import 'widgets/pointer_move_indicator_widget.dart';
import 'widgets/use_absorbPointer_listener_widget.dart';

class HomePage extends StatelessWidget {
  const HomePage({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('原始指针事件处理')),
      body: const UseAbsorbPointerListenerWidget(),
    );
  }
}
