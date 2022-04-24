import 'package:flutter/material.dart';
import 'widgets/switch_and_check_box_widget.dart';

class HomePage extends StatelessWidget {
  const HomePage({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('switch and checkbox demo'),
      ),
      body: const SwitchAndCheckBoxWidget(),
    );
  }
}
