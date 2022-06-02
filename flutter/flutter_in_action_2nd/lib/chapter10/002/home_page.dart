import 'package:flutter/material.dart';

import 'widgets/use_turn_box_widget.dart';

class HomePage extends StatelessWidget {
  const HomePage({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('custom TrunBox widget')),
      body: const UseTurnBoxWidget(),
    );
  }
}
